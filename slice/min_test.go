package slice

import (
	"fmt"
	"github.com/Gnoloayoul/GYLgg"
	"github.com/stretchr/testify/assert"
	"testing"
)

// testMinTypes
// 测试辅助函数
// 用来测试一下满足 Min 方法约束的所有类型
func testMinTypes[T gylgg.RealNumber](t *testing.T) {
	res := Min[T]([]T{1, 2, 3})
	assert.Equal(t, T(1), res)
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name string
		input []Integer
		want Integer
	} {
		{
			name: "value",
			input: []Integer{1},
			want: 1,
		},
		{
			name: "values",
			input: []Integer{2, 3, 1, 4},
			want: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Min[Integer](tc.input)
			assert.Equal(t, tc.want, res)
		})
	}

	assert.Panics(t, func() {
		Min[int](nil)
	})
	assert.Panics(t, func() {
		Min[int]([]int{})
	})

	testMinTypes[uint](t)
	testMinTypes[uint8](t)
	testMinTypes[uint16](t)
	testMinTypes[uint32](t)
	testMinTypes[uint64](t)
	testMinTypes[int](t)
	testMinTypes[int8](t)
	testMinTypes[int16](t)
	testMinTypes[int32](t)
	testMinTypes[int64](t)
	testMinTypes[float32](t)
	testMinTypes[float64](t)
}

func ExampleMin() {
	src := []int{1, 2, 3}
	fmt.Println(Min(src))
	// output: 1
}

