package kitgo

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
)

type HttpClient struct {
	client *http.Client
	cookie *cookiejar.Jar
	option *HttpClientOption
}

type HttpClientOption struct {
	Proxy         string
	ProxySock5    string
	DefaultHeader map[string]string
	DefaultCookie map[string]map[string]string
	UseAgent      string
	DefaultRefer  string
	Debug         bool
}

func NewHttpClient(o *HttpClientOption) *HttpClient {
	j, _ := cookiejar.New(&cookiejar.Options{})
	c := &http.Client{
		Jar:j,
	}

	//
	if o.Proxy != "" {
		p, _ := url.Parse(o.Proxy)
		c.Transport = &http.Transport{
			Proxy: http.ProxyURL(p),
		}
	} else if o.ProxySock5 != "" {
		dialer, err := proxy.SOCKS5("tcp", o.ProxySock5,
			nil,
			&net.Dialer{
				Timeout: 30 * time.Second,
				KeepAlive: 30 * time.Second,
			},
		)
		if err != nil {
			log.Println("sock5 连接失败 => ", err)
		}
		c.Transport = &http.Transport{
			Proxy: nil,
			Dial: dialer.Dial,
			TLSHandshakeTimeout: 10 * time.Second,
		}
	}
	hc := &HttpClient{c, j, o}
	if o.DefaultCookie != nil {
		hc.SetCookie(o.DefaultCookie)
	}
	return hc
}

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
	//log.Println(c.cookie)
}

//
//
//
//
//
func (c *HttpClient) Get(urlStr string) *HttpResult {
	req, _ := http.NewRequest("GET", urlStr, nil)
	return c.doRequest(req)
}

//-------------------------------------
//
//
//
//-------------------------------------
func (c *HttpClient) GetReply(urlStr string, reply int) *HttpResult {
	r := c.Get(urlStr)
	for i := 0; i < reply; i++ {
		if r != nil {
			return r
		} else {
			r = c.Get(urlStr)
		}
	}
	return r
}

//
//
//
//
//
func (c *HttpClient) Post(urlStr string, dataMap map[string]string) *HttpResult {
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
func (c *HttpClient) PostReply(urlStr string, dataMap map[string]string, reply int) *HttpResult {
	r := c.Post(urlStr, dataMap)
	for i := 0; i < reply; i++ {
		if r != nil {
			return r
		} else {
			r = c.Post(urlStr, dataMap)
		}
	}
	return r
}

//
//
//
//
//
func (c *HttpClient) PostJson(urlStr string, data interface{}) *HttpResult {
	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", urlStr, bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	return c.doRequest(req)
}

//
//
//
//
//
func (c *HttpClient) PostFile(urlStr string, dataMap map[string]interface{}) *HttpResult {
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

//
//
//
//
//
func (c *HttpClient) doRequest(req *http.Request) *HttpResult {
	if c.option.DefaultHeader != nil {
		for k, v := range c.option.DefaultHeader {
			req.Header.Add(k, v)
		}
	}
	if c.option.UseAgent != "" {
		req.Header.Add("User-Agent", c.option.UseAgent)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		if c.option.Debug {
			log.Println(req.URL.String(), " 获取失败 => ", err)
		}
		return nil
	}
	return NewHttpResult(resp)
}
