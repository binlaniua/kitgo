package http

import (
	"time"
)

var (
	defaultUseAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.94 Safari/537.36"
	defaultHeader   = map[string]string{
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Accept-Language":           "zh-CN,zh;q=0.8,en-US;q=0.6,en;q=0.4",
		"Cache-Control":             "max-age=0",
	}
)

//-------------------------------------
//
//
//
//-------------------------------------
type HttpClientOption struct {
	Proxy      string
	ProxySock5 string

	NoSsl       bool
	SSLKeyPath  string
	SSLCertPath string
	SSLKeyData  []byte
	SSLCertData []byte

	Timeout       time.Duration
	DefaultHeader map[string]string
	DefaultCookie map[string]map[string]string
	UseAgent      string
	DefaultRefer  string
	Debug         bool

	IsLazy bool
}
