package main

import (
	"douyin/cmd/rpc/video/dao"
	"douyin/cmd/rpc/video/global"
	"douyin/cmd/rpc/video/initialize"
	"douyin/cmd/rpc/video/mq"
	"douyin/cmd/rpc/video/pkg"
	video "douyin/shared/kitex_gen/video/videoservice"
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
	videoDao := dao.NewVideo(global.DB)
	publisher := mq.NewPublisher(global.AmqpConn, "video")
	subscriber := mq.NewSubscriber(global.AmqpConn, "video")
	go func() {
		if err := pkg.SubscribeVideo(subscriber, videoDao); err != nil {
			klog.Errorf("publish action goroutine error: %s", err.Error())
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
	svr := video.NewServer(&VideoServiceImpl{
		Dao:       videoDao,
		Publisher: publisher,
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
