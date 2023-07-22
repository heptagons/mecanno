package angles

import (
	"math"
)

// PrimesList returns a list of the first primes
// Use the Sieve of Erathostenes
// If value = 0xffff, the returns a list of primes < 32 bits)
func PrimesList(value int) []uint {
    f := make([]bool, value)
    for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
        if f[i] == false {
            for j := i * i; j < value; j += i {
                f[j] = true
            }
        }
    }
    primes := make([]uint, 0)
    for i := uint(2); i < uint(value); i++ {
        if f[i] == false {
            primes = append(primes, i)
        }
    }
    return primes
}
