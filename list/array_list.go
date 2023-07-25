package list

import (
	"git.acwing.com/Gnoloayoul/gylgg/internal/errs"
	"git.acwing.com/Gnoloayoul/gylgg/internal/slice"
)

var (
	_ List[any] = &ArrayList[any]{}
)

// ArrayList
// 切片封装
type ArrayList[T any] struct {
	vals []T
}

// NewArrayList
// 初始化出长度为0，容量为cap的ArrayList
func NewArrayList[T any](cap int) *ArrayList[T] {
	return &ArrayList[T]{vals: make([]T, 0, cap)}
}


func NewArray