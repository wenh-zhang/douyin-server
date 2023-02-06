// Code generated by Kitex v0.4.4. DO NOT EDIT.

package coreservice

import (
	"context"
	core "douyin/kitex_gen/core"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Feed(ctx context.Context, Req *core.DouyinFeedRequest, callOptions ...callopt.Option) (r *core.DouyinFeedResponse, err error)
	UserRegister(ctx context.Context, Req *core.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *core.DouyinFeedResponse, err error)
	UserLogin(ctx context.Context, Req *core.DouyinUserLoginRequest, callOptions ...callopt.Option) (r *core.DouyinUserLoginResponse, err error)
	UserInfo(ctx context.Context, Req *core.DouyinUserRequest, callOptions ...callopt.Option) (r *core.DouyinUserResponse, err error)
	PublishAction(ctx context.Context, Req *core.DouyinPublishActionRequest, callOptions ...callopt.Option) (r *core.DouyinPublishActionResponse, err error)
	PublishList(ctx context.Context, Req *core.DouyinPublishListRequest, callOptions ...callopt.Option) (r *core.DouyinPublishListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kCoreServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCoreServiceClient struct {
	*kClient
}

func (p *kCoreServiceClient) Feed(ctx context.Context, Req *core.DouyinFeedRequest, callOptions ...callopt.Option) (r *core.DouyinFeedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Feed(ctx, Req)
}

func (p *kCoreServiceClient) UserRegister(ctx context.Context, Req *core.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *core.DouyinFeedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserRegister(ctx, Req)
}

func (p *kCoreServiceClient) UserLogin(ctx context.Context, Req *core.DouyinUserLoginRequest, callOptions ...callopt.Option) (r *core.DouyinUserLoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserLogin(ctx, Req)
}

func (p *kCoreServiceClient) UserInfo(ctx context.Context, Req *core.DouyinUserRequest, callOptions ...callopt.Option) (r *core.DouyinUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserInfo(ctx, Req)
}

func (p *kCoreServiceClient) PublishAction(ctx context.Context, Req *core.DouyinPublishActionRequest, callOptions ...callopt.Option) (r *core.DouyinPublishActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishAction(ctx, Req)
}

func (p *kCoreServiceClient) PublishList(ctx context.Context, Req *core.DouyinPublishListRequest, callOptions ...callopt.Option) (r *core.DouyinPublishListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishList(ctx, Req)
}
