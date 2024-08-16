package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryCaseTable(t *testing.T) {
	testCases := []struct {
		name  string
		slice []int
		value int
		exp   int
	}{
		{
			name:  "testBinarySearch",
			slice: []int{4, 7, 12, 17, 23, 25, 38, 42},
			value: 38,
			exp:   38,
		},
		{
			name:  "testBinarySearch",
			slice: []int{24, 41, 52, 66, 74, 80, 91, 102, 225},
			value: 91,
			exp:   91,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := binaryS(tC.slice, tC.value)
			assert.Equal(t, tC.exp, actual)
		})
	}
}
