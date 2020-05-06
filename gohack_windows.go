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
)

func ptrToHex(ptr uintptr) string {
	s := fmt.Sprintf("%d", ptr)
	n, _ := strconv.Atoi(s)
	h := fmt.Sprintf("0x%x", n)

	return h
}

func Instrument() (*gohack.Client, error) {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})

	offsets, err := gohack.GetOffsets()
	if err != nil {
		return nil, errors.New("Failed to get offsets " + err.Error())
	}
	log.WithFields(log.Fields{"url": gohack.OffsetsURL}).Info("GetOffsets")

	process, err := gomem.GetFromProcessName("csgo.exe")
	if err != nil {
		return nil, errors.New("Failed to get pid csgo.exe")
	}
	log.WithFields(log.Fields{"pid": process.ID}).Info("GetFromProcessName csgo.exe")

	client, err := gohack.ClientFrom(process, offsets)
	if err != nil {
		return nil, err
	}
	log.WithFields(log.Fields{"handle": process.Handle}).Info("OpenProcess ", process.ID)
	log.WithFields(log.Fields{"value": ptrToHex(client.Offset)}).Info("- Offset")
	log.WithFields(log.Fields{"value": ptrToHex(client.OffsetForceJump())}).Info("- OffsetForceJump")
	log.WithFields(log.Fields{"value": ptrToHex(client.OffsetPlayer())}).Info("- OffsetPlayer")
	log.WithFields(log.Fields{"value": ptrToHex(client.OffsetPlayerFlags())}).Info("- OffsetPlayerFlags")

	return client, err
}

func Execute(c *gohack.Client) {
	gohack.RunBHOP(c)
}
