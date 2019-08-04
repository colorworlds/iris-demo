package helper

import (
	"encoding/json"
	"github.com/kataras/iris"
)

// 将struct转换为map
func StToMap(s interface{}) iris.Map {
	var v = make(iris.Map)

	j, _ := json.Marshal(s)

	json.Unmarshal(j, &v)

	return v
}
