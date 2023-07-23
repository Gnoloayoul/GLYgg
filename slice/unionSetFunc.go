package slice

// UnionSetFunc
// 并集，泛型支持，返回的是已经去重的切片
// 不保证顺序
// 【最好优先使用 UnionSet】
func UnionSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	var res = make([]T, 0, len(src) + len(dst))
	res = append(res, dst...)
	res = append(res, src...)
	return deduplicateFunc[T](res, equal)
}