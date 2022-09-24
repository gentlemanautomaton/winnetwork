package netresource

import "syscall"

// Possible network resource errors.
var (
	ErrNotConnected    = syscall.Errno(2250) // ERROR_NOT_CONNECTED
	ErrAlreadyAssigned = syscall.Errno(85)   // ERROR_ALREADY_ASSIGNED
)
