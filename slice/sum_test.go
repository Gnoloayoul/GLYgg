package slice

import (
	"fmt"
	"git.acwing.com/Gnoloayoul/gylgg"
	"github.com/stretchr/testify/assert"
	"testing"
)

// testSumTypes
// 测试辅助函数
// 用来测试一下满足 Sum 方法约束的所有类型
func testSumTypes[T gylgg.RealNumber](t *testing.T) {
	res := Sum[T]([]T{1, 2, 3})
	assert.Equal(t, T(6), res)
}

func TestSum(t *testing.T) {
	testCases := []struct {
		name string
		input []Integer
		want Integer
	} {
		{
			name: "nil",
		},
		{
			name: "empty",
			input: []Integer{},
		},
		{
			name: "value",
			input: []Integer{1},
			want: 1,
		},
		{
			name: "values",
			input: []Integer{1, 2, 3},
			want: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Sum[Integer](tc.input)
			assert.Equal(t, tc.want, res)
		})
	}

	testSumTypes[uint](t)
	testSumTypes[uint8](t)
	testSumTypes[uint16](t)
	testSumTypes[uint32](t)
	testSumTypes[uint64](t)
	testSumTypes[int](t)
	testSumTypes[int8](t)
	testSumTypes[int16](t)
	testSumTypes[int32](t)
	testSumTypes[int64](t)
	testSumTypes[float32](t)
	testSumTypes[float64](t)
}

func ExampleSum() {
	src := []int{1, 2, 3}
	fmt.Println(Sum(src))
	// output: 6
}

