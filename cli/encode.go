package cli

import (
	"strings"

	"github.com/skeptycal/gosimple/convert"
)

var (
	blank       []byte  = []byte{ansiPrefixByte1, ansiPrefixByte2, 0, ansiSuffixByte}
	blankArray4 [4]byte = [4]byte{ansiPrefixByte1, ansiPrefixByte2, 0, ansiSuffixByte}
	blankArray5 [5]byte = [5]byte{ansiPrefixByte1, ansiPrefixByte2, 0, 0, ansiSuffixByte}
	blankArray6 [6]byte = [6]byte{ansiPrefixByte1, ansiPrefixByte2, 0, 0, 0, ansiSuffixByte}
)

func BasicEncode(in string) string {
	switch len(in) {
	case 2:
		return encode5(in)
	case 1:
		return encode4(in)
	case 3:
		return encode6(in)
	default:
		return ""
	}
}

func encode4(in string) string {
	newBlank := blankArray4
	newBlank[2] = in[0]
	return convert.UnsafeBytesToString(newBlank[:])
}

func encode5(in string) string {
	newBlank := blankArray5
	newBlank[2] = in[0]
	newBlank[3] = in[1]
	return convert.UnsafeBytesToString(newBlank[:])
}

func encode6(in string) string {
	newBlank := blankArray6
	newBlank[2] = in[0]
	newBlank[3] = in[1]
	newBlank[4] = in[2]
	return convert.UnsafeBytesToString(newBlank[:])
}

// // basicEncode is a slow method and should not be used
// func basicEncode(in string) string {
// 	return ansiPrefix + in + ansiSuffix
// }

// arrayEncode
func arrayEncode(b string) string {
	newBlank := blankArray4
	newBlank[2] = b[0] // not unicode safe anymore, but ANSI codes are not unicode ...
	return convert.UnsafeBytesToString(newBlank[:])
}

// var pblankArray = blankArray

// // arrayPtrEncode
// func arrayPtrEncode(b string) string {
// 	pblankArray[2] = b[0] // not unicode safe anymore, but ANSI codes are not unicode ...
// 	return convert.UnsafeBytesToString(pblankArray[:])
// }

// blankEncode is the fastest method to encode without unsafe
func blankEncode(b string) string {
	newBlank := blank
	newBlank[2] = b[0] // not unicode safe anymore, but ANSI codes are not unicode ...
	return convert.UnsafeBytesToString(newBlank)
}

// var pblank = blank

// // blankPtrEncode is the fastest method to encode without unsafe
// func blankPtrEncode(b string) string {
// 	pblank[2] = b[0] // not unicode safe anymore, but ANSI codes are not unicode ...
// 	return convert.UnsafeBytesToString(pblank)
// }

func unsafeEncode(b string) string {
	newBlank := blank
	newBlank[2] = b[0] // not unicode safe anymore, but ANSI codes are not unicode ...

	return convert.UnsafeBytesToString(newBlank)
}

// this is slow for single bytes
func encode(b ...string) string {
	sb := strings.Builder{}
	sb.WriteByte(ansiPrefixByte1)
	sb.WriteByte(ansiPrefixByte2)
	for _, c := range b {
		sb.WriteString(c)
		sb.WriteByte(ansiSepByte)
	}
	sb.WriteByte(ansiSuffixByte)
	return sb.String()
}

func newAnsiColorString(in string) string {
	return NewAnsiColor(in).String()
}

// used to change function signature for table based testing
func fakeEncode(in string) string {
	return encode(in)
}

func goEncodeMulti(in ...string) string {
	ret := ansiPrefix
	c := make(chan string)
	done := make(chan bool)
	for _, b := range in {
		go func(b string) {
			c <- BasicEncode(b)
		}(b)
	}
	done <- true
	select {
	case s := <-c:
		ret += s
	case <-done:
		break
	}
	return ret
}
