package switcher

import (
	"reflect"
	"testing"

	"github.com/fieldse/brave-profile-switcher/internal/logger"

	"github.com/stretchr/testify/assert"
)

func TestParseData(t *testing.T) {
	fp := braveConfigFilepath()

	// Read as raw bytes
	buf, err := ReadToBytes(fp)
	assert.Nil(t, err)

	res, err := parseToStruct(buf)
	assert.Nil(t, err)
	logger.Debug("parseToStruct", "profile data: %+v", res)
}

func TestBraveData(t *testing.T) {
	res, err := ReadBraveData()
	assert.Nil(t, err)

	// print map keys
	keys := reflect.ValueOf(res).MapKeys()
	logger.Debug("keys", "%s", keys)

	// pretty-print json
	x := PrintJSON(res)
	logger.Debug("data", x)
}

func Test_braveDirectory(t *testing.T) {
	res := braveDirectory()
	assert.Contains(t, res, ".config/BraveSoftware/Brave-Browser")
	assert.DirExists(t, res)
	logger.Debug("brave directory", res)
}

func Test_braveConfigFilepath(t *testing.T) {
	res := braveConfigFilepath()
	assert.FileExists(t, res)
}

func TestProfileData(t *testing.T) {
	x, err := ReadBraveData()
	assert.Nil(t, err)

	res, err := ProfileData(x)
	if err != nil {
		logger.Debug("error", "%s", err.Error())
		return
	}
	assert.Nil(t, err)
	logger.Debug("returned data", "%+v", res)
}
