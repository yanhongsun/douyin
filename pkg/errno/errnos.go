package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode         = 0
	ServiceErrCode      = 40001
	CommentParamErrCode = 40002
	UserIdErrCode       = 40003
	VideoIdErrCode      = 40004
	CommentIdErrCode    = 40005
	ActionTypeErrCode   = 40006
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
	Success         = NewErrNo(SuccessCode, "Success")
	ServiceErr      = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	CommentParamErr = NewErrNo(CommentParamErrCode, "Wrong Comment Parameter has been given")
	UserIdErr       = NewErrNo(UserIdErrCode, "Wrong UserId Parameter or Wrong UserId has been given")
	VideoIdErr      = NewErrNo(VideoIdErrCode, "Wrong VedioId Parameter or Wrong VedioId has been given")
	CommentIdErr    = NewErrNo(CommentIdErrCode, "Wrong CommentId Parameter or Wrong CommentId has been given")
	ActionTypeErr   = NewErrNo(ActionTypeErrCode, "Wrong ActionTypeErr has been given")
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
