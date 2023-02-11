package rpc

import (
	"context"
	"douyin/kitex_gen/core"
	"douyin/kitex_gen/interact"
	"douyin/kitex_gen/interact/interactservice"
	"douyin/pkg/constant"
	"douyin/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var interactClient interactservice.Client

func initInteractRPC() {
	r, err := etcd.NewEtcdResolver([]string{constant.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := interactservice.NewClient(
		constant.InteractServiceName,
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
	interactClient = c
}

func FavoriteAction(ctx context.Context, req *interact.DouyinFavoriteActionRequest) error {
	resp, err := interactClient.FavoriteAction(ctx, req)
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

func FavoriteList(ctx context.Context, req *interact.DouyinFavoriteListRequest) ([]*core.Video, error) {
	resp, err := interactClient.FavoriteList(ctx, req)
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

func CommentAction(ctx context.Context, req *interact.DouyinCommentActionRequest) (*interact.Comment, error){
	resp, err := interactClient.CommentAction(ctx, req)
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
	return resp.Comment, nil
}

func CommentList(ctx context.Context, req *interact.DouyinCommentListRequest) ([]*interact.Comment, error){
	resp, err := interactClient.CommentList(ctx, req)
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
	return resp.CommentList, nil
}