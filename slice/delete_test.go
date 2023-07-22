package slice

import (
	"fmt"
	"git.acwing.com/Gnoloayoul/gylgg/errs"
	"reflect"
	"testing"

)

func TestDelete(t *testing.T) {
	testCases := []struct {
		name string
		slice []interface{}
		index int
		wantSlice []interface{}
		wantVal interface{}
		wantErr error
	} {
		{
			name: "int index 0",
			slice: []interface{}{123, 156, 999},
			index: 0,
			wantSlice: []interface{}{156, 999},
			wantVal: 123,
		},
		{
			name: "string index 0",
			slice: []interface{}{"aaa", "bbb", "ccc"},
			index: 0,
			wantSlice: []interface{}{"bbb", "ccc"},
			wantVal: "aaa",
		},
		{
			name: "int index middle",
			slice: []interface{}{123, 156, 999, 123456},
			index: 2,
			wantSlice: []interface{}{123, 156, 123456},
			wantVal: 999,
		},
		{
			name: "string index middle",
			slice: []interface{}{"aaa", "bbb", "ccc", "ddd"},
			index: 1,
			wantSlice: []interface{}{"aaa", "ccc", "ddd"},
			wantVal: "bbb",
		},
		{
			name: "int index out of range",
			slice: []interface{}{123, 156, 999, 123456},
			index: 12,
			wantErr: errs.NewErrIndexOutOfRange(4, 12),
		},
		{
			name: "string index out of range",
			slice: []interface{}{"aaa", "bbb", "ccc", "ddd"},
			index: 12,
			wantErr: errs.NewErrIndexOutOfRange(4, 12),
		},
		{
			name: "int index less than 0",
			slice: []interface{}{123, 156, 999, 123456},
			index: -1,
			wantErr: errs.NewErrIndexOutOfRange(4, -1),
		},
		{
			name: "string index less than 0",
			slice: []interface{}{"aaa", "bbb", "ccc", "ddd"},
			index: -1,
			wantErr: errs.NewErrIndexOutOfRange(4, -1),
		},
		{
			name: "int index last",
			slice: []interface{}{123, 156, 999, 123456},
			index: 3,
			wantSlice: []interface{}{123, 156, 999},
			wantVal: 123456,
		},
		{
			name: "string index last",
			slice: []interface{}{"aaa", "bbb", "ccc", "ddd"},
			index: 3,
			wantSlice: []interface{}{"aaa", "bbb", "ccc"},
			wantVal: "ddd",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, val, err := Delete(tc.slice, tc.index)
			if !reflect.DeepEqual(res, tc.wantSlice) {
				t.Errorf("Input: %v, %d\nExpected: %v, %v\nGot: %v, %v", tc.slice, tc.index, tc.wantSlice, tc.wantVal, res, val)
			}
			if !reflect.DeepEqual(err, tc.wantErr) {
				t.Errorf("Want: %v\nbut get: %v", tc.wantErr, err)
			}
		})
	}
}

func ExampleDelete() {
	src := []int{1, -2, 3}
	idx := 0
	res, val, err := Delete(src, idx)
	fmt.Println(res, val, err)
	// output: [-2 3] 1 <nil>
}

// 测试：删除序列头 int
func BenchmarkDeleteIF(b *testing.B) {
	slice := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index := 0
	for i := 0; i < b.N; i++ {
		Delete(slice, index)
	}
}

// 测试：删除序列尾 int
func BenchmarkDeleteIL(b *testing.B) {
	slice := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index := 9
	for i := 0; i < b.N; i++ {
		Delete(slice, index)
	}
}

// 测试：删除序列中间 int
func BenchmarkDeleteIM(b *testing.B) {
	slice := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index := 4
	for i := 0; i < b.N; i++ {
		Delete(slice, index)
	}
}

// 测试：删除序列头 string
func BenchmarkDeleteSF(b *testing.B) {
	slice := []interface{}{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	index := 0
	for i := 0; i < b.N; i++ {
		Delete(slice, index)
	}
}

// 测试：删除序列尾 string
func BenchmarkDeleteSL(b *testing.B) {
	slice := []interface{}{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	index := 9
	for i := 0; i < b.N; i++ {
		Delete(slice, index)
	}
}

// 测试：删除序列中间 string
func BenchmarkDeleteSM(b *testing.B) {
	slice := []interface{}{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	index := 4
	for i := 0; i < b.N; i++ {
		Delete(slice, index)
	}
}