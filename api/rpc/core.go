package rpc

import (
	"context"
	"douyin/kitex_gen/core"
	"douyin/kitex_gen/core/coreservice"
	"douyin/pkg/constant"
	"douyin/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var coreClient coreservice.Client

func initCoreRPC() {
	r, err := etcd.NewEtcdResolver([]string{constant.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := coreservice.NewClient(
		constant.CoreServiceName,
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

// Feed Get the list of videos from response by calling rpc service
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

func UserRegister(ctx context.Context, req *core.DouyinUserRegisterRequest) (int64, error) {
	resp, err := coreClient.UserRegister(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != errno.SuccessCode {
		var statusMsg string
		if resp.StatusMsg != nil {
			statusMsg = *resp.StatusMsg
		}
		return 0, errno.NewErrNo(resp.StatusCode, statusMsg)
	}
	return resp.UserId, nil
}

func UserLogin(ctx context.Context, req *core.DouyinUserLoginRequest) (int64, error) {
	resp, err := coreClient.UserLogin(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != errno.SuccessCode {
		var statusMsg string
		if resp.StatusMsg != nil {
			statusMsg = *resp.StatusMsg
		}
		return 0, errno.NewErrNo(resp.StatusCode, statusMsg)
	}
	return resp.UserId, nil
}

func UserInfo(ctx context.Context, req *core.DouyinUserRequest) (*core.User, error) {
	resp, err := coreClient.UserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != errno.SuccessCode {
		var statusMsg string
		if resp.StatusMsg != nil {
			statusMsg = *resp.StatusMsg
		}
		return nil, errno.NewErrNo(resp.StatusCode, statusMsg)
	}
	return resp.User, nil
}

func PublishAction(ctx context.Context, req *core.DouyinPublishActionRequest) error {
	resp, err := coreClient.PublishAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != errno.SuccessCode {
		var statusMsg string
		if resp.StatusMsg != nil {
			statusMsg = *resp.StatusMsg
		}
		return errno.NewErrNo(resp.StatusCode, statusMsg)
	}
	return nil
}

func PublishList(ctx context.Context, req *core.DouyinPublishListRequest) ([]*core.Video, error) {
	resp, err := coreClient.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != errno.SuccessCode {
		var statusMsg string
		if resp.StatusMsg != nil {
			statusMsg = *resp.StatusMsg
		}
		return nil, errno.NewErrNo(resp.StatusCode, statusMsg)
	}
	return resp.VideoList, nil
}
