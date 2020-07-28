package gohack

import (
	"errors"
	"github.com/jamesmoriarty/gomem"
	"unsafe"
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

	process.Open()

	if client.OffsetPlayer() == 0 {
		return nil, errors.New("Failed to get player offset")
	}

	return client, nil
}

func (a *Client) OffsetForceJump() uintptr {
	return a.Address + a.Offsets.Signatures.OffsetdwForceJump
}

func (a *Client) OffsetPlayer() uintptr {
	var (
		readValue    uintptr
		readValuePtr = (uintptr)(unsafe.Pointer(&readValue))
	)

	a.Process.Read(a.Address + a.Offsets.Signatures.OffsetdwLocalPlayer, readValuePtr, 4)

	return readValue
}

func (a *Client) OffsetPlayerFlags() uintptr {
	return a.OffsetPlayer() + a.Offsets.Netvars.Offsetm_fFlags
}
