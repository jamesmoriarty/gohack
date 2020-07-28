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

func (c *Client) OffsetForceJump() uintptr {
	return c.Address + c.Offsets.Signatures.dwForceJump
}

func (c *Client) OffsetPlayer() uintptr {
	var (
		readValue    uintptr
		readValuePtr = (uintptr)(unsafe.Pointer(&readValue))
	)

	c.Process.Read(c.Address + c.Offsets.Signatures.dwLocalPlayer, readValuePtr, 4)

	return readValue
}

func (c *Client) OffsetPlayerFlags() uintptr {
	return c.OffsetPlayer() + c.Offsets.Netvars.m_fFlags
}
