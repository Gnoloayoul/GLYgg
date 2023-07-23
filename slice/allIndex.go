package slice

// AllIndex
// 返回和 dst 相等的所有元素的下标
func AllIndex[T comparable](src []T, dst T) []int {
	return AllIndexFunc[T](src, dst, func(src, dst T) bool {
		return src == dst
	})
}
