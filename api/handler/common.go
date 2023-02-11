package handler

import (
	"douyin/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func NewResponse(err error) Response{
	errNo := errno.ConvertErr(err)
	return Response{
		StatusCode: errNo.ErrCode,
		StatusMsg: errNo.ErrMsg,
	}
}

func SendResponse(c *app.RequestContext, obj interface{}){
	c.JSON(consts.StatusOK, obj)
}