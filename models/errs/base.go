package errs

import "github.com/kataras/iris"

// 没有错误
func NoError(data... interface{}) iris.Map {
	return iris.Map{"code": 1000, "data": data}
}

// 系统性错误
func SystemError(err ...error) *Error {
	return NewError(1001, "system error", err...)
}

// 参数错误错误
func ParamError(err ...error) *Error {
	return NewError(1002, "param error", err...)
}

// DB错误错误
func DBError(err ...error) *Error {
	return NewError(1003, "db error", err...)
}

// Auth错误错误
func AuthError(err ...error) *Error {
	return NewError(1004, "auth error", err...)
}

// 响应空数据错误错误
func NoDataError(err ...error) *Error {
	return NewError(1005, "no data error", err...)
}
