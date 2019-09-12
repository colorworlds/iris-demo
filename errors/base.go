package errors

// 系统性错误
func SystemError(options ...interface{}) *Error {
	return NewError(1001, "system error", options...)
}

// 参数错误错误
func ParamError(options ...interface{}) *Error {
	return NewError(1002, "param error", options...)
}

// DB错误错误
func DBError(options ...interface{}) *Error {
	return NewError(1003, "db error", options...)
}

// Auth错误错误
func AuthError(options ...interface{}) *Error {
	return NewError(1004, "auth error", options...)
}

// 响应空数据错误错误
func NoDataError(options ...interface{}) *Error {
	return NewError(1005, "no data error", options...)
}
