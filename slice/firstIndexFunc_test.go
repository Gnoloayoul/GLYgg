package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFirstIndexFunc(t *testing.T) {
	testCases := []struct {
		name string
		src []int
		dst int
		want int
	} {
		{
			name: "first one",
			src: []int{1, 2, 3},
			dst: 1,
			want: 0,
		},
		{
			name: "empty src",
			src: []int{},
			dst: 1,
			want: -1,
		},
		{
			name: "src nil",
			dst: 1,
			want: -1,
		},
		{
			name: "dst not exist",
			src: []int{1, 2, 3},
			dst: 7,
			want: -1,
		},
		{
			name: "last one",
			src: []int{1, 2, 3},
			dst: 3,
			want: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			equal := func(a, b int) bool {
				return a == b
			}
			res := FirstIndexFunc(tc.src, tc.dst, equal)
			if !reflect.DeepEqual(res, tc.want) {
				t.Errorf("\nInput: %v, %v\nExpected: %v\nGot: %v", tc.src, tc.dst, tc.want, res)
			}
		})
	}
}

func ExampleFirstIndexFunc() {
	src1 := []int{1, 1, 3, 3, 3, 5, 7}
	src2 := 3
	equal := func(a, b int) bool {
		return a == b
	}
	res := FirstIndexFunc(src1, src2, equal)
	fmt.Println(res)
	// output: 2
}