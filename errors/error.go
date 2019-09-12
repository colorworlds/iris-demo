package errors

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

func NewError(code int, msg string, options ...interface{}) *Error {
	e := &Error{Code: code, Msg: msg}

	// 这里委婉的提示了出错信息的文件和行数，方便定位问题
	// 而且可以根据tip的md5值，全局查找日志文件中的请求日志
	if _, file, line, ok := runtime.Caller(2); ok {
		fileAndLine := strings.ReplaceAll(filepath.Base(file), ".go", strconv.Itoa(line))
		e.Tip = helper.MD5(fmt.Sprint(time.Now().UnixNano())) + " -- " + fileAndLine
	}

	if len(options) > 0 {
		// 如果是错误信息，则记录在日志文件中
		var optionTip string
		for _, option := range options {
			if err, ok := option.(error); ok {
				logrus.WithFields(logrus.Fields{"errMsg": err}).Error(e.Error())
			} else {
				optionTip += fmt.Sprint(option)
			}
		}

		// 否则，将他们拼接在tip的提示中
		if optionTip != "" {
			e.Tip += " -- " + optionTip
		}
	}

	return e
}
