package gylgg

// RealNumber
// 实数（接口）
// 最好使用该接口来表达数字
type RealNumber interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

// Number
// 数字（接口）
type Number interface {
	RealNumber | ~complex64 | ~complex128
}