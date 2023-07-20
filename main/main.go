package main

import (
	"fmt"
	"git.acwing.com/Gnoloayoul/gylgg/slice"
)

// 这里相当于playground

func main() {
	a := []any{1, 2, 3, "abc", "abc", 5, 5}
	input := interface{}(5)
	fmt.Println(slice.FindIdx(a, input))
}