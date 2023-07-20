package slice

// Add
// 泛型
// 在切片里，在输入index的前面插入input
// 当index小于0，则认为是往该切片的头插入input
// 当index远大于切片，则认为是往该切片的尾插入input
// 返回的是新切片
func Add[T any](a []T, index int, input T) []T {
	length := len(a)
	ans := []T{}
	switch {
	case index <= 0:
		ans = append(ans, input)
		ans = append(ans, a...)
		return ans
	case index >= length:
		ans = append(ans, a...)
		ans = append(ans, input)
		return ans
	default:
		ans = append(ans, a[:index]...)
		ans = append(ans, input)
		ans = append(ans, a[index:]...)
		return ans
	}
}