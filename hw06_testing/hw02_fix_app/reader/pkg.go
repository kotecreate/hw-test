package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/kotecreate/hw-test/hw06_testing/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	var data []types.Employee
	if err = json.Unmarshal(bytes, &data); err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	return data, nil
}
