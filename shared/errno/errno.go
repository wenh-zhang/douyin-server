package errno

import (
	"douyin/shared/kitex_gen/errno"
	"errors"
	"fmt"
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func (e ErrNo) ErrorDetail() (int32, string) {
	return e.ErrCode, e.ErrMsg
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

var (
	Success              = NewErrNo(int32(errno.Code_Success), "success")
	ServiceErr           = NewErrNo(int32(errno.Code_ServiceErr), "service error")
	ParamErr             = NewErrNo(int32(errno.Code_ParamsErr), "parameter error")
	RPCInteractionErr    = NewErrNo(int32(errno.Code_RPCInteractionErr), "call rpc interaction error")
	InteractionServerErr = NewErrNo(int32(errno.Code_InteractionServerErr), "interaction server error")
	RPCSocialityErr      = NewErrNo(int32(errno.Code_RPCSocialityErr), "call rpc sociality error")
	SocialityServerErr   = NewErrNo(int32(errno.Code_SocialityServerErr), "sociality server error")
	RPCUserErr           = NewErrNo(int32(errno.Code_RPCUserErr), "call rpc user error")
	UserServerErr        = NewErrNo(int32(errno.Code_UserServerErr), "user server error")
	UserAlreadyExistErr  = NewErrNo(int32(errno.Code_UserAlreadyExistErr), "user already exist")
	UserNotFoundErr      = NewErrNo(int32(errno.Code_UserNotFoundErr), "user not found")
	AuthorizeFailErr     = NewErrNo(int32(errno.Code_AuthorizeFailErr), "authorize fail")
	RPCVideoErr          = NewErrNo(int32(errno.Code_RPCVideoErr), "call rpc video error")
	VideoServerErr       = NewErrNo(int32(errno.Code_VideoServerErr), "video server error")
	RPCMessageErr        = NewErrNo(int32(errno.Code_RPCMessageErr), "call rpc message error")
	MessageServerErr     = NewErrNo(int32(errno.Code_MessageServerErr), "message server error")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr.WithMessage(err.Error())
	return s
}
