package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAllIndexFunc(t *testing.T) {
	testCases := []struct {
		name string
		src []int
		dst int
		want []int
	} {
		{
			name: "normal test",
			src: []int{1, 1, 3, 5},
			dst: 1,
			want: []int{0, 1},
		},
		{
			name: "src empty",
			src: []int{},
			dst: 1,
			want: []int{},
		},
		{
			name: "dst not exist",
			src: []int{1, 1, 3, 5},
			dst: 7,
			want: []int{},
		},
		{
			name: "normal test2",
			src: []int{0, 1, 1, 3, 5, 0},
			dst: 0,
			want: []int{0, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			equal := func(a, b int) bool {
				return a == b
			}
			res := AllIndexFunc(tc.src, tc.dst, equal)
			if !reflect.DeepEqual(res, tc.want) {
				t.Errorf("\nInput: %v, %v\nExpected: %v\nGot: %v", tc.src, tc.dst, tc.want, res)
			}
		})
	}
}

func ExampleAllIndexFunc() {
	src := []int{1, -2, 3, 5, 3}
	dst := 3
	equal := func(a, b int) bool {
		return a == b
	}
	res := AllIndexFunc(src, dst, equal)
	fmt.Println(res)
	// output: [2, 4]
}