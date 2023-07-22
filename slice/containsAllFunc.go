package slice

// ContainsAllFunc
// 判断 src 里面是否存在 dst 中的所有元素
// 【最好优先使用 ContainsAll】
func ContainsAllFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	for _, valDst := range dst {
		if !ContainsFunc[T](src, valDst, equal) {
			return false
		}
	}
	return true
}