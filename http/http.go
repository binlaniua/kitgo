package http

import (
	"bytes"
	"encoding/json"
	"github.com/binlaniua/kitgo"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

//
//
//
//
//
func HttpGet(urlStr string) *HttpResult {
	resp, err := http.Get(urlStr)
	if err != nil {
		kitgo.Log(urlStr, " => ", err)
		return nil
	}
	return NewHttpResult(resp, false)
}

//
//
//
//
//
func HttpGetReply(urlStr string, times int) *HttpResult {
	resp, err := http.Get(urlStr)
	if err != nil {
		for i := 0; i < times; i++ {
			r := HttpGet(urlStr)
			if r != nil {
				return r
			}
		}
		return nil
	}
	return NewHttpResult(resp, false)
}

//
//
//
//
//
func HttpPostFrom(urlStr string, dataMap map[string]string) *HttpResult {
	reqParams := url.Values{}
	if dataMap != nil {
		for k, v := range dataMap {
			reqParams.Add(k, v)
		}
	}
	resp, err := http.PostForm(urlStr, reqParams)
	if err != nil {
		kitgo.Log(urlStr, " => ", err)
		return nil
	}
	return NewHttpResult(resp, false)
}

//-------------------------------------
//
//
//
//-------------------------------------
func HttpPost(urlStr string, body string) *HttpResult {
	resp, err := http.Post(urlStr, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(body)))
	if err != nil {
		kitgo.Log(urlStr, " => ", err)
		return nil
	}
	return NewHttpResult(resp, false)
}

//
//
//
//
//
func HttpPostJson(urlStr string, data interface{}) *HttpResult {
	jsonData, _ := json.Marshal(data)
	resp, err := http.Post(urlStr, "application/json;charset=utf-8", bytes.NewBuffer(jsonData))
	if err != nil {
		kitgo.Log(urlStr, " => ", err)
		return nil
	}
	return NewHttpResult(resp, false)
}

//
//
//
//
//
func HttpPostFile(urlStr string, dataMap map[string]interface{}) *HttpResult {
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
	resp, err := http.Post(urlStr, write.FormDataContentType(), buff)
	if err != nil {
		kitgo.Log(urlStr, " => ", err)
		return nil
	}
	return NewHttpResult(resp, false)
}
