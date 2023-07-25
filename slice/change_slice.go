package slice

// ChangeSlice
// 其实这是切片变换框架
// 怎么换就看你怎么写
func ChangeSlice[Src, Dst any](src []Src, m func(idx int, src Src) Dst) []Dst {
	dst := make([]Dst, len(src))
	for i, s := range src {
		dst[i] = m(i, s)
	}
	return dst
}
