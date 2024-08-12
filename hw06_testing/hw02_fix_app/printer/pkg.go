package printer

import (
	"fmt"

	"github.com/kotecreate/hw-test/hw06_testing/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	for i := 0; i < len(staff); i++ {
		str, err := fmt.Println(staff[i].String())
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		fmt.Println(str)
	}
}
