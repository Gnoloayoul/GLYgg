package slice

// DiffSetFunc
// 计算两个切片的差集,返回去重的切片
// 第一个切片必须长于第二个切片
// 【最好优先使用 DiffSet 】
func DiffSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	res := make([]T, 0, len(src))
	for _, val := range src {
		if !ContainsFunc[T](dst, val, equal) {
			res = append(res, val)
		}
	}
	return deduplicateFunc[T](res, equal)
}