package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/skeptycal/gosimple/types/convert"
	"golang.org/x/net/html"
)

var (
	ComicURL string = "http://xkcd.com/"
	B2S             = convert.UnsafeBytesToString
	S2B             = convert.UnsafeStringToBytes
)

type Comic struct {
	Year       string `json:"year"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Image      string `json:"img"`
}

func main() {
	for i := 1; i < 50; i++ {
		result, err := GetComic(ComicURL + strconv.Itoa(i) + "/info.0.json")
		if err != nil {
			fmt.Printf("Error %v", err)
		}

		extension := result.Image[len(result.Image)-4:]
		img, err := os.Create("./xkcd/" + result.SafeTitle + extension)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.Get(result.Image)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		_, err = io.Copy(img, resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		defer img.Close()
		fmt.Printf("URL : %s\n", result.Image)
	}

}

func ReadHTML(r io.Reader) (string, error) {

	// bs, err := ioutil.ReadFile(fileName)
	bs, err := io.ReadAll(r)

	if err != nil {
		return "", err
	}

	return convert.UnsafeBytesToString(bs), nil
}

func ParseTags(r io.Reader, tag string) (data []string) {
	tkn := html.NewTokenizer(r)
	var vals []string
	var isTag bool

	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			return vals

		case tt == html.StartTagToken:
			t := tkn.Token()
			isTag = t.Data == tag

		case tt == html.TextToken:
			t := tkn.Token()
			if isTag {
				vals = append(vals, t.Data)
			}
			isTag = false
		}
	}
}

func GetURLContent(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {

		fmt.Println("ERROR")
		err = fmt.Errorf("error trying to access the provided URL: %s", resp.Status)
		resp.Body.Close()
		return nil, err
	}

	return resp.Body, err

}

func GetComicURL(url string) (string, error) {
	r, err := GetURLContent(url)
	if err != nil {
		return "", err
	}

}

func GetComic(url string) (*Comic, error) {
	r, err := GetURLContent(url)

	var comicBlock Comic
	result := json.NewDecoder(r).Decode(&comicBlock)

	if result != nil {
		return nil, err
	}

	return &comicBlock, nil
}
