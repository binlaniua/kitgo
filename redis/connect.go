package rc

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
	"time"
)

//-------------------------------------
//
//
//
//-------------------------------------
func ConnectPoolConfig(info *RedisConfig) *redis.Pool {
	dest := fmt.Sprintf("%s:%s", info.Ip, info.Port)
	r := &redis.Pool{
		MaxIdle:     info.MaxIdle,
		MaxActive:   info.MaxActive,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", dest, redis.DialPassword(info.Password))
			if err != nil {
				return nil, err
			} else {
				return c, nil
			}
		},
	}
	return r
}

//-------------------------------------
//
//
//
//-------------------------------------
func ConnectConfig(info *RedisConfig) (redis.Conn, error) {
	dest := fmt.Sprintf("%s:%s", info.Ip, info.Port)
	conn, err := net.DialTimeout("tcp", dest, time.Second)
	if err != nil {
		return nil, err
	}
	r := redis.NewConn(conn, time.Second, time.Second)
	return r, nil
}
