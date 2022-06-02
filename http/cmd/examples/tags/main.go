package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	. "github.com/skeptycal/gosimple/cli/basic"
	"github.com/skeptycal/gosimple/types/convert"
	"golang.org/x/net/html"
)

const (
	defaultTag   = "div"
	defaultDEBUG = true
)

var (
	log     = Log
	allFlag bool
	tag     string
)

func init() {
	flag.BoolVar(&allFlag, "all", false, "print all HTML contents")
	flag.BoolVar(&DEBUG, "debug", defaultDEBUG, "show debug information")
	flag.StringVar(&tag, "tag", defaultTag, "name of tag to retrieve")

	flag.Parse()
}

func dbInfo() {
	DbEcho("DEBUG: ", DEBUG)
	DbEcho("tag: ", tag)
	DbEcho("allFlag: ", allFlag)
	DbEcho("isTerminal: ", IsTerminal)
	DbEcho("column width: ", COLUMNS)
	DbEcho("args before flag.Parse(): ", os.Args[1:])
	DbEcho("args after flag.Parse(): ", flag.Args())
}

func main() {
	// flags()
	dbInfo()
	// DEBUG = true

	url := "https://xkcd.com/1/"
	r, err := GetURLReader(url)
	if err != nil {
		log.Fatal(err)
	}

	if allFlag {
		s, err := ReadHTML(r)
		if err != nil {
			log.Errorf("error reading html: %v", err)
		}
		Box(s)
		fmt.Println(s)
	}

	data := ParseTags(r, tag)
	if len(data) == 0 {
		log.Fatalf("no tags(%s) found: %d", tag, len(data))
	}

	for i, item := range data {
		fmt.Printf("%3d: %q\n", i, item)
	}

}

// ReadHTML reads all HTML from the io.Reader and returns
// the string representation. The caller is responsible
// for closing the io.Reader if necessary.
func ReadHTML(r io.Reader) (string, error) {
	// bs, err := ioutil.ReadFile(fileName)
	bs, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	return convert.UnsafeBytesToString(bs), nil
}

// ParseTags parses the HTML stream from the io.Reader
// and returns a list of locations of the requested tags.
func ParseTags(r io.Reader, tag string) (data []string) {
	tkn := html.NewTokenizer(r)
	var vals []string
	var isTag bool

	for {
		tt := tkn.Next()
		switch tt {
		case html.ErrorToken:
			return vals

		case html.StartTagToken:
			t := tkn.Token()
			isTag = t.Data == tag

		case html.TextToken:
			t := tkn.Token()
			if isTag {
				if strings.TrimSpace(t.Data) != "" {
					vals = append(vals, t.Data)
				}
			}
			isTag = false
		}
	}
}

// GetURLReader connects to the URL and returns an
// io.ReadCloser that reads from the URL.
//
// If any error occurs, or if the response status code
// is not 200 (StatusOK), then an error is returned.
func GetURLReader(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("error trying to access the provided URL: %s", resp.Status)
		resp.Body.Close()
		return nil, err
	}

	return resp.Body, err
}
