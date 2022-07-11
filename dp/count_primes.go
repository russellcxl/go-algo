package dp

// return no. of prime numbers < n
// e.g. n = 10, primes = 2, 3, 5, 7
func countPrimes(n int) int {
	nonPrimes := make([]bool, n)

	// start i as the first prime number: 2
	// keep multiplying the first number with the second to get all non-primes
	for i := 2 ; i * i < n ; i++ {
		if !nonPrimes[i] {
			for j := i ; i * j < n ; j++ {
				nonPrimes[i * j] = true
			}
		}
	}

	// count the number of primes
	var count int
	for i := 2 ; i < n ; i++ {
		if !nonPrimes[i] { count++ }
	}

	return count
}