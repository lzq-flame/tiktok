package pack

import (
	"errors"
	"example/cmd/video/kitex_gen/video"
	"example/pkg/errno"
	"time"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/22 17:22
 **/

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *video.VideoBaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *video.VideoBaseResp {
	return &video.VideoBaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}
