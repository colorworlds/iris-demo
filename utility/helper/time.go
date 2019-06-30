package helper

import (
	"time"
)

// 全局时区
var location *time.Location

// 标准时间模板
const stdTmpl = "2006-01-02 15:04:05"

func init() {
	// 默认的东八区
	location = time.FixedZone("CST", 8 * 60 * 60)
}

// 格式化时间，时间 => 当地的时间
func Date(format string, tm time.Time) string {
	return tm.In(location).Format(format)
}

// 解析时间，当地时间 => 时间
func StrToTime(format string, strTime string) (time.Time, error) {
	return time.ParseInLocation(format, strTime, location)
}

// 获取当前的时间戳
func Time() int64 {
	return time.Now().Unix()
}

// 获取当前的时间字符串
func Now() string {
	return time.Now().In(location).Format(stdTmpl)
}
