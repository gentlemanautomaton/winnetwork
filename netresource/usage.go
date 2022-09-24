package netresource

// Usage is a bitmask that provides guidance for how to enumerate network
// resources.
type Usage uint32

// Network resource usage.
const (
	UsageConnectable   = 1 << 0 // RESOURCEUSAGE_CONNECTABLE
	UsageContainer     = 1 << 1 // RESOURCEUSAGE_CONTAINER
	UsageNoLocalDevice = 1 << 2 // RESOURCEUSAGE_NOLOCALDEVICE
	UsageSibling       = 1 << 3 // RESOURCEUSAGE_SIBLING
	UsageAttached      = 1 << 4 // RESOURCEUSAGE_ATTACHED
)
