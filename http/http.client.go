package http

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"github.com/binlaniua/kitgo"
	"github.com/binlaniua/kitgo/file"
	"golang.org/x/net/proxy"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

//-------------------------------------
//
//
//
//-------------------------------------
type HttpClient struct {
	client *http.Client
	cookie *cookiejar.Jar
	option *HttpClientOption
}

//
//
//
//
//
func NewHttpClient(o *HttpClientOption) *HttpClient {
	hc := &HttpClient{
		option: o,
	}

	//
	j, _ := cookiejar.New(&cookiejar.Options{})
	hc.cookie = j
	c := &http.Client{
		Jar: j,
	}
	hc.client = c

	//
	if o.UseAgent == "" {
		o.UseAgent = defaultUseAgent
	}

	//
	if o.DefaultHeader == nil {
		o.DefaultHeader = defaultHeader
	}

	//
	if o.Proxy != "" {
		hc.SetProxy(o.Proxy)
	} else if o.ProxySock5 != "" {
		hc.SetProxySock5(o.ProxySock5)
	}

	//
	if o.DefaultCookie != nil {
		hc.SetCookie(o.DefaultCookie)
	}

	//
	if o.Timeout > 0 {
		hc.SetTimeout(o.Timeout)
	}

	if o.NoSsl {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		c.Transport = tr
	}

	//
	if o.SSLCertPath != "" && o.SSLKeyPath != "" {
		err := hc.SetSSL(o.SSLCertPath, o.SSLKeyPath)
		if err != nil {
			panic(err)
			return nil
		}
	} else if len(o.SSLKeyData) > 0 && len(o.SSLCertData) > 0 {
		err := hc.SetSSLData(o.SSLCertData, o.SSLKeyData)
		if err != nil {
			panic(err)
			return nil
		}
	}

	//
	return hc
}

//
//
//
//
//
func (hc *HttpClient) SetProxy(proxy string) error {
	p, err := url.Parse(proxy)
	if err != nil {
		kitgo.ErrorLog.Println("设置代理[ %s ]失败 => [ %v ]", proxy, err)
		return err
	}
	t := &http.Transport{
		Proxy:             http.ProxyURL(p),
		DisableKeepAlives: true,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
	}
	hc.client.Transport = t
	return nil
}

//
//
//
//
//
func (hc *HttpClient) SetProxySock5(proxyString string) error {
	dialer, err := proxy.SOCKS5("tcp", proxyString, nil,
		&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		},
	)
	if err != nil {
		kitgo.ErrorLog.Println("设置SOCK5代理[ %s ]失败 => [ %v ]", proxyString, err)
		return err
	}
	t := &http.Transport{
		Proxy:               nil,
		Dial:                dialer.Dial,
		DisableKeepAlives:   true,
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	hc.client.Transport = t
	return nil
}

//
//
//
//
//
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

//
//
//
//
//
func (c *HttpClient) SetSSLData(certData, keyData []byte) error {
	cert, err := tls.X509KeyPair(certData, keyData)
	if err != nil {
		return err
	}
	var ht *http.Transport
	if c.client.Transport == nil {
		ht = &http.Transport{
		}
	} else {
		ht = c.client.Transport.(*http.Transport)
	}
	ht.TLSClientConfig = &tls.Config{Certificates: []tls.Certificate{cert}}
	c.client.Transport = ht
	return nil
}

//
//
//
//
//
func (c *HttpClient) SetTimeout(to time.Duration) error {
	c.client.Timeout = to * time.Second
	return nil
}

//
//
//
//
//
func (c *HttpClient) NotFollowRedirect() {
	c.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
}

//
//
//
//
//
func (c *HttpClient) SetCookie(mm map[string]map[string]string) {
	for site, siteCs := range mm {
		u, _ := url.Parse(site)
		cs := make([]*http.Cookie, 0)
		for key, val := range siteCs {
			c := &http.Cookie{Name: key, Value: val, Secure: true}
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

//
//
//
//
//
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

//
//
//
//
//
func (c *HttpClient) Delete(urlStr string, dataMap map[string]string) (*HttpResult, error) {
	reqParams := url.Values{}
	if dataMap != nil {
		for k, v := range dataMap {
			reqParams.Add(k, v)
		}
	}
	req, _ := http.NewRequest("DELETE", urlStr, strings.NewReader(reqParams.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return c.doRequest(req)
}

//
//
//
//
//
func (c *HttpClient) PostJson(urlStr string, data interface{}) (*HttpResult, error) {
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
func (c *HttpClient) PutJson(urlStr string, data interface{}) (*HttpResult, error) {
	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequest("PUT", urlStr, bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	return c.doRequest(req)
}

//
//
//
//
//
func (c *HttpClient) PostString(urlStr string, body string) (*HttpResult, error) {
	req, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(body))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	return c.doRequest(req)
}

//
//
//
//
//
func (c *HttpClient) PostXMLString(urlStr string, src string) (*HttpResult, error) {
	req, _ := http.NewRequest("POST", urlStr, bytes.NewBuffer([]byte(src)))
	req.Header.Add("Content-Type", "application/xml;charset=utf-8")
	return c.doRequest(req)
}

//
//
//
//
//
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

//
//
//
//
//
func (c *HttpClient) PostGzip(urlString string, data string) (*HttpResult, error) {
	var b bytes.Buffer
	w, _ := gzip.NewWriterLevel(&b, gzip.NoCompression)
	defer w.Close()
	_, err := w.Write([]byte(data))
	if err != nil {
		return nil, err
	}
	w.Flush()
	req, _ := http.NewRequest("POST", urlString, &b)
	req.Header.Add("Content-Encoding", "gzip")
	return c.doRequest(req)
}

//
//
//
//
//
func (c *HttpClient) doRequest(req *http.Request) (*HttpResult, error) {
	if c.option.DefaultHeader != nil {
		for k, v := range c.option.DefaultHeader {
			req.Header.Add(k, v)
		}
	}
	if c.option.UseAgent != "" {
		req.Header.Add("User-Agent", c.option.UseAgent)
	}

	//
	req.Close = true
	resp, err := c.client.Do(req)
	if err != nil {
		if _, ok := err.(*url.Error); ok {
			return nil, err
		}
	}

	//
	return NewHttpResult(resp, c.option.IsLazy), nil
}
