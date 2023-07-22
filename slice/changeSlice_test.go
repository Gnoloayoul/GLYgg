package slice

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestChangeSlice(t *testing.T) {
	testCases := []struct {
		name string
		src []int
		want []string
		wantSlice []interface{}
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
			src: []int{1, 2, 3},
			want: []string{"1", "2", "3"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ChangeSlice(tc.src, func(idx int, src int) string {
				return strconv.Itoa(src)
			})
			if !reflect.DeepEqual(res, tc.want) {
				t.Errorf("\nInput: %v\nExpected: %v\nGot: %v", tc.src, tc.want, res)
			}
		})
	}
}

func ExampleChangeSlice() {
	src := []int{1, 2, 3}
	dst := slice.ChangeSlice(src, func(idx, src int) string {
		return strconv.Itoa(src)
	})
	fmt.Println(dst)
	// output: [1, 2, 3]
}

// 还需改进，目前测不出速度
func BenchmarkSmap(b *testing.B) {
	src := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < b.N; i++ {
		dst := ChangeSlice(src, func(idx int, src int) string {
			return strconv.Itoa(src)
		})
		println(dst)
	}
}