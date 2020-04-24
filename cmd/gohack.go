package main

import (
	"github.com/jamesmoriarty/gohack"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	gohack.PrintBanner()

	process, client, err := gohack.Instrument()

	if err != nil {
		log.Fatal(err)

		os.Exit(1)
	}

	gohack.Execute(process, client)
}
