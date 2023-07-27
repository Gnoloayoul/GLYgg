package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShrink(t *testing.T) {
	testCases := []struct {
		name        string
		originCap   int
		enqueueLoop int
		expectCap   int
	}{
		{
			name:        "小于64",
			originCap:   4000,
			enqueueLoop: 999,
			expectCap:   2500,
		},
		{
			name:        "小于2048, 不足1/4",
			originCap:   2047,
			enqueueLoop: 10,
			expectCap:   1023,
		},
		{
			name:        "小于2048, 超过1/4",
			originCap:   2049,
			enqueueLoop: 10,
			expectCap:   1280,
		},
		{
			name:        "大于2048，不足一半",
			originCap:   3000,
			enqueueLoop: 60,
			expectCap:   1875,
		},
		{
			name:        "大于2048，大于一半",
			originCap:   3000,
			enqueueLoop: 2000,
			expectCap:   3000,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := make([]int, 0, tc.originCap)

			for i := 0; i < tc.enqueueLoop; i++ {
				l = append(l, i)
			}
			l = Shrink[int](l)
			assert.Equal(t, tc.expectCap, cap(l))
		})
	}
}

