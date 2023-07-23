package main

import (
	"fmt"
	"git.acwing.com/Gnoloayoul/gylgg/slice"
)

// 这里相当于playground

func main() {
	src1 := []int{1, 1, 3, 5, 7}
	src2 := []int{1, 3, 5, 5}
	res := slice.DiffSet(src1, src2)
	fmt.Println(res)
}