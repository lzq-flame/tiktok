package db

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/30 10:39
 **/

type Comment struct {
	gorm.Model
	ID      int64  `gorm:"type:int;primaryket;autoIncrement"`
	VideoId int64  `gorm:"type:int"`
	UserId  int64  `gorm:"type:int"`
	Content string `gorm:"type:text"`
}

// CreateComment 创建一条评论
func CreateComment(ctx context.Context, comment *Comment) error {
	return DB.WithContext(ctx).Create(comment).Error
}

// DeleteComment 删除一条评论
func DeleteComment(ctx context.Context, comment *Comment) error {
	return DB.WithContext(ctx).Delete(comment).Error
}

// GetCommentByVideoId 根据视频id获取评论列表
func GetCommentByVideoId(ctx context.Context, videoId int64) ([]*Comment, error) {
	list := make([]*Comment, 0)
	fmt.Println("video_id = ", videoId)
	err := DB.WithContext(ctx).Where(&Comment{VideoId: videoId}).Find(&list).Order("create_at desc").Error
	if err != nil {
		return nil, err
	}
	fmt.Println("list: ", list)
	return list, nil
}
