package list

import (
	"sort"

	"github.com/skeptycal/gosimple/types/constraints"
)

// iLoc returns the location where needle should be inserted
// in haystack to maintain the sort order.
func iLoc[T constraints.Ordered, E ~[]T](haystack E, needle T) int {
	return sort.Search(len(haystack), func(i int) bool { return haystack[i] >= needle })
}

// iLoc returns the location where needle should be inserted
// in haystack to maintain the sort order.
func iLoc2[T constraints.Ordered, E ~[]T](haystack E, needle T) int {
	return sort.Search(len(haystack), func(i int) bool { return haystack[i] > needle })
}

func isItemSorted[T constraints.Ordered, E ~[]T](haystack E, needle T) bool {
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
