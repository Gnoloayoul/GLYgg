package errs

import "fmt"

// NewErrIndexOutOfRange
// 下标越界警告
func NewErrIndexOutOfRange(length, index int) error {
	return fmt.Errorf("GYLgg: 下标越界, 长度 %d, 当前下标 %d\n", length, index)
}