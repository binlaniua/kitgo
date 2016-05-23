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
	Timeout       time.Duration
	DefaultHeader map[string]string
	DefaultCookie map[string]map[string]string
	UseAgent      string
	DefaultRefer  string
	Debug         bool
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
	if o.Proxy != "" {
		p, _ := url.Parse(o.Proxy)
		t := &http.Transport{
			Proxy: http.ProxyURL(p),
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
			log.Fatal("证书加载出错 " + o.SSLCertPath + " : " + o.SSLKeyPath)
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
	if c.transport == nil {
		c.transport = &http.Transport{}
		c.client.Transport = c.transport
	}
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return err
	}
	c.transport.TLSClientConfig = &tls.Config{Certificates: []tls.Certificate{cert}}
	//kitgo.Log("设置证书成功 " + certPath + " : " + keyPath)
	return nil
}

//-------------------------------------
//
// 
//
//-------------------------------------
func (c *HttpClient) SetTimeout(to time.Duration) error {
	if c.transport == nil {
		c.transport = &http.Transport{}
		c.client.Transport = c.transport
	}
	c.client.Timeout = to * time.Second
	c.transport.Dial = func(netw, addr string) (net.Conn, error) {
		c, err := net.DialTimeout(netw, addr, time.Second * to)
		if err != nil {
			return nil, err
		}
		c.SetDeadline(time.Now().Add(to * time.Second))
		return c, nil
	}
	return nil
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
	req.Header.Add("Connection", "close")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return NewHttpResult(resp, req.URL.String()), nil
}
