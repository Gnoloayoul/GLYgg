package slice

import (
	"git.acwing.com/Gnoloayoul/gylgg/errs"
)

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