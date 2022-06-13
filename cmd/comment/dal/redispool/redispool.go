package redispool

import (
	"example/pkg/constants"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/30 10:41
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

// IncrCommentCount 添加评论数量
func IncrCommentCount(videoId int64) error {
	videoName := "video_" + strconv.FormatInt(videoId, 10) + "_comment_count"
	_, err := RedisPool.Get().Do("INCR", videoName)
	return err
}

// DecrCommentCount 减少评论数量
func DecrCommentCount(videoId int64) error {
	videoName := "video_" + strconv.FormatInt(videoId, 10) + "_comment_count"
	_, err := RedisPool.Get().Do("DECR", videoName)
	return err
}

// GetCommentCount 获取视频评论数量
func GetCommentCount(videoId int64) int64 {
	videoName := "video_" + strconv.FormatInt(videoId, 10) + "_comment_count"
	v, err := RedisPool.Get().Do("GET", videoName)
	if err != nil || v == nil {
		return 0
	}
	count, err := strconv.ParseInt(string(v.([]byte)), 10, 64)
	if err != nil {
		return 0
	}
	return count
}
