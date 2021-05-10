package gohack

import (
	"github.com/jamesmoriarty/gomem"
	"time"
	"runtime"
	"unsafe"
)

func RunBHOP(client *Client) {
	var (
		readValue     = byte(0x0)
		readValuePtr  = (uintptr)(unsafe.Pointer(&readValue))
		writeValue    = byte(0x6)
		writeValuePtr = (uintptr)(unsafe.Pointer(&writeValue))
	)

	for {
		if gomem.IsKeyDown(0x20) { // https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
			client.Process.Read(client.OffsetPlayerFlags(), readValuePtr, unsafe.Sizeof(readValue))

			if (readValue & (1 << 0)) > 0 { // FL_ONGROUND (1<<0) // https://github.com/ValveSoftware/source-sdk-2013/blob/master/mp/src/public/const.h
				client.Process.Write(client.OffsetForceJump(), writeValuePtr, unsafe.Sizeof(writeValue))
			}

			// N.B. writing can silently fails so we need to verify the write. I suspect we might need to re-open the process handle.

			readValue = 0x0

			client.Process.Read(client.OffsetForceJump(), readValuePtr, unsafe.Sizeof(readValuePtr))

			if readValue == 0x0 {
				panic("Write Error" + string(readValue) + " != " + string(writeValue))
			}
		}

		time.Sleep(90)

		// N.B. guard against buffer gc.
		runtime.KeepAlive(&readValue)
		runtime.KeepAlive(&writeValue)
	}
}
