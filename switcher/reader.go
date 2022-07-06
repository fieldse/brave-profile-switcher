package switcher

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

// jsonMap represents basic json data structure as a map
type jsonMap map[string]interface{}

// BraveData reads the profiles.json file in our local Brave directory
// Returns as map
func BraveData() (jsonMap, error) {
	var data = make(jsonMap)

	fp := braveConfigFilepath()
	if !exists(fp) {
		return data, fmt.Errorf("brave home directory %s not found", fp)
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

// PrintJSON returns JSON 2 pretty-printed
func PrintJSON(data jsonMap) string {
	strData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal("failed to parse JSON data: %w", err)
	}
	return string(strData)
}

// exists checks a filepath exists
func exists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

// braveConfigFilepath returns filepath to the Brave config JSON file
func braveConfigFilepath() string {
	return path.Join(braveDirectory(), "Local State")
}

// braveDirectory returns the path to the user's brave data directory
func braveDirectory() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return path.Join(homeDir, ".config/BraveSoftware/Brave-Browser")
}
