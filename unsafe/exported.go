package unsafe

func S2B(s string) []byte { return unsafeStringToBytes(s) }
func B2S(b []byte) string { return unsafeBytesToString(b) }
func DataSize(v any) int  { return intDataSize(v) }
