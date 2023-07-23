package slice

import (
	"fmt"
	"reflect"
	"testing"

)

func TestDiffSet(t *testing.T) {
	testCases := []struct {
		name string
		src1 []int
		src2 []int
		want []int
	} {
		{
			name: "diff 1",
			src1: []int{1, 2, 4, 6, 7},
			src2: []int{1, 4, 6, 7},
			want: []int{2},
		},
		{
			name: "src1 less than src2",
			src1: []int{1, 2},
			src2: []int{1, 2, 5, 7},
			want: []int{},
		},
		{
			name: "diff deduplicate",
			src1: []int{1, 2, 3, 7, 7},
			src2: []int{1, 2, 3},
			want: []int{7},
		},
		{
			name: "src2 duplicate ele",
			src1: []int{1, 1, 3, 5, 7},
			src2: []int{1, 3, 5, 5},
			want: []int{7},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := DiffSet[int](tc.src1, tc.src2)
			if !reflect.DeepEqual(res, tc.want) {
				t.Errorf("Input: %v, %v\nExpected: %v\nGot: %v", tc.src1, tc.src2, tc.want, res)
			}
		})
	}
}

func ExampleDiffSet() {
	src1 := []int{1, 1, 3, 5, 7}
	src2 := []int{1, 3, 5, 5}
	res := DiffSet(src1, src2)
	fmt.Println(res)
	// output: [7]
}