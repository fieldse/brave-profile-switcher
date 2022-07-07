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
	logger.Debug("read brave data", "data: %v", data)
}

func fail(msg string, err error) {
	log.Printf("%s: error -- %s", msg, err.Error())
	os.Exit(1)
}
