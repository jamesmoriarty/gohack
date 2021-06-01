package gohack

import (
	"time"

	"github.com/jamesmoriarty/gomem"
)

func RunTrigger(client *Client) {
	for {
		if gomem.IsKeyDown(0x10) { // https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
			if client.OffsetEntityId() > 0 && client.OffsetEntityId() <= 64 {
				client.Process.WriteByte(client.OffsetForceAttack(), 0x6)
			}
		}

		time.Sleep(50)
	}

}
