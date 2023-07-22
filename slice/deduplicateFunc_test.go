package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeduplicateFunc(t *testing.T) {
	testCases := []struct {
		name string
		src []int
		want []int
	} {
		{
			name: "int",
			src: []int{1, 1, 6, 2, 6},
			want: []int{1, 2, 6},
		},
		{
			name: "not change",
			src: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			equal := func(a, b int) bool {
				return a == b
			}
			res := DeduplicateFunc(tc.src, equal)
			if !reflect.DeepEqual(res, tc.want) {
				t.Errorf("\nInput: %v\nExpected: %v\nGot: %v", tc.src, tc.want, res)
			}
		})
	}
}

func ExampleDeduplicateFunc() {
	src1 := []int{1, 2, 3, 3}
	equal := func(a, b int) bool {
		return a == b
	}
	res1 := DeduplicateFunc(src1, equal)
	fmt.Println(res1)
	// output: [1 2 3]
}