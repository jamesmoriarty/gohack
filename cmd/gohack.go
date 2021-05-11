package main

import (
	"fmt"
	"os"

	"github.com/jamesmoriarty/gohack"
	log "github.com/sirupsen/logrus"
)

func PrintBanner() {
	fmt.Printf(
		`
    ___       ___       ___       ___       ___       ___   
   /\  \     /\  \     /\__\     /\  \     /\  \     /\__\  
  /::\  \   /::\  \   /:/__/_   /::\  \   /::\  \   /:/ _/_ 
 /:/\:\__\ /:/\:\__\ /::\/\__\ /::\:\__\ /:/\:\__\ /::-"\__\
 \:\:\/__/ \:\/:/  / \/\::/  / \/\::/  / \:\ \/__/ \;:;-",-"
  \::/  /   \::/  /    /:/  /    /:/  /   \:\__\    |:|  |  
   \/__/     \/__/     \/__/     \/__/     \/__/     \|__| 
 
version: %s-%s

`, gohack.Version, gohack.Date)
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
