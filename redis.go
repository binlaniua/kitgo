package kitgo

import (
	"github.com/garyburd/redigo/redis"
)

//-------------------------------------
//
//
//
//-------------------------------------
func RedisZrange(redisConn redis.Conn, key string, start int, end int) ([]string, error)  {
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
func RedisLrange(redisConn redis.Conn, key string, start int, end int) ([]string, error)  {
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
func RedisLLen(rc redis.Conn, key string) (int64, error)  {
	r, err := rc.Do("LLEN", key)
	if err != nil {
		return -1, err
	}
	return r.(int64), nil
}