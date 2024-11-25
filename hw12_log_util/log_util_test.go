package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParceLogFile(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		level    string
		want     []LogType
	}{
		{
			name:     "Test 1",
			fileName: "test.txt",
			level:    "info",
			want:     []LogType{},
		},
	}
	for _, tC := range tests {
		t.Run(tC.name, func(t *testing.T) {
			actual, err := parseLogFile(tC.fileName, tC.level)
			if err != nil {
				fmt.Print(err)
			}
			assert.Equal(t, tC.want, actual)

		})
	}
}
