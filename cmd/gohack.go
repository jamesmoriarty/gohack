package main

import (
	"github.com/jamesmoriarty/gohack"
	log "github.com/sirupsen/logrus"
	"fmt"
	"os"
)

func PrintBanner() {
	fmt.Printf(gohack.Banner, gohack.Version, gohack.Date)

	fmt.Println()
}

func main() {
	PrintBanner()

	client, err := gohack.Instrument()

	if err != nil {
		log.Fatal(err)

		os.Exit(1)
	}

	gohack.Execute(client)
}
