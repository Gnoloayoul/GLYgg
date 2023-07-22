package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestContainsAllFunc(t *testing.T) {
	testCases := []struct {
		name string
		src1 []int
		src2 []int
		want bool
	} {
		{
			name: "src1 exist ont not in src2",
			src1: []int{1, 4, 6, 2, 6},
			src2: []int{1, 4, 6, 2},
			want: true,
		},
		{
			name: "src1 not include the whole ele",
			src1: []int{1, 4, 6, 2, 6},
			src2: []int{1, 4, 6, 2, 6, 7},
			want: false,
		},
		{
			name: "length of src1 is 0",
			src1: []int{},
			src2: []int{1},
			want: false,
		},
		{
			name: "src1 nil dst empty",
			src1: nil,
			src2: []int{},
			want: true,
		},
		{
			name: "src1 and src2 nil",
			src1: nil,
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsAllFunc(tc.src1, tc.src2, func(src, dst int) bool {
				return src == dst
			})
			if !reflect.DeepEqual(res, tc.want) {
				t.Errorf("\nInput: %v, %v\nExpected: %v\nGot: %v", tc.src1, tc.src2, tc.want, res)
			}
		})
	}
}

func ExampleContainsAllFunc() {
	src1 := []int{1, 2, 3}
	src2 := []int{3, 1}
	res1 := ContainsAllFunc(src1, src2, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res1)

	src3 := []int{1, 2, 3}
	src4 := []int{3, 1, 4}
	res2 := ContainsAllFunc(src3, src4, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(res2)
	// output:
	// true
	// false
}