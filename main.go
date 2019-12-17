package main

import (
	config "github.com/jamesmoriarty/gohack/config"
	hacks "github.com/jamesmoriarty/gohack/hacks"
	util "github.com/jamesmoriarty/gohack/util"
	win32 "github.com/jamesmoriarty/gohack/win32"
	log "github.com/sirupsen/logrus"
	"os"
	"unsafe"
)

const (
	processName = "csgo.exe"
	moduleName  = "client_panorama.dll"
)

func main() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})

	config.PrintBanner()

	offsets, err := config.GetOffsets()
	if err != nil {
		log.Fatal("Failed getting offsets ", err)
		os.Exit(1)
	}

	pid, success := win32.GetProcessID(processName)
	log.WithFields(log.Fields{"pid": pid}).Info("GetProcessID ", processName)
	if !success {
		log.Fatal("Failed to get pid ", processName)
		os.Exit(1)
	}

	_, success, address := win32.GetModule(moduleName, pid)
	log.WithFields(log.Fields{"address": address}).Info("GetModule ", moduleName)
	if !success {
		log.Fatal("Failed to get module address ", moduleName)
		os.Exit(1)
	}

	processHandle, _ := win32.OpenProcess(win32.PROCESS_ALL_ACCESS, false, pid)
	log.WithFields(log.Fields{"processHandle": processHandle}).Info("OpenProcess ", pid)

	addresses, err := config.GetAddresses(processHandle, uintptr(unsafe.Pointer(address)), offsets)

	go util.NeverExit(func() { hacks.DoBHOP(processHandle, addresses) })

	select {}
}
