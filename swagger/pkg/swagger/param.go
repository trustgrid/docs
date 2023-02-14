package swagger

type ParamPosition string

const (
	PP_Query ParamPosition = "query"
	PP_Path  ParamPosition = "path"
	PP_Body  ParamPosition = "body"
)

type ParamType string

const (
	PT_String  ParamType = "string"
	PT_Number  ParamType = "number"
	PT_Boolean ParamType = "boolean"
	PT_Array   ParamType = "array"
)

type Param struct {
	Name        string        `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
	In          ParamPosition `json:"in"`
	Required    bool          `json:"required,omitempty"`
	Type        ParamType     `json:"type,omitempty"`
	Schema      *Schema       `json:"schema,omitempty"`
	Enum        []interface{} `json:"enum,omitempty"`
}

type ParamConfig func(p *Param)

func P_Required(p *Param) {
	p.Required = true
}

func P_Array(p *Param) {
	p.Type = PT_Array
}

func P_Number(p *Param) {
	p.Type = PT_Number
}

func P_Boolean(p *Param) {
	p.Type = PT_Boolean
}

func P_Path(p *Param) {
	p.In = PP_Path
}

func P_Query(p *Param) {
	p.In = PP_Query
}

func P_Enum(values ...interface{}) ParamConfig {
	return func(p *Param) {
		p.Enum = values
	}
}

func P_Body(p *Param) {
	p.In = PP_Body
	p.Type = ""
}

func P_Schema(s *Schema) ParamConfig {
	return func(p *Param) {
		p.Schema = s
	}
}

func NewParam(name string, desc string, config ...ParamConfig) *Param {
	p := &Param{
		Name:        name,
		Description: desc,
		In:          PP_Query,
		Required:    false,
		Type:        PT_String,
	}
	for _, c := range config {
		c(p)
	}
	return p
}
