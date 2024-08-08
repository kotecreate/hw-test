package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStCompCaseTable(t *testing.T) {
	testCases := []struct {
		name   string
		d      any
		expect float64
	}{
		{
			name:   "testSC",
			d:      &Circle{8},
			expect: 201.06192982974676,
		},
		{
			name:   "testSC",
			d:      &Line{7},
			expect: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual, err := calculateArea(tC.d)
			assert.Equal(t, tC.expect, actual)
			if err != nil {
				fmt.Println(err)
			}
		})
	}
}
