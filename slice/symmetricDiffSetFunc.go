package slice

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// SymmetricDiffSetFunc
// 对称差值
// 【最好优先使用 SymmetricDiffSet 】
func SymmetricDiffSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	interSection := make([]T, 0, min(len(src), len(dst)))
	for _, valSrc := range src {
		for _, valDst := range dst {
			if equal(valSrc, valDst) {
				interSection = append(interSection, valSrc)
				break
			}
		}
	}

	res := make([]T, 0, len(src) + len(dst) - len(interSection) * 2)
	for _, v := range src {
		if !ContainsFunc[T](interSection, v, equal) {
			res = append(res, v)
		}
	}
	for _, v := range dst {
		if !ContainsFunc[T](interSection, v, equal) {
			res = append(res, v)
		}
	}
	return deduplicateFunc[T](res, equal)
}