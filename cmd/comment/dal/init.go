package dal

import (
	"example/cmd/comment/dal/db"
	"example/cmd/comment/dal/redispool"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/30 10:39
 **/

// Init init dal
func Init() {
	db.Init()             // mysql init
	redispool.InitRedis() // redis init
}
