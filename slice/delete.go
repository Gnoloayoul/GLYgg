package slice

import "github.com/Gnoloayoul/GYLgg/internal/slice"

// Delete
// 删除 index 处的元素
func Delete[Src any](src []Src, index int) ([]Src, error) {
	res, _, err := slice.Delete[Src](src, index)
	return res, err
}