package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersectSet(t *testing.T) {
	testCases := []struct {
		name string
		src []int
		dst []int
		want []int
	} {
		{
			name: "normal test",
			src: []int{1, 3, 5, 7},
			dst: []int{1, 3, 5},
			want: []int{1, 3, 5},
		},
		{
			name: "src empty",
			src: []int{},
			dst: []int{1, 3, 5, 7},
			want: []int{},
		},
		{
			name: "src nil",
			dst: []int{1, 3, 5},
			want: []int{},
		},
		{
			name: "exist the same ele in src",
			src: []int{1, 3, 5, 5},
			dst: []int{1, 3, 5},
			want: []int{1, 3, 5},
		},
		{
			name: "dst empty",
			src: []int{1, 3, 5, 5},
			dst: []int{},
			want: []int{},
		},
		{
			name: "dst nil",
			src: []int{1, 3, 5, 5},
			want: []int{},
		},
		{
			name: "exist the same ele in src and dst",
			src: []int{1, 1, 3, 5, 7},
			dst: []int{1, 3, 5, 5},
			want: []int{1, 3, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := IntersectSet(tc.src, tc.dst)
			assert.ElementsMatch(t, tc.want, res)
		})
	}
}

func ExampleIntersectSet() {
	src1 := []int{1, 1, 3, 5, 7}
	src2 := []int{1, 3, 5, 5}
	res1 := IntersectSet(src1, src2)
	fmt.Println(res1)
	// output: [1, 3, 5]
}