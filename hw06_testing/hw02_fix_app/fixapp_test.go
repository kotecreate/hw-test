package main

import (
	"fmt"
	"testing"

	"github.com/kotecreate/hw-test/hw06_testing/hw02_fix_app/reader"
	"github.com/kotecreate/hw-test/hw06_testing/hw02_fix_app/types"
	"github.com/stretchr/testify/assert"
)

func TestStCompCaseTable(t *testing.T) {
	testCases := []struct {
		name   string
		a      string
		expect []types.Employee
	}{
		{
			name:   "testReadJ",
			a:      "test data.json",
			expect: []types.Employee{{UserID: 9, Age: 25, Name: "Bob", DepartmentID: 4}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			var err error
			var actual []types.Employee
			actual, err = reader.ReadJSON(tC.a)
			if err != nil {
				fmt.Print(err)
				return
			}
			assert.Equal(t, tC.expect, actual)
		})
	}
}
