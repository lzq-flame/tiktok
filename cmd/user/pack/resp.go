package pack

import (
	"errors"
	"example/cmd/user/kitex_gen/userdemo"
	"example/pkg/errno"
	"time"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/20 16:41
 **/

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *userdemo.BaseResp {
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

func baseResp(err errno.ErrNo) *userdemo.BaseResp {
	return &userdemo.BaseResp{StatusCode: err.ErrCode, StatusMessage: err.ErrMsg, ServiceTime: time.Now().Unix()}
}
