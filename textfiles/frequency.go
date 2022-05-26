package textfiles

import (
	"fmt"
	"unicode"

	"github.com/skeptycal/gosimple/cmd/goutil_playground/sort"
	"github.com/skeptycal/gosimple/types/constraints"
)

func (t *TextFile) Frequency() map[string]int {
	t.m = make(map[string]int, len(t.words)/4) // TODO what initial size is best?
	for _, word := range t.Words() {
		if word != "" {
			t.m[word]++
		}
	}

	return t.m
}

func (t *TextFile) WordCount() (int, int) {
	m := t.Frequency()
	sum := 0
	for _, v := range m {
		sum += v
	}
	return len(m), sum
}

func (t *TextFile) Top(n int) {
	m := t.Frequency()
	ms := KeySort(m)

	fmt.Println("aid: ", m["aid"])
	sort.Sort(ms)
	fmt.Println(ms)
	sort.Sort(ms)
	fmt.Println(ms)
	for i := 0; i < n; i++ {
		c := len(ms) - i - 1
		fmt.Printf("%3d: %v\n", c, ms[c])
	}
}

type Ordered constraints.Ordered

type tuple[K Ordered, V Ordered] struct {
	k K
	v V
}

type MapSlice[K Ordered, V Ordered] []tuple[K, V]

// TODO this sort interface doesn't seem to work ... Len(), Less(), Swap() ... hmm .. not sure
func (m MapSlice[K, V]) Len() int           { return len(m) }
func (m MapSlice[K, V]) Less(i, j int) bool { return m[i].k < m[i].k }
func (m MapSlice[K, V]) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MapSlice[K, V]) Sort() {
	// if !sort.IsSorted(m) {
	sort.Sort(m)
	// }
}

func Values(m map[string]int) []int {
	values := make([]int, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	sort.Ints(values)
	return values
}

func Keys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func KeySort(m map[string]int) MapSlice[string, int] {
	keys := Keys(m)
	sort.Strings(keys)

	out := make(MapSlice[string, int], 0, len(m))

	for _, k := range keys {
		out = append(out, tuple[string, int]{k, m[k]})
	}

	return out
}

func cleanWord(s string) string {
	newword := []byte{}
	for _, c := range S2B(s) {
		if unicode.IsLetter(rune(c)) {
			newword = append(newword, c)
		}
	}
	return B2S(newword)
}

// CleanAlphaNumeric cleans the word list and returns
// only items that are alpha-numeric.
func (t *TextFile) CleanAlphaNumeric() {
	list := t.Words()
	for i := 0; i < len(list); i++ {
		w := cleanWord(list[i])
		if w != "" {
			list[i] = w
		} else {
			list[i] = ""
		}
	}

	t.words = list
}
