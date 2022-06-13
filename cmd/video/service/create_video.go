package service

import (
	"context"
	"example/cmd/video/dal/db"
	"example/cmd/video/dal/redispool"
	"example/cmd/video/kitex_gen/video"
	"example/pkg/constants"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/22 16:12
 **/

type CreateVideoService struct {
	ctx context.Context
}

func NewCreateVideoService(ctx context.Context) *CreateVideoService {
	return &CreateVideoService{ctx: ctx}
}

func (s *CreateVideoService) CreateVideo(req *video.CreateVideoRequest) (int64, error) {
	var singleVideo db.Video
	//singleVideo = new(db.Video)
	err := db.DB.WithContext(s.ctx).Transaction(func(tx *gorm.DB) error {
		singleVideo.UserId = req.UserId
		err := tx.Create(&singleVideo).Error
		if err != nil {
			fmt.Println("创建视频失败")
			return err
		}
		//视频id
		videoId := strconv.FormatInt(singleVideo.ID, 10)
		fileName := videoId + ".mp4"
		singleVideo.PlayUrl = "http://" + constants.LocalIP + ":80/video/" + fileName
		singleVideo.Title = req.Title
		fmt.Println("video url : ", singleVideo.PlayUrl)
		//保存video
		err = tx.Save(&singleVideo).Error
		if err != nil {
			fmt.Println("保存video失败")
			return err
		}

		////保存视频到./data/video
		//err = SaveUploadedFile(req.File, fileName)
		//if err != nil {
		//	return err
		//}

		// 将视频id保存到用户list （redis）
		redisListName := "user_video_list_" + strconv.FormatInt(req.UserId, 10)
		_, err = redispool.RedisPool.Get().Do("rpush", redisListName, videoId)
		if err != nil {
			fmt.Println("保存redis失败")
			return err
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return singleVideo.ID, err
}

//func SaveUploadedFile(b []byte, filename string) error {
//	if b == nil {
//		errno.NewErrNo(-1, "数据传输为空")
//	}
//	r := io.NewSectionReader(bytes.NewReader(b), 0, int64(len(b)))
//	out, err := os.Create(filename)
//	if err != nil {
//		return err
//	}
//	defer out.Close()
//	_, err = io.Copy(out, r)
//	return err
//}
