package basic

import (
	"os"

	"github.com/skeptycal/gosimple/cli/errorlogger"
	"github.com/skeptycal/gosimple/cli/terminal"
)

var (

	// Global errorlogger instance
	Log = errorlogger.New()

	// DEBUG flag to enable debug logging and features
	DEBUG = true

	// Column width of CLI terminal display
	COLUMNS int = 80
	ROWS    int = 24

	// Terminal flag to enable CLI terminal display
	IsTerminal = terminal.IsTerminal(int(os.Stdout.Fd()))
)

func init() {
	// TODO: causes seg fault in VsCode terminal window, e.g.
	/* time="2022-06-01T12:39:29-05:00" level=error msg="GetWinsize: operation not supported on socket"
	panic: runtime error: invalid memory address or nil pointer dereference
	[signal SIGSEGV: segmentation violation code=0x2 addr=0x2 pc=0x104f87b44]
	*/
	winSize, err := terminal.GetWinsize()
	if err != nil {
		log.Error(err)
	}

	// COLUMNS = envvars.COLUMNS // TODO not working ... should be working ...
	COLUMNS = int(winSize.Col)
	ROWS = int(winSize.Row)

}
