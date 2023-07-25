package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestContainsAny(t *testing.T) {
	testCases := []struct {
		name string
		src1 []int
		src2 []int
		want bool
	} {
		{
			name: "exist two ele",
			src1: []int{1, 4, 6, 2, 6},
			src2: []int{1, 6},
			want: true,
		},
		{
			name: "not exist the same",
			src1: []int{1, 4, 6, 2, 6},
			src2: []int{7, 0},
			want: false,
		},
		{
			name: "exist two same ele",
			src1: []int{1, 1, 6, 2, 6},
			src2: []int{1, 1},
			want: true,
		},
		{
			name: "length of src is 0",
			src1: []int{},
			src2: []int{},
			want: false,
		},
		{
			name: "src nil",
			src2: []int{1},
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsAny(tc.src1, tc.src2)
			if !reflect.DeepEqual(res, tc.want) {
				t.Errorf("\nInput: %v, %v\nExpected: %v\nGot: %v", tc.src1, tc.src2, tc.want, res)
			}
		})
	}
}

func ExampleContainsAny() {
	src1 := []int{1, 2, 3}
	src2 := []int{3, 1}
	res1 := ContainsAny(src1, src2)
	fmt.Println(res1)

	src3 := []int{1, 2, 3}
	src4 := []int{4, 7, 6}
	res2 := ContainsAny(src3, src4)
	fmt.Println(res2)
	// output:
	// true
	// false
}