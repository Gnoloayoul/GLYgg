package slice

// Deduplicate
// 给输入的切片去掉重复元素，返回新的切片
// 与Deduplicate不同的是，不需要输入equal的逻辑。而是在内部使用Map的功能
// 要求输入的是comparable
// 返回到新切片是不保证顺序的
func deduplicate[T comparable](data []T) []T {
	dataMap := toMap[T](data)
	var newData = make([]T, 0, len(dataMap))
	for key := range dataMap {
		newData = append(newData, key)
	}
	return newData
}