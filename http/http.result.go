package http

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"net/http"
	"io/ioutil"
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"github.com/binlaniua/kitgo/file"
	"github.com/binlaniua/kitgo"
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
		kitgo.Log(hr.Url, " 转换JSON失败 => ", err);
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
func (hr *HttpResult) ToJsonData() (*simplejson.Json, bool) {
	r, err := simplejson.NewJson(hr.Body)
	if err != nil {
		kitgo.Log(hr.Url, " 转换JSON失败 => ", err);
		return nil, false
	} else {
		return r, true
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
func (hr *HttpResult) ToQuery() (*goquery.Document, bool) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(hr.Body))
	if err != nil {
		kitgo.Log(hr.Url, " 转换Document失败  => ", err);
		return nil, false
	} else {
		return doc, true
	}
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