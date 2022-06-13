package service

import (
	"context"
	"example/cmd/user/kitex_gen/userdemo"
	"example/cmd/video/dal/db"
	"example/cmd/video/kitex_gen/video"
	"example/cmd/video/pack"
	"example/cmd/video/rpc"
	"time"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/22 20:39
 **/

type GetVideoFeedService struct {
	ctx context.Context
}

func NewGetVideoFeedService(ctx context.Context) *GetVideoFeedService {
	return &GetVideoFeedService{ctx: ctx}
}

func (s *GetVideoFeedService) GetVideoFeed(req *video.GetVideoFeedRequest) ([]*video.Video, error) {
	timeStamp := req.TimeStamp
	t := time.Unix(timeStamp, 0)
	timeFormat := t.Format("2006-01-02 15:04:05")
	videos, err := db.GetVideoList(s.ctx, timeFormat)
	if err != nil {
		return nil, err
	}
	vs := pack.Videos(0, videos)
	ids := make([]int64, 0)

	for _, v := range vs {
		ids = append(ids, v.User.Id)
	}
	userMap, err := rpc.GetUsers(s.ctx, &userdemo.GetUsersRequest{
		UserIds: ids,
	})
	if err != nil {
		return nil, err
	}

	for _, v := range vs {
		v.User.UserName = userMap[v.User.Id].UserName
		//fmt.Println("userMap[v.User.Id].UserName:", userMap[v.User.Id].UserName)
		v.User.IsFollow = userMap[v.User.Id].IsFollow
		v.User.FollowerCount = userMap[v.User.Id].FollowerCount
		v.User.FollowCount = userMap[v.User.Id].FollowCount
	}

	return vs, err
}
