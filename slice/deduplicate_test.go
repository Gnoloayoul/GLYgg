package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeduplicate(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		want []int
	}{
		{
			name: "int",
			src:  []int{1, 1, 6, 2, 6},
			want: []int{2, 1, 6},
		},
		{
			name: "not change",
			src:  []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := deduplicate(tc.src)
			if !reflect.DeepEqual(res, tc.want) {
				t.Errorf("\nInput: %v\nExpected: %v\nGot: %v", tc.src, tc.want, res)
			}
		})
	}
}

func ExampleDeduplicate() {
	src := []int{1, 2, 3, 3}
	res := deduplicate(src)
	fmt.Println(res)

	src2 := []string{"a", "b", "b", "c"}
	res2 := deduplicate(src2)
	fmt.Println(res2)
	// output:
	// [1 2 3]
	// [a b c]
}
