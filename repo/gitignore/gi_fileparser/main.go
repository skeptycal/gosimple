package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/skeptycal/gosimple/repo/gitignore/cli"
	"github.com/skeptycal/gosimple/repo/gitignore/file"
)

//go:generate goyacc -o gopher.go -p parser gopher.y

const (
	newline = cli.Newline
)

// FieldsFlag  bool
// LinesFlag   bool

var (
	Flags = cli.Flags
	V     = cli.Vprintln
	NL    = cli.Br
	Flag  = cli.Flag

	// P   = fmt.Println
	B2S = cli.B2S
	S2B = cli.S2B
)

func Head[E any](s []E) []E {
	// TODO: an alias would not work with generic type.
	// var Head = cli.Head
	return cli.Head(s)
}

var (
	packageName = ""
)

func init() {
	// flag.BoolVar(O.FieldsFlag, "fields", false, "print file contents as fields")
	// flag.BoolVar(O.LinesFlag, "lines", false, "print file contents as lines")
	Flags.StringVar(&packageName, "package", "main", "package name for generated files")

	Flags.Parse(os.Args[1:])
}

func main() {
	s, w := getIoCLI(*cli.InFile, *cli.OutFile, true)
	defer w.Close()

	V("io.Writer: ", w, newline)

	/// Cleanup input
	s = Cleanup(s)

	list := file.Fields(s, ",")
	V("Head of cleaned list:", Head(list))
	s = strings.Join(list, ", ")
	Head(S2B(s))
	V("final cleaned, joined list: ", s)

	NL()
	out := wrapContents("%q, ", fileHeader(), fileFooter(), list, false)

	V("wrapped contents ready for file output: ", out)

	cli.WriteString(w, out)

	fi, err := os.Stat(*cli.OutFile)
	if err != nil {
		log.Fatal(err)
	}

	V("filesize: ", fi.Size(), newline)

	// write stuff to output file here ...

}

// getIo returns the string contents of the input file
// and an io.WriteCloser to the output file.
// The truncate parameter truncates the output file
// upon opening if set to true. If false, the output file
// is returned in append mode.
// Any error is returned unchanged.
func getIo(in, out string, truncate bool) (string, io.WriteCloser, error) {
	s, err := file.GetFileData(in)
	if err != nil {
		return "", nil, err
	}
	_, wc, err := file.NewGoFile(out)
	if err != nil {
		return "", nil, err
	}

	// f, err := os.OpenFile(name string, flag int, perm os.FileMode)

	return s, wc, nil
}

// getIoCLI returns the string contents of the input file
// and an io.WriteCloser to the output file.
// In the CLI version, any error results in a
// log.Fatal(err).
func getIoCLI(in, out string, truncate bool) (string, io.WriteCloser) {
	s, wc, err := getIo(in, out, truncate)
	if err != nil {
		log.Fatal(err)
	}

	return s, wc
}

// Cleanup is a data cleaning function specific to
// the current data set. It will likely need to be
// revised any time a new data set is processed.
func Cleanup(s string) string {
	V(Head(S2B(s)))
	s = file.AddTrailingSep(s, ",", false)
	V(Head(S2B(s)))
	s = file.NormalizeWhitespace(s)
	V(Head(S2B(s)))
	return s
}

// wrapContents wraps repeated formatted fields with a
// file header and footer and returns the resulting
// string.
func wrapContents(format, header, footer string, fields []string, addNewlines bool) string {
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

// func printFields(s string) {
// 	if O.FieldsFlag {
// 		fields := Fields(NormalizeWhitespace(s), ",")
// 		if fields == nil {
// 			log.Fatal("error getting file lines")
// 		}
// 		V("fields: ", fields)
// 		for i, v := range fields {
// 			V("%3d: %v\n", i, v)
// 		}
// 	}
// }

// func printLines(s string) {
// 	if O.LinesFlag {
// 		lines := Lines(s)
// 		if lines == nil {
// 			log.Fatal("error getting file lines")
// 		}
// 		for i, v := range lines {
// 			fmt.Printf("%3d: %v\n", i, v)
// 		}
// 	}
// }
