package handler

import (
	"douyin/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Response The return json format of basic request
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// NewResponse Build Response base on error type
func NewResponse(err error) Response {
	errNo := errno.ConvertErr(err)
	return Response{
		StatusCode: errNo.ErrCode,
		StatusMsg:  errNo.ErrMsg,
	}
}

// SendResponse Send response in json format
func SendResponse(c *app.RequestContext, obj interface{}) {
	c.JSON(consts.StatusOK, obj)
}
