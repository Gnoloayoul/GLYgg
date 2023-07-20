// 这里相当于playground

package main
import (
	"fmt"
	"git.acwing.com/Gnoloayoul/gylgg/slice"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	idx := 0
	res, val, err := slice.Delete(a, idx)
	if err == nil {
		fmt.Printf("befort: \narr: %v, len: %d, cap: %d\nafter: \narr: %v, len: %d, cap: %d\n%d", a, len(a), cap(a), res, len(res), cap(res), val)
	} else {
		fmt.Println(err)
	}
}