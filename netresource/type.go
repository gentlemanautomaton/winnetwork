package netresource

// Type is a network resource type.
type Type uint32

// Network resource types.
const (
	TypeAny   = 0 // RESOURCETYPE_ANY
	TypeDisk  = 1 // RESOURCETYPE_DISK
	TypePrint = 2 // RESOURCETYPE_PRINT
)
