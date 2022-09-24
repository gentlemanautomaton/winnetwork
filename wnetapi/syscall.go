package wnetapi

import (
	"syscall"
	"unsafe"

	"github.com/gentlemanautomaton/winnetwork/netresource"
	"golang.org/x/sys/windows"
)

var (
	modmpr = windows.NewLazySystemDLL("Mpr.dll")

	procWNetGetConnectionW     = modmpr.NewProc("WNetGetConnectionW")
	procWNetAddConnection2W    = modmpr.NewProc("WNetAddConnection2W")
	procWNetCancelConnection2W = modmpr.NewProc("WNetCancelConnection2W")
)

// GetConnection returns the remote resource that the given local name
// is mapped to. It calls the WNetGetConnectionW windows API function.
//
// https://learn.microsoft.com/en-us/windows/win32/api/winnetwk/nf-winnetwk-wnetgetconnectionw
func GetConnection(local string) (remote string, err error) {
	utf16Local, err := syscall.UTF16PtrFromString(local)
	if err != nil {
		return "", err
	}

	var utf16RemoteBuffer [4096]uint16
	length := uint32(len(utf16RemoteBuffer))

	r0, _, e := syscall.SyscallN(
		procWNetGetConnectionW.Addr(),
		uintptr(unsafe.Pointer(utf16Local)),
		uintptr(unsafe.Pointer(&utf16RemoteBuffer[0])),
		uintptr(unsafe.Pointer(&length)))
	if r0 != 0 {
		return "", e
	}

	return syscall.UTF16ToString(utf16RemoteBuffer[:length]), err
}

// AddConnection2 attempts to map the given local name to a remote
// resource. If persistent is true it creates a persistent connection
// that survives logon/logoff cycles. It calls the WNetAddConnection2
// windows API function.
//
// https://learn.microsoft.com/en-us/windows/win32/api/winnetwk/nf-winnetwk-wnetaddconnection2w
func AddConnection2(userName, password string, flags netresource.Option, resource netresource.Data) (err error) {
	var utf16UserName, utf16Password *uint16
	if userName != "" {
		if utf16UserName, err = syscall.UTF16PtrFromString(userName); err != nil {
			return err
		}
	}

	if password != "" {
		if utf16Password, err = syscall.UTF16PtrFromString(password); err != nil {
			return err
		}
	}

	raw, err := resource.Raw()
	if err != nil {
		return err
	}

	r0, _, e := syscall.SyscallN(
		procWNetAddConnection2W.Addr(),
		uintptr(unsafe.Pointer(&raw)),
		uintptr(unsafe.Pointer(utf16UserName)),
		uintptr(unsafe.Pointer(utf16Password)),
		uintptr(flags))
	if r0 != 0 {
		return e
	}

	return nil
}

// CancelConnection2 attempts to unmap the given local resource name that
// has previously been mapped to a remote resource. It calls the
// WNetCancelConnection2W windows API function.
//
// If the ConnectUpdateProfile flag is provided, the connection will be
// permanently removed from the user profile.
//
// If force is true, it will forcefully close open files if necessary.
//
// https://learn.microsoft.com/en-us/windows/win32/api/winnetwk/nf-winnetwk-wnetcancelconnection2w
func CancelConnection2(name string, flags netresource.Option, force bool) (err error) {
	var utf16Name *uint16
	if name != "" {
		if utf16Name, err = syscall.UTF16PtrFromString(name); err != nil {
			return err
		}
	}

	var uintForce uintptr
	if force {
		uintForce = 1
	}

	r0, _, e := syscall.SyscallN(
		procWNetCancelConnection2W.Addr(),
		uintptr(unsafe.Pointer(utf16Name)),
		uintptr(flags),
		uintptr(uintForce))
	if r0 != 0 {
		return e
	}

	return nil
}
