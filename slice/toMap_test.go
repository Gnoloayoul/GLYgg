package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToMap(t *testing.T) {
	testCases := []struct {
		name string
		src []interface{}
		want map[interface{}]struct{}
	} {
		{
			name: "int",
			src: []interface{}{123, 156, 999},
			want: map[interface{}]struct{}{123:{}, 156:{}, 999:{}},
		},
		{
			name: "string",
			src: []interface{}{"a", "b", "c"},
			want: map[interface{}]struct{}{"a":{}, "b":{}, "c":{}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := ToMap(tc.src)
			if !reflect.DeepEqual(res, tc.want) {
				t.Errorf("\nInput: %v\nExpected: %v\nGot: %v", tc.src, tc.want, res)
			}
		})
	}
}

func ExampleToMap() {
	src := []string{"a", "b", "c"}
	fmt.Println(ToMap(src))
	// output: map[a:{} b:{} c:{}]
}

func BenchmarkToMap1(b *testing.B) {
	src := []int{11, 3, 444, 56, 67, 8, 777, 89, 342, 123, 99}
	for i := 0; i < b.N; i++ {
		ToMap(src)
	}
}

func BenchmarkToMap2(b *testing.B) {
	src := []string{"abs", "asdfasd", "cccc", "ddd", "ffff", "asdfa", "gsrg", "dtgsdfg", "dgsedgsrty"}
	for i := 0; i < b.N; i++ {
		ToMap(src)
	}
}