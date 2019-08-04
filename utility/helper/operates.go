package helper

import "fmt"

// 三元运算符
func IF(condition bool, val1, val2 interface{}) interface{} {
	if condition {
		return val1
	} else {
		return val2
	}
}

// 两元运算符
func OR(val1, val2 interface{}) interface{} {
	v := fmt.Sprintf("%v", val1)

	if v == "" || v == "0" || v == "<nil>" {
		return val2
	} else {
		return val1
	}
}