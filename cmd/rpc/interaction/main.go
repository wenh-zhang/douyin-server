package main

import (
	"douyin/cmd/rpc/interaction/dao"
	"douyin/cmd/rpc/interaction/global"
	"douyin/cmd/rpc/interaction/initialize"
	"douyin/cmd/rpc/interaction/mq"
	"douyin/cmd/rpc/interaction/pkg"
	"douyin/cmd/rpc/interaction/redis"
	interaction "douyin/shared/kitex_gen/interaction/interactionserver"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"log"
	"net"
)

func main() {
	initialize.Init()
	favoriteDao := dao.NewFavorite(global.DB)
	commentDao := dao.NewComment(global.DB)
	favoritePublisher := mq.NewFavoritePublisher(global.AmqpConn, "favorite")
	favoriteSubscriber := mq.NewFavoriteSubscriber(global.AmqpConn, "favorite")
	commentPublisher := mq.NewCommentPublisher(global.AmqpConn, "comment")
	CommentSubscriber := mq.NewCommentSubscriber(global.AmqpConn, "comment")
	go func() {
		if err := pkg.SubscribeFavorite(favoriteSubscriber, favoriteDao); err != nil {
			klog.Errorf("favorite action goroutine error: %s", err.Error())
			panic(err)
		}
	}()
	go func() {
		if err := pkg.SubscribeComment(CommentSubscriber, commentDao); err != nil {
			klog.Errorf("comment action goroutine error: %s", err.Error())
			panic(err)
		}
	}()

	etcdConfig := global.EtcdConfig
	etcdAddr := fmt.Sprintf("%s:%d", etcdConfig.Host, etcdConfig.Port)
	r, err := etcd.NewEtcdRegistry([]string{etcdAddr})
	if err != nil {
		panic(err)
	}
	rpcConfig := global.RPCConfig
	rpcAddr := fmt.Sprintf("%s:%d", rpcConfig.Host, rpcConfig.Port)
	addr, err := net.ResolveTCPAddr("tcp", rpcAddr)
	if err != nil {
		panic(err)
	}
	svr := interaction.NewServer(&InteractionServerImpl{
		FavoriteDao:       favoriteDao,
		CommentDao:        commentDao,
		FavoriteRedisDao:  redis.NewFavorite(global.RedisFavoriteClient),
		CommentRedisDao:   redis.NewComment(global.RedisCommentClient),
		FavoritePublisher: favoritePublisher,
		CommentPublisher:  commentPublisher,
	},
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: rpcConfig.Name}),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		server.WithRegistry(r),
	)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
