package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestContainsAll(t *testing.T) {
	testCases := []struct {
		name string
		src1 []int
		src2 []int
		want bool
	} {
		{
			name: "src1 exist one not in src2",
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
			name: "src1 nil src2 empty",
			src1: nil,
			src2: []int{},
			want: true,
		},
		{
			name: "src1 and src2 nil",
			src1: nil,
			src2: nil,
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsAll(tc.src1, tc.src2)
			if !reflect.DeepEqual(res, tc.want) {
				t.Errorf("\nInput: %v, %v\nExpected: %v\nGot: %v", tc.src1, tc.src2, tc.want, res)
			}
		})
	}
}

func ExampleContainsAll() {
	src1 := []int{1, 2, 3}
	src2 := []int{3, 1}
	res1 := ContainsAll(src1, src2)
	fmt.Println(res1)

	src3 := []int{1, 2, 3}
	src4 := []int{3, 1, 4}
	res2 := ContainsAll(src3, src4)
	fmt.Println(res2)
	// output:
	// true
	// false
}