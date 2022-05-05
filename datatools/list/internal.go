package list

// Append appends an item to the end of a list. This a redundant
// wrapper around the built-in append and is used for testing purposes.
func Append[T any, E ~[]T](list E, item T) E { return append(list, item) }
