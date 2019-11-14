package main

import (
	"fmt"
	win32 "github.com/jamesmoriarty/gohack/win32"
	"time"
	"unsafe"
)

func main() {
	var (
		processName          = "csgo.exe"
		moduleName           = "client_panorama.dll"
		constJump            = uintptr(0x6)
		// https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
		vk_space             = 0x20
		// https://github.com/frk1/hazedumper/blob/master/csgo.cs
		offsetM_fFlags       = uintptr(0x104)
		offsetDwLocalPlayer  = uintptr(0xCFAA3C)
		offsetDwForceJump    = uintptr(0x51B0758)
		addressDwLocalPlayer uintptr
		addressDwForceJump   uintptr
		addressBase          uintptr
	)

	pid, success := win32.GetProcessID(processName)
	fmt.Println("Info: GetProcessID", success, pid)
	module, success, address := win32.GetModule(moduleName, pid)
	fmt.Println("Info: GetModule", module.ModBaseSize, success, address)
	process, err := win32.OpenProcess(win32.PROCESS_ALL_ACCESS, false, pid)
	fmt.Println("Info: OpenProcess", process, err)
	addressBase = uintptr(unsafe.Pointer(address))
	fmt.Printf("Info: addressBase %v\n", addressBase)
	addressDwForceJump = addressBase + offsetDwForceJump
	fmt.Printf("Info: addressDwForceJump %v\n", addressDwForceJump)
	win32.ReadProcessMemory(process, win32.LPCVOID(addressBase+offsetDwLocalPlayer), &addressDwLocalPlayer, 4)
	fmt.Printf("Info: addressDwLocalPlayer %v\n", addressDwLocalPlayer)

	var buffer uintptr

	for {
		if win32.GetAsyncKeyState(vk_space) > 0 {
			win32.ReadProcessMemory(process, win32.LPCVOID(addressDwLocalPlayer+offsetM_fFlags), &buffer, 1)

			if buffer != 0 {
				win32.WriteProcessMemory(process, addressDwForceJump, unsafe.Pointer(&constJump), 1)
			}
		}
		time.Sleep(5)
	}
}
