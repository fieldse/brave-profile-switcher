package switcher

import (
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
