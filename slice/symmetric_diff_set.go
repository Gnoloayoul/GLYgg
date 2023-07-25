package slice

// SymmetricDiffSet
// 对称差集
// 已去重，返回的元素顺序不一定
func SymmetricDiffSet[T comparable](src, dst []T) []T {
	srcMap, dstMap := toMap[T](src), toMap[T](dst)
	for dstKey := range dstMap {
		if _, exist := srcMap[dstKey]; exist {
			// 删除了共同元素，两者剩余的并集就是对称差
			delete(dstMap, dstKey)
			delete(srcMap, dstKey)
		}
	}

	for k, v := range dstMap {
		srcMap[k] = v
	}
	res := make([]T, 0, len(srcMap))
	for k := range srcMap {
		res = append(res, k)
	}
	return res
}