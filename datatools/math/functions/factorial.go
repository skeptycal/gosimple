package functions

func Factorial(n int) int {
	return factorial(n)
}

func factorial(n int) (retval int) {
	retval = 1
	for i := 1; i <= n; i++ {
		retval *= i
	}
	return
}

func factorialIteration(n int) (retval int) {
	retval = 1
	for n >= 1 {
		retval *= n
		n -= 1
	}
	return
}

func factorialIteration2(n int) (retval int) {
	if n < 2 {
		return 1
	}
	a, b := 1, 1
	_ = a
	for n >= 1 {
		a, b = b, n*b
		n -= 1
	}
	return b
}

func factorialRecursive(n int) int {
	if n < 2 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func factorialPtr(n int) int {
	// var nptr *int
	// var retvalptr *int
	retval := 1
	retvalptr := &retval
	nptr := &n

	for i := 1; i <= *nptr; i++ {
		*retvalptr *= i
	}
	return retval
}

func factorialLoop(n int) int {
	if n < 2 {
		return n
	}
	sum := 0
	for i := 1; i <= n; i++ {
		sum *= i
	}
	return sum
}

func FactorialClosure() func(n int) int {
	var a, b = 1, 1
	_ = a
	return func(n int) int {
		if n > 1 {
			a, b = b, n*b
		} else {
			return 1
		}

		return b
	}
}

const LIM = 256

var facts [LIM]int

func FactorialMemoization(n int) (retval int) {
	if n < 2 {
		return 1
	}
	if facts[n] != 0 {
		retval = facts[n]
		return retval
	}

	if n > 0 {
		retval = n * FactorialMemoization(n-1)
		return retval
	}

	return 1
}
