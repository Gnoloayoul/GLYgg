package slice

// UnionSet
// 并集 只支持 comparable ，返回切片
// 已去重
// 不保证返回的元素顺序不变
func UnionSet[T comparable](src, dst []T) []T {
	srcMap, dstMap := toMap[T](src), toMap[T](dst)
	for key := range srcMap {
		dstMap[key] = struct{}{}
	}
	var res = make([]T, 0, len(dstMap))
	for key := range dstMap {
		res = append(res, key)
	}
	return res
}