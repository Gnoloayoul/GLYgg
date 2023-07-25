package slice

import "reflect"

// FindIdx
// 泛型
// 根据输入找出在切片里的位置
// 找不到，就返回-1
func FindIdx[T any](a []T, input T) int {
	length := len(a)
	if length == 0 {
		return -1
	}

	for idx, v := range a {
		if reflect.DeepEqual(v, input) {
			return idx
		}
	}

	return -1
}