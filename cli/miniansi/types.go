package miniansi

import (
	"github.com/skeptycal/gosimple/constraints"
)

type AnsiConstraint interface {
	~string | constraints.Integer // | []byte
}

type Config interface {
	SetDebug(debug bool)
	SetEnabled(enabled bool)
	DbPrint(args ...interface{}) (n int, err error)
}

type control[T AnsiConstraint] struct {
	debug        bool
	enabled      bool
	InfoColor    ansi[T]
	NoticeColor  ansi[T]
	WarningColor ansi[T]
	ErrorColor   ansi[T]
	DebugColor   ansi[T]
	dbcolor      ansi[T]
	ResetColor   ansi[T]
}
