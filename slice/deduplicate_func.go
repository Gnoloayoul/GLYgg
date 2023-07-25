package slice

// DeduplicateFunc
// 给输入的切片去掉重复元素，返回新的切片
// 需要输入equal的逻辑
func deduplicateFunc[T any](data []T, equal equalFunc[T]) []T {
	var newData = make([]T, 0, len(data))
	for k, v := range data {
		if !ContainsFunc[T](data[k + 1:], v, equal) {
			newData = append(newData, v)
		}
	}
	return newData
}