package http

import (
	"net/http/cookiejar"
	"net/http"
	"net/url"
	"strings"
	"encoding/json"
	"bytes"
	"mime/multipart"
	"io"
	"golang.org/x/net/proxy"
	"net"
	"time"
	"log"
	"crypto/tls"
	"github.com/binlaniua/kitgo/file"
	"errors"
)

type HttpClient struct {
	client    *http.Client
	transport *http.Transport
	cookie    *cookiejar.Jar
	option    *HttpClientOption
}

type HttpClientOption struct {
	Proxy         string
	ProxySock5    string

	SSLKeyPath    string
	SSLCertPath   string
	SSLKeyData    []byte
	SSLCertData   []byte

	Timeout       time.Duration
	DefaultHeader map[string]string
	DefaultCookie map[string]map[string]string
	UseAgent      string
	DefaultRefer  string
	Debug         bool

	IsLazy        bool
}


//-------------------------------------
//
//
//
//-------------------------------------
func NewHttpClient(o *HttpClientOption) *HttpClient {
	hc := &HttpClient{}
	hc.option = o
	j, _ := cookiejar.New(&cookiejar.Options{})
	hc.cookie = j
	c := &http.Client{
		Jar:j,
	}
	hc.client = c

	//
	if o.UseAgent == "" {
		o.UseAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.94 Safari/537.36"
	}
	if o.DefaultHeader == nil {
		o.DefaultHeader = map[string]string{
			"Accept":"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
			"Accept-Language":"zh-CN,zh;q=0.8,en-US;q=0.6,en;q=0.4",
			"Cache-Control":"max-age=0",
			"Connection":"keep-alive",
			"Upgrade-Insecure-Requests":"1",
		}
	}

	//
	if o.Proxy != "" {
		p, _ := url.Parse(o.Proxy)
		t := &http.Transport{
			Proxy: http.ProxyURL(p),
			DisableKeepAlives: true,
		}
		c.Transport = t
		hc.transport = t
	} else if o.ProxySock5 != "" {
		dialer, err := proxy.SOCKS5("tcp", o.ProxySock5,
			nil,
			&net.Dialer{
				Timeout: o.Timeout * time.Second,
				KeepAlive: o.Timeout * time.Second,
			},
		)
		if err != nil {
			log.Println("sock5 连接失败 => ", err)
		}
		t := &http.Transport{
			Proxy: nil,
			Dial: dialer.Dial,
			DisableKeepAlives: true,
			TLSHandshakeTimeout: o.Timeout * time.Second,
		}
		c.Transport = t
		hc.transport = t
	}
	if o.DefaultCookie != nil {
		hc.SetCookie(o.DefaultCookie)
	}
	if o.Timeout > 0 {
		hc.SetTimeout(o.Timeout)
	}
	if o.SSLCertPath != "" && o.SSLKeyPath != "" {
		err := hc.SetSSL(o.SSLCertPath, o.SSLKeyPath)
		if err != nil {
			panic(err)
			return nil
		}
	}
	if len(o.SSLKeyData) > 0 && len(o.SSLCertData) > 0 {
		err := hc.SetSSLData(o.SSLCertData, o.SSLKeyData)
		if err != nil {
			panic(err)
			return nil
		}
	}
	return hc
}

//-------------------------------------
//
//
//
//-------------------------------------
func (c *HttpClient) SetSSL(certPath string, keyPath string) error {
	cd, err := file.ReadBytes(certPath)
	if err != nil {
		return err
	}
	kd, err := file.ReadBytes(keyPath)
	if err != nil {
		return err
	}
	return c.SetSSLData(cd, kd)
}

func (c *HttpClient) SetSSLData(certData, keyData []byte) error {
	if c.transport == nil {
		c.transport = &http.Transport{
			DisableKeepAlives: true,
		}
		c.client.Transport = c.transport
	}
	cert, err := tls.X509KeyPair(certData, keyData)
	if err != nil {
		return err
	}
	c.transport.TLSClientConfig = &tls.Config{Certificates: []tls.Certificate{cert}}
	return nil
}

