package slice

func FilterDelete[Src any](src []Src, m func(idx int, src Src) bool) []Src {
	// 记录被删除的元素位置，也就是空缺的位置
	emptyPos := 0
	for idx := range src {
		// 判断是否满足删除条件
		if m(idx, src[idx]) {
			continue
		}
		// 移动元素
		src[emptyPos] = src[idx]
		emptyPos++

	}
	return src[:emptyPos]
}