// Code generated by Kitex v0.4.4. DO NOT EDIT.

package coreservice

import (
	"context"
	core "douyin/kitex_gen/core"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return coreServiceServiceInfo
}

var coreServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CoreService"
	handlerType := (*core.CoreService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Feed":          kitex.NewMethodInfo(feedHandler, newFeedArgs, newFeedResult, false),
		"UserRegister":  kitex.NewMethodInfo(userRegisterHandler, newUserRegisterArgs, newUserRegisterResult, false),
		"UserLogin":     kitex.NewMethodInfo(userLoginHandler, newUserLoginArgs, newUserLoginResult, false),
		"UserInfo":      kitex.NewMethodInfo(userInfoHandler, newUserInfoArgs, newUserInfoResult, false),
		"PublishAction": kitex.NewMethodInfo(publishActionHandler, newPublishActionArgs, newPublishActionResult, false),
		"PublishList":   kitex.NewMethodInfo(publishListHandler, newPublishListArgs, newPublishListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "douyin.core",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func feedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core.DouyinFeedRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core.CoreService).Feed(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FeedArgs:
		success, err := handler.(core.CoreService).Feed(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FeedResult)
		realResult.Success = success
	}
	return nil
}
func newFeedArgs() interface{} {
	return &FeedArgs{}
}

func newFeedResult() interface{} {
	return &FeedResult{}
}

type FeedArgs struct {
	Req *core.DouyinFeedRequest
}

