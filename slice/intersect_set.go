package slice

// IntersectSet
// 取交集
// 只支持 Comparable 类型，返回已经去重的切片。
func IntersectSet[T comparable](src []T, dst []T) []T {
	srcMap := toMap(src)
	res := make([]T, 0, len(src))
	// 交集长度小于两个集合任意一个长度
	for _, valDst := range dst {
		if _, exist := srcMap[valDst]; exist {
			res = append(res, valDst)
		}
	}

	return deduplicate[T](res)
}