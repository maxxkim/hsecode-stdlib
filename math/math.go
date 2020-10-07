package math

import "math"

func NthPrime(n int) int {
	prime := 2
	for i, count := 3, 1; count < n; i = i + 2 {
		if IsPrime(i) {
			prime = i
			count = count + 1
		}
	}
	return prime
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	max := int(math.Sqrt(float64(n))) + 1
	for i := 1; 6*i-1 < max; i++ {
		if n%(6*i+1) == 0 || n%(6*i-1) == 0 {
			return false
		}
	}
	return true
}