package gohack

import (
	"github.com/jamesmoriarty/gomem"
)

type Addresses struct {
	Process *gomem.Process
	Local   uintptr
	Offsets *Offsets
}

func (a *Addresses) LocalForceJump() uintptr {
	return a.Local + a.Offsets.Signatures.OffsetForceJump
}

func (a *Addresses) LocalPlayer() uintptr {
	var localPlayer uintptr

	a.Process.Read(a.Local+a.Offsets.Signatures.OffsetLocalPlayer, &localPlayer, 4)

	return localPlayer
}

func (a *Addresses) LocalPlayerFlags() uintptr {
	return a.LocalPlayer() + a.Offsets.Netvars.OffsetLocalPlayerFlags
}
