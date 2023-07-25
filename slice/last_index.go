package slice

// LastIndex
// 返回和 dst 相等的最后一个元素下标
// -1 表示没找到
func LastIndex[T comparable](src []T, dst T) int {
	return LastIndexFunc[T](src, dst, func(src, dst T) bool {
		return src == dst
	})
}