func (p *FeedArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core.DouyinFeedRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FeedArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FeedArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FeedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FeedArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FeedArgs) Unmarshal(in []byte) error {
	msg := new(core.DouyinFeedRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FeedArgs_Req_DEFAULT *core.DouyinFeedRequest

func (p *FeedArgs) GetReq() *core.DouyinFeedRequest {
	if !p.IsSetReq() {
		return FeedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FeedArgs) IsSetReq() bool {
	return p.Req != nil
}

type FeedResult struct {
	Success *core.DouyinFeedResponse
}

var FeedResult_Success_DEFAULT *core.DouyinFeedResponse

func (p *FeedResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core.DouyinFeedResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FeedResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FeedResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FeedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FeedResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FeedResult) Unmarshal(in []byte) error {
	msg := new(core.DouyinFeedResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FeedResult) GetSuccess() *core.DouyinFeedResponse {
	if !p.IsSetSuccess() {
		return FeedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FeedResult) SetSuccess(x interface{}) {
	p.Success = x.(*core.DouyinFeedResponse)
}

func (p *FeedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func userRegisterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core.DouyinUserRegisterRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core.CoreService).UserRegister(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UserRegisterArgs:
		success, err := handler.(core.CoreService).UserRegister(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UserRegisterResult)
		realResult.Success = success
	}
	return nil
}
func newUserRegisterArgs() interface{} {
	return &UserRegisterArgs{}
}

func newUserRegisterResult() interface{} {
	return &UserRegisterResult{}
}

type UserRegisterArgs struct {
	Req *core.DouyinUserRegisterRequest
}

func (p *UserRegisterArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core.DouyinUserRegisterRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UserRegisterArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UserRegisterArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UserRegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UserRegisterArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UserRegisterArgs) Unmarshal(in []byte) error {
	msg := new(core.DouyinUserRegisterRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UserRegisterArgs_Req_DEFAULT *core.DouyinUserRegisterRequest

func (p *UserRegisterArgs) GetReq() *core.DouyinUserRegisterRequest {
	if !p.IsSetReq() {
		return UserRegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UserRegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

type UserRegisterResult struct {
	Success *core.DouyinFeedResponse
}

var UserRegisterResult_Success_DEFAULT *core.DouyinFeedResponse

func (p *UserRegisterResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core.DouyinFeedResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UserRegisterResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UserRegisterResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UserRegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UserRegisterResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UserRegisterResult) Unmarshal(in []byte) error {
	msg := new(core.DouyinFeedResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UserRegisterResult) GetSuccess() *core.DouyinFeedResponse {
	if !p.IsSetSuccess() {
		return UserRegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UserRegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*core.DouyinFeedResponse)
}

func (p *UserRegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func userLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core.DouyinUserLoginRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core.CoreService).UserLogin(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UserLoginArgs:
		success, err := handler.(core.CoreService).UserLogin(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UserLoginResult)
		realResult.Success = success
	}
	return nil
}
func newUserLoginArgs() interface{} {
	return &UserLoginArgs{}
}

func newUserLoginResult() interface{} {
	return &UserLoginResult{}
}

type UserLoginArgs struct {
	Req *core.DouyinUserLoginRequest
}

func (p *UserLoginArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core.DouyinUserLoginRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UserLoginArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UserLoginArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UserLoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UserLoginArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UserLoginArgs) Unmarshal(in []byte) error {
	msg := new(core.DouyinUserLoginRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UserLoginArgs_Req_DEFAULT *core.DouyinUserLoginRequest

func (p *UserLoginArgs) GetReq() *core.DouyinUserLoginRequest {
	if !p.IsSetReq() {
		return UserLoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UserLoginArgs) IsSetReq() bool {
	return p.Req != nil
}

type UserLoginResult struct {
	Success *core.DouyinUserLoginResponse
}

var UserLoginResult_Success_DEFAULT *core.DouyinUserLoginResponse

func (p *UserLoginResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core.DouyinUserLoginResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UserLoginResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UserLoginResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UserLoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UserLoginResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UserLoginResult) Unmarshal(in []byte) error {
	msg := new(core.DouyinUserLoginResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UserLoginResult) GetSuccess() *core.DouyinUserLoginResponse {
	if !p.IsSetSuccess() {
		return UserLoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UserLoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*core.DouyinUserLoginResponse)
}

func (p *UserLoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func userInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core.DouyinUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core.CoreService).UserInfo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UserInfoArgs:
		success, err := handler.(core.CoreService).UserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UserInfoResult)
		realResult.Success = success
	}
	return nil
}
func newUserInfoArgs() interface{} {
	return &UserInfoArgs{}
}

func newUserInfoResult() interface{} {
	return &UserInfoResult{}
}

type UserInfoArgs struct {
	Req *core.DouyinUserRequest
}

func (p *UserInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core.DouyinUserRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UserInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UserInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UserInfoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UserInfoArgs) Unmarshal(in []byte) error {
	msg := new(core.DouyinUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UserInfoArgs_Req_DEFAULT *core.DouyinUserRequest

func (p *UserInfoArgs) GetReq() *core.DouyinUserRequest {
	if !p.IsSetReq() {
		return UserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

type UserInfoResult struct {
	Success *core.DouyinUserResponse
}

var UserInfoResult_Success_DEFAULT *core.DouyinUserResponse

func (p *UserInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core.DouyinUserResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UserInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UserInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UserInfoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UserInfoResult) Unmarshal(in []byte) error {
	msg := new(core.DouyinUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UserInfoResult) GetSuccess() *core.DouyinUserResponse {
	if !p.IsSetSuccess() {
		return UserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*core.DouyinUserResponse)
}

func (p *UserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func publishActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core.DouyinPublishActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core.CoreService).PublishAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PublishActionArgs:
		success, err := handler.(core.CoreService).PublishAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PublishActionResult)
		realResult.Success = success
	}
	return nil
}
func newPublishActionArgs() interface{} {
	return &PublishActionArgs{}
}

func newPublishActionResult() interface{} {
	return &PublishActionResult{}
}

type PublishActionArgs struct {
	Req *core.DouyinPublishActionRequest
}

func (p *PublishActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core.DouyinPublishActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PublishActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PublishActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PublishActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PublishActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PublishActionArgs) Unmarshal(in []byte) error {
	msg := new(core.DouyinPublishActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PublishActionArgs_Req_DEFAULT *core.DouyinPublishActionRequest

func (p *PublishActionArgs) GetReq() *core.DouyinPublishActionRequest {
	if !p.IsSetReq() {
		return PublishActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PublishActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type PublishActionResult struct {
	Success *core.DouyinPublishActionResponse
}

var PublishActionResult_Success_DEFAULT *core.DouyinPublishActionResponse

func (p *PublishActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core.DouyinPublishActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PublishActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PublishActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PublishActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PublishActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PublishActionResult) Unmarshal(in []byte) error {
	msg := new(core.DouyinPublishActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PublishActionResult) GetSuccess() *core.DouyinPublishActionResponse {
	if !p.IsSetSuccess() {
		return PublishActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PublishActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*core.DouyinPublishActionResponse)
}

func (p *PublishActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func publishListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core.DouyinPublishListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core.CoreService).PublishList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PublishListArgs:
		success, err := handler.(core.CoreService).PublishList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PublishListResult)
		realResult.Success = success
	}
	return nil
}
func newPublishListArgs() interface{} {
	return &PublishListArgs{}
}

func newPublishListResult() interface{} {
	return &PublishListResult{}
}

type PublishListArgs struct {
	Req *core.DouyinPublishListRequest
}

func (p *PublishListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core.DouyinPublishListRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PublishListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PublishListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PublishListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PublishListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PublishListArgs) Unmarshal(in []byte) error {
	msg := new(core.DouyinPublishListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PublishListArgs_Req_DEFAULT *core.DouyinPublishListRequest

func (p *PublishListArgs) GetReq() *core.DouyinPublishListRequest {
	if !p.IsSetReq() {
		return PublishListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PublishListArgs) IsSetReq() bool {
	return p.Req != nil
}

type PublishListResult struct {
	Success *core.DouyinPublishListResponse
}

var PublishListResult_Success_DEFAULT *core.DouyinPublishListResponse

func (p *PublishListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core.DouyinPublishListResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PublishListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PublishListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PublishListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PublishListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PublishListResult) Unmarshal(in []byte) error {
	msg := new(core.DouyinPublishListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PublishListResult) GetSuccess() *core.DouyinPublishListResponse {
	if !p.IsSetSuccess() {
		return PublishListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PublishListResult) SetSuccess(x interface{}) {
	p.Success = x.(*core.DouyinPublishListResponse)
}

func (p *PublishListResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Feed(ctx context.Context, Req *core.DouyinFeedRequest) (r *core.DouyinFeedResponse, err error) {
	var _args FeedArgs
	_args.Req = Req
	var _result FeedResult
	if err = p.c.Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserRegister(ctx context.Context, Req *core.DouyinUserRegisterRequest) (r *core.DouyinFeedResponse, err error) {
	var _args UserRegisterArgs
	_args.Req = Req
	var _result UserRegisterResult
	if err = p.c.Call(ctx, "UserRegister", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserLogin(ctx context.Context, Req *core.DouyinUserLoginRequest) (r *core.DouyinUserLoginResponse, err error) {
	var _args UserLoginArgs
	_args.Req = Req
	var _result UserLoginResult
	if err = p.c.Call(ctx, "UserLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserInfo(ctx context.Context, Req *core.DouyinUserRequest) (r *core.DouyinUserResponse, err error) {
	var _args UserInfoArgs
	_args.Req = Req
	var _result UserInfoResult
	if err = p.c.Call(ctx, "UserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishAction(ctx context.Context, Req *core.DouyinPublishActionRequest) (r *core.DouyinPublishActionResponse, err error) {
	var _args PublishActionArgs
	_args.Req = Req
	var _result PublishActionResult
	if err = p.c.Call(ctx, "PublishAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishList(ctx context.Context, Req *core.DouyinPublishListRequest) (r *core.DouyinPublishListResponse, err error) {
	var _args PublishListArgs
	_args.Req = Req
	var _result PublishListResult
	if err = p.c.Call(ctx, "PublishList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
