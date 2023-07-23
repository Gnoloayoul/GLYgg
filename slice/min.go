package slice

import "git.acwing.com/Gnoloayoul/gylgg"

// Min
// 返回最大值
// 该方法默认至少传入一个值
// 使用浮点型注意精度问题
func Min[T gylgg.RealNumber](src []T) T {
	res := src[0]
	for i := 1; i < len(src); i++ {
		if src[i] < res {
			res = src[i]
		}
	}
	return res
}