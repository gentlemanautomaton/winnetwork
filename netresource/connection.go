package netresource

import (
	"github.com/gentlemanautomaton/winnetwork/wnetapi"
)

// GetConnection returns the remote resource for the local drive letter.
func GetConnection(local string) (remote string, err error) {
	letter, err := MakeDriveLetter(local)
	if err != nil {
		return "", err
	}
	return wnetapi.WNetGetConnection(letter + ":")
}
