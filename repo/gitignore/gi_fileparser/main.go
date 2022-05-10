package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/skeptycal/gosimple/repo/gitignore/gi_fileparser/cli"
)

//go:generate goyacc -o gopher.go -p parser gopher.y

const (
	NormalMode os.FileMode = 0644
	DirMode    os.FileMode = 0755
	newline                = cli.Newline
)

// FieldsFlag  bool
// LinesFlag   bool

var (
	flag = cli.Flag
	O    = &cli.Options
	V    = cli.Vprintln
	NL   = cli.Br
	Head = cli.Head
	Tail = cli.Tail
	// P   = fmt.Println
	B2S = cli.B2S
	S2B = cli.S2B
)

var (
	packageName = ""
)

func init() {
	flag.BoolVar(O.FieldsFlag, "fields", false, "print file contents as fields")
	flag.BoolVar(O.LinesFlag, "lines", false, "print file contents as lines")
	flag.StringVar(&packageName, "package", "main", "package name for generated files")

	flag.Parse(os.Args[1:])
}

func main() {
	s, w := getIo(O.InFile, O.OutFile)
	defer w.Close()

	V("io.Writer: ", w, newline)

	/// Cleanup input
	s = Cleanup(s)

	list := Fields(s, ",")
	V("Head of cleaned list:", Head(strings.Join(list, ", "), 0))
	s = strings.Join(list, ", ")
	V("final cleaned, joined list: ", s)

	NL()
	out := wrapContents(w, "%q, ", fileHeader(), fileFooter(), list, false)

	V("wrapped contents ready for file output: ", out)

	writeFile(w, out)

	fi, err := os.Stat(O.OutFile)
	if err != nil {
		log.Fatal(err)
	}

	V("filesize: ", fi.Size(), newline)

	// write stuff to output file here ...

}

func getIo(in, out string) (string, io.WriteCloser) {
	s, err := getFileData(in)
	if err != nil {
		log.Fatal(err)
	}

	w, err := getWriter(out)
	if err != nil {
		log.Fatal(err)
	}
	return s, w
}

// Cleanup is a data cleaning function specific to
// the current data set. It will likely need to be
// revised any time a new data set is processed.
func Cleanup(s string) string {
	V(Head(s, 0))
	s = AddTrailingSep(s, ",", false)
	V(Head(s, 0))
	s = NormalizeWhitespace(s)
	V(Head(s, 0))
	return s
}

// writeFile writes the results to w.
// As a precaution, the writer uses os.Stdout unless
// the -force CLI option is enabled.
func writeFile(w io.Writer, s string) (n int, err error) {
	if !O.ForceFlag {
		w = os.Stdout
	}
	return w.Write(S2B(s))
}

// wrapContents wraps repeated formatted fields with a
// file header and footer and returns the resulting
// string.
func wrapContents(w io.Writer, format, header, footer string, fields []string, addNewlines bool) string {
	if format == "" {
		format = "%q, " // quoted string items in a list
	}
	// used to estimate length of contents
	// for rough estimate of filesize
	s := strings.Join(fields, format)
	length := len(s) + len(header) + len(footer)
	s = ""

	V("length: ", length, newline)

	sb := &strings.Builder{}
	defer sb.Reset()
	sb.Grow(length)

	fmt.Fprint(sb, header)
	for i := 0; i < len(fields); i++ {
		if fields[i] == "" {
			continue
		}
		// V(format, fields[i])
		fmt.Fprintf(sb, format, fields[i])
		if addNewlines {
			sb.WriteByte('\n')
		}
	}
	fmt.Fprintln(sb, footer) // trailing newline

	return sb.String()
}

// AddTrailingSep adds a trailing separator sep to each
// newline in s.
// The actual newline character can be kept or discarded
// based on the keepNewLines bool.
func AddTrailingSep(s, sep string, keepNewLines bool) string {
	if keepNewLines {
		sep = sep + newline
	}
	return strings.ReplaceAll(s, "\n", sep)
}

func fileHeader() string {
	return `package main

// Copyright (c) 2022 Michael Treanor (skeptycal@gmail.com)
// https://github.com/skeptycal
// MIT License

var giParams = []string{`
}

func fileFooter() string {
	return `}
	`
}

func getFileData(filename string) (string, error) {
	fi, err := os.Stat(O.InFile)
	if err != nil {
		return "", err
	}
	inFileName := fi.Name()
	V("file stat ok: ", inFileName)

	b, err := os.ReadFile(inFileName)
	if err != nil {
		return "", err
	}
	V("file opened: ", inFileName)

	return B2S(b), nil
}

func getWriter(filename string) (io.WriteCloser, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	f, err := os.OpenFile(fi.Name(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, NormalMode)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func printFields(s string) {
	if O.FieldsFlag {
		fields := Fields(NormalizeWhitespace(s), ",")
		if fields == nil {
			log.Fatal("error getting file lines")
		}
		V("fields: ", fields)
		for i, v := range fields {
			V("%3d: %v\n", i, v)
		}
	}
}

func printLines(s string) {
	if O.LinesFlag {
		lines := Lines(s)
		if lines == nil {
			log.Fatal("error getting file lines")
		}
		for i, v := range lines {
			fmt.Printf("%3d: %v\n", i, v)
		}
	}
}

// Lines returns s separated on occurrences of newline.
func Lines(s string) []string {
	return strings.Split(s, newline)
}

// Fields returns s separated on occurrences of sep.
func Fields(s string, sep string) []string {
	return strings.Split(s, ",")
}

// NormalizeWhitespace splits s on any whitespace
// and returns each element as a single space
// separated string.
func NormalizeWhitespace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
