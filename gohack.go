package gohack

import (
	"fmt"
	"errors"
	"github.com/jamesmoriarty/gomem"
	log "github.com/sirupsen/logrus"
	"time"
	"strconv"
	"unsafe"
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

const (
	processName = "csgo.exe"
	moduleName  = "client_panorama.dll"
)

func Instrument() (*gomem.Process, *Addresses, error) {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})

	offsets, err := GetOffsets()
	if err != nil {
		return nil, nil, errors.New("Failed to get offsets " + err.Error())
	}

	process, err := gomem.GetFromProcessName(processName)
	if err != nil {
		return nil, nil, errors.New("Failed to get pid " + processName)
	}
	log.WithFields(log.Fields{"pid": process.ID}).Info("GetFromProcessName ", processName)

	address, err := process.GetModule(moduleName)
	if err != nil {
		return nil, nil, errors.New("Failed to get module address " + moduleName)
	}
	log.WithFields(log.Fields{"address": address}).Info("GetModule ", moduleName)

	process.Open()
	log.WithFields(log.Fields{"handle": process.Handle}).Info("OpenProcess ", process.ID)

	addresses := &Addresses{Process: process, Local: address, Offsets: offsets}
	if addresses.LocalPlayer() == 0 {
		return process, addresses, errors.New("Failed to get LocalPlayer address")
	}
	log.WithFields(log.Fields{"value": ptrToHex(addresses.Local)}).Info("- addressLocal")
	log.WithFields(log.Fields{"value": ptrToHex(addresses.LocalForceJump())}).Info("- addressLocalForceJump")
	log.WithFields(log.Fields{"value": ptrToHex(addresses.LocalPlayer())}).Info("- addressLocalPlayer")
	log.WithFields(log.Fields{"value": ptrToHex(addresses.LocalPlayerFlags())}).Info("- addressLocalPlayerFlags")

	return process, addresses, err
}

func RunBHOP(p *gomem.Process, addresses *Addresses) {
	var (
		readValue     byte
		readValuePtr  = (*uintptr)(unsafe.Pointer(&readValue))
		writeValue    = byte(0x6)
		writeValuePtr = (*uintptr)(unsafe.Pointer(&writeValue))
	)

	for {
		if gomem.IsKeyDown(0x20) { // https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
			p.Read(addresses.LocalPlayerFlags(), readValuePtr, unsafe.Sizeof(readValue))

			if (readValue & (1 << 0)) > 0 { // FL_ONGROUND (1<<0) // https://github.com/ValveSoftware/source-sdk-2013/blob/master/mp/src/public/const.h
				p.Write(addresses.LocalForceJump(), writeValuePtr, unsafe.Sizeof(writeValue))
			}
		}

		time.Sleep(90) // 15ms tick
	}
}

