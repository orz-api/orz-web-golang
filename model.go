package orzweb

type ErrorRsp struct {
	Code   string          `json:"code,omitempty"`
	Reason string          `json:"reason,omitempty"`
	Traces []*ErrorTraceTo `json:"traces,omitempty"`
}

type ErrorTraceTo struct {
	Service  string `json:"service,omitempty"`
	Endpoint string `json:"endpoint,omitempty"`
	Details  string `json:"details,omitempty"`
}

type ProtocolBo struct {
	Version int
	Code    string
	Notice  string
}

func NewErrorProtocol(version *int, code string, notice string) *ProtocolBo {
	var protocol = &ProtocolBo{
		Version: VERSION_CURRENT,
		Code:    CODE_UNDEFINED,
		Notice:  notice,
	}
	if version != nil {
		protocol.Version = *version
	}
	if code != "" {
		protocol.Code = code
	}
	return protocol
}
