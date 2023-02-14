package swagger

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Header struct {
	Name        string `json:"-"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Response struct {
	Description string   `json:"description,omitempty"`
	Schema      *Schema  `json:"schema,omitempty"`
	Headers     []Header `json:"headers,omitempty"`
}

type Action struct {
	method      string
	permissions map[string]bool
	tags        map[string]bool
	produces    map[string]bool
	consumes    map[string]bool
	description string
	body        *Schema
	parent      *Path
	parameters  []*Param
	responses   map[int]*Response
}

func (a *Action) Permission(perm string) *Action {
	a.permissions[perm] = true
	return a
}

func (a *Action) Param(param *Param) *Action {
	a.parameters = append(a.parameters, param)
	return a
}

func (a *Action) computeTags() []string {
	tagset := make(map[string]bool)
	for t := range a.tags {
		tagset[t] = true
	}

	p := a.parent
	for p != nil {
		for t := range p.tags {
			tagset[t] = true
		}
		p = p.parent
	}

	tags := make([]string, 0, len(tagset))
	for t := range tagset {
		tags = append(tags, t)
	}

	return tags
}

func (a *Action) computeConsumes() []string {
	cset := make(map[string]bool)
	for t := range a.consumes {
		cset[t] = true
	}

	p := a.parent
	for p != nil {
		for t := range p.consumes {
			cset[t] = true
		}
		p = p.parent
	}

	consumes := make([]string, 0, len(cset))
	for t := range cset {
		consumes = append(consumes, t)
	}

	return consumes
}

func (a *Action) computePermissions() []string {
	pset := make(map[string]bool)
	for t := range a.permissions {
		pset[t] = true
	}

	p := a.parent
	for p != nil {
		for t := range p.permissions {
			pset[t] = true
		}
		p = p.parent
	}

	permissions := make([]string, 0, len(pset))
	for t := range pset {
		permissions = append(permissions, t)
	}

	return permissions
}

func (a *Action) computeProduces() []string {
	pset := make(map[string]bool)
	for t := range a.produces {
		pset[t] = true
	}

	p := a.parent
	for p != nil {
		for t := range p.produces {
			pset[t] = true
		}
		p = p.parent
	}

	produces := make([]string, 0, len(pset))
	for t := range pset {
		produces = append(produces, t)
	}

	return produces
}

func (a *Action) Response(code int, description string, schema *Schema, headers ...Header) *Action {
	a.responses[code] = &Response{
		Description: description,
		Schema:      schema,
		Headers:     headers,
	}
	return a
}

func (r *Response) MarshalJSON() ([]byte, error) {
	var out struct {
		Description string            `json:"description,omitempty"`
		Schema      *Schema           `json:"schema,omitempty"`
		Headers     map[string]Header `json:"headers,omitempty"`
	}

	out.Description = r.Description
	out.Schema = r.Schema
	if len(r.Headers) > 0 {
		out.Headers = make(map[string]Header)
		out.Headers = make(map[string]Header)
		for _, h := range r.Headers {
			out.Headers[h.Name] = h
		}
	}

	return json.Marshal(out)
}

func (a *Action) MarshalJSON() ([]byte, error) {
	var out struct {
		Description string            `json:"description,omitempty"`
		Tags        []string          `json:"tags,omitempty"`
		Produces    []string          `json:"produces,omitempty"`
		Responses   map[int]*Response `json:"responses,omitempty"`
		Parameters  []*Param          `json:"parameters,omitempty"`
		Consumes    []string          `json:"consumes,omitempty"`
	}

	out.Description = a.description
	out.Tags = a.computeTags()
	out.Produces = a.computeProduces()
	out.Responses = a.responses
	out.Parameters = a.parameters

	if a.method == http.MethodPost || a.method == http.MethodPut {
		out.Consumes = a.computeConsumes()
	}

	permissions := a.computePermissions()
	for i := range permissions {
		permissions[i] = "`" + permissions[i] + "`"
	}

	if len(permissions) > 1 {
		out.Description += "\n\n---\n\nRequires " + strings.Join(permissions, ",") + " permissions."
	} else if len(permissions) == 1 {
		out.Description += "\n\n---\n\nRequires " + permissions[0] + " permission."
	}

	return json.Marshal(out)
}
