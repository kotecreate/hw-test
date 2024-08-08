package hw03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardCaseTable(t *testing.T) {
	testCases := []struct {
		name   string
		size   uint8
		expect string
	}{
		{
			name:   "empty",
			size:   2,
			expect: "  ##\n##  \n",
		},
		{
			name:   "empty",
			size:   0,
			expect: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actual := board(tC.size)
			assert.Equal(t, tC.expect, actual)
		})
	}
}
