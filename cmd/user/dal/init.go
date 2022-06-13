package dal

import (
	"example/cmd/user/dal/db"
	"example/cmd/user/dal/redispool"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/20 16:02
 **/

// Init init dal
func Init() {
	db.Init()             // mysql init
	redispool.InitRedis() // redis init
}
