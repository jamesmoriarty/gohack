package gohack

import (
	"errors"

	"github.com/jamesmoriarty/gomem"
)

type Client struct {
	Process *gomem.Process
	Address uintptr
	Offsets *Offsets
}

func GetClientFrom(process *gomem.Process, offsets *Offsets) (*Client, error) {
	address, err := process.GetModule("client.dll")

	if err != nil {
		return nil, errors.New("Failed to get module offset")
	}

	client := &Client{Process: process, Address: address, Offsets: offsets}

	if client.OffsetPlayer() == 0 {
		return nil, errors.New("Failed to get player offset")
	}

	return client, nil
}

func (a *Client) OffsetForceJump() uintptr {
	return a.Address + a.Offsets.Signatures.OffsetdwForceJump
}

func (a *Client) OffsetForceAttack() uintptr {
	return a.Address + a.Offsets.Signatures.OffsetdwForceAttack
}

func (a *Client) OffsetPlayer() uintptr {
	ptr, _ := a.Process.ReadUInt32(a.Address + a.Offsets.Signatures.OffsetdwLocalPlayer)

	return (uintptr)(ptr)
}

func (a *Client) OffsetPlayerFlags() uintptr {
	return a.OffsetPlayer() + a.Offsets.Netvars.Offsetm_fFlags
}

func (a *Client) OffsetEntityId() uintptr {
	ptr, _ := a.Process.ReadUInt32(a.OffsetPlayer() + a.Offsets.Netvars.Offsetm_iCrosshairId)

	return (uintptr)(ptr)
}
