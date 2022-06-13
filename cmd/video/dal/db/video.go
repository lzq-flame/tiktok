package db

import (
	"context"
	"gorm.io/gorm"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/22 16:15
 **/

type Video struct {
	gorm.Model
	ID            int64  `gorm:"type:int;primarykey;autoIncrement"`
	PlayUrl       string `gorm:"type:text"`
	CoverUrl      string `gorm:"type:text"`
	FavoriteCount int64  `gorm:"type:int;default:0"`
	CommentCount  int64  `gorm:"type:int;default:0"`
	IsFavorite    bool   `gorm:"type:tinyint;default:0"`
	UserId        int64  `gorm:"type:int"`
	Title         string `gorm:"type:text"`
}

// CreateVideo 创建video
func CreateVideo(ctx context.Context, video *Video) error {
	return DB.WithContext(ctx).Create(video).Error
}

func GetVideoList(ctx context.Context, timeStamp string) ([]*Video, error) {
	videos := make([]*Video, 0)
	err := DB.WithContext(ctx).Where("created_at < ?", timeStamp).Limit(30).Find(&videos).Error
	return videos, err
}

func GetVideoListByVideoIds(ctx context.Context, ids []int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if len(ids) == 0 {
		return res, nil
	}
	if err := DB.WithContext(ctx).Where("id in ?", ids).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
