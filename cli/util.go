package cli

import (
	"io"

	"github.com/buger/goterm"
	"github.com/skeptycal/gosimple/cli/terminal"
)

// var GetWinSize = terminal.GetWinsize

func CheckIfTerminal(w io.Writer) bool {
	// TODO not working ... see issues
	return terminal.CheckIfTerminal(w)
}

// Columns returns the number of columns in the terminal,
// similar to the COLUMNS environment variable on macOS
// and Linux systems.
func Cols() int {
	return goterm.Width()
}

// Rows returns the number of rows in the terminal,
func Rows() int {
	return goterm.Height()
}

func XPixels() int {
	ws := terminal.GetWinSize()

	if ws == nil {
		return -1
	}

	return ws.Xpixel()
}

func YPixels() int {
	ws := terminal.GetWinSize()

	if ws == nil {
		return -1
	}

	return ws.Ypixel()
}
