package orzweb

type Props struct {
	BaseUri           string                     `json:"baseUri,omitempty" yaml:"baseUri,omitempty" toml:"baseUri,omitempty"`
	ExposeErrorReason bool                       `json:"exposeErrorReason,omitempty" yaml:"exposeErrorReason,omitempty" toml:"exposeErrorReason,omitempty"`
	ExposeErrorTraces bool                       `json:"exposeErrorTraces,omitempty" yaml:"exposeErrorTraces,omitempty" toml:"exposeErrorTraces,omitempty"`
	RequestHeader     PropsRequestHeadersConfig  `json:"requestHeader,omitempty" yaml:"requestHeader,omitempty" toml:"requestHeader,omitempty"`
	ResponseHeader    PropsResponseHeadersConfig `json:"responseHeader,omitempty" yaml:"responseHeader,omitempty" toml:"responseHeader,omitempty"`
	Page              PropsPageConfig            `json:"page,omitempty" yaml:"page,omitempty" toml:"page,omitempty"`
	Cors              PropsCorsConfig            `json:"cors,omitempty" yaml:"cors,omitempty" toml:"cors,omitempty"`
}

type PropsRequestHeadersConfig struct {
	// TODO
}

type PropsResponseHeadersConfig struct {
	Version string `json:"version,omitempty" yaml:"version,omitempty" toml:"version,omitempty"`
	Code    string `json:"code,omitempty" yaml:"code,omitempty" toml:"code,omitempty"`
	Notice  string `json:"notice,omitempty" yaml:"notice,omitempty" toml:"notice,omitempty"`
}

type PropsPageConfig struct {
	DefaultSize int `json:"defaultSize,omitempty" yaml:"defaultSize,omitempty" toml:"defaultSize,omitempty"`
	MaxSize     int `json:"maxSize,omitempty" yaml:"maxSize,omitempty" toml:"maxSize,omitempty"`
}

type PropsCorsConfig struct {
	// TODO
}

var PropsObj = Props{
	BaseUri:           "",
	ExposeErrorReason: false,
	ExposeErrorTraces: false,
	ResponseHeader: PropsResponseHeadersConfig{
		Version: "Orz-Version",
		Code:    "Orz-Code",
		Notice:  "Orz-Notice",
	},
	Page: PropsPageConfig{
		DefaultSize: 50,
		MaxSize:     100,
	},
}
