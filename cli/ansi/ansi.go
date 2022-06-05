package ansi

// ANSI implements the interface for ANSI encoded values typically used
// for control and output in CLI applications.
type ANSI interface {
	String() string
}

func NewAnsiColor(in string) ANSI {
	return &AnsiColor{in, BasicEncode(in)}
}

// AnsiColor is a buffered, encoded ANSI color string typically used
// for CLI output. The encoded ANSI color code string is JIT
// buffered at the time of the first output request to eliminate
// repeated fmt.Sprintf (or similar) calls.
type AnsiColor struct {
	Color string
	out   string
}

// String returns the ANSI formatted string representation of the AnsiColor byte.
func (a AnsiColor) String() string {
	if a.out == "" {
		a.out = BasicEncode(a.Color)
	}
	return a.out
}
