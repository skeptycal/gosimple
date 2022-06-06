package ansi

// ANSI escape codes for text effects
//
// These are the most commonly used. ANSI codes above 9 are very
// rare and many are not fully implemented.
//  Normal      = 0
//  Bold        = 1
//  Faint       = 2
//  Italics     = 3
//  Underline   = 4
//  Inverse     = 7
//  Conceal     = 8
//  Strikeout   = 9
//
const (
	Normal byte = iota
	Bold        // bold or increased intensity
	Faint       // faint, decreased intensity or second color
	Italics
	Underline
	Blink
	FastBlink
	Inverse
	Conceal
	Strikeout
	// ANSI codes above 9 are very rare and many are not fully implemented.
	PrimaryFont
	AltFont1
	AltFont2
	AltFont3
	AltFont4
	AltFont5
	AltFont6
	AltFont7
	AltFont8
	AltFont9
	Gothic // fraktur
	DoubleUnderline
	NormalColor // normal color or normal intensity (neither bold nor faint)
	NotItalics  // not italicized, not fraktur
	NotUnderlined
	Steady     // not Blink or FastBlink
	Reserved26 // reserved for proportional spacing as specified in CCITT Recommendation T.61
	NotInverse // Positive
	NotHidden  // Revealed
	NotStrikeout
	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	SetForeground     // Next arguments are 5;n or 2;r;g;b, see below
	DefaultForeground // default display color (implementation-defined)
	BlackBackground
	RedBackground
	GreenBackground
	YellowBackground
	BlueBackground
	MagentaBackground
	CyanBackground
	WhiteBackground
	SetBackground              // Next arguments are 5;n or 2;r;g;b, see below
	DefaultBackground          // default background color (implementation-defined)
	DisableProportionalSpacing // reserved for cancelling the effect of parameter value 26
	Framed
	Encircled
	Overlined
	NotFramed // NotEncircled
	NotOverlined
	Reserved56
	Reserved57
	SetUnderlineColor // Next arguments are 5;n or 2;r;g;b, see below
	DefaultUnderlineColor
	IdeogramUnderline       // ideogram underline or right side line
	IdeogramDoubleUnderline // ideogram double underline or double line on the right side
	IdeogramOverline        // ideogram overline or left side line
	IdeogramDoubleOverline  // ideogram double overline or double line on the left side
	IdeogramStress          // ideogram stress marking
	IdeogramCancel          // reset the effects of all of 60–64
	Superscript             = 73
	Subscript               = 74
)
