package slice

// ContainsAnyFunc
// 判断 src 里面是否存在 dst 中的任何一个元素
// 【最好优先使用ContainsAny】
func ContainsAnyFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	for _, valDst := range dst {
		for _, valSrc := range src {
			if equal(valSrc, valDst) {
				return true
			}
		}
	}
	return false
}