package kitgo

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"net/http"
	"io/ioutil"
	"github.com/PuerkitoBio/goquery"
	"log"
	"bytes"
)

type HttpResult struct {
	Status int
	Body   []byte
	Origin *http.Response
}

//-------------------------------------
//
//
//
//-------------------------------------
func NewHttpResult(res *http.Response) *HttpResult {
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	r := &HttpResult{res.StatusCode, bytes, res}
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
func (hr *HttpResult) ToJsonData() *simplejson.Json {
	r, err := simplejson.NewJson(hr.Body)
	if err != nil {
		return nil
	} else {
		return r
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
func (hr *HttpResult) ToQuery() *goquery.Document {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(hr.Body))
	if err != nil {
		log.Println(err);
		return nil
	} else {
		return doc
	}
}

//-------------------------------------
//
//
//
//-------------------------------------
func (hr *HttpResult) ToFile(filePath string) bool {
	return FileWriteBytes(filePath, hr.Body)
}

//-------------------------------------
//
//
//
//-------------------------------------
func (hr *HttpResult) IsEmpty() bool {
	return len(hr.Body) == 0
}