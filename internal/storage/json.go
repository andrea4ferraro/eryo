package storage

import (
	"encoding/json"
	"os"
)

func Load(file string, target any) error {

	data, err := os.ReadFile(file)

	if err != nil {

		return err

	}

	return json.Unmarshal(data, target)

}
