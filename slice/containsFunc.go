package slice

// ContainsFunc
// 判断 src 里面是否存在 dst
// 【最好优先使用Contains】
func ContainsFunc[T any](src []T, dst T, equal equalFunc[T]) bool {
	// 遍历调用equal函数进行判断
	for _, v := range src {
		if equal(v, dst) {
			return true
		}
	}
	return false
}