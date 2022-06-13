package pack

import (
	"context"
	"example/cmd/comment/kitex_gen/comment"
	"example/cmd/video/dal/db"
	"example/cmd/video/dal/redispool"
	"example/cmd/video/kitex_gen/video"
	"example/cmd/video/rpc"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/22 20:48
 **/

func Video(myId int64, v *db.Video) *video.Video {
	if v == nil {
		return nil
	}
	if myId == 0 {
		return &video.Video{
			Id:            v.ID,
			User:          &video.VideoUser{Id: v.UserId},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			Title:         v.Title,
			IsFavorite:    false,
			FavoriteCount: redispool.GetFavoriteCount(v.ID),
			CommentCount: rpc.GetCommentCount(context.Background(), &comment.GetCommentCountRequest{
				VideoId: v.ID,
			}),
		}
	}
	return &video.Video{
		Id:            v.ID,
		User:          &video.VideoUser{Id: v.UserId},
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		Title:         v.Title,
		IsFavorite:    redispool.IsFavoriteExist(myId, v.ID),
		FavoriteCount: redispool.GetFavoriteCount(v.ID),
		CommentCount: rpc.GetCommentCount(context.Background(), &comment.GetCommentCountRequest{
			VideoId: v.ID,
		}),
	}
}

func Videos(myId int64, vs []*db.Video) []*video.Video {
	videos := make([]*video.Video, 0)
	for _, v := range vs {
		if v2 := Video(myId, v); v2 != nil {
			videos = append(videos, v2)
		}
	}
	return videos
}
