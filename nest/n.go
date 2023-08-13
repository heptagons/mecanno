package nest

import (
	"fmt"
	"math"
)

var ErrOverflow = fmt.Errorf("Overflow")
var ErrInfinite = fmt.Errorf("Infinite")
var ErrInvalid  = fmt.Errorf("Invalid")

type N uint64

func Ngcd(a, b N) N {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a > b {
		return Ngcd(b, a % b)
	}
	return Ngcd(a, b % a)
}

func (a *N) Reduce2(b *Z) N {
	var bb N
	if *b < 0 {
		bb = N(-*b)
	} else {
		bb = N(*b)
	}
	g := Ngcd(*a, bb)
	if g > 1 {
		*a /= g
		*b /= Z(g)
	}
	return g
}

func (a *N) Reduce3(b, c *Z) N {
	var bb, cc N
	if *b < 0 {
		bb = N(-*b)
	} else {
		bb = N(*b)
	}
	if *c < 0 {
		cc = N(-*c)
	} else {
		cc = N(*c)
	}
	g := Ngcd(Ngcd(*a, bb), cc)
	if g > 1 {
		*a /= g
		*b /= Z(g)
		*c /= Z(g)
	}
	return g
}

func (a *N) Reduce4(b, c, e *Z) N {
	var bb, cc, ee N
	if *b < 0 {
		bb = N(-*b)
	} else {
		bb = N(*b)
	}
	if *c < 0 {
		cc = N(-*c)
	} else {
		cc = N(*c)
	}
	if *e < 0 {
		ee = N(-*e)
	} else {
		ee = N(*e)
	}
	g := Ngcd(Ngcd(Ngcd(*a, bb), cc), ee)
	if g > 1 {
		*a /= g
		*b /= Z(g)
		*c /= Z(g)
		*e /= Z(g)
	}
	return g
}


const N32_MAX = N(0xffffffff)

type N32 uint32 // range 0 - 0xffffffff

// gcd returns the greatest common divisor of 
// this natural and the other given
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

func (a *N32) Reduce2(b *N32) {
	if g := NatGCD(*a, *b); g > 1 {
		*a /= g
		*b /= g
	}
}

func (a *N32) Reduce3(b, c *N32) {
	if g := NatGCD(NatGCD(*a, *b), *c); g > 1 {
		*a /= g
		*b /= g
		*c /= g
	}
}


// N32s is factory with a primes list to speed up
// some 32-bit nested algebraic rational numbers
type N32s struct {
	primes []N32
}

func NewN32s() *N32s {
	value := 0xffff
    f := make([]bool, value)
    for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
        if f[i] == false {
            for j := i * i; j < value; j += i {
                f[j] = true
            }
        }
    }
    primes := make([]N32, 0)
    for i := N32(2); i < N32(value); i++ {
        if f[i] == false {
            primes = append(primes, i)
        }
    }
	return &N32s{
		primes: primes,
	}
}
