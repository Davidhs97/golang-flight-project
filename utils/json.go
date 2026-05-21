package utils

import (
	"encoding/json"
	"os"
)

func LoadJSON(path string, target interface{}) error {
	file, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	return json.Unmarshal(file, target)
}
