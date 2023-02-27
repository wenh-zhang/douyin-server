package pkg

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// SendResponse 服务端响应
func SendResponse(c *app.RequestContext, obj interface{}) {
	c.JSON(consts.StatusOK, obj)
}

