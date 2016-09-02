package lib

import "github.com/garyburd/redigo/redis"

//-------------------------------------
//
// redis 订阅消息
//
//-------------------------------------
func RedisSub(redisConn redis.Conn, channel string, handle func(err error, channel string, msg string)) *redis.PubSubConn  {
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
	ss := [len(is)]string{}
	for index, item := range r.([]interface{}) {
		ss[index] = string(item.([]byte))
	}
	return ss, nil
}
