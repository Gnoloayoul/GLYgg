package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLastIndex(t *testing.T) {
	testCases := []struct {
		name string
		src []int
		dst int
		want int
	} {
		{
			name: "first one",
			src: []int{1, 1, 3, 5},
			dst: 1,
			want: 1,
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
			src: []int{1, 4, 6},
			dst: 7,
			want: -1,
		},
		{
			name: "last one",
			src: []int{1, 4, 6},
			dst: 1,
			want: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := LastIndex(tc.src, tc.dst)
			if !reflect.DeepEqual(res, tc.want) {
				t.Errorf("\nInput: %v, %v\nExpected: %v\nGot: %v", tc.src, tc.dst, tc.want, res)
			}
		})
	}
}

func ExampleLastIndex() {
	src := []int{1, -2, 3, 5, 3}
	dst := 3
	res := LastIndex(src, dst)
	fmt.Println(res)
	// output: 4
}