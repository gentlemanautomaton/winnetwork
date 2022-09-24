package netresource

// Scope is a network resource scope.
type Scope uint32

// Network resource scopes.
const (
	ScopeConnected  = 1 // RESOURCE_CONNECTED
	ScopeGlobalNet  = 2 // RESOURCE_GLOBALNET
	ScopeRemembered = 3 // RESOURCE_REMEMBERED
	ScopeRecent     = 4 // RESOURCE_RECENT
	ScopeContext    = 5 // RESOURCE_CONTEXT
)
