package textfiles

import (
	"io"
	"os"
	"sort"
	"strings"
	"unicode"
)

func OpenTextFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}

func String(b []byte) string            { return B2S(b) }
func Bytes(s string) []byte             { return S2B(s) }
func Lines(s string) []string           { return strings.Split(s, newline) }
func Words(s string) []string           { return strings.Fields(s) }
func ToValidUTF8(s string) string       { return strings.ToValidUTF8(s, string(unicode.ReplacementChar)) }
func RemoveNewlines(text string) string { return strings.Join(strings.Split(text, newline), " ") }

func SortLines(lines []string) []string {
	if !sort.StringsAreSorted(lines) {
		sort.Strings(lines)
	}
	return lines
}

func SortWords(words []string) []string {
	if !sort.StringsAreSorted(words) {
		sort.Strings(words)
	}
	return words
}

// LineSet returns a map of line numbers to lines.
func LineMap(s string) map[int]string {
	lines := Lines(s)
	m := make(map[int]string, len(lines))

	for i, line := range lines {
		m[i] = line
	}

	return m
}
