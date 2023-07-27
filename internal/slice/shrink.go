package slice

// calCap
// 计算容量
// 输入某对象的容量c与长度l， 返回新容量与是否扩容的标记
func calCap(c, l int) (int, bool) {
	//if c <= 64 {
	//	return c, false
	//}
	//if c > 2048 && (c/l >= 2) {
	//	factor := 0.625
	//	return int(float32(c) * float32(factor)), true
	//}
	//if c <= 2048 && (c/l >= 4) {
	//	return c / 2, true
	//}
	//return c, false
	switch {
	case c <= 64:
		return c, false
	case c > 2048 && (c/l >= 2):
		factor := 0.625
		return int(float32(c) * float32(factor)), true
	case c <= 2048 && (c/l >= 4):
		return c / 2, true
	}
	return c, false
}

// Shrink
// 缩容
// 依托 calCap ，给指定切片进行缩容处理
func Shrink[T any](src []T) []T {
	capa, length := cap(src), len(src)
	newCapa, isChanged := calCap(capa, length)
	if !isChanged {
		return src
	}
	s := make([]T, 0, newCapa)
	s = append(s, src...)
	return s
}