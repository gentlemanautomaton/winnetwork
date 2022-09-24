package winnetwork

import (
	"github.com/gentlemanautomaton/winnetwork/netresource"
	"github.com/gentlemanautomaton/winnetwork/wnetapi"
)

// GetConnection returns the remote resource for the local drive letter.
//
// Connections are unique to each user.
func GetConnection(local string) (remote string, err error) {
	letter, err := netresource.MakeDriveLetter(local)
	if err != nil {
		return "", err
	}
	return wnetapi.GetConnection(letter + ":")
}

// AddConnection attempts to map the given local name to a remote resource.
// If persistent is true it creates a persistent connection that survives
// logon/logoff cycles.
//
// Connections are unique to each user.
func AddConnection(local, remote string, persistent bool) (err error) {
	local, err = netresource.MakeDriveLetter(local)
	if err != nil {
		return err
	}
	local += ":"

	flags := netresource.Option(netresource.ConnectRedirect)
	if persistent {
		flags |= netresource.ConnectUpdateProfile
	} else {
		flags |= netresource.ConnectTemporary
	}

	return wnetapi.AddConnection2("", "", flags, netresource.Data{
		Type:       netresource.TypeDisk,
		LocalName:  local,
		RemoteName: remote,
	})
}

// RemoveConnection attempts to unmap the given local resource name that
// has previously been mapped to a remote resource.
//
// If persistent is true, it permanently removes the connection from the user
// profile.
//
// If force is true, it will forcefully close open files if necessary.
//
// Connections are unique to each user.
func RemoveConnection(name string, persistent, force bool) (err error) {
	var flags netresource.Option
	if persistent {
		flags |= netresource.ConnectUpdateProfile
	}

	return wnetapi.CancelConnection2(name, flags, force)
}
