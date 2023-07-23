package slice

// Union
// 求切片a和切片b的并集
// 不去重
func Union[T any](a, b []T) []T {
	a = append(a, b...)
	return a
}