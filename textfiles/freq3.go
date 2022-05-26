package textfiles

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
)

type byFreq []WordFreq

func (a byFreq) Len() int           { return len(a) }
func (a byFreq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFreq) Less(i, j int) bool { return a[i].freq < a[j].freq }

func Frequency3(filename string) {

	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	text := string(bs)

	re := regexp.MustCompile("[a-zA-Z']+")
	matches := re.FindAllString(text, -1)

	words := make(map[string]int)

	for _, match := range matches {
		words[match]++
	}

	var wordFreqs []WordFreq
	for k, v := range words {
		wordFreqs = append(wordFreqs, WordFreq{k, v})
	}

	sort.Sort(sort.Reverse(byFreq(wordFreqs)))

	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", wordFreqs[i])
	}
}
