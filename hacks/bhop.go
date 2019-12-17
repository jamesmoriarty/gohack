package hacks

import (
	win32 "github.com/jamesmoriarty/gohack/win32"
	"time"
	"unsafe"
)

const (
	vkSpace = 0x20 // https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
)

func DoBHOP(processHandle win32.HANDLE, addressLocalPlayerFlags uintptr, addressLocalForceJump uintptr) {
	var (
		flagsCurrent    uintptr
		playerFlagsJump = uintptr(0x6)
	)

	for {
		if win32.GetAsyncKeyState(vkSpace) > 0 {
			win32.ReadProcessMemory(processHandle, win32.LPCVOID(addressLocalPlayerFlags), &flagsCurrent, 1)

			if flagsCurrent != 0 {
				win32.WriteProcessMemory(processHandle, addressLocalForceJump, unsafe.Pointer(&playerFlagsJump), 1)
			}
		}
		time.Sleep(35)
	}
}
