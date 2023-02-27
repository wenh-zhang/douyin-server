// Code generated by Kitex v0.4.4. DO NOT EDIT.

package messageservice

import (
	"context"
	"douyin/shared/kitex_gen/message"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SendMessage(ctx context.Context, req *message.DouyinMessageActionRequest, callOptions ...callopt.Option) (r *message.DouyinMessageActionResponse, err error)
	GetMessageHistory(ctx context.Context, req *message.DouyinGetMessageChatRequest, callOptions ...callopt.Option) (r *message.DouyinGetMessageChatResponse, err error)
	BatchGetLatestMessage(ctx context.Context, req *message.DouyinBatchGetLatestMessageRequest, callOptions ...callopt.Option) (r *message.DouyinBatchGetLatestMessageResponse, err error)
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
	return &kMessageServiceClient{
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

type kMessageServiceClient struct {
	*kClient
}

func (p *kMessageServiceClient) SendMessage(ctx context.Context, req *message.DouyinMessageActionRequest, callOptions ...callopt.Option) (r *message.DouyinMessageActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendMessage(ctx, req)
}

func (p *kMessageServiceClient) GetMessageHistory(ctx context.Context, req *message.DouyinGetMessageChatRequest, callOptions ...callopt.Option) (r *message.DouyinGetMessageChatResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMessageHistory(ctx, req)
}

func (p *kMessageServiceClient) BatchGetLatestMessage(ctx context.Context, req *message.DouyinBatchGetLatestMessageRequest, callOptions ...callopt.Option) (r *message.DouyinBatchGetLatestMessageResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BatchGetLatestMessage(ctx, req)
}