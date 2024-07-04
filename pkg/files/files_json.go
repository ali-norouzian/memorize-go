package files

import (
	"encoding/json"
	"os"
)

func ReadJsonFile(path string, resultRef any) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &resultRef)
	if err != nil {
		return err
	}

	return nil
}
