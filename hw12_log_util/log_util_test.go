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
			fileName: "testdata.txt",
			level:    "",
			want: []LogType{
				{Time: "2024-01-01T00:00:01", Level: "INFO", Message: "First message"},
				{Time: "2024-01-02T00:00:05", Level: "WARNING", Message: "Second message"},
			},
		},
		{
			name:     "Test 2",
			fileName: "testdata.txt",
			level:    "WARNING",
			want: []LogType{
				{Time: "2024-01-02T00:00:05", Level: "WARNING", Message: "Second message"},
			},
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

func TestAnalyze(t *testing.T) {
	tests := []struct {
		name  string
		input []LogType
		want  map[string]int
	}{
		{
			name: "Test 3",
			input: []LogType{
				{Time: "2024-01-01T00:00:01", Level: "INFO", Message: "First message"},
				{Time: "2024-01-02T00:00:05", Level: "WARNING", Message: "Second message"},
			},
			want: map[string]int{
				"INFO":    1,
				"WARNING": 1,
			},
		},
	}
	for _, tC := range tests {
		t.Run(tC.name, func(t *testing.T) {
			actual := analyze(tC.input)
			assert.Equal(t, tC.want, actual)
		})
	}
}

func TestPrintStats(t *testing.T) {
	tests := []struct {
		name       string
		input      map[string]int
		outputPath string
		want       []string
	}{
		{
			name: "Test 4",
			input: map[string]int{
				"INFO":    3,
				"WARNING": 2,
			},
			outputPath: "",
			want: []string{
				"Статистика:",
				"INFO: 3",
				"WARNING: 2",
			},
		},
	}
	for _, tC := range tests {
		t.Run(tC.name, func(t *testing.T) {
			err := printStats(tC.input, tC.outputPath)
			if err != nil {
				t.Errorf("не верные значения")
			}
		})
	}
}
