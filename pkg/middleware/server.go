package middleware

import (
	"context"
	"douyin/api/handler"
	"douyin/pkg/constant"
	"douyin/pkg/errno"
	"douyin/pkg/util"
	"github.com/cloudwego/hertz/pkg/app"
)

func AuthenticationMiddleware() []app.HandlerFunc {
	return []app.HandlerFunc{func(ctx context.Context, c *app.RequestContext) {
		token := c.Query(constant.Token)
		if len(token) == 0 {
			token = c.PostForm(constant.Token)
		}
		if len(token) == 0 {
			handler.SendResponse(c, handler.NewResponse(errno.LackTokenErr))
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			handler.SendResponse(c, handler.NewResponse(err))
			return
		}
		userID := claims.UserID
		c.Set(constant.TokenUserIdentifyKey, userID)
		c.Next(ctx)
	}}
}
