package main

import (
	"fmt"
	"hw02/printer"
	"hw02/reader"
	"hw02/types"
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
