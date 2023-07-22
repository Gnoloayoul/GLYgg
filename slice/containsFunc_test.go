package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestContainsFunc(t *testing.T) {
	testCases := []struct {
		name string
		src []int
		target int
		want bool
	} {
		{
			name: "target exist",
			src: []int{1, 4, 6, 2, 6},
			target: 4,
			want: true,
		},
		{
			name: "target not exist",
			src: []int{1, 4, 6, 2, 6},
			target: 3,
			want: false,
		},
		{
			name: "empty src",
			src: []int{},
			target: 3,
			want: false,
		},
		{
			name: "src nil",
			target: 3,
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ContainsFunc(tc.src, tc.target, func(src, dst int) bool {
				return src == dst
			})
			if !reflect.DeepEqual(res, tc.want) {
				t.Errorf("\nInput: %v, %v\nExpected: %v\nGot: %v", tc.src, tc.target, tc.want, res)
			}
		})
	}
}

func ExampleContainsFunc() {
	src := []int{1, 4, 6, 2, 6}
	target := 4
	check := ContainsFunc(src, target, func(src, dst int) bool {
		return src == dst
	})
	fmt.Println(check)
	// output: true
}