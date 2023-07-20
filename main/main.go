package main

import (
	"fmt"
	"git.acwing.com/Gnoloayoul/gylgg/slice"
)

// 这里相当于playground

func main() {
	a := []any{123, 156, 999, 123456}
	fmt.Println("before len:", len(a), "cap:", cap(a))
	index := 2
	input := interface{}(666)
	res := slice.Add(a, index, input)
	fmt.Println(res)
	fmt.Println("after len:", len(res), "cap:", cap(res))
}