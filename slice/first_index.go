package slice

// FirstIndex
// 返回和 dst 相同的第一个元素下标。
// -1 表示没找到。
func FirstIndex[T comparable](src []T, dst T) int {
	return FirstIndexFunc[T](src, dst, func(src, dst T) bool {
		return src == dst
	})
}
