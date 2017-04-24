package rc

import (
	"github.com/garyburd/redigo/redis"
)

//-------------------------------------
//
//
//
//-------------------------------------
func RedisZrange(redisConn redis.Conn, key string, start int, end int) ([]string, error) {
	r, err := redisConn.Do("ZRANGE", key, start, end)
	if err != nil {
		return nil, err
	}
	rs := r.([]interface{})
	result := []string{}
	for _, item := range rs {
		result = append(result, string(item.([]byte)))
	}
	return result, nil
}

//-------------------------------------
//
//
//
//-------------------------------------
func RedisLrange(redisConn redis.Conn, key string, start int, end int) ([]string, error) {
	r, err := redisConn.Do("LRANGE", key, start, end)
	if err != nil {
		return nil, err
	}
	rs := r.([]interface{})
	result := []string{}
	for _, item := range rs {
		result = append(result, string(item.([]byte)))
	}
	return result, nil
}

//
//
//
//
//
func RedisLLen(rc redis.Conn, key string) (int64, error) {
	r, err := rc.Do("LLEN", key)
	if err != nil {
		return -1, err
	}
	return r.(int64), nil
}

//-------------------------------------
//
// redis 订阅消息
//
//-------------------------------------
func RedisSub(redisConn redis.Conn, channel string, handle func(err error, channel string, msg string)) *redis.PubSubConn {
	src := &redis.PubSubConn{redisConn}
	src.Subscribe(channel)
	defer src.Close()
	for {
		switch v := src.Receive().(type) {
		case redis.Message:
			handle(nil, channel, string(v.Data))
		case error:
			handle(v, "", "")
		}
	}
	return src
}

//-------------------------------------
//
// Lrange
//
//-------------------------------------
func RedisLRangeStrings(redisConn redis.Conn, key string, start, end int) ([]string, error) {
	r, err := redisConn.Do("LRANGE", key, start, end)
	if err != nil {
		return nil, err
	}
	is := r.([]interface{})
	ss := make([]string, len(is))
	for index, item := range r.([]interface{}) {
		ss[index] = string(item.([]byte))
	}
	return ss, nil
}
