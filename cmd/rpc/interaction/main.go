package main

import (
	"douyin/cmd/rpc/interaction/dao"
	"douyin/cmd/rpc/interaction/global"
	"douyin/cmd/rpc/interaction/initialize"
	"douyin/cmd/rpc/interaction/redis"
	interaction "douyin/shared/kitex_gen/interaction/interactionserver"
	"fmt"
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
		FavoriteDao:      dao.NewFavorite(global.DB),
		CommentDao:       dao.NewComment(global.DB),
		FavoriteRedisDao: redis.NewFavorite(global.RedisFavoriteClient),
		CommentRedisDao:  redis.NewComment(global.RedisCommentClient),
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
