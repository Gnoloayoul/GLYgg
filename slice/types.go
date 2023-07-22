package slice

// equalFunc
// 比较两个元素是否相等
type equalFunc[T any] func(src, dst T) bool