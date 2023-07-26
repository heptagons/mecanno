package alg

import (
	"math"
)

type Int int32

const MAXNAT = uint64(0xffffffff)

// N32 represents a small natural number in the range 1 - 0xffffffff
type N32 uint32


// gcd returns the greatest common divisor of two numbers
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


func NewNats() *N32s {
	return &N32s{
		primes: NatPrimes(),
	}
}

// SqrtMul returns the square root of the product of two naturals a,b
// Returns two naturals c,d as the simplification:
//	√(a*b) = (c)√(d)
// ok is false for overflow of out or in.
func (n *N32s) sqrtMul(a, b N32) (N32, N32, bool) {
	in := uint64(a) * uint64(b)
	return n.sqrt(in)
}

func (n *N32s) sqrt(in uint64) (c N32, d N32, ok bool) {
	if in == 0 {
		ok = true
		return // zero 0√0 ok
	}
	if in == 1 {
		c = 1
		d = 1
		ok = true
		return // one 1√1 ok
	}
	out := uint64(1)
	for _, prime := range n.primes {
		p := uint64(prime)
		if pp := p*p; in >= pp {
			for {
				if in % pp == 0 {
					// product has a prime factor squared
					// move from in to out
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
	if in > MAXNAT {
		return // not ok OVERFLOW
	} else if out > MAXNAT {
		return // not ok OVERFLOW
	} else {
		c = N32(out)
		d = N32(in)
		ok = true
		return // (out)√(in) ok
	}
}

// Sqrt returns the square root of the rational r as an algebraic.
// returns nil when rational is negative (imaginary) or natural overflow.
//
//	            sqrt(rat.Num*rat.Den)     out
//	sqrt(rat) = --------------------- = ------- * sqrt(in)
//	                  rat.Den           rat.Den
func (n *N32s) Sqrt(r *Rat) *Alg {
	if r == nil {
		return nil // invalid rational
	}
	if r.Neg {
		return nil // Imaginary
	}
	if out, in, ok := n.sqrtMul(r.Num, r.Den); !ok {
		return nil // overflow
	} else if r2 := NewRat(int(out), int(r.Den)); r2 == nil {
		// update rational since sqrtMul updated numerator which
		// can be simplied with previos denominator
		return nil // rare
	} else {
		return NewAlg(r2, in)
	}
}

// Mul multiply two algebraic numbers a and b to return another algrebraic
// returns nil if a or b are nil or after operation overflow
func (n *N32s) Mul(a, b *Alg) *Alg {
	if a == nil || b == nil {
		return nil
	} else if out, in, ok := n.sqrtMul(a.In, b.In); !ok {
		return nil
	} else {
		return &Alg{
			Rat: a.Mul(NewRat(int(out), 1)).Mul(b.Rat),
			In: in,
		}
	}
}








