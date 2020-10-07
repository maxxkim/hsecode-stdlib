package math

import (
	"math"
)

func NthPrime(n int) int {
	switch {
	case n <= 0:
		panic("Panicking")
	case n <= 141:
		var firstPrimes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
			73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179,
			181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283,
			293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419,
			421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541, 547,
			557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659, 661,
			673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811}
		return firstPrimes[n-1]

	default:
		logarifm := math.Log(float64(n))
		maxSize := int(float64(n) * (logarifm + math.Log(logarifm)))
		isPrime := make([]bool, (maxSize+2)/2)
		for i := 3; i < int(math.Sqrt(float64(maxSize))); i += 2 {
			if !isPrime[i/2-1] {
				for j := i * i; j < maxSize; j += i {
					if j%2 == 0 {
						continue
					}
					isPrime[j/2-1] = true
				}
			}

		}
		count := 1
		for i, v := range isPrime {
			if !v {
				count++
			}
			if n == count {
				return (i+1)*2 + 1
			}
		}
	}
	return 0
}
