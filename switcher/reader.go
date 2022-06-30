package switcher

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

// jsonMap represents basic json data structure as a map
type jsonMap map[string]interface{}

// BraveData reads the profiles.json file in our local Brave directory
func BraveData() (data jsonMap, err error) {
	data = make(jsonMap)
	fp := braveHomeDirectory()
	if exists(fp) {
		return data, fmt.Errorf("brave home directory %s not found", fp)
	}
	buf, err := os.ReadFile(fp)
	if exists(fp) {
		return data, fmt.Errorf("error reading file: %w", err)
	}
	err = json.Unmarshal(buf, &data)
	if err != nil {
		return data, fmt.Errorf("error loading JSON data: %w", err)
	}
	return data, nil
}

// exists checks a filepath exists
func exists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

// braveHomeDirectory returns the path to the user's brave data directory
func braveHomeDirectory() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return path.Join(homeDir, ".config/BraveSoftware/Brave-Browser")
}
