package slice

import (
	"fmt"
	"reflect"
	"testing"

)

func TestFilterDelete(t *testing.T) {
	testCases := []struct {
		name string
		src []int
		deleteCondition func(idx, src int) bool
		wantRes []int
	} {
		{
			name: "empty src",
			src: []int{},
			deleteCondition: func(idx, src int) bool {
				return false
			},
			wantRes: []int{},
		},
		{
			name: "no delete",
			src: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx, src int) bool {
				return false
			},
			wantRes: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "delete the front",
			src: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx, src int) bool {
				return idx == 0
			},
			wantRes: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "delete front two",
			src: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx, src int) bool {
				return idx == 0 || idx == 1
			},
			wantRes: []int{2, 3, 4, 5, 6, 7},
		},
		{
			name: "delete middle single",
			src: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx, src int) bool {
				return idx == 3
			},
			wantRes: []int{0, 1, 2, 4, 5, 6, 7},
		},
		{
			name: "delete middle els not slice",
			src: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx, src int) bool {
				return idx == 2 || idx == 4
			},
			wantRes: []int{0, 1, 3, 5, 6, 7},
		},
		{
			name: "delete middle els slice",
			src: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx, src int) bool {
				return idx == 3 || idx == 4
			},
			wantRes: []int{0, 1, 2, 5, 6, 7},
		},
		{
			name: "delete middle els case1",
			src: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx, src int) bool {
				return idx == 2 || idx == 4 || idx == 5
			},
			wantRes: []int{0, 1, 3, 6, 7},
		},
		{
			name: "delete middle els case2",
			src: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx, src int) bool {
				return idx == 2 || idx == 3 || idx == 5
			},
			wantRes: []int{0, 1, 4, 6, 7},
		},
		{
			name: "delete middle els case3",
			src: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx, src int) bool {
				return idx == 6 || idx == 7
			},
			wantRes: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "delete middle els case4",
			src: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx, src int) bool {
				return idx == 7
			},
			wantRes: []int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			name: "delete all",
			src: []int{0, 1, 2, 3, 4, 5, 6, 7},
			deleteCondition: func(idx, src int) bool {
				return true
			},
			wantRes: []int{},
		},

	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := FilterDelete(tc.src, tc.deleteCondition)
			if !reflect.DeepEqual(res, tc.wantRes) {
				t.Errorf("Input: %v\nExpected: %v\nGot: %v", tc.src, tc.wantRes, res)
			}
		})
	}
}

func ExampleFilterDelete() {
	src := []int{1, -2, 3}
	res := FilterDelete(src, func(idx, src int) bool {
		return idx == 0
	})
	fmt.Println(res)
	// output: [-2 3]
}