package slice

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionSet(t *testing.T) {
	testCases := []struct {
		name string
		src1 []int
		src2 []int
		want []int
	} {
		{
			name: "not empty",
			src1: []int{1, 2, 3},
			src2: []int{4, 5, 6, 1},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "src1 empty",
			src1: []int{},
			src2: []int{4, 5, 6, 1},
			want: []int{4, 5, 6, 1},
		},
		{
			name: "src2 is empty",
			src1: []int{1, 2, 3},
			src2: []int{},
			want: []int{1, 2, 3},
		},
		{
			name: "all empty",
			src1: []int{},
			src2: []int{},
			want: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := UnionSet[int](tc.src1, tc.src2)
			assert.ElementsMatch(t, tc.want, res)
		})
	}
}

func ExampleUnionSet() {
	src1 := []int{1, 3, 4, 5}
	src2 := []int{1, 4, 7}
	res1 := UnionSet(src1, src2)
	sort.Ints(res1) // 给结果排下序
	fmt.Println(res1)
	// output: [1 3 4 5 7]
}