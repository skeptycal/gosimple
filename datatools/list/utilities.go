package list

func prependPreAlloc[E any, S ~[]E](element E, list S) S {
	var chain = make(S, 1, len(list)+1)
	chain[0] = element
	chain = append(chain, list...)
	return chain
}

func prependOneLinePreAlloc[E any, S ~[]E](element E, list S) S {
	chain := make(S, 1, len(list)+1)
	chain[0] = element
	return append(chain, list...)
}

func prependOneLine[E any, S ~[]E](element E, list S) S {
	return append(S{element}, list...)
}
