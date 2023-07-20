package main

import (
	"fmt"
	"git.acwing.com/Gnoloayoul/gylgg/slice"
)

// 这里相当于playground

func main() {
	a := []any{123, 156, 999, 123456}
	index := 2
	input := interface{}(666)
	fmt.Println(slice.Add(a, index, input))
	fmt.Println(a[index:])
}