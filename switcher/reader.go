package switcher

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/mnogu/go-dig"
)

// BraveData reads the profiles.json file in our local Brave directory
// Returns as map
func BraveData() (map[string]interface{}, error) {
	var data = make(map[string]interface{})

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

// ProfileData returns slice of raw profile data maps
func ProfileData(data map[string]interface{}) (interface{}, error) {
	var res interface{}
	res, err := dig.Dig(data, "profile", "info_cache")
	if err != nil {
		return res, fmt.Errorf("error reading profile data: %s", err.Error())
	}
	return res, nil
}

// PrintJSON returns JSON 2 pretty-printed
func PrintJSON(data map[string]interface{}) string {
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
