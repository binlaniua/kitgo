package kitgo

import "github.com/garyburd/redigo/redis"


//-------------------------------------
//
//
//
//-------------------------------------
func RedisLrange(redisConn redis.Conn, key string, start int, end int) ([]string, error)  {
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
