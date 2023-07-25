package slice

// FirstIndexFunc
// 返回和 dst 相等的第一个元素下标。
// 返回 -1 表示没找到。
// 【最好优先使用 FirstIndex 】
func FirstIndexFunc[T any](src []T, dst T, equal equalFunc[T]) int {
	for k, v := range src {
		if equal(v, dst) {
			return k
		}
	}
	return -1
}