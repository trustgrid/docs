package swagger

import (
	"encoding/json"
	"net/http"
)

type Path struct {
	tags map[string]bool

	path        string
	permissions map[string]bool
	produces    map[string]bool
	consumes    map[string]bool
	parameters  []*Param
	api         *API
	parent      *Path

	get    *Action
	post   *Action
	put    *Action
	patch  *Action
	delete *Action
}

func (p *Path) Permission(perm string) *Path {
	p.permissions[perm] = true
	return p
}

func (p *Path) Produces(contentType string) *Path {
	p.produces[contentType] = true
	return p
}

func (p *Path) Consumes(contentType string) *Path {
	p.consumes[contentType] = true
	return p
}

func (p *Path) Tag(tag string) *Path {
	p.tags[tag] = true
	return p
}

func (p *Path) Path(path string) *Path {
	c := p.api.Path(p.path + path)
	c.parent = p
	return c
}

func (p *Path) PathParam(param *Param) *Path {
	c := p.api.Path(p.path + "/{" + param.Name + "}")
	c.parameters = []*Param{param}
	c.parent = p
	return c
}

func (p *Path) HasActions() bool {
	return p.get != nil || p.post != nil || p.delete != nil || p.put != nil || p.patch != nil
}

func (p *Path) MarshalJSON() ([]byte, error) {
	var out struct {
		Get        *Action  `json:"get,omitempty"`
		Post       *Action  `json:"post,omitempty"`
		Put        *Action  `json:"put,omitempty"`
		Patch      *Action  `json:"patch,omitempty"`
		Delete     *Action  `json:"delete,omitempty"`
		Parameters []*Param `json:"parameters,omitempty"`
	}

	out.Get = p.get
	out.Post = p.post
	out.Put = p.put
	out.Delete = p.delete
	out.Patch = p.patch

	out.Parameters = make([]*Param, len(p.parameters))
	copy(out.Parameters, p.parameters)
	parent := p.parent
	for parent != nil {
		out.Parameters = append(out.Parameters, parent.parameters...)
		parent = parent.parent
	}

	return json.Marshal(out)
}

func (p *Path) newAction(method, description string) *Action {
	return &Action{
		method:      method,
		description: description,
		parent:      p,
		tags:        make(map[string]bool),
		produces:    make(map[string]bool),
		permissions: make(map[string]bool),
		consumes:    make(map[string]bool),
		responses:   make(map[int]*Response),
	}
}

func (p *Path) Get(description string) *Action {
	p.get = p.newAction(http.MethodGet, description)
	return p.get
}

func (p *Path) Post(description string) *Action {
	p.post = p.newAction(http.MethodPost, description)
	return p.post
}

func (p *Path) Put(description string) *Action {
	p.put = p.newAction(http.MethodPut, description)
	return p.put
}

func (p *Path) Patch(description string) *Action {
	p.patch = p.newAction(http.MethodPatch, description)
	return p.patch
}

func (p *Path) Delete(description string) *Action {
	p.delete = p.newAction(http.MethodDelete, description)
	return p.delete
}
