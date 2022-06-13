package pack

import (
	"example/cmd/user/dal/db"
	"example/cmd/user/kitex_gen/userdemo"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/20 16:41
 **/

// User pack user info
func User(myId int64, u *db.User) *userdemo.User {
	if u == nil {
		return nil
	}
	if myId == 0 {
		return &userdemo.User{
			Id:            u.ID,
			UserName:      u.Username,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      false,
		}
	}
	return &userdemo.User{
		Id:            u.ID,
		UserName:      u.Username,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      db.IsFollowerExist(u.ID, myId),
	}
}

// Users pack list of user info
func Users(myId int64, us []*db.User) []*userdemo.User {
	users := make([]*userdemo.User, 0)
	for _, u := range us {
		if user2 := User(myId, u); user2 != nil {
			users = append(users, user2)
		}
	}
	return users
}
