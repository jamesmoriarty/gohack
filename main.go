package main

import (
	config "github.com/jamesmoriarty/gohack/config"
	util "github.com/jamesmoriarty/gohack/util"
	win32 "github.com/jamesmoriarty/gohack/win32"
	log "github.com/sirupsen/logrus"
	"os"
	"unsafe"
)

const (
	url         = "https://raw.githubusercontent.com/frk1/hazedumper/master/csgo.yaml"
	processName = "csgo.exe"
	moduleName  = "client_panorama.dll"
)

func main() {
	// Dynamic adresses
	var (
		addressLocal            uintptr
		addressLocalForceJump   uintptr
		addressLocalPlayer      uintptr
		addressLocalPlayerFlags uintptr
	)

	log.SetFormatter(&log.TextFormatter{ForceColors: true})

	config.PrintBanner()

	log.WithFields(log.Fields{"url": url}).Info("GetLatestOffsets")
	offsets, err := config.GetLatestOffsets(url)
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

	addressLocal = uintptr(unsafe.Pointer(address))
	log.WithFields(log.Fields{"value": util.ConvertPtrToHex(addressLocal)}).Info("- addressLocal")

	addressLocalForceJump = addressLocal + offsets.Signatures.OffsetForceJump
	log.WithFields(log.Fields{"value": util.ConvertPtrToHex(addressLocalForceJump)}).Info("- addressLocalForceJump")

	win32.ReadProcessMemory(processHandle, win32.LPCVOID(addressLocal+offsets.Signatures.OffsetLocalPlayer), &addressLocalPlayer, 4)
	log.WithFields(log.Fields{"value": util.ConvertPtrToHex(addressLocalPlayer)}).Info("- addressLocalPlayer")

	addressLocalPlayerFlags = addressLocalPlayer + offsets.Netvars.OffsetLocalPlayerFlags
	log.WithFields(log.Fields{"value": util.ConvertPtrToHex(addressLocalPlayerFlags)}).Info("- addressLocalPlayerFlags")

	go util.NeverExit(func() { util.DoBHOP(processHandle, addressLocalPlayerFlags, addressLocalForceJump) })

	select {}
}
