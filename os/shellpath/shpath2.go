package shpath

import (
	"fmt"
	"os"
	"strings"
)

const (
	ListSep = ":"
)

func (p ShPath) Load() error {

	s, err := GetEnvValue("path")
	if err != nil {
		return err
	}

	for _, v := range strings.Split(s, ListSep) {
		if IsDir(v) {
			p.list = append(p.list, v)
		} else {
			fmt.Fprintf(os.Stderr, "the path (%v) is not a valid directory\n", v)
		}
	}

	_ = p.list

	return nil
}

func (p ShPath) Save(path string) error {

	s := strings.Join(p.list, ListSep)

	return os.Setenv("PATH", s)
}
