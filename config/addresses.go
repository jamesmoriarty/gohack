package config

import (
	util "github.com/jamesmoriarty/gohack/util"
	win32 "github.com/jamesmoriarty/gohack/win32"
	log "github.com/sirupsen/logrus"
)

type Addresses struct {
	Local            uintptr
	LocalForceJump   uintptr
	LocalPlayer      uintptr
	LocalPlayerFlags uintptr
}

func GetAddresses(processHandle win32.HANDLE, address uintptr, offsets *Offsets) (*Addresses, error) {
	addresses := Addresses{}

	addresses.Local = address
	log.WithFields(log.Fields{"value": util.ConvertPtrToHex(addresses.Local)}).Info("- addressLocal")

	addresses.LocalForceJump = addresses.Local + offsets.Signatures.OffsetForceJump
	log.WithFields(log.Fields{"value": util.ConvertPtrToHex(addresses.LocalForceJump)}).Info("- addressLocalForceJump")

	win32.ReadProcessMemory(processHandle, win32.LPCVOID(addresses.Local+offsets.Signatures.OffsetLocalPlayer), &addresses.LocalPlayer, 4)
	log.WithFields(log.Fields{"value": util.ConvertPtrToHex(addresses.LocalPlayer)}).Info("- addressLocalPlayer")

	addresses.LocalPlayerFlags = addresses.LocalPlayer + offsets.Netvars.OffsetLocalPlayerFlags
	log.WithFields(log.Fields{"value": util.ConvertPtrToHex(addresses.LocalPlayerFlags)}).Info("- addressLocalPlayerFlags")

	return &addresses, nil
}
