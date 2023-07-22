package slice

// FilterSlice
// 按照输入条件过滤切片，返回过滤后的切片（新切片）
// 即使有不符合条件的元素，后续遍历仍会继续
func FilterSlice[Src, Dst any](src []Src, m func(idx int, src Src) (Dst, bool)) []Dst {
	res := make([]Dst, 0, len(src))
	for i, s := range src {
		dst, ok := m(i, s)
		if ok {
			res = append(res, dst)
		}
	}
	return res
}