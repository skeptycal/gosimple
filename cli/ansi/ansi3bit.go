package ansi

// Premade ANSI basic 3 bit ANSI color strings
//
// Reference: https://en.wikipedia.org/wiki/ANSI_escape_code
const (
	BlackText          string = "\033[30m"
	RedText            string = "\033[31m"
	GreenText          string = "\033[32m"
	YellowText         string = "\033[33m"
	BlueText           string = "\033[34m"
	PurpleText         string = "\033[35m"
	CyanText           string = "\033[36m"
	WhiteText          string = "\033[37m"
	DefaultColorText   string = "\033[39m" // Normal foreground color
	BgBlackText        string = "\033[40m"
	BgRedText          string = "\033[41m"
	BgGreenText        string = "\033[42m"
	BgYellowText       string = "\033[43m"
	BgBlueText         string = "\033[44m"
	BgPurpleText       string = "\033[45m"
	BgCyanText         string = "\033[46m"
	BgWhiteText        string = "\033[47m"
	BhDefaultColorText string = "\033[49m" // Normal background color
)

// Premade bold and dim 3 bit ANSI color strings
//
// Reference: https://en.wikipedia.org/wiki/ANSI_escape_code
const (
	BoldText         string = "\033[1m"
	BoldBlackText    string = "\033[1;30m"
	BoldRedText      string = "\033[1;31m"
	BoldGreenText    string = "\033[1;32m"
	BoldYellowText   string = "\033[1;33m"
	BoldBlueText     string = "\033[1;34m"
	BoldMagentaText  string = "\033[1;35m"
	BoldCyanText     string = "\033[1;36m"
	BoldWhiteText    string = "\033[1;37m"
	FaintText        string = "\033[2m"
	FaintBlackText   string = "\033[2;30m"
	FaintRedText     string = "\033[2;31m"
	FaintGreenText   string = "\033[2;32m"
	FaintYellowText  string = "\033[2;33m"
	FaintBlueText    string = "\033[2;34m"
	FaintMagentaText string = "\033[2;35m"
	FaintCyanText    string = "\033[2;36m"
	FaintWhiteText   string = "\033[2;37m"
)

// All possible colors (3 bit ANSI)
const (
	BLACK = iota
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)
