package slice

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestFilterSlice(t *testing.T) {
	testCases := []struct {
		name string
		src []int
		want []string
	} {
		{
			name: "src nil",
			want: []string{},
		},
		{
			name: "src empty",
			src: []int{},
			want: []string{},
		},
		{
			name: "src has element",
			src: []int{1, -2, 3},
			want: []string{"1", "3"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := FilterSlice(tc.src, func(idx, src int) (string, bool) {
				return strconv.Itoa(src), src >= 0
			})
			if !reflect.DeepEqual(res, tc.want) {
				t.Errorf("\nInput: %v\nExpected: %v\nGot: %v", tc.src, tc.want, res)
			}
		})
	}
}

func ExampleFilterSlice() {
	src := []int{1, -2, 3}
	dst := FilterSlice(src, func(idx, src int) (string, bool) {
		return strconv.Itoa(src), src >= 0
	})
	fmt.Println(dst)
	// output: [1, 3]
}