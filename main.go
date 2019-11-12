package main

import (
	"fmt"
	win32 "github.com/jamesmoriarty/gohack/win32"
	"time"
	"unsafe"
)

func main() {
	var (
		processName         = "csgo.exe"
		moduleName          = "client_panorama.dll"
		constJump           = uintptr(0x6)
		// https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
		vk_space            = 0x20
		// https://github.com/frk1/hazedumper/blob/master/csgo.cs
		offsetM_fFlags      = uintptr(0x104)
		offsetDwLocalPlayer = uintptr(0xCFAA3C)
		offsetDwForceJump   = uintptr(0x51B0758)
		addressDwForceJump  uintptr
		addressBase         uintptr
		nSize               uint32
	)

	pid, success := win32.GetProcessID(processName)
	fmt.Println("Info: GetProcessID", success, pid)
	module, success, addr := win32.GetModule(moduleName, pid)
	fmt.Println("Info: GetModule", module.ModBaseSize, success, addr)
	process, err := win32.OpenProcess(win32.PROCESS_ALL_ACCESS, false, pid)
	fmt.Println("Info: OpenProcess", process, err)

	addressBase = uintptr(unsafe.Pointer(addr))
	addressDwForceJump = addressBase + offsetDwForceJump

	win32.ReadProcessMemory(process, win32.LPCVOID(addressBase+offsetDwLocalPlayer), &addressBase, unsafe.Sizeof(nSize))

	var buffer uintptr

	for {
		if win32.GetAsyncKeyState(vk_space) > 0 {
			win32.ReadProcessMemory(process, win32.LPCVOID(addressBase+offsetM_fFlags), &buffer, 1)

			if buffer != 0 {
				win32.WriteProcessMemory(process, addressDwForceJump, unsafe.Pointer(&constJump), 1)
			}
		}
		time.Sleep(5)
	}
}
