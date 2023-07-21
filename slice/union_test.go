package slice

import (
	"reflect"
	"testing"
)

func TestUnion(t *testing.T) {
	testCases := []struct {
		name string
		sliceA []interface{}
		sliceB []interface{}
		wantSlice []interface{}
	} {
		{
			name: "all int",
			sliceA: []interface{}{123, 156, 999},
			sliceB: []interface{}{123, 156, 999},
			wantSlice: []interface{}{123, 156, 999, 123, 156, 999},
		},
		{
			name: "all string",
			sliceA: []interface{}{"aaa", "bbb", "ccc"},
			sliceB: []interface{}{"aaa", "bbb", "ccc"},
			wantSlice: []interface{}{"aaa", "bbb", "ccc", "aaa", "bbb", "ccc"},
		},
		{
			name: "int + string",
			sliceA: []interface{}{123, 156, 999},
			sliceB: []interface{}{"aaa", "bbb", "ccc"},
			wantSlice: []interface{}{123, 156, 999, "aaa", "bbb", "ccc"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Union(tc.sliceA, tc.sliceB)
			if !reflect.DeepEqual(res, tc.wantSlice) {
				t.Errorf("\nInput: %v, %v\nExpected: %v\nGot: %v", tc.sliceA, tc.sliceB, tc.wantSlice, res)
			}
		})
	}
}

func BenchmarkUnion(b *testing.B) {
	sliceA := []interface{}{1, "bb", 3, 4, "ccc", 6, 7, "ddd", "ddd", 9, 9, 10, 10, "bb", 10}
	sliceB := []interface{}{1, "bbn", 43, 24, "vccc", 76, 77, "ffddd", "eddd", 19, 19, 110, 20, "vbb", 10}
	for i := 0; i < b.N; i++ {
		Union(sliceA, sliceB)
	}
}