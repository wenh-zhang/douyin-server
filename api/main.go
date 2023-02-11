package main

import (
	"douyin/api/handler"
	"douyin/api/rpc"
	"douyin/pkg/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/hertz-contrib/gzip"
)

func Init() {
	rpc.InitRPC()
}

func main() {
	Init()
	r := server.Default(
		server.WithMaxRequestBodySize(1000<<20), // 1000MB
		server.WithTransport(standard.NewTransporter),
		server.WithStreamBody(true),
	)
	r.Static("/video", "../public")
	r.Static("/cover", "../public")

	douyin := r.Group("/douyin")
	douyin.Use(gzip.Gzip(gzip.DefaultCompression))
	douyin.GET("/feed/", handler.Feed)
	user := douyin.Group("/user")
	user.GET("/", append(middleware.AuthenticationMiddleware(), handler.UserInfo)...)
	user.POST("/register/", handler.UserRegister)
	user.POST("/login/", handler.UserLogin)
	publish := douyin.Group("/publish")
	publish.POST("/action/", append(middleware.AuthenticationMiddleware(), handler.PublishAction)...)
	publish.GET("/list/", handler.PublishList)
	favorite := douyin.Group("/favorite")
	favorite.POST("/action/", append(middleware.AuthenticationMiddleware(), handler.FavoriteAction)...)
	favorite.GET("/list/", handler.FavoriteList)
	comment := douyin.Group("/comment")
	comment.POST("/action/", append(middleware.AuthenticationMiddleware(), handler.CommentAction)...)
	comment.GET("/list/", handler.CommentList)
	r.Spin()
}
