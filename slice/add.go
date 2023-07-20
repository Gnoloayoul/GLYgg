package slice

// Add
// 泛型
// 在切片里，在输入index的前面插入input
// 当index小于0，则认为是往该切片的头插入input
// 当index远大于切片，则认为是往该切片的尾插入input
// 当触发扩容时，返回的是新切片，否则就是原切片
func Add[T any](a []T, index int, input T) []T {
	length := len(a)
	capt := cap(a)
	switch {
	case index < 0:
		index = 0
	case index > length:
		index = length
	}

	if length + 1 >= capt {
		tmp := make([]T, 0, length * 2)
		tmp = append(tmp, a[:index]...)
		tmp = append(tmp, input)
		tmp = append(tmp, a[index:]...)
		return tmp
	}

	var zero T
	a = append(a, zero)
	copy(a[index + 1:], a[index:])
	a[index] = input
	return a
}