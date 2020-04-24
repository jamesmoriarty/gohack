package gohack

import (
	"errors"
	"github.com/jamesmoriarty/gomem"
)

type Client struct {
	Process *gomem.Process
	Offset  uintptr
	Offsets *Offsets
}

func ClientFrom(process *gomem.Process, offsets *Offsets) (*Client, error) {
	offset, err := process.GetModule("client_panorama.dll")

	if err != nil {
		return nil, errors.New("Failed to get offset")
	}

	client := &Client{Process: process, Offset: offset, Offsets: offsets}

	process.Open()

	if client.OffsetPlayer() == 0 {
		return nil, errors.New("Failed to get OffsetPlayer")
	}

	return client, nil
}

func (a *Client) OffsetForceJump() uintptr {
	return a.Offset + a.Offsets.Signatures.OffsetForceJump
}

func (a *Client) OffsetPlayer() uintptr {
	var buffer uintptr

	a.Process.Read(a.Offset+a.Offsets.Signatures.OffsetPlayer, &buffer, 4)

	return buffer
}

func (a *Client) OffsetPlayerFlags() uintptr {
	return a.OffsetPlayer() + a.Offsets.Netvars.OffsetOffsetPlayerFlags
}
