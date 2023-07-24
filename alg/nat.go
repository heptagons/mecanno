package alg

import (
	"math"
)

const MAXNAT = uint64(0xffffffff)

// Nat represents a small natural number in the range 1 - 0xffffffff
type Nat uint32

// NatGCD returns the greatest common divisor of two naturals
func NatGCD(a, b Nat) Nat {
	if b == 0 {
		return a
	}
	return NatGCD(b, a % b)
}

// NatPrimes returns a list of the first primes
// Use the Sieve of Erathostenes
func NatPrimes() []Nat {
	value := 0xffff
    f := make([]bool, value)
    for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
        if f[i] == false {
            for j := i * i; j < value; j += i {
                f[j] = true
            }
        }
    }
    list := make([]Nat, 0)
    for i := Nat(2); i < Nat(value); i++ {
        if f[i] == false {
            list = append(list, i)
        }
    }
    return list
}

type Nats struct {
	primes []Nat
}


func NewNats() *Nats {
	return &Nats{
		primes: NatPrimes(),
	}
}

// SqrtMul returns the square root of the product of two naturals a,b
// Returns two naturals c,d as the simplification:
//	√(a*b) = (c)√(d)
// ok is false for overflow of out or in.
func (n *Nats) sqrtMul(a, b Nat) (c Nat, d Nat, ok bool) {
	in := uint64(a) * uint64(b)
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
		c = Nat(out)
		d = Nat(in)
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
func (n *Nats) Sqrt(r *Rat) *Alg {
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
func (n *Nats) Mul(a, b *Alg) *Alg {
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