//-------------------------------------
//
// 
//
//-------------------------------------
func (c *HttpClient) SetTimeout(to time.Duration) error {
	if c.transport == nil {
		c.transport = &http.Transport{
			DisableKeepAlives: true,
		}
		c.client.Transport = c.transport
	}
	c.client.Timeout = to * time.Second
	return nil
}

//-------------------------------------
//
//
//
//-------------------------------------
var errorNotFollowRedirect = errors.New("not follow redirect")

func (c *HttpClient) NotFollowRedirect() {
	c.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errorNotFollowRedirect
	}
}

//-------------------------------------
//
// 
//
//-------------------------------------
func (c *HttpClient) SetCookie(mm map[string]map[string]string) {
	for site, siteCs := range mm {
		u, _ := url.Parse(site)
		cs := make([]*http.Cookie, 0)
		for key, val := range siteCs {
			c := &http.Cookie{Name: key, Value: val}
			cs = append(cs, c)
		}
		c.cookie.SetCookies(u, cs)
	}
}

//
//
//
//
//
func (c *HttpClient) Get(urlStr string) (*HttpResult, error) {
	req, _ := http.NewRequest("GET", urlStr, nil)
	return c.doRequest(req)
}

//-------------------------------------
//
//
//
//-------------------------------------
func (c *HttpClient) GetReply(urlStr string, reply int) (*HttpResult, error) {
	r, err := c.Get(urlStr)
	for i := 0; i < reply; i++ {
		if err == nil {
			return r, nil
		} else {
			r, err = c.Get(urlStr)
		}
	}
	return r, err
}

//-------------------------------------
//
//
//
//-------------------------------------
func (c *HttpClient) Post(urlStr string, dataMap map[string]string) (*HttpResult, error) {
	reqParams := url.Values{}
	if dataMap != nil {
		for k, v := range dataMap {
			reqParams.Add(k, v)
		}
	}
	req, _ := http.NewRequest("POST", urlStr, strings.NewReader(reqParams.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return c.doRequest(req)
}

//-------------------------------------
//
//
//
//-------------------------------------
func (c *HttpClient) PostReply(urlStr string, dataMap map[string]string, reply int) (*HttpResult, error) {
	r, err := c.Post(urlStr, dataMap)
	for i := 0; i < reply; i++ {
		if err == nil {
			return r, nil
		} else {
			r, err = c.Post(urlStr, dataMap)
		}
	}
	return r, err
}

//-------------------------------------
//
//
//
//-------------------------------------
func (c *HttpClient) PostJson(urlStr string, data interface{}) (*HttpResult, error) {
	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", urlStr, bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	return c.doRequest(req)
}

//-------------------------------------
//
//
//
//-------------------------------------
func (c *HttpClient) PostXMLString(urlStr string, src string) (*HttpResult, error) {
	req, _ := http.NewRequest("POST", urlStr, bytes.NewBuffer([]byte(src)))
	req.Header.Add("Content-Type", "application/xml;charset=utf-8")
	return c.doRequest(req)
}

//-------------------------------------
//
//
//
//-------------------------------------
func (c *HttpClient) PostFile(urlStr string, dataMap map[string]interface{}) (*HttpResult, error) {
	buff := &bytes.Buffer{}
	write := multipart.NewWriter(buff)
	if dataMap != nil {
		for key, val := range dataMap {
			switch val.(type) {
			case string:
				write.WriteField(key, val.(string))
			default:
				w, _ := write.CreateFormField(key)
				io.Copy(w, val.(io.Reader))
			}
		}
	}
	req, _ := http.NewRequest("POST", urlStr, buff)
	req.Header.Add("Content-Type", write.FormDataContentType())
	return c.doRequest(req)
}

//-------------------------------------
//
// 
//
//-------------------------------------
func (c *HttpClient) doRequest(req *http.Request) (*HttpResult, error) {
	if c.option.DefaultHeader != nil {
		for k, v := range c.option.DefaultHeader {
			req.Header.Add(k, v)
		}
	}
	if c.option.UseAgent != "" {
		req.Header.Add("User-Agent", c.option.UseAgent)
	}
	req.Close = true
	resp, err := c.client.Do(req)
	if err != nil {
		if e, ok := err.(*url.Error); ok && e.Err != errorNotFollowRedirect {
			return nil, err
		}
	}
	return NewHttpResult(resp, c.option.IsLazy), nil
}
