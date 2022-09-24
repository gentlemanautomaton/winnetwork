package wnetapi

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	modmpr = windows.NewLazySystemDLL("Mpr.dll")

	procWNetGetConnectionW = modmpr.NewProc("WNetGetConnectionW")
)

// WNetGetConnection returns the remote resource that the given local name
// is mapped to. It calls the WNetGetConnectionW windows API function.
//
// https://learn.microsoft.com/en-us/windows/win32/api/winnetwk/nf-winnetwk-wnetgetconnectionw
func WNetGetConnection(local string) (remote string, err error) {
	utf16Local, err := syscall.UTF16PtrFromString(local)
	if err != nil {
		return "", err
	}

	var utf16RemoteBuffer [4096]uint16
	length := uint32(len(utf16RemoteBuffer))

	r0, _, _ := syscall.SyscallN(
		procWNetGetConnectionW.Addr(),
		uintptr(unsafe.Pointer(utf16Local)),
		uintptr(unsafe.Pointer(&utf16RemoteBuffer[0])),
		uintptr(unsafe.Pointer(&length)))
	if r0 != 0 {
		return "", syscall.Errno(r0)
	}

	return syscall.UTF16ToString(utf16RemoteBuffer[:length]), err
}
