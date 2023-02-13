package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode                = 0
	ServiceErrCode             = 10001
	ParamErrCode               = 10002
	UserAlreadyExistErrCode    = 10003
	AuthorizationFailedErrCode = 10004
	LackTokenErrCode           = 10005
	TokenTimeOutErrCode        = 10006
	UserNotExistErrCode        = 10007
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
	Success                = NewErrNo(SuccessCode, "Success")
	ServiceErr             = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr               = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
	LackTokenErr           = NewErrNo(LackTokenErrCode, "Lack parameter token")
	TokenTimeOutErr        = NewErrNo(TokenTimeOutErrCode, "Token time out")
	UserNotExistErr        = NewErrNo(UserNotExistErrCode, "User not exists")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
