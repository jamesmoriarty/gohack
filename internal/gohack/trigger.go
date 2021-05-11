package gohack

import (
	"runtime"
	"time"
	"unsafe"

	"github.com/jamesmoriarty/gomem"
)

func RunTrigger(client *Client) {
	var (
		writeValue    = byte(0x6)
		writeValuePtr = (uintptr)(unsafe.Pointer(&writeValue))
	)

	for {
		if gomem.IsKeyDown(0x10) { // https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
			if client.OffsetEntityId() > 0 && client.OffsetEntityId() <= 64 {
				client.Process.Write(client.OffsetForceAttack(), writeValuePtr, unsafe.Sizeof(writeValue))
			}
		}

		time.Sleep(50)

		// N.B. guard against buffer gc.
		runtime.KeepAlive(&writeValue)
	}

}
