package list

import (
	"github.com/Gnoloayoul/GYLgg/internal/errs"
	"github.com/Gnoloayoul/GYLgg/internal/slice"
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

// NewArrayListof
// 将输入的切片ts转成arrayLIst
// 直接使用，没有复制
func NewArrayListOf[T any](ts []T) *ArrayList[T] {
	return &ArrayList[T]{
		vals: ts,
	}
}

// Get
// 获取名为a的ArrayList，在下标index的元素
func (a *ArrayList[T]) Get(index int) (t T, e error) {
	length := len(a.vals)
	if index >= length || index < 0 {
		return  t, errs.NewErrIndexOutOfRange(length, index)
	}
	return a.vals[index], e
}
// Append
// 往ArrayList的尾部追加数据
func (a *ArrayList[T]) Append(ts ...T) error {
	a.vals = append(a.vals, ts...)
	return nil
}

// Add
// 在ArrayList下标为index的位置插入一个元素
// 当index等于ArrayList长度，就是append
func (a *ArrayList[T]) Add(index int, t T) error {
	length := len(a.vals)
	if index > length || index < 0 {
		return  errs.NewErrIndexOutOfRange(length, index)
	}

	a.vals = append(a.vals, t)
	copy(a.vals[index + 1:], a.vals[index:])
	a.vals[index] = t
	return nil
}

// Set
// 将ArrayList中index位置的元素设置成t
func (a *ArrayList[T]) Set(index int, t T) error {
	length := len(a.vals)
	if index >= length || index < 0 {
		return  errs.NewErrIndexOutOfRange(length, index)
	}
	a.vals[index] = t
	return nil
}

// Delete
// 删除ArrayList中index位置的元素
// 会按照规则自动缩容
func (a *ArrayList[T]) Delete(index int) (T, error) {
	res, t, err := slice.Delete(a.vals, index)
	if err != nil {
		return t, err
	}
	a.vals = res
	a.shrink()
	return t, nil
}

func (a *ArrayList[T]) shrink() {
	a.vals = slice.Shrink(a.vals)
}

func (a *ArrayList[T]) Len() int {
	return len(a.vals)
}

func (a *ArrayList[T]) Cap() int {
	return cap(a.vals)
}

func (a *ArrayList[T]) Range(fn func(index int, t T) error) error {
	for key, value := range a.vals {
		e := fn(key, value)
		if e != nil {
			return e
		}
	}
	return nil
}

func (a *ArrayList[T]) ChangeSlice() []T {
	res := make([]T, len(a.vals))
	copy(res, a.vals)
	return res
}








