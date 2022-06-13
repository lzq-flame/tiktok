package pack

import (
	"example/cmd/comment/dal/db"
	"example/cmd/comment/kitex_gen/comment"
	"example/cmd/user/kitex_gen/userdemo"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/30 11:24
 **/

func Comment(u *userdemo.User, cmt *db.Comment) *comment.Comment {
	if cmt == nil {
		return nil
	}
	return &comment.Comment{
		CommentId:  0,
		User:       user(u),
		Content:    cmt.Content,
		CreateDate: cmt.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func Comments(us []*userdemo.User, comments []*db.Comment) []*comment.Comment {
	allComment := make([]*comment.Comment, 0)
	for k, c := range comments {
		if c1 := Comment(us[k], c); c1 != nil {
			allComment = append(allComment, c1)
		}
	}
	return allComment
}

func user(u *userdemo.User) *comment.CommentUser {
	if u == nil {
		return nil
	}
	return &comment.CommentUser{
		Id:            u.Id,
		UserName:      u.UserName,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
}
