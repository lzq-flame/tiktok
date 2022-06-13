package dal

import (
	"example/cmd/video/dal/db"
	"example/cmd/video/dal/redispool"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/22 16:15
 **/

// Init init dal
func Init() {
	db.Init()
	redispool.InitRedis()
}
