package calculator

// IsPrime returns true if the number is prime
func IsPrime(n int) (ok bool) {
	// constants
	if n == -1 || n == 0 || n == 1 {
		return
	}

	// loop
	for i := 2; i*i <= n; i++ {
		// check if the number is divisible by i
		if n%i == 0 {
			return
		}
	}

	ok = true
	return
}