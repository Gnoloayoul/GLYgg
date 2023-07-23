package main

import (
	"fmt"
	"git.acwing.com/Gnoloayoul/gylgg/slice"
)

// 这里相当于playground

func main() {
	src := []int{1, -2, 3, 5, 3}
	dst := 3
	res := slice.AllIndex(src, dst)
	fmt.Println(res)
}