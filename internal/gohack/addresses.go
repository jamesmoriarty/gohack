package gohack

import (
	"github.com/jamesmoriarty/gomem"
)

type Addresses struct {
	Process *gomem.Process
	Offset   uintptr
	Offsets *Offsets
}

func (a *Addresses) OffsetForceJump() uintptr {
	return a.Offset + a.Offsets.Signatures.OffsetForceJump
}

func (a *Addresses) OffsetPlayer() uintptr {
	var buffer uintptr

	a.Process.Read(a.Offset+a.Offsets.Signatures.OffsetPlayer, &buffer, 4)

	return buffer
}

func (a *Addresses) OffsetPlayerFlags() uintptr {
	return a.OffsetPlayer() + a.Offsets.Netvars.OffsetOffsetPlayerFlags
}
