package main

import (
	"errors"
	config "github.com/jamesmoriarty/gohack/config"
	hacks "github.com/jamesmoriarty/gohack/hacks"
	win32 "github.com/jamesmoriarty/gohack/win32"
	log "github.com/sirupsen/logrus"
	"os"
	"unsafe"
)

const (
	processName = "csgo.exe"
	moduleName  = "client_panorama.dll"
)

func instrument() (*win32.HANDLE, *config.Addresses, error) {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})

	offsets, err := config.GetOffsets()
	if err != nil {
		return nil, nil, errors.New("Failed to get offsets " + err.Error())
	}

	pid, success := win32.GetProcessID(processName)
	log.WithFields(log.Fields{"pid": pid}).Info("GetProcessID ", processName)
	if !success {
		return nil, nil, errors.New("Failed to get pid " + processName)
	}

	_, success, address := win32.GetModule(moduleName, pid)
	log.WithFields(log.Fields{"address": address}).Info("GetModule ", moduleName)
	if !success {
		return nil, nil, errors.New("Failed to get module address " + moduleName)
	}

	handle, err := win32.OpenProcess(win32.PROCESS_ALL_ACCESS, false, pid)
	log.WithFields(log.Fields{"handle": handle}).Info("OpenProcess ", pid)

	addresses, err := config.GetAddresses(handle, uintptr(unsafe.Pointer(address)), offsets)

	return &handle, addresses, err
}

func main() {
	config.PrintBanner()

	handle, addresses, err := instrument()

	if err != nil {
		log.Fatal(err)

		os.Exit(1)
	}

	hacks.RunBHOP(*handle, addresses)

	select {}
}
