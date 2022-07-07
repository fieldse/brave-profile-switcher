package main

import (
	"log"
	"os"

	"github.com/fieldse/brave-profile-switcher/internal/logger"
	"github.com/fieldse/brave-profile-switcher/switcher"
)

func main() {
	data, err := switcher.ReadBraveData()
	if err != nil {
		fail("failed to read brave data", err)
	}
	res, err := switcher.ProfileData(data)
	if err != nil {
		fail("failed to read profile data", err)
	}
	strData := switcher.PrintJSON(res)
	logger.Info("profile data", strData)
}

func fail(msg string, err error) {
	log.Printf("%s: error -- %s", msg, err.Error())
	os.Exit(1)
}
