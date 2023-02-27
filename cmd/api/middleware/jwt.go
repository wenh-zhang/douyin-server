package middleware

import (
	"context"
	"douyin/cmd/api/pkg"
	"douyin/shared/constant"
	"douyin/shared/errno"
	"douyin/shared/util"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func AuthenticationMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Query(constant.Token)
		if len(token) == 0 {
			token = c.PostForm(constant.Token)
		}
		if len(token) == 0 {
			pkg.SendResponse(c, utils.H{
				"status_code": errno.ParamErr.ErrCode,
				"status_msg":  "Token not found",
			})
			c.Abort()
			return
		}
		claims, err := util.NewJWT(constant.TokenSignedKey).ParseToken(token)
		if err != nil {
			pkg.SendResponse(c, utils.H{
				"status_code": errno.ParamErr.ErrCode,
				"status_msg":  "Token parse error",
			})
			return
		}
		userID := claims.UserID
		c.Set(constant.TokenUserIdentifyKey, userID)
		c.Next(ctx)
	}
}
