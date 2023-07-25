package main

import (
	"fmt"
	"git.acwing.com/Gnoloayoul/gylgg/slice"
	"sort"
)

// 这里相当于playground

func main() {
	src := []int{1, 3, 4, 2}
	dst := []int{2, 5, 7, 3}
	res := slice.SymmetricDiffSet[int](src, dst)
	sort.Ints(res) // 给结果排个序
	fmt.Println(res)
}