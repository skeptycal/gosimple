package goshell

import (
	"os"
	"os/user"
	"path/filepath"
	"testing"
)

// TODO: check this test ...
// func Test_PWD_defaults(t *testing.T) {
// 	home := &homedir{}
// 	home.disableCache = defaultDisableCache

// 	if home.disableCache != defaultDisableCache {
// 		t.Fatalf("pwd.DisabledCache: %v  want: %v\n", home.disableCache, defaultDisableCache)
// 	}

// 	if home.homedirCache != "" {
// 		t.Fatalf("pwd.homedirCache: %q  want: %q\n", home.homedirCache, "")
// 	}

// 	if HOME().Abs() == "" {
// 		t.Fatalf("HOME().Abs(): %q want: not the empty string\n", HOME().Abs())
// 	}

// 	if HOME().Base() == "" {
// 		t.Fatalf("HOME().Base(): %q want: not the empty string\n", HOME().Base())
// 	}
// }

func patchEnv(key, value string) func() {
	bck := os.Getenv(key)
	deferFunc := func() {
		os.Setenv(key, bck)
	}

	if value != "" {
		os.Setenv(key, value)
	} else {
		os.Unsetenv(key)
	}

	return deferFunc
}

func BenchmarkDir(b *testing.B) {
	// We do this for any "warmups"
	for i := 0; i < 10; i++ {
		GetHomeDir()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetHomeDir()
	}
}

func TestDir(t *testing.T) {
	u, err := user.Current()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	dir, err := GetHomeDir()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if u.HomeDir != dir {
		t.Fatalf("%#v != %#v", u.HomeDir, dir)
	}

	DisableCache = true
	defer func() { DisableCache = false }()
	defer patchEnv("HOME", "")()
	dir, err = GetHomeDir()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if u.HomeDir != dir {
		t.Fatalf("%#v != %#v", u.HomeDir, dir)
	}
}

func TestExpand(t *testing.T) {
	u, err := user.Current()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	cases := []struct {
		Input  string
		Output string
		Err    bool
	}{
		{
			"/foo",
			"/foo",
			false,
		},

		{
			"~/foo",
			filepath.Join(u.HomeDir, "foo"),
			false,
		},

		{
			"",
			"",
			false,
		},

		{
			"~",
			u.HomeDir,
			false,
		},

		{
			"~foo/foo",
			"",
			true,
		},
	}

	for _, tc := range cases {
		actual, err := Expand(tc.Input)
		if (err != nil) != tc.Err {
			t.Fatalf("Input: %#v\n\nErr: %s", tc.Input, err)
		}

		if actual != tc.Output {
			t.Fatalf("Input: %#v\n\nOutput: %#v", tc.Input, actual)
		}
	}

	DisableCache = true
	defer func() { DisableCache = false }()
	defer patchEnv("HOME", "/custom/path/")()
	expected := filepath.Join("/", "custom", "path", "foo/bar")
	actual, err := Expand("~/foo/bar")

	if err != nil {
		t.Errorf("No error is expected, got: %v", err)
	} else if actual != expected {
		t.Errorf("Expected: %v; actual: %v", expected, actual)
	}
}