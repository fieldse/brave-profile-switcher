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

// ReadToBytes reads file to bytes
func ReadToBytes(fp string) (buf []byte, err error) {
	if !exists(fp) {
		return buf, fmt.Errorf("path %s not found", fp)
	}
	buf, err = os.ReadFile(fp)
	if err != nil {
		return buf, fmt.Errorf("error reading file: %w", err)
	}
	return buf, nil
}

// ReadJSONFile reads and unmarshals data from a json file to a map
func ReadJSONFile(fp string) (data map[string]interface{}, err error) {
	data = make(map[string]interface{})
	buf, err := ReadToBytes(fp)
	if err != nil {
		return data, err
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
