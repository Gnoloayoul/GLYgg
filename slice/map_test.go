package slice

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSmap(t *testing.T) {
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
			res := Smap(tc.src, func(idx int, src int) string {
				return strconv.Itoa(src)
			})
			if !reflect.DeepEqual(res, tc.want) {
				t.Errorf("\nInput: %v\nExpected: %v\nGot: %v", tc.src, tc.want, res)
			}
		})
	}
}

func BenchmarkSmap(b *testing.B) {
	src := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < b.N; i++ {
		Smap(src, func(idx int, src Src) Dst)
	}
}