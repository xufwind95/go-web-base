package app_error

import "fmt"

type AppError struct {
	Code    int
	Message string
}

// 实现Error，使之成为 系统定义的 error类型
func (err AppError) Error() string {
	return err.Message
}

// 添加额外信息
func (err *AppError) AddMsg(msg string) error {
	err.Message = fmt.Sprintf("%s %s", err.Message, msg)
	return err
}

func DecodeError(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch errType := err.(type) {
	case *AppError:
		return errType.Code, errType.Message
	default:
	}

	return InternalServerError.Code, InternalServerError.Message
}
