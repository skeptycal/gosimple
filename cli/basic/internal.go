package basic

import "strings"

var (
	log          = Log // alias for local logging
	headerChar   = "*"
	footerChar   = "-"
	headerBorder = headerString()
	footerBorder = footerString()
)

func headerString() string {
	return strings.Repeat(headerChar, COLUMNS/len(headerChar))
}

func footerString() string {
	return strings.Repeat(headerChar, COLUMNS/len(headerChar))
}
