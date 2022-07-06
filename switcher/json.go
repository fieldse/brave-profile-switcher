// Contains json-related utility functions
package switcher

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// PrintJSON returns pretty-printed JSON with 2-space indentation
func PrintJSON(data map[string]interface{}) string {
	strData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal("failed to parse JSON data: %w", err)
	}
	return string(strData)
}

func ReadJSONFile(fp string) (map[string]interface{}, error) {
	var data = make(map[string]interface{})
	if !exists(fp) {
		return data, fmt.Errorf("path %s not found", fp)
	}
	buf, err := os.ReadFile(fp)
	if err != nil {
		return data, fmt.Errorf("error reading file: %w", err)
	}
	err = json.Unmarshal(buf, &data)
	if err != nil {
		return data, fmt.Errorf("error loading JSON data: %w", err)
	}
	return data, nil
}
