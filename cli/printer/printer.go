package printer

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

// Printer implements common printer functionality.
// It provides Print, Println, and Printf to os.Stdout
// as well as the more general FPrint, etc methods
// for writing to a specified io.Writer. In addition,
// the standard Sprint methods are available.
//
// For convenience, a DbPrinter is provided through
// the methods DbPrint, DbPrintln, and DbPrintf.
// These 'debug' printer functions print to os.Stderr
// by default and wrap the output in given ANSI codes.
type Printer interface {
	NormalPrinter
	Fprinter
	Sprinter
	DbPrinter
	SetWriter(w io.Writer)
	SetDbWriter(w io.Writer)
	SetPrefix(prefix []byte)
	SetSuffix(suffix []byte)
}

type print struct {
	w       io.Writer
	db      io.Writer
	prefix  []byte
	dbColor []byte
	suffix  []byte
}

const (
	NewLine    = "\n"         // Newline
	DbColor    = "\033[1;31m" // ANSI dbecho code
	ResetColor = "\033[0m"    // ANSI reset code
)

var (
	bReset   = []byte(ResetColor)
	bNewline = []byte(NewLine)
	bDbColor = []byte(DbColor)
)

func NewPrinter() Printer {
	return NewPrinterWithOptions(os.Stdout, os.Stderr, nil, nil, bDbColor)
}

func NewPrinterWithOptions(w io.Writer, db io.Writer, prefix, dbColor, suffix []byte) Printer {
	return &print{w, db, prefix, dbColor, suffix}
}

func (p *print) Fprintf(w io.Writer, format string, args ...interface{}) (n int, err error) {
	w.Write(p.prefix)
	n, err = fmt.Fprintf(w, format, args...)
	w.Write(p.suffix)
	w.Write(bReset)
	return
}
func (p *print) Fprint(w io.Writer, args ...interface{}) (n int, err error) {
	w.Write(p.prefix)
	n, err = fmt.Fprint(w, args...)
	w.Write(p.suffix)
	w.Write(bReset)
	return
}
func (p *print) Fprintln(w io.Writer, args ...interface{}) (n int, err error) {
	w.Write(p.prefix)
	n, err = fmt.Fprint(w, args...)
	w.Write(p.suffix)
	w.Write(bNewline)
	w.Write(bReset)
	return
}

func (p *print) Printf(format string, args ...interface{}) (n int, err error) {
	return p.Fprintf(os.Stdout, format, args...)
}
func (p *print) Print(args ...interface{}) (n int, err error) {
	return p.Fprint(os.Stdout, args...)
}
func (p *print) Println(args ...interface{}) (n int, err error) {
	return p.Fprintln(os.Stdout, args...)
}

func (p *print) DbPrintf(format string, args ...interface{}) (n int, err error) {
	p.db.Write(bDbColor)
	n, err = fmt.Fprintf(p.db, format, args...)
	p.db.Write(bReset)
	return
}
func (p *print) DbPrint(args ...interface{}) (n int, err error) {
	p.db.Write(p.dbColor)
	n, err = fmt.Fprint(p.db, args...)
	p.db.Write(bReset)
	return
}
func (p *print) DbPrintln(args ...interface{}) (n int, err error) {
	p.db.Write(p.dbColor)
	n, err = fmt.Fprint(p.db, args...)
	p.db.Write(bNewline)
	p.db.Write(bReset)
	return
}

var sbPool = sync.Pool{
	New: func() interface{} {
		return &strings.Builder{}
	},
}

func (p *print) Sprintf(format string, args ...interface{}) string {
	return p.sprintf1(format, args...)
}

func (p *print) sprintf1(format string, args ...interface{}) string {
	sb := &strings.Builder{}
	defer sb.Reset()
	sb.Write(p.prefix)
	fmt.Fprintf(sb, format, args...)
	sb.Write(p.suffix)
	return sb.String()
}

func (p *print) sprintf2(format string, args ...interface{}) string {
	sb := &strings.Builder{}
	sb.Write(p.prefix)
	fmt.Fprintf(sb, format, args...)
	sb.Write(p.suffix)
	s := sb.String()
	sb.Reset()
	return s
}
func (p *print) sprintf3(format string, args ...interface{}) (s string) {
	sb := &strings.Builder{}
	sb.Write(p.prefix)
	fmt.Fprintf(sb, format, args...)
	sb.Write(p.suffix)
	s = sb.String()
	sb.Reset()
	return
}

func (p *print) sprintf4(format string, args ...interface{}) (s string) {
	sb := sbPool.Get().(*strings.Builder)
	sb.Reset()
	// defer sbPool.Put(sb)
	sb.Write(p.prefix)
	fmt.Fprintf(sb, format, args...)
	sb.Write(p.suffix)
	s = sb.String()
	sbPool.Put(sb)
	return
}

func (p *print) sprintf5(format string, args ...interface{}) string {
	sb := sbPool.Get().(*strings.Builder)
	sb.Reset()
	// defer sbPool.Put(sb)
	sb.Write(p.prefix)
	fmt.Fprintf(sb, format, args...)
	sb.Write(p.suffix)
	s := sb.String()
	sbPool.Put(sb)
	return s
}

func (p *print) sprintf6(format string, args ...interface{}) string {
	sb := sbPool.Get().(*strings.Builder)
	sb.Reset()
	defer sbPool.Put(sb)

	sb.Write(p.prefix)
	fmt.Fprintf(sb, format, args...)
	sb.Write(p.suffix)
	return sb.String()
}

func (p *print) Sprint(args ...interface{}) string {
	sb := strings.Builder{}
	defer sb.Reset()
	sb.Write(p.prefix)
	sb.WriteString(fmt.Sprint(args...))
	sb.Write(p.suffix)
	return sb.String()
}
func (p *print) Sprintln(args ...interface{}) string {
	sb := strings.Builder{}
	defer sb.Reset()
	sb.Write(p.prefix)
	sb.WriteString(fmt.Sprint(args...))
	sb.Write(p.suffix)
	sb.Write(bNewline)
	return sb.String()
}

func (p *print) SetWriter(w io.Writer)    { p.w = w }
func (p *print) SetDbWriter(db io.Writer) { p.db = db }
func (p *print) SetPrefix(prefix []byte)  { p.prefix = prefix }
func (p *print) SetSuffix(suffix []byte)  { p.suffix = suffix }

type NormalPrinter interface {
	Printf(format string, args ...interface{}) (int, error)
	Print(args ...interface{}) (int, error)
	Println(args ...interface{}) (int, error)
}

type DbPrinter interface {
	DbPrintf(format string, args ...interface{}) (int, error)
	DbPrint(args ...interface{}) (int, error)
	DbPrintln(args ...interface{}) (int, error)
}

type Fprinter interface {
	Fprintf(w io.Writer, format string, args ...interface{}) (int, error)
	Fprint(w io.Writer, args ...interface{}) (int, error)
	Fprintln(w io.Writer, args ...interface{}) (int, error)
}

type Sprinter interface {
	Sprintf(format string, args ...interface{}) string
	Sprint(args ...interface{}) string
	Sprintln(args ...interface{}) string
}
