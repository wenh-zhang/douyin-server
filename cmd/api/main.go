package main

import (
	"douyin/cmd/api/handler"
	"douyin/cmd/api/rpc"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Init() {
	rpc.InitRPC()
}

func main() {
	Init()
	r := server.Default()

	//r.Use(recovery.Recovery(recovery.WithRecoveryHandler(
	//	func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
	//		hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
	//		c.JSON(consts.StatusInternalServerError, map[string]interface{}{
	//			"code":    errno.ServiceErrCode,
	//			"message": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
	//		})
	//	})))

	douyin := r.Group("/douyin")
	douyin.GET("/feed/", handler.Feed)

	//r.NoRoute(func(ctx context.Context, c *app.RequestContext) {
	//	c.String(consts.StatusOK, "no route")
	//})
	//r.NoMethod(func(ctx context.Context, c *app.RequestContext) {
	//	c.String(consts.StatusOK, "no method")
	//})
	r.Spin()
}

//func main() {
//	Init()
//	r := gin.Default()
//
//	apiRouter := r.Group("/douyin")
//
//	// basic apis
//	apiRouter.GET("/feed/", handler.Feed)
//
//	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
//}
