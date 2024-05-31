package main

import (
	"fmt"
	"hw-test/printer"
	"hw-test/reader"
	"hw-test/types"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)
	if err != nil {
		fmt.Print(err)
		return
	}

	printer.PrintStaff(staff)
}
