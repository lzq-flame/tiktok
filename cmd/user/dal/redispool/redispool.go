package redispool

import (
	"example/pkg/constants"
	"github.com/gomodule/redigo/redis"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/22 17:30
 **/

// RedisPool 创建全局redis 连接池 句柄
var RedisPool redis.Pool

//InitRedis 初始化Redis连接池
func InitRedis() {
	RedisPool = redis.Pool{
		MaxIdle:         20,
		MaxActive:       50,
		MaxConnLifetime: 60 * 5,
		IdleTimeout:     60,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", constants.RedisAddress)
		},
	}
}
