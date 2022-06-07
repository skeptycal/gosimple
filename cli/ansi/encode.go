package ansi

import (
	"sync"

	"github.com/skeptycal/gosimple/reallyunsafe"
)

const (
	blankStr = ansiPrefix + "0" + ansiSuffix
)

var (
	s2b                 = reallyunsafe.S2B
	b2s                 = reallyunsafe.B2S
	blank       []byte  = []byte(blankStr)
	blankArray4 [4]byte = [4]byte{ansiPrefixByte1, ansiPrefixByte2, 0, ansiSuffixByte}
	blankArray5 [5]byte = [5]byte{ansiPrefixByte1, ansiPrefixByte2, 0, 0, ansiSuffixByte}
	blankArray6 [6]byte = [6]byte{ansiPrefixByte1, ansiPrefixByte2, 0, 0, 0, ansiSuffixByte}
)

type ansiArray [4]byte

func (a *ansiArray) Set(b byte) { a[2] = b }

func (a *ansiArray) String() string {
	return b2s(a[:])
}

var ansiPool = sync.Pool{
	New: func() any {
		v := new(ansiArray)
		*v = blankArray4
		return v
	},
}

func BasicEncode(in string) string {
	return arrayEncode(in)
}

// blankEncode is the fastest method to encode without unsafe.
func blankEncode(b string) string {
	newBlank := blank
	newBlank[2] = b[0]
	return string(newBlank)
}

// uses sync.Pool with []byte
func poolArrayEncode(b string) string {
	newBlank := ansiPool.Get().(*ansiArray)
	defer ansiPool.Put(newBlank)

	newBlank[2] = b[0]
	return string(newBlank[:])
}

func poolArrayStringerEncode(b string) string {
	newBlank := ansiPool.Get().(*ansiArray)
	defer ansiPool.Put(newBlank)
	newBlank[2] = b[0]
	return newBlank.String()
}

func poolArraySetterEncode(b string) string {
	newBlank := ansiPool.Get().(*ansiArray)
	defer ansiPool.Put(newBlank)
	newBlank.Set(b[0])
	return newBlank.String()
}

func unsafeBlankEncode(b string) string {
	newBlank := blank
	newBlank[2] = b[0]
	return b2s(newBlank)
}

// arrayEncode
func arrayEncode(b string) string {
	var newBlank = blankArray6[:]
	for i := 0; i < len(b); i++ {
		newBlank[i+2] = b[i]
	}
	// newBlank[2] = b[0]
	return string(newBlank[:])
}

// arrayEncode
func arrayEncodeUnsafe(b string) string {
	newBlank := blankArray4
	newBlank[2] = b[0]
	return b2s(newBlank[:])
}

func b2sSafe(in []byte) string   { return string(in) }
func s2bSafe(in string) []byte   { return []byte(in) }
func b2sUnSafe(in []byte) string { return b2s(in) }
func s2bUnSafe(in string) []byte { return s2b(in) }

//////// The following are alternative implementations that were benchmarked and eliminated

// func sprintencode(in string) string {
// 	return fmt.Sprintf(fmtANSI, in[0])
// }

// // basicEncode is a slow method and should not be used
// func basicEncode(in string) string {
// 	return ansiPrefix + in + ansiSuffix
// }

// // used to change function signature for table based testing
// func sbencode(in string) string {
// 	return encode(in)
// }

// // this is slow for single bytes
// func encode(b ...string) string {
// 	sb := strings.Builder{}
// 	sb.WriteByte(ansiPrefixByte1)
// 	sb.WriteByte(ansiPrefixByte2)
// 	for _, c := range b {
// 		sb.WriteString(c)
// 		sb.WriteByte(ansiSepByte)
// 	}
// 	sb.WriteByte(ansiSuffixByte)
// 	return sb.String()
// }

// func newAnsiColorString(in string) string {
// 	return NewAnsiColor(in).String()
// }

// func BasicStringEncode(in string) string {
// 	switch len(in) {
// 	case 2:
// 		return encode5(in)
// 	case 1:
// 		return encode4(in)
// 	case 3:
// 		return encode6(in)
// 	default:
// 		return ""
// 	}
// }

// func encode4(in string) string {
// 	newBlank := blankArray4
// 	newBlank[2] = in[0]
// 	return b2s(newBlank[:])
// }

// func encode5(in string) string {
// 	newBlank := blankArray5
// 	newBlank[2] = in[0]
// 	newBlank[3] = in[1]
// 	return b2s(newBlank[:])
// }

// func encode6(in string) string {
// 	newBlank := blankArray6
// 	newBlank[2] = in[0]
// 	newBlank[3] = in[1]
// 	newBlank[4] = in[2]
// 	return b2s(newBlank[:])
// }

// func appendencode(in string) string {
// 	return string(append(ansiPrefixByte, in[0], ansiSuffixByte))
// }

// func appendencodeUnsafe(in string) string {
// 	return b2s(append(ansiPrefixByte, in[0], ansiSuffixByte))
// }

// func appendencode2(in string) (s string) {
// 	b := append(ansiPrefixByte, in[0], ansiSuffixByte)
// 	return string(b)
// }

// func appendencodeprealloc(in string) (s string) {
// 	b := make([]byte, 6)
// 	_ = b
// 	b = append(ansiPrefixByte, in[0], ansiSuffixByte)
// 	return string(b)
// }

// func unsafeappendencodeprealloc(in string) string {
// 	b := make([]byte, 6)
// 	_ = b
// 	b = append(ansiPrefixByte, in[0], ansiSuffixByte)
// 	return b2s(b)
// }
