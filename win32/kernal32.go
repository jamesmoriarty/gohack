package win32

import (
	"log"
	"syscall"
	"unsafe"
)

var (
	kernel32                     = syscall.MustLoadDLL("kernel32.dll")
	procCloseHandle              = kernel32.MustFindProc("CloseHandle")
	procCreateToolhelp32Snapshot = kernel32.MustFindProc("CreateToolhelp32Snapshot")
	procGetLastError             = kernel32.MustFindProc("GetLastError")
	procGetModuleHandle          = kernel32.MustFindProc("GetModuleHandleW")
	procProcess32First           = kernel32.MustFindProc("Process32First")
	procProcess32Next            = kernel32.MustFindProc("Process32Next")
	procModule32First            = kernel32.MustFindProc("Module32First")
	procModule32Next             = kernel32.MustFindProc("Module32Next")
	procOpenProcess              = kernel32.MustFindProc("OpenProcess")
	procReadProcessMemory        = kernel32.MustFindProc("ReadProcessMemory")
	procWriteProcessMemory       = kernel32.MustFindProc("WriteProcessMemory")
	psapi                        = syscall.MustLoadDLL("psapi.dll") //kern32 didnt work
	procEnumProcessModules       = psapi.MustFindProc("EnumProcessModules")
)

// https://msdn.microsoft.com/b4088506-2f69-4cf0-9bab-3e6a7185f5b2
func EnumProcessModules(hProcess HANDLE, cb uintptr, lpcbNeeded uintptr) (uintptr, []uint16, error) {

	defer func() {
		log.Println("done") // Println executes normally even if there is a panic
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()

	lphModuleBuffer := make([]uint16, uintptr(lpcbNeeded))

	ret, _, err := procEnumProcessModules.Call(
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&lphModuleBuffer[0])),
		uintptr(cb),
		uintptr(lpcbNeeded),
	)

	return ret, lphModuleBuffer, err
}

// https://msdn.microsoft.com/9cd91f1c-58ce-4adc-b027-45748543eb06
func WriteProcessMemory(hProcess HANDLE, lpBaseAddress uintptr, lpBuffer unsafe.Pointer, nSize uintptr) (uintptr, error) {
	ret, _, err := procWriteProcessMemory.Call(
		uintptr(hProcess),
		uintptr(lpBaseAddress),
		uintptr(lpBuffer),
		uintptr(nSize),
	)
	return ret, err
}

// https://msdn.microsoft.com/8774e145-ee7f-44de-85db-0445b905f986
func ReadProcessMemory(hProcess HANDLE, lpBaseAddress LPCVOID, lpBuffer *uintptr, nSize uintptr) (uintptr, error) {
	ret, _, str := procReadProcessMemory.Call(
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpBaseAddress)),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(nSize),
		0,
	)
	return ret, str
}

// https://msdn.microsoft.com/df643c25-7558-424c-b187-b3f86ba51358
func CreateToolhelp32Snapshot(dwFlags uintptr, th32ProcessID uint32) HANDLE {
	ret, _, _ := procCreateToolhelp32Snapshot.Call(
		uintptr(dwFlags),
		uintptr(th32ProcessID),
	)
	return HANDLE(ret)
}

//ret != 0 is a bool itself returning true or false

// https://msdn.microsoft.com/097790e8-30c2-4b00-9256-fa26e2ceb893
func Process32First(hSnapshot HANDLE, pe *PROCESSENTRY32) bool {
	ret, _, _ := procProcess32First.Call(
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(pe)),
	)
	return ret != 0
}

// https://msdn.microsoft.com/843a95fd-27ae-4215-83d0-82fc402b82b6
func Process32Next(hSnapshot HANDLE, pe *PROCESSENTRY32) bool {
	ret, _, _ := procProcess32Next.Call(
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(pe)),
	)
	return ret != 0
}

// https://msdn.microsoft.com/bb41cab9-13a1-469d-bf76-68c172e982f6
func Module32First(hSnapshot HANDLE, me *MODULEENTRY32) bool {
	ret, _, _ := procModule32First.Call(
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(me)),
	)
	return ret != 0
}

// https://msdn.microsoft.com/88ec1af4-bae7-4cd7-b830-97a98fb337f4
func Module32Next(hSnapshot HANDLE, me *MODULEENTRY32) bool {
	ret, _, _ := procModule32Next.Call(
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(me)),
	)
	return ret != 0
}

// https://msdn.microsoft.com/29514410-89fe-4888-8b34-0c30d5af237f
func GetModuleHandle(lpModuleName string) HMODULE {
	ret, _, _ := procGetModuleHandle.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpModuleName))),
	)
	return HMODULE(ret)
}

// https://msdn.microsoft.com/8f695c38-19c4-49e4-97de-8b64ea536cb1
func OpenProcess(dwDesiredAccess uint32, bInheritHandle bool, dwProcessId uint32) (HANDLE, error) {
	inHandle := 0
	if bInheritHandle {
		inHandle = 1
	}

	ret, _, err := procOpenProcess.Call(
		uintptr(dwDesiredAccess),
		uintptr(inHandle),
		uintptr(dwProcessId),
	)
	return HANDLE(ret), err
}

// https://msdn.microsoft.com/9b84891d-62ca-4ddc-97b7-c4c79482abd9
func CloseHandle(hObject HANDLE) bool {
	ret, _, _ := procCloseHandle.Call(
		uintptr(hObject),
	)
	return ret != 0
}

func GetLastError() uint32 {
	ret, _, _ := procGetLastError.Call()
	return uint32(ret)
}
