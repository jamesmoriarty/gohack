package win32

import (
	"syscall"
)

var (
	moduser32            = syscall.MustLoadDLL("user32.dll")
	procGetAsyncKeyState = moduser32.MustFindProc("GetAsyncKeyState")
)

func GetAsyncKeyState(vKey int) uint16 {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(vKey))
	return uint16(ret)
}
