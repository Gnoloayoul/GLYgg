package slice

import (
	"fmt"
	"github.com/Gnoloayoul/GYLgg"
	"github.com/stretchr/testify/assert"
	"testing"
)

// testMaxTypes
// 测试辅助函数
// 用来测试一下满足 Max 方法约束的所有类型
func testMaxTypes[T gylgg.RealNumber](t *testing.T) {
	res := Max[T]([]T{1, 2, 3})
	assert.Equal(t, T(3), res)
}

func TestMax(t *testing.T) {
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
			want: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Max[Integer](tc.input)
			assert.Equal(t, tc.want, res)
		})
	}

	assert.Panics(t, func() {
		Max[int](nil)
	})
	assert.Panics(t, func() {
		Max[int]([]int{})
	})

	testMaxTypes[uint](t)
	testMaxTypes[uint8](t)
	testMaxTypes[uint16](t)
	testMaxTypes[uint32](t)
	testMaxTypes[uint64](t)
	testMaxTypes[int](t)
	testMaxTypes[int8](t)
	testMaxTypes[int16](t)
	testMaxTypes[int32](t)
	testMaxTypes[int64](t)
	testMaxTypes[float32](t)
	testMaxTypes[float64](t)
}

func ExampleMax() {
	src := []int{1, 2, 3}
	fmt.Println(Max(src))
	// output: 3
}

