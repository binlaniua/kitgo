package http

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"net/http"
	"io/ioutil"
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"github.com/binlaniua/kitgo/file"
	"io"
	"compress/gzip"
	"log"
	"github.com/binlaniua/kitgo"
)

type HttpResult struct {
	Status   int
	Body     []byte
	Response *http.Response
}

//-------------------------------------
//
//
//
//-------------------------------------
func NewHttpResult(res *http.Response, isLazy bool) *HttpResult {
	r := &HttpResult{Status:res.StatusCode, Response:res}
	if !isLazy {
		r.readBody()
	}
	return r
}

func (hr *HttpResult) readBody() {
	var reader io.ReadCloser
	switch hr.Response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(hr.Response.Body)
	default:
		reader = hr.Response.Body
	}
	defer reader.Close()
	byteData, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Println(hr.GetUrl(), " => 读取内容出错, ", err)
	}
	hr.Body = byteData
}

//-------------------------------------
//
//
//
//-------------------------------------
func (hr *HttpResult) ToJson(data interface{}) error {
	err := json.Unmarshal(hr.Body, data)
	return err
}

//-------------------------------------
//
//
//
//-------------------------------------
func (hr *HttpResult) ToJsonData() (*simplejson.Json, error) {
	r, err := simplejson.NewJson(hr.Body)
	if err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func (hr *HttpResult) ToString() string {
	if (hr.Body == nil) {
		hr.readBody()
	}
	return string(hr.Body)
}

//-------------------------------------
//
//
//
//-------------------------------------
func (hr *HttpResult) IsSuccess() bool {
	return hr.Status == http.StatusOK
}

//-------------------------------------
//
//
//
//-------------------------------------
func (hr *HttpResult) ToQuery() (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(hr.Body))
	return doc, err
}

//-------------------------------------
//
//
//
//-------------------------------------
func (hr *HttpResult) ToQuerySelect(exp string) (*goquery.Selection, error) {
	doc, err := hr.ToQuery()
	if err != nil {
		return nil, err
	}
	return doc.Find(exp), nil
}

//-------------------------------------
//
//
//
//-------------------------------------
func (hr *HttpResult) ToFile(filePath string) bool {
	return file.WriteBytes(filePath, hr.Body)
}

//-------------------------------------
//
//
//
//-------------------------------------
func (hr *HttpResult) IsEmpty() bool {
	return len(hr.Body) == 0
}

//-------------------------------------
//
//
//
//-------------------------------------
func (hr *HttpResult) GetUrl() string {
	return hr.Response.Request.URL.String()
}

//-------------------------------------
//
//
//
//-------------------------------------
func (hr *HttpResult) Close() {
	defer kitgo.ExceptionCatch()
	hr.Response.Body.Close()
}