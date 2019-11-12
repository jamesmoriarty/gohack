package win32

import "unsafe"

type (
	//BOOL      int32
	//DWORD     uint32
	//ULONG_PTR uintptr
	HANDLE uintptr
	//LPVOID    unsafe.Pointer
	LPCVOID unsafe.Pointer
	//SIZE_T    uintptr
	HMODULE HANDLE
	//BYTE      byte
)

// https://msdn.microsoft.com/9e2f7345-52bf-4bfc-9761-90b0b374c727
type PROCESSENTRY32 struct {
	DwSize              uint32
	CntUsage            uint32
	Th32ProcessID       uint32
	Th32DefaultHeapID   uintptr
	Th32ModuleID        uint32
	CntThreads          uint32
	Th32ParentProcessID uint32
	PcPriClassBase      uint32
	DwFlags             uint32
	SzExeFile           [MAX_PATH]uint8
}

// https://msdn.microsoft.com/305fab35-625c-42e3-a434-e2513e4c8870
type MODULEENTRY32 struct {
	DwSize        uint32
	Th32ModuleID  uint32
	Th32ProcessID uint32
	GlblcntUsage  uint32
	ProccntUsage  uint32
	ModBaseAddr   *uintptr
	ModBaseSize   uint32
	HModule       HMODULE
	SzModule      [MAX_MODULE_NAME32 + 1]uint8
	SzExePath     [MAX_PATH]uint8
}
