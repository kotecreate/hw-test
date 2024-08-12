package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCaseTable(t *testing.T) {
	var err error
	testCases := []struct {
		name string
		str  string
		exp  map[string]int
	}{
		{
			name: "testCount",
			str:  "Привет мир! привет",
			exp:  (map[string]int{"мир": 1, "привет": 2}),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := countWords(tC.str)
			if err != nil {
				fmt.Print(err)
			}
			assert.Equal(t, tC.exp, actual)
		})
	}
}

func TestCleanCaseTable(t *testing.T) {
	testCases := []struct {
		name string
		str  string
		exp  []string
	}{
		{
			name: "testCount",
			str:  "Привет мир! привет",
			exp:  ([]string{"привет", "мир", "привет"}),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := cleanWords(tC.str)
			assert.Equal(t, tC.exp, actual)
		})
	}
}
