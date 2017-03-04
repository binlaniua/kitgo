package kitgo

import "orcaman/concurrent-map"

//-------------------------------------
//
//
//
//-------------------------------------
type concurrentMap struct {
	m cmap.ConcurrentMap
}

//
//
//
//
//
func NewMap() *concurrentMap {
	c := &concurrentMap{
		m: cmap.New(),
	}
	return c
}

//
//
//
//
//
func (c *concurrentMap) Put(key string, value interface{}) *concurrentMap {
	c.m.Set(key, value)
	return c
}

//
//
//
//
//
func (c *concurrentMap) GetString(k string) (string, bool) {
	r, ok := c.m.Get(k)
	if ok {
		return r.(string), true
	} else {
		return "", false
	}
}

//
//
//
//
//
func (c *concurrentMap) GetInt(k string) (int, bool) {
	r, ok := c.m.Get(k)
	if ok {
		return r.(int), true
	} else {
		return -1, false
	}
}

//
//
//
//
//
func (c *concurrentMap) GetBool(k string) (bool, bool) {
	r, ok := c.m.Get(k)
	if ok {
		return r.(bool), true
	} else {
		return false, false
	}
}

//
//
//
//
//
func (c *concurrentMap) String() string {
	r, _ := c.m.MarshalJSON()
	return string(r)
}
