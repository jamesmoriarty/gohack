package gohack

import (
	"time"

	"github.com/jamesmoriarty/gomem"
)

const (
	VK_SPACE         = 0x20   // https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
	CSGO_FL_ONGROUND = 1 << 0 // https://github.com/ValveSoftware/source-sdk-2013/blob/master/mp/src/public/const.h
	CSGO_FORCEJUMP   = 0x6    // https://github.com/ValveSoftware/source-sdk-2013/blob/0d8dceea4310fde5706b3ce1c70609d72a38efdf/sp/src/game/shared/sdk/sdk_playeranimstate.cpp#L517
)

func RunHop(client *Client) {
	for {
		if gomem.IsKeyDown(VK_SPACE) {
			flags, _ := client.Process.ReadByte(client.OffsetPlayerFlags())

			if (flags & CSGO_FL_ONGROUND) > 0 {
				client.Process.WriteByte(client.OffsetForceJump(), CSGO_FORCEJUMP)
			}
		}

		time.Sleep(1 * time.Microsecond)
	}
}
