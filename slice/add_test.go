package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name string
		slice []interface{}
		index int
		input interface{}
		wantSlice []interface{}
	} {
		{
			name: "front",
			slice: []interface{}{123, 156, 999},
			index: 0,
			input: interface{}("aaa"),
			wantSlice: []interface{}{"aaa", 123, 156, 999},
		},
		{
			name: "back",
			slice: []interface{}{123, 156, 999},
			index: 3,
			input: interface{}("aaa"),
			wantSlice: []interface{}{123, 156, 999, "aaa"},
		},
		{
			name: "middle",
			slice: []interface{}{123, 156, 999, 123456},
			index: 2,
			input: interface{}(666),
			wantSlice: []interface{}{123, 156, 666, 999, 123456},
		},
		{
			name: "out of range back",
			slice: []interface{}{123, 156, 999, 123456},
			index: 12,
			input: interface{}(666),
			wantSlice: []interface{}{123, 156, 999, 123456, 666},
		},
		{
			name: "out of range from",
			slice: []interface{}{123, 156, 999, 123456},
			index: -3,
			input: interface{}("aaa"),
			wantSlice: []interface{}{"aaa", 123, 156, 999, 123456},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Add(tc.slice, tc.index, tc.input)
			if !reflect.DeepEqual(res, tc.wantSlice) {
				t.Errorf("Input: %v, %d, %v\nExpected: %v\nGot: %v", tc.slice, tc.index, tc.input, tc.wantSlice, res)
			}
		})
	}
}

func ExampleAdd() {
	src := []int{1, -2, 3}
	input := 6
	idx := 2
	fmt.Println(slice.Add(src, idx, input))
	// output: [1 -2 6 3]
}

// 测试：在序列头添加
func BenchmarkAddF(b *testing.B) {
	slice := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index := 0
	input := interface{}("aa")
	for i := 0; i < b.N; i++ {
		Add(slice, index, input)
	}
}

// 测试：在序列尾添加
func BenchmarkAddL(b *testing.B) {
	slice := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index := 10
	input := interface{}("aa")
	for i := 0; i < b.N; i++ {
		Add(slice, index, input)
	}
}

// 测试：在序列中间添加
func BenchmarkAddM(b *testing.B) {
	slice := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index := 5
	input := interface{}("aa")
	for i := 0; i < b.N; i++ {
		Add(slice, index, input)
	}
}