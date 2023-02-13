package handler

import (
	"context"
	"douyin/api/rpc"
	"douyin/kitex_gen/core"
	"douyin/pkg/constant"
	"douyin/pkg/errno"
	"douyin/pkg/util"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

// UserRegisterResponse The returned json format of register request
type UserRegisterResponse struct {
	Response
	UserID int64  `json:"user_id, omitempty"`
	Token  string `json:"token,omitempty"`
}

// UserLoginResponse The returned json format of login request
type UserLoginResponse struct {
	Response
	UserID int64  `json:"user_id, omitempty"`
	Token  string `json:"token,omitempty"`
}

// UserInfoResponse The returned json format of querying user information request
type UserInfoResponse struct {
	Response
	*core.User
}

// UserRegister The handler of request for user to register
func UserRegister(ctx context.Context, c *app.RequestContext) {
	request := new(core.DouyinUserRegisterRequest)
	username := c.Query(constant.UserName)
	password := c.Query(constant.Password)
	if len(username) == 0 || len(password) == 0 {
		SendResponse(c, NewResponse(errno.ParamErr))
		return
	}
	request.Username, request.Password = username, password

	// call RPC service
	userID, err := rpc.UserRegister(context.Background(), request)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	var response UserRegisterResponse
	response.Response = NewResponse(errno.Success)
	response.UserID = userID

	// generate token after register
	token, err := util.GenerateToken(userID)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	response.Token = token
	SendResponse(c, response)
}

// UserLogin Handler of request for user to login
func UserLogin(ctx context.Context, c *app.RequestContext) {
	request := new(core.DouyinUserLoginRequest)
	username := c.Query(constant.UserName)
	password := c.Query(constant.Password)
	if len(username) == 0 || len(password) == 0 {
		SendResponse(c, NewResponse(errno.ParamErr))
		return
	}
	request.Username, request.Password = username, password

	// call RPC service
	userID, err := rpc.UserLogin(context.Background(), request)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	var response UserLoginResponse
	response.Response = NewResponse(errno.Success)
	response.UserID = userID

	// generate token after check user's qualification
	token, err := util.GenerateToken(userID)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	response.Token = token
	SendResponse(c, response)
}

// UserInfo Handler of request for user to query another user's information
func UserInfo(ctx context.Context, c *app.RequestContext) {
	request := new(core.DouyinUserRequest)
	userID, _ := c.Get(constant.TokenUserIdentifyKey)
	queryUserIDStr := c.Query(constant.UserIdentityKey)
	queryUserID, err := strconv.ParseInt(queryUserIDStr, 10, 64)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	request.QueryUserId, request.UserId = queryUserID, userID.(int64)

	// call RPC service
	user, err := rpc.UserInfo(context.Background(), request)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	SendResponse(c, UserInfoResponse{
		Response: NewResponse(errno.Success),
		User:     user,
	})
}
