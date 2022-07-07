package switcher

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/fieldse/brave-profile-switcher/internal/logger"
	"github.com/mnogu/go-dig"
)

// BraveData represents a portion of the Brave settings JSON, which we need
// to extract the Brave profiles data
type BraveData struct {
	Profile struct {
		// InfoCache map[string]interface{} `json:"info_cache"`
		InfoCache map[string]BraveProfile `json:"info_cache"`
	}
}

type BraveProfile struct {
	Name                  string `json:"name"` // user-defined profile name
	AvatarIcon            string `json:"avatar_icon"`
	ProfileHighlightColor int    `json:""`
	ProfileDirName        string // brave auto-defined directory name. eg: "Profile 8"
}

// ReadBraveData reads the profiles.json file in our local Brave directory
// Returns as map
func ReadBraveData() (map[string]interface{}, error) {
	fp := braveConfigFilepath()
	return ReadJSONFile(fp)
}

// parseToStruct will try to parse json file data to a BraveData struct
func parseToStruct(buf []byte) (BraveData, error) {
	var b1 BraveData

	// Attempt to unmarshal
	err := json.Unmarshal(buf, &b1)
	if err != nil {
		logger.Error(err, "unmarshal JSON data")
		return b1, err
	}
	logger.Success("parse json data", "data: %s", PrintJSON(b1))
	return b1, nil
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
