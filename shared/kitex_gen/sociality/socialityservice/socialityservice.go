// Code generated by Kitex v0.4.4. DO NOT EDIT.

package socialityservice

import (
	"context"
	"douyin/shared/kitex_gen/sociality"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return socialityServiceServiceInfo
}

var socialityServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "SocialityService"
	handlerType := (*sociality.SocialityService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Relation":              kitex.NewMethodInfo(relationHandler, newSocialityServiceRelationArgs, newSocialityServiceRelationResult, false),
		"GetRelationUserIdList": kitex.NewMethodInfo(getRelationUserIdListHandler, newSocialityServiceGetRelationUserIdListArgs, newSocialityServiceGetRelationUserIdListResult, false),
		"BatchGetSocialInfo":    kitex.NewMethodInfo(batchGetSocialInfoHandler, newSocialityServiceBatchGetSocialInfoArgs, newSocialityServiceBatchGetSocialInfoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "sociality",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func relationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*sociality.SocialityServiceRelationArgs)
	realResult := result.(*sociality.SocialityServiceRelationResult)
	success, err := handler.(sociality.SocialityService).Relation(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialityServiceRelationArgs() interface{} {
	return sociality.NewSocialityServiceRelationArgs()
}

func newSocialityServiceRelationResult() interface{} {
	return sociality.NewSocialityServiceRelationResult()
}

func getRelationUserIdListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*sociality.SocialityServiceGetRelationUserIdListArgs)
	realResult := result.(*sociality.SocialityServiceGetRelationUserIdListResult)
	success, err := handler.(sociality.SocialityService).GetRelationUserIdList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialityServiceGetRelationUserIdListArgs() interface{} {
	return sociality.NewSocialityServiceGetRelationUserIdListArgs()
}

func newSocialityServiceGetRelationUserIdListResult() interface{} {
	return sociality.NewSocialityServiceGetRelationUserIdListResult()
}

func batchGetSocialInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*sociality.SocialityServiceBatchGetSocialInfoArgs)
	realResult := result.(*sociality.SocialityServiceBatchGetSocialInfoResult)
	success, err := handler.(sociality.SocialityService).BatchGetSocialInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialityServiceBatchGetSocialInfoArgs() interface{} {
	return sociality.NewSocialityServiceBatchGetSocialInfoArgs()
}

func newSocialityServiceBatchGetSocialInfoResult() interface{} {
	return sociality.NewSocialityServiceBatchGetSocialInfoResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Relation(ctx context.Context, req *sociality.DouyinRelationActionRequest) (r *sociality.DouyinRelationActionResponse, err error) {
	var _args sociality.SocialityServiceRelationArgs
	_args.Req = req
	var _result sociality.SocialityServiceRelationResult
	if err = p.c.Call(ctx, "Relation", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetRelationUserIdList(ctx context.Context, req *sociality.DouyinGetRelationUserIdListRequest) (r *sociality.DouyinGetRelationUserIdListResponse, err error) {
	var _args sociality.SocialityServiceGetRelationUserIdListArgs
	_args.Req = req
	var _result sociality.SocialityServiceGetRelationUserIdListResult
	if err = p.c.Call(ctx, "GetRelationUserIdList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BatchGetSocialInfo(ctx context.Context, req *sociality.DouyinBatchGetSocialInfoRequest) (r *sociality.DouyinBatchGetSocialInfoResponse, err error) {
	var _args sociality.SocialityServiceBatchGetSocialInfoArgs
	_args.Req = req
	var _result sociality.SocialityServiceBatchGetSocialInfoResult
	if err = p.c.Call(ctx, "BatchGetSocialInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}