package gohack

import (
	"github.com/jamesmoriarty/gomem"
	"time"
	"unsafe"
)

func RunBHOP(client *Client) {
	var (
		readValue     byte
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
		}

		time.Sleep(90) // 15ms tick
	}
}
