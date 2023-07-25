package slice

// AllIndexFunc
// 返回所有和 dst 相同的元素下标
// 【最好优先使用 AllIndex】
func AllIndexFunc[T any](src []T, dst T, equal equalFunc[T]) []int {
	idxes := make([]int, 0, len(src))
	for k, v := range src {
		if equal(v, dst) {
			idxes = append(idxes, k)
		}
	}
	return idxes
}