package netresource

// Option is a bitmask that can contain network resource connection options.
type Option uint32

// Network resource connection options.
const (
	ConnectUpdateProfile    = 1 << 0  // CONNECT_UPDATE_PROFILE
	ConnectUpdateRecent     = 1 << 1  // CONNECT_UPDATE_RECENT
	ConnectTemporary        = 1 << 2  // CONNECT_TEMPORARY
	ConnectInteractive      = 1 << 3  // CONNECT_INTERACTIVE
	ConnectPrompt           = 1 << 4  // CONNECT_PROMPT
	ConnectRedirect         = 1 << 7  // CONNECT_REDIRECT
	ConnectCurrentMedia     = 1 << 9  // CONNECT_CURRENT_MEDIA
	ConnectCommandLine      = 1 << 11 // CONNECT_COMMANDLINE
	ConnectSaveCredentials  = 1 << 12 // CONNECT_CMD_SAVECRED
	ConnectResetCredentials = 1 << 13 // CONNECT_CRED_RESET
)
