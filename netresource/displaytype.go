package netresource

// DisplayType is a network resource display type.
type DisplayType uint32

// Network resource display types.
const (
	DisplayTypeGeneric      = 0  // RESOURCEDISPLAYTYPE_GENERIC
	DisplayTypeDomain       = 1  // RESOURCEDISPLAYTYPE_DOMAIN
	DisplayTypeServer       = 2  // RESOURCEDISPLAYTYPE_SERVER
	DisplayTypeShare        = 3  // RESOURCEDISPLAYTYPE_SHARE
	DisplayTypeFile         = 4  // RESOURCEDISPLAYTYPE_FILE
	DisplayTypeGroup        = 5  // RESOURCEDISPLAYTYPE_GROUP
	DisplayTypeNetwork      = 6  // RESOURCEDISPLAYTYPE_NETWORK
	DisplayTypeRoot         = 7  // RESOURCEDISPLAYTYPE_ROOT
	DisplayTypeShareAdmin   = 8  // RESOURCEDISPLAYTYPE_SHAREADMIN
	DisplayTypeDirectory    = 9  // RESOURCEDISPLAYTYPE_DIRECTORY
	DisplayTypeTree         = 10 // RESOURCEDISPLAYTYPE_TREE
	DisplayTypeNDSContainer = 11 // RESOURCEDISPLAYTYPE_NDSCONTAINER
)
