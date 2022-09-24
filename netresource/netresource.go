package netresource

import (
	"syscall"
)

// Data is used in wnetapi calls to describe network resources.
//
// https://learn.microsoft.com/en-us/windows/win32/api/winnetwk/ns-winnetwk-netresourcew
type Data struct {
	Scope       Scope
	Type        Type
	DisplayType DisplayType
	Usage       Usage
	LocalName   string
	RemoteName  string
	Comment     string
	Provider    string
}

// Raw returns the network resource as raw data ready to be used in API calls.
func (nr Data) Raw() (out RawData, err error) {
	var localName, remoteName, comment, provider *uint16
	if nr.LocalName != "" {
		if localName, err = syscall.UTF16PtrFromString(nr.LocalName); err != nil {
			return
		}
	}
	if nr.RemoteName != "" {
		if remoteName, err = syscall.UTF16PtrFromString(nr.RemoteName); err != nil {
			return
		}
	}
	if nr.Comment != "" {
		if comment, err = syscall.UTF16PtrFromString(nr.Comment); err != nil {
			return
		}
	}
	if nr.Provider != "" {
		if provider, err = syscall.UTF16PtrFromString(nr.Provider); err != nil {
			return
		}
	}
	return RawData{
		Scope:       nr.Scope,
		Type:        nr.Type,
		DisplayType: nr.DisplayType,
		Usage:       nr.Usage,
		LocalName:   localName,
		RemoteName:  remoteName,
		Comment:     comment,
		Provider:    provider,
	}, nil
}

// Data is the raw form of data used in wnetapi calls to describe network
// resources.
type RawData struct {
	Scope       Scope
	Type        Type
	DisplayType DisplayType
	Usage       Usage
	LocalName   *uint16
	RemoteName  *uint16
	Comment     *uint16
	Provider    *uint16
}
