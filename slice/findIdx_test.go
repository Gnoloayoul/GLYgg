package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindIdx(t *testing.T) {
	testCases := []struct {
		name string
		slice []interface{}
		input interface{}
		wantVal interface{}
	} {
		{
			// 数组正常搜索
			name: "int normal case",
			slice: []interface{}{123, 156, 999},
			input: 156,
			wantVal: 1,
		},
		{
			// 字符串正常搜索
			name: "string normal case",
			slice: []interface{}{"aaa", "bbb", "ccc", "ddd"},
			input: "ddd",
			wantVal: 3,
		},
		{
			// 当数组有重复元素
			name: "repetition of figures",
			slice: []interface{}{123, 156, 156, 999},
			input: 156,
			wantVal: 1,
		},
		{
			// 当字符串有重复元素
			name: "repeating character",
			slice: []interface{}{"aaa", "bbb", "ccc", "ddd", "ddd", "ddd"},
			input: "ddd",
			wantVal: 3,
		},
		{
			// 切片混杂，还有重复元素，找数字
			name: "hybrid element find number",
			slice: []interface{}{123, "bbb", 123, "ddd", "ddd", "ddd"},
			input: 123,
			wantVal: 0,
		},
		{
			// 切片混杂，还有重复元素，找字符串
			name: "hybrid element find number",
			slice: []interface{}{123, "bbb", 123, "ddd", "ddd", "ddd"},
			input: "ddd",
			wantVal: 3,
		},
		{
			// 切片混杂，还有重复元素，找不到
			name: "hybrid element find number",
			slice: []interface{}{123, "bbb", 123, "ddd", "ddd", "ddd"},
			input: "dddd",
			wantVal: -1,
		},

	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := FindIdx(tc.slice, tc.input)
			if !reflect.DeepEqual(res, tc.wantVal) {
				t.Errorf("Input: %v, %v\nExpected: %v\nGot: %v", tc.slice, tc.input, tc.wantVal, res)
			}
		})
	}
}

func ExampleFindIdx() {
	src := []string{"a", "b", "c"}
	input := "c"
	fmt.Println(FindIdx(src, input))
	// output: 2
}

func BenchmarkFindIdx(b *testing.B) {
	slice := []interface{}{1, "bb", 3, 4, "ccc", 6, 7, "ddd", "ddd", 9, 9, 10, 10, "bb", 10}
	input := interface{}("ccc")
	for i := 0; i < b.N; i++ {
		FindIdx(slice, input)
	}
}