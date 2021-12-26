package errno

import (
	"fmt"
)

// 自定义错误码
type ErrNo struct {
	Code    int
	Message string
}

func (err ErrNo) Error() string {
	return err.Message
}

// 自定义错误
type Err struct {
	ErrNo
	Err error
}

func New(errNo *ErrNo, err error) *Err {
	return &Err{
		ErrNo: ErrNo{
			Code:    errNo.Code,
			Message: errNo.Message,
		},
		Err: err,
	}
}

func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *ErrNo:
		return typed.Code, typed.Message
	default:
	}
	return InternaleServerError.Code, err.Error()
}

func IsErrUSerNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == UserNotFoundError.Code
}
