package slice

import (
	"fmt"
	"testing"

	"github.com/Gnoloayoul/GYLgg/internal/errs"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	// Delete 主要依赖于 internal/slice.Delete 来保证正确性
	testCases := []struct {
		name      string
		slice     []int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{123, 100},
			index:     0,
			wantSlice: []int{100},
		},
		{
			name:    "index -1",
			slice:   []int{123, 100},
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Delete(tc.slice, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}

func ExampleDelete() {
	res, _ := Delete[int]([]int{1, 2, 3, 4}, 2)
	fmt.Println(res)
	_, err := Delete[int]([]int{1, 2, 3, 4}, -1)
	fmt.Println(err)
	// Output:
	// [1 2 4]
	// GYLgg: 下标超出范围，长度 4, 下标 -1
}
