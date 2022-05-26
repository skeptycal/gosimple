package textfiles

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

func OpenTextFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}

func String(b []byte) string            { return B2S(b) }
func Bytes(s string) []byte             { return S2B(s) }
func Lines(s string) []string           { return strings.Split(s, newline) }
func Words(s string) []string           { return strings.Fields(s) }
func ToValidUTF8(s string) string       { return strings.ToValidUTF8(s, string(unicode.ReplacementChar)) }
func RemoveNewlines(text string) string { return strings.Join(strings.Split(text, newline), " ") }

func SortLines(lines []string) []string {
	if !sort.StringsAreSorted(lines) {
		sort.Strings(lines)
	}
	return lines
}

func SortWords(words []string) []string {
	if !sort.StringsAreSorted(words) {
		sort.Strings(words)
	}
	return words
}

// LineSet returns a map of line numbers to lines.
func LineMap(s string) map[int]string {
	lines := Lines(s)
	m := make(map[int]string, len(lines))

	for i, line := range lines {
		m[i] = line
	}

	return m
}

type word_struct struct {
	freq int
	word string
}

// word_struct will be displayed in this format
func (p word_struct) String() string {
	return fmt.Sprintf("%3d   %s", p.freq, p.word)
}

// by_freq implements sort.Interface for []word_struct based on the freq field
type by_freq []word_struct

func (a by_freq) Len() int           { return len(a) }
func (a by_freq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a by_freq) Less(i, j int) bool { return a[i].freq < a[j].freq }

// by_word implements sort.Interface for []word_struct based on the word field
type by_word []word_struct

func (a by_word) Len() int           { return len(a) }
func (a by_word) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a by_word) Less(i, j int) bool { return a[i].word < a[j].word }

func WordFrequency(words []string) map[string]int {
	m := make(map[string]int, len(words)/10) // TODO: what is the best initial size?
	for _, word := range words {
		m[word]++
	}
	return m
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

//CleanWords matches whole words and removes any punctuation/whitespace.
func CleanWords(text string) []string {
	re := regexp.MustCompile("\\w+")
	return re.FindAllString(text, -1)
}

func stuff(text string) {

	text = ToLower(text)
	words := CleanWords(text)

	// word_map() returns a map of word:frequency pairs
	word_map := WordFrequency(words)

	// convert the map to a slice of structures for sorting
	// create pointer to an instance of word_struct
	pws := new(word_struct)
	struct_slice := make([]word_struct, len(word_map))
	index := 0
	for k, v := range word_map {
		pws.freq = v
		pws.word = k
		// test, %+v shows field names
		//fmt.Printf("%v %v  %+v\n", pws.freq, pws.word, *pws)
		struct_slice[index] = *pws
		index++
	}
	// testing ...
	//fmt.Printf("%+v\n", struct_slice[0])
	//fmt.Printf("%+v\n", struct_slice[1])
	//fmt.Printf("%v\n", struct_slice[1].freq)
	//fmt.Printf("%v\n", struct_slice[1].word)

	//fmt.Println("-------------")

	fmt.Println("Words in text sorted by frequency:")
	// sorting slice of structers by field freq in place
	sort.Sort(by_freq(struct_slice))
	for i := 0; i < len(struct_slice); i++ {
		fmt.Printf("%v\n", struct_slice[i])
	}

	fmt.Println("-------------")

	fmt.Println("Words in text sorted by word:")
	// sorting slice of structures by field word in place
	sort.Sort(by_word(struct_slice))
	for i := 0; i < len(struct_slice); i++ {
		fmt.Printf("%v\n", struct_slice[i])
	}
}
