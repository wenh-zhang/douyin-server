package initialize

import (
	"douyin/cmd/rpc/interaction/global"
	"douyin/shared/constant"
	"douyin/shared/kitex_gen/user/userservice"
	"douyin/shared/kitex_gen/video/videoservice"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

func InitRPC(){
	initUserRPC()
	initVideoRPC()
}

func initUserRPC() {
	etcdConfig := global.EtcdConfig
	etcdAddr := fmt.Sprintf("%s:%d", etcdConfig.Host, etcdConfig.Port)
	r, err := etcd.NewEtcdResolver([]string{etcdAddr})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constant.RPCUserName,
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	global.UserClient = c
}

func initVideoRPC() {
	etcdConfig := global.EtcdConfig
	etcdAddr := fmt.Sprintf("%s:%d", etcdConfig.Host, etcdConfig.Port)
	r, err := etcd.NewEtcdResolver([]string{etcdAddr})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		constant.RPCVideoName,
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	global.VideoClient = c
}