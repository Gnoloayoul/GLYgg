package slice

// LastIndexFunc
// 返回和 dst 相等的一个元素下标
// -1 表示没找到
// 【最好优先使用 LastIndex 】
func LastIndexFunc[T any](src []T, dst T, equal equalFunc[T]) int {
	for i := len(src) - 1; i >= 0; i-- {
		if equal(dst, src[i]) {
			return i
		}
	}
	return -1
}