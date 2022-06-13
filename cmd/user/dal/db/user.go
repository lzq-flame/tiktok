package db

import (
	"context"
	"example/cmd/user/dal/redispool"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
	"strconv"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/20 16:02
 **/

type User struct {
	gorm.Model
	ID            int64  `gorm:"type:int;primarykey;autoIncrement"`
	Username      string `gorm:"type:varchar(20);index"`
	Password      string `gorm:"type:varchar(20)"`
	Token         string `gorm:"type:text"`
	FollowCount   int64  `gorm:"type:int;default:0"`
	FollowerCount int64  `gorm:"type:int;default:0"`
}

// CreateUser 创建用户
func CreateUser(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).Create(user).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}
	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetUserName(ctx context.Context, userId int64) string {
	name := ""
	user := User{}
	id := strconv.Itoa(int(userId))
	err := DB.WithContext(ctx).Where("id = ?", id).Find(&user).Error
	if err != nil {
		return name
	}
	return user.Username
}

func SetFollowerCountById(ctx context.Context, id int64, followerCount int64) error {
	err := DB.WithContext(ctx).Model(&User{}).Where("id = ?", id).Update("follower_count", followerCount).Error
	return err
}

func SetFollowCountById(ctx context.Context, id int64, followCount int64) error {
	err := DB.WithContext(ctx).Model(&User{}).Where("id = ?", id).Update("follow_count", followCount).Error
	return err
}

func GetUser(ctx context.Context, userId int64) (*User, error) {
	user := User{}
	if err := DB.WithContext(ctx).First(&user, userId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func IsFollowerExist(userid int64, followerId int64) bool {
	redisFollowerSetName := "user_follower_set_" + strconv.FormatInt(userid, 10)
	val, err := redispool.RedisPool.Get().Do("SISMEMBER", redisFollowerSetName, followerId)
	if err != nil {
		return false
	}

	isExist, _ := strconv.ParseBool(string(val.(int64)))

	return isExist
}

// GetFollowerIds 获取粉丝id
func GetFollowerIds(id int64) ([]interface{}, error) {
	redisFollowerSetName := "user_follower_set_" + strconv.FormatInt(id, 10)
	userIds, err := redis.Values(redispool.RedisPool.Get().Do("SMEMBERS", redisFollowerSetName))
	return userIds, err
}

// GetFollowIds 获取关注id
func GetFollowIds(id int64) ([]interface{}, error) {
	redisFollowSetName := "user_follow_set_" + strconv.FormatInt(id, 10)
	userIds, err := redis.Values(redispool.RedisPool.Get().Do("SMEMBERS", redisFollowSetName))

	return userIds, err
}

// AddUserId 添加关注列表
func AddUserId(setName string, id int64) error {
	_, err := redispool.RedisPool.Get().Do("SADD", setName, id)
	return err
}

// RemoveUserId 移除userid
func RemoveUserId(setName string, id int64) error {
	_, err := redispool.RedisPool.Get().Do("SREM", setName, id)

	return err
}
