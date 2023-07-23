package main

import (
	"fmt"
	"git.acwing.com/Gnoloayoul/gylgg/slice"
)

// 这里相当于playground

func main() {
	src := []int{1, -2, 3}
	res := slice.FilterDelete(src, func(idx, src int) bool {
		return idx == 0
	})
	fmt.Println(res)
}