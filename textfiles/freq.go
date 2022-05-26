package textfiles

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func Frequency(fileName string) {
	bs, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	text := B2S(bs)

	fields := strings.FieldsFunc(text, func(r rune) bool {
		// return !('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || r == '\'')
		return ('a' > r || r > 'z') && ('A' > r || r > 'Z') && r != '\''
	})

	wordsCount := make(map[string]int)
	for _, field := range fields {
		wordsCount[field]++
	}

	keys := make([]string, 0, len(wordsCount))
	for key := range wordsCount {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return wordsCount[keys[i]] > wordsCount[keys[j]]
	})

	for idx, key := range keys {
		fmt.Printf("%s %d\n", key, wordsCount[key])
		if idx == 10 {
			break
		}
	}
}
