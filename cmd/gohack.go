package main

import (
	"github.com/jamesmoriarty/gohack"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	gohack.PrintBanner()

	process, addresses, err := gohack.Instrument()

	if err != nil {
		log.Fatal(err)

		os.Exit(1)
	}

	gohack.RunBHOP(process, addresses)
}
