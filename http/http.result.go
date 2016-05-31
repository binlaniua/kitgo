package http

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"net/http"
	"io/ioutil"
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"github.com/binlaniua/kitgo/file"
)

type HttpResult struct {
	Status   int
	Body     []byte
	Url      string
	Response *http.Response
}

//-------------------------------------
//
//
//
//-------------------------------------
func NewHttpResult(res *http.Response, urlStr string) *HttpResult {
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	r := &HttpResult{res.StatusCode, bytes, urlStr, res}
	return r
}

//-------------------------------------
//
//
//
//-------------------------------------
func (hr *HttpResult) ToJson(data interface{}) bool {
	err := json.Unmarshal(hr.Body, data)
	if err != nil {
		//kitgo.Log(hr.Url, " 转换JSON失败 => ", err);
		return false
	} else {
		return true
	}
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