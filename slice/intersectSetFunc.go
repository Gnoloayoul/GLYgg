package slice

// IntersectSetFunc
// 取交集
// 支持任何类型，返回已经去重的切片。
// 【最好优先使用 Intersect】
func IntersectSetFunc[T any](src []T, dst []T, equal equalFunc[T]) []T {
	res := make([]T, 0, len(src))
	for _, valSrc := range src {
		for _, valDst := range dst {
			if equal(valDst, valSrc) {
				res = append(res, valSrc)
				break
			}
		}
	}
	return deduplicateFunc[T](res, equal)
}