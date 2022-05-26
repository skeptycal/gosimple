package textfiles

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
)

type WordFreq struct {
	word string
	freq int
}

func (p WordFreq) String() string {
	return fmt.Sprintf("%s %d", p.word, p.freq)
}

func Frequency2(filename string) {

	// filename = "the-king-james-bible.txt"

	reg := regexp.MustCompile("[a-zA-Z']+")
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	text := string(bs)
	matches := reg.FindAllString(text, -1)

	words := make(map[string]int)

	for _, match := range matches {
		words[match]++
	}

	var wordFreqs []WordFreq
	for k, v := range words {
		wordFreqs = append(wordFreqs, WordFreq{k, v})
	}

	sort.Slice(wordFreqs, func(i, j int) bool {

		return wordFreqs[i].freq > wordFreqs[j].freq
	})

	for i := 0; i < 10; i++ {
		fmt.Println(wordFreqs[i])
	}
}
