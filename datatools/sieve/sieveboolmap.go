package sieve

func BoolIt(max int) (primes []int) {
	size := MaxSearchValue(max)
	var boolMap map[int]bool = make(map[int]bool, size)

	// ... do stuff

	// ...
	// count true to preallocate
	counter := 0
	for _, v := range boolMap {
		if v {
			counter++
		}
	}
	outList := make([]int, 0, counter)
	for k, v := range boolMap {
		if v {
			outList = append(outList, k)
		}
	}
	return outList

	// if we see a number on the channel, mark it ...
}
