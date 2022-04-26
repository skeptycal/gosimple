package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"

	"github.com/skeptycal/gosimple/cli/errorlogger"
	"github.com/skeptycal/gosimple/os/gofile"
)

var (
	log       = errorlogger.New()
	forceFlag bool
)

func FindDirs(dir, pattern string) ([]fs.DirEntry, error) {

	return nil, errors.New("not implemented")
}

func init() {
	flag.BoolVar(&forceFlag, "force", false, "force deletion of located files")

}

func main() {

	// find . -type f -wholename "./*/go.sum" -exec rm -rf {} +

	// FindDirs locates directories from the root 'dir' that
	// match the pattern.

	dir := gofile.PWD()

	patterns := []string{"go.mod", "go.sum", "go.sum", ".git"}

	for _, pattern := range patterns {
		list, err := FindDirs(dir, pattern)
		if err != nil {
			log.Errorf("error finding directories: %v", err)
		}

		for _, file := range list {
			fmt.Printf("%8v %s\n", file.Type().Perm(), file.Name())
			if forceFlag {
				os.Remove(file.Name())
			}
		}
	}
}
