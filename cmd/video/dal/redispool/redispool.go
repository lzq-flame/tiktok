package redispool

import (
	"example/pkg/constants"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/22 17:39
 **/

// RedisPool 创建全局redis 连接池 句柄
var RedisPool redis.Pool

//InitRedis 初始化Redis连接池
func InitRedis() {
	RedisPool = redis.Pool{
		MaxIdle:         50,
		MaxActive:       100,
		MaxConnLifetime: 60 * 5,
		IdleTimeout:     60,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", constants.RedisAddress)
		},
	}
}

// IncrFavoriteCount 增加视频点赞数
func IncrFavoriteCount(videoId int64) error {
	videoName := "video_" + strconv.FormatInt(videoId, 10) + "_favorite_count"
	_, err := RedisPool.Get().Do("INCR", videoName)
	return err
}

// DecrFavoriteCount 减少视频点赞数
func DecrFavoriteCount(videoId int64) error {
	videoName := "video_" + strconv.FormatInt(videoId, 10) + "_favorite_count"
	_, err := RedisPool.Get().Do("DECR", videoName)
	return err
}

// GetVideoIdsByUserId 获取用户发布过的视频列表
func GetVideoIdsByUserId(userId int64) ([]int64, error) {
	redisListName := "user_video_list_" + strconv.FormatInt(userId, 10)
	videoList, err := redis.Values(RedisPool.Get().Do("lrange", redisListName, 0, -1))
	if err != nil {
		return nil, err
	}
	ids := make([]int64, len(videoList))
	for k, v := range videoList {
		ids[k], _ = strconv.ParseInt(string(v.([]byte)), 10, 64)
	}
	return ids, nil
}

// IsFavoriteExist 判断是否点赞
func IsFavoriteExist(userId int64, videoId int64) bool {
	redisActionName := "user_favorite_set_" + strconv.FormatInt(userId, 10)
	val, err := RedisPool.Get().Do("SISMEMBER", redisActionName, videoId)
	if err != nil {
		//logger.L().Warn("IsFavoriteExist redis查询失败! err: ", err)
		return false
	}

	isExist, _ := strconv.ParseBool(strconv.FormatInt(val.(int64), 10))

	return isExist
}

// AddVideoId 添加点赞
func AddVideoId(userId int64, videoId int64) error {
	redisActionName := "user_favorite_set_" + strconv.FormatInt(userId, 10)
	_ = IncrFavoriteCount(videoId)
	_, err := RedisPool.Get().Do("SADD", redisActionName, videoId)

	return err
}

// RemoveVideoId 取消点赞
func RemoveVideoId(userId int64, videoId int64) error {
	redisActionName := "user_favorite_set_" + strconv.FormatInt(userId, 10)
	_ = DecrFavoriteCount(videoId)
	_, err := RedisPool.Get().Do("SREM", redisActionName, videoId)
	return err
}

// GetFavoriteCount 获取视频点赞数量
func GetFavoriteCount(videoId int64) int64 {
	videoName := "video_" + strconv.FormatInt(videoId, 10) + "_favorite_count"
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

// GetFavoriteVideo 获取用户点赞列表
func GetFavoriteVideo(userId int64) ([]int64, error) {
	redisListName := "user_favorite_set_" + strconv.FormatInt(userId, 10)
	videoIds, err := redis.Values(RedisPool.Get().Do("SMEMBERS", redisListName))
	if err != nil {
		return nil, err
	}
	ids := make([]int64, len(videoIds))
	for k, v := range videoIds {
		ids[k], _ = strconv.ParseInt(string(v.([]byte)), 10, 64)
	}
	return ids, nil
}
