package main

import (
	"fmt"
	"git.acwing.com/Gnoloayoul/gylgg/slice"
	"sort"
)

// 这里相当于playground

func main() {
	src1 := []int{1, 3, 4, 5}
	src2 := []int{1, 4, 7}
	res1 := slice.UnionSet(src1, src2)
	sort.Ints(res1) // 给结果排下序
	fmt.Println(res1)
}