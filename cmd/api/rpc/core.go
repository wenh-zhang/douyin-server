package rpc

import (
	"context"
	"douyin/kitex_gen/core"
	"douyin/kitex_gen/core/coreservice"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var coreClient coreservice.Client

func initCoreRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := coreservice.NewClient(
		constants.CoreServiceName,
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
	coreClient = c
}

func Feed(ctx context.Context, req *core.DouyinFeedRequest) ([]*core.Video, *int64, error) {
	resp, err := coreClient.Feed(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode != errno.SuccessCode {
		var statusMsg string
		if resp.StatusMsg != nil {
			statusMsg = *resp.StatusMsg
		}
		return nil, nil, errno.NewErrNo(resp.StatusCode, statusMsg)
	}
	return resp.VideoList, resp.NextTime, nil
}
