package hacks

import (
	config "github.com/jamesmoriarty/gohack/config"
	win32 "github.com/jamesmoriarty/gohack/win32"
	"time"
	"unsafe"
)

const (
	vkSpace = 0x20 // https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
)

func DoBHOP(processHandle win32.HANDLE, addresses *config.Addresses) {
	var (
		flagsCurrent       uintptr
		playerFlagsJump    = uintptr(0x6)
		playerFlagsJumpPtr = unsafe.Pointer(&playerFlagsJump)
	)

	for {
		if win32.GetAsyncKeyState(vkSpace) > 0 {
			win32.ReadProcessMemory(processHandle, win32.LPCVOID(addresses.LocalPlayerFlags), &flagsCurrent, 1)

			if flagsCurrent != 0 {
				win32.WriteProcessMemory(processHandle, addresses.LocalForceJump, playerFlagsJumpPtr, 1)
			}
		}
		time.Sleep(35)
	}
}
