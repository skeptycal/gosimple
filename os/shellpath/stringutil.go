package shpath

import (
	"bytes"
	"strings"

	"github.com/skeptycal/gosimple/types/convert"
)

// NormalizeNL normalizes newlines to the popular
// standard '\n' convention.
//
// Any '\r\n' (a Windows convention) or '\r' (an
// older Apple convention) sequences are replaced
// with a single '\n'.
//
// Several methods are profiled and the most efficient
// one is aliased to this function.
var NormalizeNL = normalizeNewlinesString

// DropDupeSeps replaces consecutive duplicates of only specific 'sep' strings with a single 'sep'
func DropDupeSeps(s string, sep string) string {
	return dropDupe(s, sep)
}

const (
	rConst  = "\r"   // single carriage return (older Mac newline)
	rnConst = "\r\n" // older Windows newline
	rrConst = "\r\r" // duplicate carriage return (no reason for existing)
	nnConst = "\n\n" // duplicate newlines (may be intended)
	newline = "\n"   // normalized newline
)

var oldList = []string{"\r", "\r\n", "\n\n"}

// byte slices of convertion string constants
var (
	r  = []byte(rConst)
	rr = []byte(rConst + rConst)
	rn = []byte(rnConst)
	nn = []byte(newline + newline)
	nl = []byte(newline)
)

// conversion function shortcut aliases
var (
	b2s = convert.UnsafeBytesToString
	s2b = convert.UnsafeStringToBytes
)

// normalize is string version of the simplest, and likely
// quickest, line ending normalization function.
//
// Benchmarks in comments.
func normalize(in string) string {
	return b2s(normalizeBytes(s2b(in)))
}

func normalizeBytesTester(s string) string {
	b := []byte(s)
	b = normalizeBytes(b)
	return string(b)
}

// normalizeBytes is the byte slice version of the
// simplest, and likely quickest, line ending
// normalization function.
//
// Benchmarks in comments.
func normalizeBytes(b []byte) []byte {
	// remove \r\n
	for bytes.Contains(b, rn) {
		b = bytes.ReplaceAll(b, rn, nl)
	}

	// dedupe any remaining duplicate r's
	for bytes.Contains(b, rr) {
		b = bytes.ReplaceAll(b, r, nl)
	}

	// replace any remaining \r with \n
	for bytes.Contains(b, r) {
		b = bytes.ReplaceAll(b, r, nl)
	}

	return b
}

// normalizeNewlinesString uses strings.Replace
//
// for benchmarking comparisons
func normalizeNewlinesString(d string) string {
	// replace CR LF \r\n (windows) with LF \n (unix)
	d = strings.ReplaceAll(d, rnConst, newline)
	// replace CF \r (mac) with LF \n (unix)
	d = strings.ReplaceAll(d, rConst, newline)
	return d
}

func normalizeWrapped(s string) string {
	return normalizeNewlinesString(s)
}

// for benchmarking comparisons
func normalizeNLForLoop(s string) string {

	old := []byte(s)
	new := make([]byte, 0, len(s))

	for i, c := range old {
		switch c {
		case '\r':
			if i >= len(old) || old[i+1] != '\n' {
				new = append(new, '\n')
			}
		case '\n':
			new = append(new, '\n')
		default:
			new = append(new, c)
		}
	}
	return string(new)
}

func dropDupe(s, sub string) string {
	for strings.Contains(s, sub+sub) {
		return strings.ReplaceAll(s, sub+sub, sub)
	}
	return s
}

func normalizeLoop(s string) string {
	for _, olds := range oldList {
		s = dropDupe(s, olds)
		s = strings.Replace(s, olds, newline, -1)
	}
	return s
}

func normalizeStringsBuilder(s string) string {
	sb := strings.Builder{}
	var start = 0
	var current = start
	for i := strings.Index(s[start:], "\r"); i > 0; {
		sb.WriteString(s[start : start+current]) // write part without \r
		start = i + 1                            // skip \r
		if start > len(s) {                      // check for done
			break
		}
		if s[i+1] != '\n' {
			sb.WriteByte('\n') // add in nl to replace isolated \r
		}
	}
	sb.WriteString(s[start:])

	return sb.String()
}

// normalizeNewlines normalizes \r\n (windows) and \r (mac)
// into \n (unix)
//
// Reference: https://www.programming-books.io/essential/go/normalize-newlines-1d3abcf6f17c4186bb9617fa14074e48
func normalizeNewlines(d []byte) []byte {
	// replace CR LF \r\n (windows) with LF \n (unix)
	d = bytes.ReplaceAll(d, []byte{13, 10}, []byte{10})
	// replace CF \r (mac) with LF \n (unix)
	d = bytes.ReplaceAll(d, []byte{13}, []byte{10})
	return d
}

// for benchmarking comparisons
func normalizeNewlinesBytesWrapper(d string) string {
	return string(normalizeNewlines([]byte(d)))
}
