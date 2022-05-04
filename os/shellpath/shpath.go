package shpath

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	NL        = "\n"
	nlWindows = "\r\n"
	TAB       = "\t"
	PATHSEP   = string(os.PathListSeparator)
)

var Verbose bool = false

type ShPath struct{ list []string }

func NewPath() (*ShPath, error) {
	p := &ShPath{}
	err := p.load()
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Clean removes invalid directories,
// and returns the number removed, if any.
// The order of the directories is maintained.
func (p *ShPath) Clean() (n int) {
	for i, v := range p.list {
		if v == "" || !IsDir(v) {
			p.list = RemoveOrdered(p.list, i)
			n += 1
			if Verbose {
				fmt.Fprintf(os.Stderr, "the path (%v) is not a valid directory\n", v)
			}
		}
	}

	if n > 0 && Verbose {
		fmt.Fprintf(os.Stdout, "directories checked (%v removed)\n", n)
	}

	return n
}

func (p *ShPath) load() error {
	s, err := GetEnvValue("path")
	if err != nil {
		return err
	}
	s = DropDupeSeps(s, PATHSEP)
	s = strings.ReplaceAll(s, nlWindows, NL)
	p.list = strings.Split(s, PATHSEP)
	return nil
}

// Out returns the path in delimited format ready
// for OS use. Out runs Clean() on the list.
func (p *ShPath) Out() string {
	p.Clean()
	return strings.Join(p.list, PATHSEP)
}

// Add checks that the directory exists and
// adds element s to the path at position n.
// If the position is not valid, s will be
// placed at the end of the list.
func (p *ShPath) Add(s string, n int) error {
	if s == "" {
		return errors.New("path cannot be empty")
	}
	if !IsDir(s) {
		v := fmt.Sprintf("the path (%v) is not a valid directory\n", s)
		if Verbose {
			fmt.Fprint(os.Stderr, v)
		}
		return errors.New(v)
	}

	// if n is out of bounds, append
	// s to end of list
	if n < 0 || n >= len(p.list) {
		p.list = Append(p.list, s)
		return nil
	}

	p.list = Insert(p.list, s, n)

	return nil
}

// Len returns the number of items in the list.
func (p *ShPath) Len() int { return len(p.list) }

// String returns the path in newline delimited format. (pretty print)
func (p *ShPath) String() string { return strings.Join(p.list, NL) }

// DebugPrint prints a numbered list of items.
func (p *ShPath) DebugPrint() {
	fmt.Println("path.DebugPrint()")
	for i, v := range p.list {
		fmt.Printf("%3d: %s\n", i, v)
	}
}
