package swagger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"

	"gopkg.in/yaml.v3"
)

type API struct {
	Swagger string
	Info    struct {
		Version string
		Title   string
	}
	Schemes []string
	Host    string

	rawSwaggers []string
	paths       map[string]*Path
	models      map[string]*Schema
}

func NewAPI() *API {
	return &API{
		paths:  make(map[string]*Path),
		models: make(map[string]*Schema),
	}
}

type rawSwagger struct {
	Swagger string `json:"swagger"`
	Info    struct {
		Version string `json:"version"`
		Title   string `json:"title"`
	} `json:"info"`
	Schemes     []string               `json:"schemes"`
	Host        string                 `json:"host"`
	Paths       map[string]interface{} `json:"paths" yaml:"paths"`
	Definitions map[string]interface{} `json:"definitions" yaml:"definitions"`
	Parameters  map[string]interface{} `json:"parameters" yaml:"parameters"`

	SecurityDefinitions map[string]interface{} `json:"securityDefinitions" yaml:"securityDefinitions"`
	Security            []interface{}          `json:"security" yaml:"security"`
}

func (api *API) MarshalJSON() ([]byte, error) {
	var nativeOut struct {
		Paths       map[string]*Path   `json:"paths"`
		Definitions map[string]*Schema `json:"definitions,omitempty"`
	}

	nativeOut.Definitions = api.models
	nativeOut.Paths = make(map[string]*Path)
	for p, path := range api.paths {
		if path.HasActions() {
			nativeOut.Paths[p] = path
		}
	}

	native, err := json.Marshal(nativeOut)
	if err != nil {
		return native, fmt.Errorf("error marshaling native out: %w", err)
	}

	var raw rawSwagger
	err = json.Unmarshal(native, &raw)
	if err != nil {
		return native, fmt.Errorf("error unmarshaling raw swagger: %w", err)
	}
	raw.Parameters = make(map[string]interface{})
	raw.SecurityDefinitions = make(map[string]interface{})
	raw.Security = make([]interface{}, 0)

	for _, file := range api.rawSwaggers {
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			return native, err
		}
		var y rawSwagger
		err = yaml.Unmarshal(bytes, &y)
		if err != nil {
			return native, fmt.Errorf("error unmarshaling swagger yml %s: %w", file, err)
		}
		for p, path := range y.Paths {
			_, ok := raw.Paths[p]
			if ok {
				return native, fmt.Errorf("path %s from swagger file %s already exists", p, file)
			}
			for k, v := range path.(map[string]interface{}) {
				if k == "get" || k == "post" || k == "put" || k == "delete" || k == "options" || k == "head" || k == "patch" || k == "trace" {
					responses := make(map[string]interface{})
					if _, ok := v.(map[string]interface{})["responses"].(map[interface{}]interface{}); ok {
						for code, response := range v.(map[string]interface{})["responses"].(map[interface{}]interface{}) {
							if reflect.TypeOf(code).Kind() == reflect.String {
								responses[code.(string)] = response
							} else if reflect.TypeOf(code).Kind() == reflect.Int {
								responses[fmt.Sprintf("%d", code.(int))] = response
							}
						}
						path.(map[string]interface{})[k].(map[string]interface{})["responses"] = responses
					}
				}
			}
			raw.Paths[p] = path
		}
		for d, def := range y.Definitions {
			_, ok := raw.Definitions[d]
			if ok {
				return native, fmt.Errorf("definition %s from swagger file %s already exists", d, file)
			}
			raw.Definitions[d] = def
		}
		for s, sec := range y.SecurityDefinitions {
			raw.SecurityDefinitions[s] = sec
		}
		raw.Security = append(raw.Security, y.Security...)
		for p, param := range y.Parameters {
			_, ok := raw.Parameters[p]
			if ok {
				return native, fmt.Errorf("parameter %s from swagger file %s already exists", p, file)
			}
			raw.Parameters[p] = param
		}
	}

	raw.Swagger = api.Swagger
	raw.Info.Version = api.Info.Version
	raw.Info.Title = api.Info.Title
	raw.Schemes = api.Schemes
	raw.Host = api.Host

	return json.Marshal(raw)
}

func (api *API) AddSwaggerYAML(path string) {
	api.rawSwaggers = append(api.rawSwaggers, path)
}

func (api *API) Path(path string) *Path {
	p, ok := api.paths[path]
	if ok {
		return p
	}

	p = &Path{
		api:         api,
		path:        path,
		permissions: make(map[string]bool),
		parameters:  make([]*Param, 0),
		tags:        make(map[string]bool),
		produces:    make(map[string]bool),
		consumes:    make(map[string]bool),
	}
	api.paths[path] = p
	return p
}

func (api *API) Model(name string) *Schema {
	m := NewSchema(name)
	m.Name = name

	key := m.SafeName()

	_, ok := api.models[key]
	if ok {
		panic("Model already exists: " + key)
	}

	api.models[key] = m
	return m
}
