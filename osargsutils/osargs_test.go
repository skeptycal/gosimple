package osargs

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Benchmark Results
/*
/// It does not matter ... in fact, simply calling the argzero function to get the initial
/// path is what takes up all of the time.

/// Will any of this matter? Not at all ... this function is likely to be called only once
/// during the runtime of the program. a few hundred nanoseconds will not matter ...

/// however, if 100,000 of these are started up in a scaling architecture, run quickly,
/// and exit nearly immediately ... well then the startup code has a huge effect.

//* As is recommended since Go 1.8, os.Executable is fastest.

BenchmarkHereMe2-8            	   23202	     55725 ns/op	    3696 B/op	      42 allocs/op
BenchmarkHereMe-8             	   24092	     48869 ns/op	    3696 B/op	      42 allocs/op
BenchmarkZeroOsExecutable-8   	   25843	     46830 ns/op	    3696 B/op	      42 allocs/op
BenchmarkZeroOsArgs-8         	   25575	     47170 ns/op	    3696 B/op	      42 allocs/op
BenchmarkRawOsArgsZero-8      	   25155	     48109 ns/op	    3695 B/op	      42 allocs/op
*/

func TestOsArgs_App(t *testing.T) {
	got := OsArgs.App()
	want, _ := filepath.EvalSymlinks(os.Args[0])
	if got != want {
		t.Errorf("OsArgs.App != os.Args[0]: %v != %v", got, want)
	}
}

func TestArgs(t *testing.T) {
	got := OsArgs.Args()
	want := os.Args[1:]
	if len(got) != len(want) {
		t.Errorf("len(OsArgs.Args) != len(os.Args[1:]): %v != %v", len(got), len(want))
	}
	if len(got) > 0 && len(want) > 0 {
		if got[0] != want[0] {
			t.Errorf("OsArgs.Args[0] != os.Args[0]): %v != %v", got[0], want[0])
		}
	}
}

func TestArgString(t *testing.T) {
	got := OsArgs.ArgString()
	want := strings.Join(os.Args[1:], " ")
	if got != want {
		t.Errorf("Argstrings not equal: %v != %v", got, want)
	}
}
func TestBase(t *testing.T) {
	got := OsArgs.Base()
	want, _ := filepath.EvalSymlinks(os.Args[0])
	want = filepath.Base(want)
	if got != want {
		t.Errorf("Basenames do not match: %v != %v", got, want)
	}
}

func TestDir(t *testing.T) {
	got := OsArgs.Dir()
	want, _ := filepath.EvalSymlinks(os.Args[0])
	want = filepath.Dir(want)
	if got != want {
		t.Errorf("Paths do not match: %v != %v", got, want)
	}
}

func TestHereMe(t *testing.T) {
	got1, got2 := HereMe()
	want, _ := filepath.EvalSymlinks(os.Args[0])
	want1 := filepath.Dir(want)
	want2 := filepath.Base(want)
	if got1 != want1 {
		t.Errorf("Paths do not match: %v != %v", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("Paths do not match: %v != %v", got2, want2)
	}
}
