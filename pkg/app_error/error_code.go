package app_error

/*
	正确返回的直接返回 0
	出现错误的，错误码为5位
	第一位为服务级别错误类型：1 为系统级错误；2 为普通错误，通常是由用户非法操作引起的
	二三位为服务模块：一个大型系统的服务模块通常不超过两位数，如果超过，说明这个系统该拆分了
	四五位为错误码：只用两位错误码以防止一个模块定制过多的错误码，后期不好维护
	code = 0 说明是正确返回，code > 0 说明是错误返回
	错误通常包括系统级错误码和服务级错误码
	建议代码中按服务模块将错误分类
	错误码均为 >= 0 的数
	在 apiserver 中 HTTP Code 固定为 http.StatusOK，错误码通过 code 来表示。
*/

var (
	// Common errors
	OK                  = &AppError{Code: 0, Message: "OK"}
	InternalServerError = &AppError{Code: 10001, Message: "Internal server error"}
	ErrBind             = &AppError{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation = &AppError{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &AppError{Code: 20002, Message: "Database error."}
	ErrToken      = &AppError{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrEncrypt           = &AppError{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &AppError{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &AppError{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &AppError{Code: 20104, Message: "The password was incorrect."}
)
