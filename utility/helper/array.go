package helper

func StrArrContains(arr []string, e string) bool {
	for _, str := range arr {
		if str == e {
			return true
		}
	}
	return false
}

func IntArrContains(arr []int, e int) bool {
	for _, n := range arr {
		if n == e {
			return true
		}
	}
	return false
}

