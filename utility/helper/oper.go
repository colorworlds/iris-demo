package helper

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
	if IsEmpty(val1){
		return val2
	} else {
		return val1
	}
}