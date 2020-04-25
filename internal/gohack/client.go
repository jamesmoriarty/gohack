package gohack

import (
	"errors"
	"github.com/jamesmoriarty/gomem"
	"unsafe"
)

type Client struct {
	Process *gomem.Process
	Offset  uintptr
	Offsets *Offsets
}

func ClientFrom(process *gomem.Process, offsets *Offsets) (*Client, error) {
	offset, err := process.GetModule("client_panorama.dll")

	if err != nil {
		return nil, errors.New("Failed to get module offset")
	}

	client := &Client{Process: process, Offset: offset, Offsets: offsets}

	process.Open()

	if client.OffsetPlayer() == 0 {
		return nil, errors.New("Failed to get player offset")
	}

	return client, nil
}

func (a *Client) OffsetForceJump() uintptr {
	return a.Offset + a.Offsets.Signatures.OffsetForceJump
}

func (a *Client) OffsetPlayer() uintptr {
	var (
		readValue    uintptr
		readValuePtr = (uintptr)(unsafe.Pointer(&readValue))
	)

	a.Process.Read(a.Offset+a.Offsets.Signatures.OffsetPlayer, readValuePtr, 4)

	return readValue
}

func (a *Client) OffsetPlayerFlags() uintptr {
	return a.OffsetPlayer() + a.Offsets.Netvars.OffsetOffsetPlayerFlags
}
