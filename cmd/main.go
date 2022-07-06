package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fieldse/brave-profile-switcher/switcher"
)

func main() {
	data, err := switcher.BraveData()
	if err != nil {
		fail("failed to read brave data", err)
	}
	res, err := switcher.ProfileData(data)
	if err != nil {
		fail("failed to read profile data", err)
	}
	fmt.Print("profile data: ", res)
}

func fail(msg string, err error) {
	log.Printf("%s: error -- %s", msg, err.Error())
	os.Exit(1)
}
