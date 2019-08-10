package helper

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"io/ioutil"
	"net/http"
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

func Post(url string, header map[string]string, data iris.Map) ([]byte, error) {
	var err error
	var req *http.Request
	var resp *http.Response
	var client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression: true,
		},
	}

	if b, err := json.Marshal(data); err != nil {
		return nil, err
	} else {
		if req, err = http.NewRequest("POST", url, bytes.NewReader(b)); err != nil {
			return nil , err
		}
	}

	if len(header) != 0 {
		for key, value := range header {
			req.Header.Add(key, value)
		}
	}

	if resp, err = client.Do(req); err != nil {
		return nil , err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%d", resp.StatusCode))
	}

	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	} else {
		return body, nil
	}
}

func Get(url string, header map[string]string, data iris.Map) ([]byte, error) {
	var err error
	var req *http.Request
	var resp *http.Response
	var client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression: true,
		},
	}

	if len(data) > 0 {
		url += "?"
		for k, v := range data {
			url += fmt.Sprintf("%v=%v&", k, v)
		}
		url = strings.TrimRight(url, "&")
	}

	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return nil , err
	}

	if len(header) != 0 {
		for key, value := range header {
			req.Header.Add(key, value)
		}
	}

	if resp, err = client.Do(req); err != nil {
		return nil , err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%d", resp.StatusCode))
	}

	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	} else {
		return body, nil
	}
}
