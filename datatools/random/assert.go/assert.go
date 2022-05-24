package assert

func LT[O Ordered](got, want O) bool      { return got < want }
func GT[O Ordered](got, want O) bool      { return got > want }
func NE[O Ordered](got, want O) bool      { return got != want }
func EQ[O Ordered](got, want O) bool      { return got == want }
func Success[O Ordered](got, want O) bool { return got+want == want+got }
