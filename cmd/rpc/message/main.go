package main

import (
	"douyin/cmd/rpc/message/dao"
	"douyin/cmd/rpc/message/global"
	"douyin/cmd/rpc/message/initialize"
	"douyin/cmd/rpc/message/mq"
	"douyin/cmd/rpc/message/pkg"
	message "douyin/shared/kitex_gen/message/messageservice"
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
	messageDao := dao.NewMessage(global.DB)
	publisher := mq.NewPublisher(global.AmqpConn, "message")
	subscriber := mq.NewSubscriber(global.AmqpConn, "message")
	go func() {
		if err := pkg.SubscribeMessage(subscriber, messageDao); err != nil {
			klog.Errorf("send message goroutine error: %s", err.Error())
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
	svr := message.NewServer(&MessageServiceImpl{
		Dao:       messageDao,
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
