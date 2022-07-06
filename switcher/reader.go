package switcher

import (
	"fmt"
	"os"
	"path"

	"github.com/mnogu/go-dig"
)

// BraveData reads the profiles.json file in our local Brave directory
// Returns as map
func BraveData() (map[string]interface{}, error) {
	fp := braveConfigFilepath()
	return ReadJSONFile(fp)
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
