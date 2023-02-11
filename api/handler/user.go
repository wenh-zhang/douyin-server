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

type UserRegisterResponse struct {
	Response
	UserID int64  `json:"user_id, omitempty"`
	Token  string `json:"token,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserID int64  `json:"user_id, omitempty"`
	Token  string `json:"token,omitempty"`
}

type UserInfoResponse struct {
	Response
	*core.User
}

func UserRegister(ctx context.Context, c *app.RequestContext) {
	request := new(core.DouyinUserRegisterRequest)
	username := c.Query(constant.UserName)
	password := c.Query(constant.Password)
	if len(username) == 0 || len(password) == 0 {
		SendResponse(c, NewResponse(errno.ParamErr))
		return
	}
	request.Username, request.Password = username, password
	userID, err := rpc.UserRegister(context.Background(), request)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	var response UserRegisterResponse
	response.Response = NewResponse(errno.Success)
	response.UserID = userID
	token, err := util.GenerateToken(userID)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	response.Token = token
	SendResponse(c, response)
}

func UserLogin(ctx context.Context, c *app.RequestContext) {
	request := new(core.DouyinUserLoginRequest)
	username := c.Query(constant.UserName)
	password := c.Query(constant.Password)
	if len(username) == 0 || len(password) == 0 {
		SendResponse(c, NewResponse(errno.ParamErr))
		return
	}
	request.Username, request.Password = username, password
	userID, err := rpc.UserLogin(context.Background(), request)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	var response UserLoginResponse
	response.Response = NewResponse(errno.Success)
	response.UserID = userID
	token, err := util.GenerateToken(userID)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	response.Token = token
	SendResponse(c, response)
}

func UserInfo(ctx context.Context, c *app.RequestContext) {
	request := new(core.DouyinUserRequest)
	userID, _ := c.Get(constant.TokenUserIdentifyKey)
	queryUserIDStr := c.Query(constant.UserIdentityKey)
	queryUserID, err := strconv.ParseInt(queryUserIDStr, 10, 64)
	if err != nil{
		SendResponse(c, NewResponse(err))
		return
	}

	request.QueryUserId, request.UserId = queryUserID, userID.(int64)
	user, err := rpc.UserInfo(context.Background(), request)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	SendResponse(c, UserInfoResponse{
		Response: NewResponse(errno.Success),
		User: user,
	})
}
