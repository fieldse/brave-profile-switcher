package switcher

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBraveData(t *testing.T) {
	res, err := BraveData()
	assert.Nil(t, err)

	// print map keys
	keys := reflect.ValueOf(res).MapKeys()
	log.Printf("\n\n\n =================> keys: %s\n", keys)

	// pretty-print json
	x := PrintJSON(res)
	fmt.Printf("\n\n\n =================> data: %s\n", x)
}

func Test_braveDirectory(t *testing.T) {
	res := braveDirectory()
	assert.Contains(t, res, ".config/BraveSoftware/Brave-Browser")
	assert.DirExists(t, res)
}

func Test_braveConfigFilepath(t *testing.T) {
	res := braveConfigFilepath()
	assert.FileExists(t, res)
}
