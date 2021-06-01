package gohack

import (
	"time"

	"github.com/jamesmoriarty/gomem"
)

func RunBHOP(client *Client) {
	for {
		if gomem.IsKeyDown(0x20) { // https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
			flags, _ := client.Process.ReadByte(client.OffsetPlayerFlags())

			if (flags & (1 << 0)) > 0 { // FL_ONGROUND (1<<0) // https://github.com/ValveSoftware/source-sdk-2013/blob/master/mp/src/public/const.h
				client.Process.WriteByte(client.OffsetForceJump(), 0x6)
			}
		}

		time.Sleep(100)
	}
}
