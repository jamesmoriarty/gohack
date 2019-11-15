package main

import (
	"fmt"
	win32 "github.com/jamesmoriarty/gohack/win32"
	"time"
	"unsafe"
)

func main() {
	var (
		processName            = "csgo.exe"
		moduleName             = "client_panorama.dll"
		// Flags
		flagsJump              = uintptr(0x6)
		// https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
		vk_space               = 0x20
		// https://github.com/frk1/hazedumper/blob/master/csgo.cs
		offsetLocalPlayerFlags = uintptr(0x104)
		offsetLocalPlayer      = uintptr(0xCFAA3C)
		offsetForceJump        = uintptr(0x51B0758)
		// Dynamic adresses
		addressLocalPlayer       uintptr
		addressForceJump         uintptr
		addressLocalPlayerFlags  uintptr
		addressBase              uintptr
	)

	pid, success := win32.GetProcessID(processName)
	fmt.Println("Info: GetProcessID", success, pid)
	module, success, address := win32.GetModule(moduleName, pid)
	fmt.Println("Info: GetModule", module.ModBaseSize, success, address)
	process, err := win32.OpenProcess(win32.PROCESS_ALL_ACCESS, false, pid)
	fmt.Println("Info: OpenProcess", process, err)
	addressBase = uintptr(unsafe.Pointer(address))
	fmt.Printf("Info: addressBase %v\n", addressBase)
	addressForceJump = addressBase + offsetForceJump
	fmt.Printf("Info: addressForceJump %v\n", addressForceJump)
	win32.ReadProcessMemory(process, win32.LPCVOID(addressBase+offsetLocalPlayer), &addressLocalPlayer, 4)
	fmt.Printf("Info: addressLocalPlayer %v\n", addressLocalPlayer)
	addressLocalPlayerFlags = addressLocalPlayer+offsetLocalPlayerFlags
	fmt.Printf("Info: addressLocalPlayerFlags %v\n", addressLocalPlayerFlags)
	
	var flagsCurrent uintptr

	for {
		if win32.GetAsyncKeyState(vk_space) > 0 {
			win32.ReadProcessMemory(process, win32.LPCVOID(addressLocalPlayerFlags), &flagsCurrent, 1)

			if flagsCurrent != 0 {
				win32.WriteProcessMemory(process, addressForceJump, unsafe.Pointer(&flagsJump), 1)
			}
		}
		time.Sleep(5)
	}
}
