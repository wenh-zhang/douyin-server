package main

import (
	"douyin/cmd/rpc/sociality/dao"
	"douyin/cmd/rpc/sociality/global"
	"douyin/cmd/rpc/sociality/initialize"
	"douyin/cmd/rpc/sociality/mq"
	"douyin/cmd/rpc/sociality/pkg"
	"douyin/cmd/rpc/sociality/redis"
	sociality "douyin/shared/kitex_gen/sociality/socialityservice"
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
	followDao := dao.NewFollow(global.DB)
	publisher := mq.NewPublisher(global.AmqpConn, "follow")
	subscriber := mq.NewSubscriber(global.AmqpConn, "follow")
	go func() {
		if err := pkg.SubscribeFollow(subscriber, followDao); err != nil {
			klog.Errorf("follow action goroutine error: %s", err.Error())
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
	svr := sociality.NewServer(&SocialityServiceImpl{
		Dao:            followDao,
		FollowRedisDao: redis.NewFollow(global.RedisFollowClient),
		Publisher:      publisher,
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
