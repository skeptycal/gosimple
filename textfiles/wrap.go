package textfiles

import "bytes"

func ReWrap(text string, size int) string {
	// words := bytes.Fields(S2B(RemoveNewlines(text)))
	words := bytes.Split(S2B(text), []byte(" "))

	newtext := make([]byte, len(text)+8)

	counter := 0
	for i := 0; i < len(words); i++ {
		word := words[i]
		newtext = append(newtext, word...)
		newtext = append(newtext, ' ')
		counter += len(word) + 1
		if counter >= size {
			newtext = append(newtext, '\n')
			counter = 0
		}
	}

	return B2S(newtext)
}
