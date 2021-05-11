package gohack

import (
	"errors"

	"github.com/jamesmoriarty/gohack/internal/gohack"
	"github.com/jamesmoriarty/gomem"
	log "github.com/sirupsen/logrus"
)

var (
	Version string
	Date    string
)

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

	client, err := gohack.GetClientFrom(process, offsets)
	if err != nil {
		return nil, err
	}
	log.WithFields(log.Fields{"handle": process.Handle}).Info("OpenProcess ", process.ID)
	log.WithFields(log.Fields{"value": gohack.ToHexString(client.Address)}).Info("- Address")
	log.WithFields(log.Fields{"value": gohack.ToHexString(client.OffsetForceJump())}).Info("- OffsetForceJump")
	log.WithFields(log.Fields{"value": gohack.ToHexString(client.OffsetForceAttack())}).Info("- OffsetForceAttack")
	log.WithFields(log.Fields{"value": gohack.ToHexString(client.OffsetPlayer())}).Info("- OffsetPlayer")
	log.WithFields(log.Fields{"value": gohack.ToHexString(client.OffsetPlayerFlags())}).Info("- OffsetPlayerFlags")
	log.WithFields(log.Fields{"value": gohack.ToHexString(client.OffsetEntityId())}).Info("- OffsetEntityId")

	return client, err
}

func Execute(c *gohack.Client) {
	go gohack.RunTrigger(c)
	gohack.RunBHOP(c)
}
