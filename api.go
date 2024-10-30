package orzweb

import "strconv"

type ApiDef[Req any, Rsp any] struct {
	Scope       string
	Domain      string
	Resource    string
	Action      string
	Variant     int
	Query       bool
	Description string
	Errors      ApiErrorsDef
	Handler     ApiHandlerDef[Req, Rsp]
}

type ApiHandlerDef[Req any, Rsp any] func(req *Req) (*Rsp, error)

type ApiErrorDef struct {
	Reason      string
	Notice      string
	Alarm       bool
	Logging     bool
	Description string
}

type ApiErrorsDef map[string]*ApiErrorDef

func (def *ApiDef[Req, Rsp]) buildPath() string {
	// TODO: check def.Scope, def.Domain, def.Action, def.Variant
	path := "/" + def.Scope + "/" + def.Domain + "/"
	if def.Resource != "" {
		path += def.Resource
	}
	path += def.Action + "V" + strconv.Itoa(def.Variant)
	return path
}
