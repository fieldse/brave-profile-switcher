// Contains json-related utility functions
package switcher

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/fieldse/brave-profile-switcher/internal/logger"
)

// PrintJSON returns pretty-printed JSON with 2-space indentation
func PrintJSON(data interface{}) string {
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

// PrintKeys prints the keys from a JSON map
func PrintKeys(data interface{}) {
	keys := reflect.ValueOf(data).MapKeys()

	logger.Info("keys", "\n")
	for i, k := range keys {
		fmt.Printf("  %d -- %s\n", i+1, k)
	}
}
