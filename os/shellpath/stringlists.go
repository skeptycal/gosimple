package shpath

import "sort"

func RemoveOrdered(s []string, pos int) []string {
	return append(s[:pos], s[pos+1:]...)
}

func RemoveUnOrdered(s []string, n int) []string {
	s[n] = s[len(s)-1]
	// s[len(s)-1] = "" // clear the last element before removing it ... maybe helps with GC?
	return s[:len(s)-1]
}

// Append appends a string to the end of a list of strings.
func Append(list []string, s string) []string { return append(list, s) }

// Insert inserts a string at pos position. If pos is invalid,
// s is appended to the end of the list. If s is the empty string,
// the original list is returned unchanged.
func Insert(list []string, s string, pos int) []string {
	if s == "" {
		return list
	}
	if pos < 0 || pos >= len(list) {
		return append(list, s)
	}

	temp := make([]string, 0, len(list)+1) // preallocating is faster on most benchmarks

	copy(temp, list[:pos]) // copy is faster than append on most benchmarks
	// temp = append(temp, s)
	copy(temp, []string{s})
	copy(temp, list[pos+1:])

	list = nil // maybe helps GC?

	return temp
}

// InsertSorted inserts s into the list in sorted order.
// func InsertSorted(list []string, s string) []string {
// 	for i, str := range list {
// 		if str > s {
// 			break
// 		}
// 	}
// }

func insertSort(haystack []string, needle string) []string {
	index := sort.Search(len(haystack), func(i int) bool { return haystack[i] > needle }) // or >= ??
	haystack = append(haystack, "")
	copy(haystack[index+1:], haystack[index:])
	haystack[index] = needle
	return haystack
}

func insertSort2(haystack []string, needle string) []string {
	newlist := make([]string, 0, len(haystack)+1)
	index := iLoc(haystack, needle)
	// haystack = append(haystack, "")
	// copy(haystack[index+1:], haystack[index:])
	// haystack[index] = needle
	copy(newlist, haystack[:index])
	copy(newlist, []string{needle})
	copy(newlist, haystack[index+1:])

	return haystack
}

// iLoc returns the location where needle should be inserted
// in haystack to maintain the sort order.
func iLoc(haystack []string, needle string) int {
	return sort.Search(len(haystack), func(i int) bool { return haystack[i] >= needle })
}

// iLoc returns the location where needle should be inserted
// in haystack to maintain the sort order.
func iLoc2(haystack []string, needle string) int {
	return sort.Search(len(haystack), func(i int) bool { return haystack[i] > needle })
}

func isItemSorted(haystack []string, needle string) bool {
	i := sort.Search(len(haystack), func(i int) bool { return haystack[i] >= needle })
	if i < len(haystack) && haystack[i] == needle {
		// x is present at data[i]
		return true
	} else {
		// x is not present in data,
		// but i is the index where it would be inserted.
		return false
	}
}
