package pack

import (
	"errors"
	"example/cmd/comment/kitex_gen/comment"
	"example/pkg/errno"
	"time"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/30 14:06
 **/

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *comment.CommentBaseResp {
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

func baseResp(err errno.ErrNo) *comment.CommentBaseResp {
	return &comment.CommentBaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}
