package slice

import (
	"git.acwing.com/Gnoloayoul/gylgg/errs"
)
// Delete
// 泛型
// 能根据下标删除切片中的元素
// 正常情况能返回删除了元素的切片（非原切片）, 被删除的数, 为nil的警告
// 当输入的下标不合法时，会直接返回越界警告，其他为nil
func Delete[T any](a []T, idx int) ([]T, T, error) {
	length := len(a)
	if idx >= length || idx < 0 {
		var zero T
		return nil, zero, errs.NewErrIndexOutOfRange(length, idx)
	}
	switch idx {
	case 0:
		val := a[0]
		res := make([]T, length - 1, length - 1)
		copy(res[:], a[1:])
		return res, val, nil
	case length - 1:
		val := a[idx]
		res := make([]T, length - 1, length - 1)
		copy(res[:], a[:idx])
		return res, val, nil
	default:
		val := a[idx]
		res := make([]T, length - 1, length - 1)
		copy(res[:idx], a[:idx])
		copy(res[idx:], a[idx + 1:])
		return res, val, nil
	}
}