package errs

import (
	"IRIS_WEB/utility/helper"
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Tip  string `json:"tip,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, msg: %s, tip: %s", e.Code, e.Msg, e.Tip)
}

func NewError(code int, msg string, err ...error) *Error {
	e := &Error{Code: code, Msg: msg}

	if _, file, line, ok := runtime.Caller(2); ok {
		e.Tip = strings.ReplaceAll(filepath.Base(file), ".go", strconv.Itoa(line))
		e.Tip += helper.MD5(fmt.Sprint(time.Now().UnixNano()))
	}

	if len(err) > 0 {
		logrus.Errorf("%v, err: %v", e, strings.ReplaceAll(err[0].Error(), "\n", " "))
	}

	return e
}
