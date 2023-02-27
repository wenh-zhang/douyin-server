package util

import (
	"douyin/shared/errno"
	"douyin/shared/kitex_gen/base"
)

func BuildBaseResp(err error) *base.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	return baseResp(errno.ConvertErr(err))
}

func baseResp(err errno.ErrNo) *base.BaseResp {
	return &base.BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
}
