package slice

// Contains
// 判断 src 里面是否存在 dst
func Contains[T comparable](src []T, dst T) bool {
	return ContainsFunc[T](src, dst, func(src, dst T) bool {
		return src == dst
	})
}