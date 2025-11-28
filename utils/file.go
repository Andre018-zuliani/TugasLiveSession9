package utils

import (
	"encoding/json"
	"os"
)

// ReadJSON reads JSON file into v
func ReadJSON(path string, v interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		// if file does not exist, treat as empty data
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(v)
}

// WriteJSON writes v into JSON file
func WriteJSON(path string, v interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}
