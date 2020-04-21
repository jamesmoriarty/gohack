package config

import (
	"errors"
	"fmt"
	"github.com/jamesmoriarty/gomem"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type Addresses struct {
	Local            uintptr
	LocalForceJump   uintptr
	LocalPlayer      uintptr
	LocalPlayerFlags uintptr
}

func GetAddresses(p *gomem.Process, address uintptr, offsets *Offsets) (*Addresses, error) {
	addresses := Addresses{}

	addresses.Local = address
	log.WithFields(log.Fields{"value": convertPtrToHex(addresses.Local)}).Info("- addressLocal")

	addresses.LocalForceJump = addresses.Local + offsets.Signatures.OffsetForceJump
	log.WithFields(log.Fields{"value": convertPtrToHex(addresses.LocalForceJump)}).Info("- addressLocalForceJump")

	p.Read(addresses.Local+offsets.Signatures.OffsetLocalPlayer, &addresses.LocalPlayer, 4)
	log.WithFields(log.Fields{"value": convertPtrToHex(addresses.LocalPlayer)}).Info("- addressLocalPlayer")

	addresses.LocalPlayerFlags = addresses.LocalPlayer + offsets.Netvars.OffsetLocalPlayerFlags
	log.WithFields(log.Fields{"value": convertPtrToHex(addresses.LocalPlayerFlags)}).Info("- addressLocalPlayerFlags")

	if addresses.LocalPlayer == 0x0 {
		return nil, errors.New("Failed to get LocalPlayer address")
	}

	return &addresses, nil
}

func convertPtrToHex(ptr uintptr) string {
	s := fmt.Sprintf("%d", ptr)
	n, _ := strconv.Atoi(s)
	h := fmt.Sprintf("0x%x", n)
	return h
}
