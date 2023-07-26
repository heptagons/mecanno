package alg

import (
	"math"
)

type Int int32

const MAXNAT = uint64(0xffffffff)

// N32 represents a small natural number in the
// range 1 - 0xffffffff
type N32 uint32


// gcd returns the greatest common divisor of 
// this natural and the given
func (a N32) gcd(b N32) N32 {
	if b == 0 {
		return a
	}
	return b.gcd(a % b)
}

// NatGCD returns the greatest common divisor of two naturals
func NatGCD(a, b N32) N32 {
	if b == 0 {
		return a
	}
	return NatGCD(b, a % b)
}

// NatPrimes returns a list of the first primes
// Use the Sieve of Erathostenes
func NatPrimes() []N32 {
	value := 0xffff
    f := make([]bool, value)
    for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
        if f[i] == false {
            for j := i * i; j < value; j += i {
                f[j] = true
            }
        }
    }
    list := make([]N32, 0)
    for i := N32(2); i < N32(value); i++ {
        if f[i] == false {
            list = append(list, i)
        }
    }
    return list
}

type N32s struct {
	primes []N32
}

func NewN32s() *N32s {
	return &N32s{
		primes: NatPrimes(),
	}
}

/*func (n *N32s) f1(c, d Z) *R {
	out, in, ok := n.Sqrt32(d)
	if !ok {
		return nil
	}
	return nil
}*/

// Sqrt32 reduces the radical (out)√(in) in in two parts
// Example: -3√(20) returns -6√(5)
// Return ok as false when returned values are larger than 32 bits (overflow).
func (n *N32s) Sqrt32(out, in uint64) (o N32, i N32, ok bool) {
	if out == 0 {
		return 0, 0, true
	}
	if in == 0 {
		return 0, 0, true
	}
	if in > 1 {
		// Try to modify out and in
		for _, prime := range n.primes {
			p := uint64(prime)
			if pp := p*p; in >= pp {
				for {
					if in % pp == 0 {
						// product has a prime factor squared move from in to out
						in  /= pp
						out *= p
						// look for more prime factor squared repeated in in.
						continue
					} else {
						// check with next prime squared
						break
					}
				}
			} else {
				// no more factors to check
				break
			}
		}
	}
	if out > MAXNAT {
		return 0, 0, false
	}
	if in > MAXNAT {
		return 0, 0, false
	}
	return N32(out), N32(in), true
}
