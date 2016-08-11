package json

import "github.com/bitly/go-simplejson"

//-------------------------------------
//
//
//
//-------------------------------------
type JSON struct {
	data    []byte
	jsonObj *simplejson.Json
}

func NewJson(data []byte) (*JSON, error) {
	jo, err := simplejson.NewJson(data)
	if err != nil {
		return nil, err
	}
	j := &JSON{data, jo}
	return j
}

//-------------------------------------
//
//
//
//-------------------------------------
func (j *JSON) GetPathString(p string) string {
	s, _ := j.jsonObj.GetPath(p).String()
	return s
}

//-------------------------------------
//
//
//
//-------------------------------------
func (j *JSON) GetPathInt(p string) int64 {
	s, _ := j.jsonObj.GetPath(p).Int64()
	return s
}
