package slice

import "github.com/Gnoloayoul/GYLgg"

// Sum
// 返回输入的切片的和
// 该方法默认至少传入一个值
// 使用浮点型注意精度问题
func Sum[T gylgg.RealNumber](src []T) T {
	var res T
	for _, n := range src {
		res += n
	}
	return res
}