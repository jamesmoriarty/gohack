package gohack

import (
	"errors"
	"fmt"
	"github.com/jamesmoriarty/gohack/internal/gohack"
	"github.com/jamesmoriarty/gomem"
	log "github.com/sirupsen/logrus"
	"strconv"
)

var (
	Version string
	Date    string
	Banner  = `
    ___       ___       ___       ___       ___       ___   
   /\  \     /\  \     /\__\     /\  \     /\  \     /\__\  
  /::\  \   /::\  \   /:/__/_   /::\  \   /::\  \   /:/ _/_ 
 /:/\:\__\ /:/\:\__\ /::\/\__\ /::\:\__\ /:/\:\__\ /::-"\__\
 \:\:\/__/ \:\/:/  / \/\::/  / \/\::/  / \:\ \/__/ \;:;-",-"
  \::/  /   \::/  /    /:/  /    /:/  /   \:\__\    |:|  |  
   \/__/     \/__/     \/__/     \/__/     \/__/     \|__| 
 
version: %s-%s
`
)

func PrintBanner() {
	fmt.Printf(Banner, Version, Date)

	fmt.Println()
}

func ptrToHex(ptr uintptr) string {
	s := fmt.Sprintf("%d", ptr)
	n, _ := strconv.Atoi(s)
	h := fmt.Sprintf("0x%x", n)
	return h
}

func Instrument() (*gomem.Process, *gohack.Client, error) {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})

	offsets, err := gohack.GetOffsets()
	if err != nil {
		return nil, nil, errors.New("Failed to get offsets " + err.Error())
	}

	process, err := gomem.GetFromProcessName("csgo.exe")
	if err != nil {
		return nil, nil, errors.New("Failed to get pid csgo.exe")
	}
	log.WithFields(log.Fields{"pid": process.ID}).Info("GetFromProcessName csgo.exe")

	client, err := gohack.ClientFrom(process, offsets)
	if err != nil {
		return nil, nil, err
	}

	log.WithFields(log.Fields{"handle": process.Handle}).Info("OpenProcess ", process.ID)
	log.WithFields(log.Fields{"value": ptrToHex(client.Offset)}).Info("- Offset")
	log.WithFields(log.Fields{"value": ptrToHex(client.OffsetForceJump())}).Info("- OffsetForceJump")
	log.WithFields(log.Fields{"value": ptrToHex(client.OffsetPlayer())}).Info("- OffsetPlayer")
	log.WithFields(log.Fields{"value": ptrToHex(client.OffsetPlayerFlags())}).Info("- OffsetPlayerFlags")

	return process, client, err
}

func Execute(p *gomem.Process, c *gohack.Client) {
	gohack.RunBHOP(p, c)
}
