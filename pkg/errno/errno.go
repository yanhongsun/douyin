package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode             = 0
	ServiceErrCode          = 10001
	ParamErrCode            = 10002
	LoginErrCode            = 10003
	UserNotExistErrCode     = 10004
	UserAlreadyExistErrCode = 10005
	VideoErrCode            = 10006
	TokenExpiredErrCode     = 10031
	TokenNotValidYetErrCode = 10032
	TokenMalformedErrCode   = 10033
	TokenInvalidErrCode     = 10034

	CommentServiceErrCode = 40001
	CommentParamErrCode   = 40002
	UserIdErrCode         = 40003
	VideoIdErrCode        = 40004
	CommentIdErrCode      = 40005
	ActionTypeErrCode     = 40006
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	ServiceErr          = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr            = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	LoginErr            = NewErrNo(LoginErrCode, "Wrong username or password")
	UserNotExistErr     = NewErrNo(UserNotExistErrCode, "User does not exists")
	UserAlreadyExistErr = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	VideoErr            = NewErrNo(VideoErrCode, "Video is Empty or Invalid")
	TokenExpiredErr     = NewErrNo(TokenExpiredErrCode, "Token is expired")
	TokenNotValidYetErr = NewErrNo(TokenNotValidYetErrCode, "Token is not active yet")
	TokenMalformedErr   = NewErrNo(TokenMalformedErrCode, "That's not even a token")
	TokenInvalidErr     = NewErrNo(TokenInvalidErrCode, "Couldn't handle this token")

	Success           = NewErrNo(SuccessCode, "Success")
	CommentServiceErr = NewErrNo(CommentServiceErrCode, "Service is unable to start successfully")
	CommentParamErr   = NewErrNo(CommentParamErrCode, "Wrong Comment Parameter has been given")
	UserIdErr         = NewErrNo(UserIdErrCode, "Wrong UserId Parameter or Wrong UserId has been given")
	VideoIdErr        = NewErrNo(VideoIdErrCode, "Wrong VideoId Parameter or Wrong VideoId has been given")
	CommentIdErr      = NewErrNo(CommentIdErrCode, "Wrong CommentId Parameter or Wrong CommentId has been given")
	ActionTypeErr     = NewErrNo(ActionTypeErrCode, "Wrong ActionTypeErr has been given")
)

func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
