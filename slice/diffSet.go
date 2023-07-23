package slice

// DiffSet
// 差集
// 切片A - 切片B，切片A的长度要大于切片B的。返回已经去重的切片。
// 返回值的顺序是不确定的
func DiffSet[T comparable](src, dst []T) []T {
	srcMap := toMap[T](src)
	for _, val := range dst {
		delete(srcMap, val)
	}

	res := make([]T, 0, len(srcMap))
	for key := range srcMap {
		res = append(res, key)
	}

	return res
}