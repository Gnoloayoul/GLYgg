package slice

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymmetricDiffSetFunc(t *testing.T) {
	testCases := []struct {
		name string
		src []int
		dst []int
		want []int
	} {
		{
			name: "normal test",
			src: []int{1, 2, 3, 4},
			dst: []int{4, 5, 6, 1},
			want: []int{2, 3, 5, 6},
		},
		{
			name: "deduplicate",
			src: []int{1, 1, 2, 3, 4},
			dst: []int{4, 5, 6, 1, 7, 6},
			want: []int{3, 6, 7, 5, 2},
		},
		{
			name: "src empty",
			src: []int{},
			dst: []int{1},
			want: []int{1},
		},
		{
			name: "not exist same ele",
			src: []int{1, 3, 5},
			dst: []int{2, 4},
			want: []int{1, 3, 2, 4, 5},
		},
		{
			name: "both nil",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			equal := func(a, b int) bool {
				return a == b
			}
			res := SymmetricDiffSetFunc(tc.src, tc.dst, equal)
			assert.ElementsMatch(t, tc.want, res)
		})
	}
}

func ExampleSymmetricDiffSetFunc() {
	src := []int{1, 3, 4, 2}
	dst := []int{2, 5, 7, 3}
	equal := func(a, b int) bool {
		return a == b
	}
	res := SymmetricDiffSetFunc(src, dst, equal)
	sort.Ints(res) // 给结果排个序
	fmt.Println(res)
	// output:
	// [1 4 5 7]
}