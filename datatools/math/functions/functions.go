package functions

func factorialr(n int) (retval int) {
	retval = 1
	for i := 1; i <= n; i++ {
		retval = retval * i
	}
	return
}
