package main

import (
	"fmt"
	"git.acwing.com/Gnoloayoul/gylgg/slice"
	"strconv"
)

// 这里相当于playground

func main() {
	src := []int{1, 2, 3}
	dst := slice.ChangeSlice(src, func(idx, src int) string {
		return strconv.Itoa(src)
	})
	fmt.Println(dst)
}