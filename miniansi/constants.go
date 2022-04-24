package miniansi

const (
	NewLine      = "\n"
	InfoColor    = ansiPrefix + "1;34m"
	NoticeColor  = ansiPrefix + "1;36m"
	WarningColor = ansiPrefix + "1;33m"
	ErrorColor   = ansiPrefix + "1;31m"
	DebugColor   = ansiPrefix + "0;36m"
	DbColor      = ansiPrefix + "1;31m" // ANSI dbecho code
	ResetColor   = ansiPrefix + "0m"    // ANSI reset code

	// semicolon delimited ANSI codes for %v
	// string to print for %s
	ansiFmt = "\033[%vm%s\033[0m"

	ansiSEP    = ";"
	ansiPrefix = "\033["
	ansiSuffix = "m"
)

var (
	bReset = []byte(ResetColor)
)
