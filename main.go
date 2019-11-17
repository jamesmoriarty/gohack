package main

import (
	"fmt"
	win32 "github.com/jamesmoriarty/gohack/win32"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
	"unsafe"
)

var (
	version string
	date    string
	banner  = `
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

func convertPtrToHex(ptr uintptr) string {
	s := fmt.Sprintf("%d", ptr)
	n, _ := strconv.Atoi(s)
	h := fmt.Sprintf("0x%x", n)
	return h
}

func main() {
	fmt.Printf(banner, version, date)
	fmt.Println()

	log.SetFormatter(&log.TextFormatter{ForceColors: true})

	var (
		// Constants
		PROCESSNAME = "csgo.exe"
		MODULENAME  = "client_panorama.dll"
		VKSPACE     = 0x20 // https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
		// Player flags https://github.com/ValveSoftware/source-sdk-2013/blob/master/mp/src/public/const.h#L147
		playerFlagsJump = uintptr(0x6)
		// Offsets https://github.com/frk1/hazedumper/blob/master/csgo.cs
		offsetLocalPlayerFlags = uintptr(0x104)
		offsetLocalPlayer      = uintptr(0xCFAA3C)
		offsetForceJump        = uintptr(0x51B0758)
		// Dynamic adresses
		addressLocal            uintptr
		addressLocalForceJump   uintptr
		addressLocalPlayer      uintptr
		addressLocalPlayerFlags uintptr
	)

	pid, success := win32.GetProcessID(PROCESSNAME)
	log.WithFields(log.Fields{"pid": pid}).Info("GetProcessID ", PROCESSNAME)
	if !success {
		log.Fatal("Failed to get pid ", PROCESSNAME)
		os.Exit(1)
	}

	_, success, address := win32.GetModule(MODULENAME, pid)
	log.WithFields(log.Fields{"address": address}).Info("GetModule ", MODULENAME)
	if !success {
		log.Fatal("Failed to get module address ", MODULENAME)
		os.Exit(1)
	}

	processHandle, _ := win32.OpenProcess(win32.PROCESS_ALL_ACCESS, false, pid)
	log.WithFields(log.Fields{"processHandle": processHandle}).Info("OpenProcess ", pid)

	addressLocal = uintptr(unsafe.Pointer(address))
	log.WithFields(log.Fields{"value": convertPtrToHex(addressLocal)}).Info("- addressLocal")

	addressLocalForceJump = addressLocal + offsetForceJump
	log.WithFields(log.Fields{"value": convertPtrToHex(addressLocalForceJump)}).Info("- addressLocalForceJump")

	win32.ReadProcessMemory(processHandle, win32.LPCVOID(addressLocal+offsetLocalPlayer), &addressLocalPlayer, 4)
	log.WithFields(log.Fields{"value": convertPtrToHex(addressLocalPlayer)}).Info("- addressLocalPlayer")

	addressLocalPlayerFlags = addressLocalPlayer + offsetLocalPlayerFlags
	log.WithFields(log.Fields{"value": convertPtrToHex(addressLocalPlayerFlags)}).Info("- addressLocalPlayerFlags")

	var flagsCurrent uintptr

	for {
		if win32.GetAsyncKeyState(VKSPACE) > 0 {
			win32.ReadProcessMemory(processHandle, win32.LPCVOID(addressLocalPlayerFlags), &flagsCurrent, 1)

			if flagsCurrent != 0 {
				win32.WriteProcessMemory(processHandle, addressLocalForceJump, unsafe.Pointer(&playerFlagsJump), 1)
			}
		}
		time.Sleep(5)
	}
}
