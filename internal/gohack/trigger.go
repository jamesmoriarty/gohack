package gohack

import (
	"time"

	"github.com/jamesmoriarty/gomem"
)

const (
	VK_SHIFT         = 0x10 // https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
	CSGO_FORCEATTACK = 0x6
)

func RunTrigger(client *Client) {
	for {
		if gomem.IsKeyDown(VK_SHIFT) {
			if client.OffsetEntityId() > 0 && client.OffsetEntityId() <= 64 {
				client.Process.WriteByte(client.OffsetForceAttack(), CSGO_FORCEATTACK)
			}
		}

		time.Sleep(50 * time.Millisecond)
	}
}
