package ansi

import "github.com/skeptycal/gosimple/types/convert"

var (
	blank       []byte  = []byte{ansiPrefixByte1, ansiPrefixByte2, 0, ansiSuffixByte}
	blankArray4 [4]byte = [4]byte{ansiPrefixByte1, ansiPrefixByte2, 0, ansiSuffixByte}
	blankArray5 [5]byte = [5]byte{ansiPrefixByte1, ansiPrefixByte2, 0, 0, ansiSuffixByte}
	blankArray6 [6]byte = [6]byte{ansiPrefixByte1, ansiPrefixByte2, 0, 0, 0, ansiSuffixByte}
)

func BasicEncode(in string) string {
	return basicEncode0(in)
}

func basicEncode0(in string) string {
	return string(append(ansiPrefixByte, in[0], ansiSuffixByte))
}

func basicEncode1(in string) (s string) {
	b := append(ansiPrefixByte, in[0], ansiSuffixByte)
	return string(b)
}

func basicEncode2(in string) (s string) {
	b := make([]byte, 6)
	_ = b
	b = append(ansiPrefixByte, in[0], ansiSuffixByte)
	return string(b)
}

func basicEncode4(in ...byte) string {
	return ""
}

func basicEncode5(in ...byte) string {
	return ""
}

func BasicStringEncode(in string) string {
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

func b2sSafe(in []byte) string   { return string(in) }
func s2bSafe(in string) []byte   { return []byte(in) }
func b2sUnSafe(in []byte) string { return convert.UnsafeBytesToString(in) }
func s2bUnSafe(in string) []byte { return convert.UnsafeStringToBytes(in) }

//////// The following are alternative implementations that were benchmarked and eliminated

// arrayEncode
func arrayEncode(b string) string {
	newBlank := blankArray4
	newBlank[2] = b[0] // not unicode safe anymore, but ANSI codes are not unicode ...
	return convert.UnsafeBytesToString(newBlank[:])
}

// // basicEncode is a slow method and should not be used
// func basicEncode(in string) string {
// 	return ansiPrefix + in + ansiSuffix
// }

// // blankEncode is the fastest method to encode without unsafe.
// func blankEncode(b string) string {
// 	newBlank := blank
// 	newBlank[2] = b[0] // not unicode safe anymore, but ANSI codes are not unicode ...
// 	return convert.UnsafeBytesToString(newBlank)
// }

// func unsafeEncode(b string) string {
// 	newBlank := blank
// 	newBlank[2] = b[0] // not unicode safe anymore, but ANSI codes are not unicode ...
// 	return convert.UnsafeBytesToString(newBlank)
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

// // used to change function signature for table based testing
// func fakeEncode(in string) string {
// 	return encode(in)
// }