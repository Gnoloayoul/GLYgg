package slice

import (
	"github.com/Gnoloayoul/GYLgg/internal/errs"
)
// Delete
// 泛型
// 能根据下标删除切片中的元素。
// 正常情况能返回删除了元素的切片（原切片）, 被删除的数, 为nil的警告。
// 当输入的下标不合法时，会直接返回越界警告，其他为nil。
// 当触发缩容机制时，返回的是新的切片
func Delete[T any](a []T, idx int) ([]T, T, error) {
	length := len(a)
	if idx >= length || idx < 0 {
		var zero T
		return nil, zero, errs.NewErrIndexOutOfRange(length, idx)
	}
	val := a[idx]
	switch idx {
	case 0:
		copy(a[:], a[1:])
		a = a[:length - 1]
	case length:
		a = a[:length - 1]
	default:
		copy(a[idx:], a[idx + 1:])
		a = a[:length - 1]
	}
	return a, val, nil
}