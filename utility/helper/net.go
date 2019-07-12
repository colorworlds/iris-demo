package helper

import (
	"bytes"
	"github.com/kataras/iris/context"
	"io/ioutil"
	"strings"
)

// 获取iris的请求头
func RequestHeader(ctx context.Context) string {
	var requestHeader string
	for k, v := range ctx.Request().Header {
		requestHeader += k + "=" + v[0] + ";"
	}
	return requestHeader
}

// 获取iris的请求体
func RequestBody(ctx context.Context) string {
	var requestBody string
	data, err := ioutil.ReadAll(ctx.Request().Body)
	if err == nil {
		requestBody = string(data)
		ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(data))
	}
	return requestBody
}

// 获取iris的get参数
func RequestQueries(ctx context.Context) string {
	var requestQuery string
	for k, v := range ctx.URLParams() {
		requestQuery += k + "=" + v + "&"
	}
	requestQuery = strings.Trim(requestQuery, "&")

	return requestQuery
}