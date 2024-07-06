package snowflakex

import (
	"math"
	"testing"
)

func findPrime(n int) []int {
	if n <= 1 {
		return nil
	}

	var primes []int
	for prime := 2; prime < n; prime++ {
		isPrime := true
		for j := 2; j <= int(math.Sqrt(float64(n))); j++ {
			if prime%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, prime)
		}
	}

	return primes
}

func TestFindPrimes(t *testing.T) {
	t.Logf("%v", findPrime(10000))
}
