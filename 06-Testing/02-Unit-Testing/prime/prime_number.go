package prime

func IsPrime(n int) (isPrime bool) {
	// constraint
	if n <= 1 {
		return
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return
		}
	}

	isPrime = true
	return
}