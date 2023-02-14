package swagger

import (
	"encoding/json"
	"strings"
)

type SchemaType string

const (
	ST_Object  SchemaType = "object"
	ST_Number  SchemaType = "number"
	ST_String  SchemaType = "string"
	ST_Array   SchemaType = "array"
	ST_Boolean SchemaType = "boolean"
)

type Schema struct {
	Name        string
	Required    bool
	Type        SchemaType
	Description string
	Properties  map[string]*Schema
	Items       *Schema
	Example     interface{}
	Default     interface{}
	Enum        []string
	ref         bool
}

type SchemaConfig func(s *Schema)

func S_Number(s *Schema) {
	s.Type = ST_Number
}

func S_Required(s *Schema) {
	s.Required = true
}

func S_String(s *Schema) {
	s.Type = ST_String
}

func S_Boolean(s *Schema) {
	s.Type = ST_Boolean
}

func S_Enum(values ...string) SchemaConfig {
	return func(s *Schema) {
		s.Enum = values
	}
}

func S_Example(ex interface{}) SchemaConfig {
	return func(s *Schema) {
		s.Example = ex
	}
}

func S_Default(d interface{}) SchemaConfig {
	return func(s *Schema) {
		s.Default = d
	}
}

func (s *Schema) Ref() *Schema {
	return &Schema{
		ref:  true,
		Name: s.Name,
	}
}

func NewSchema(description string, config ...SchemaConfig) *Schema {
	s := &Schema{
		Type:        ST_Object,
		Description: description,
		Properties:  make(map[string]*Schema),
	}
	for _, c := range config {
		c(s)
	}
	return s
}

func NewArraySchema(items *Schema) *Schema {
	return &Schema{
		Type:  ST_Array,
		Items: items,
	}
}

func (s *Schema) Prop(name string, schema *Schema) *Schema {
	s.Properties[name] = schema
	return s
}

func (s *Schema) SafeName() string {
	return strings.ReplaceAll(s.Name, " ", "")
}

func (s *Schema) MarshalJSON() ([]byte, error) {
	if s.ref {
		return json.Marshal(map[string]interface{}{
			"$ref": "#/definitions/" + s.SafeName(),
		})
	}

	var out struct {
		Name        string             `json:"title,omitempty"`
		Required    []string           `json:"required,omitempty"`
		Type        SchemaType         `json:"type,omitempty"`
		Description string             `json:"description,omitempty"`
		Properties  map[string]*Schema `json:"properties,omitempty"`
		Items       *Schema            `json:"items,omitempty"`
		Example     interface{}        `json:"example,omitempty"`
		Default     interface{}        `json:"default,omitempty"`
		Enum        []string           `json:"enum,omitempty"`
	}

	out.Name = s.Name
	out.Type = s.Type
	out.Description = s.Description
	out.Properties = s.Properties
	out.Items = s.Items
	out.Example = s.Example
	out.Enum = s.Enum
	out.Default = s.Default
	for k, p := range s.Properties {
		if p.Required {
			out.Required = append(out.Required, k)
		}
	}

	return json.Marshal(out)
}
