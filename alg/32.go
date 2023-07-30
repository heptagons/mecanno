package alg

import (
	"math"
//	"fmt"
)

type Int int32

const N32_MAX = N(0xffffffff)
//const MAXNAT = uint64(0xffffffff)

// N32 represents a small natural number in the
// range 1 - 0xffffffff
type N32 uint32

func N32overflowZ(z Z) bool {
	if z > 0 {
		return z > Z(N32_MAX)
	} else {
		return -z > Z(N32_MAX)
	}
}

func N32overflowN(n N) bool {
	return n > N32_MAX
}


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


// I is an integer of 32 bits
type I32 struct {
	s bool
	n N32
}

func NewI32(z Z) (*I32, bool) {
	if N32overflowZ(z) {
		return nil, false
	} else if z >= 0 {
		return newI32plus(N32(z)), true
	} else {
		return newI32minus(N32(-z)), true
	}
}

func newI32plus(n N32) *I32 {
	return &I32{
		s: false,
		n: n,
	}
}

func newI32minus(n N32) *I32 {
	return &I32{
		s: true,
		n: n,
	}	
}

func (i *I32) clone() *I32 {
	if i == nil {
		return nil
	} else {
		return &I32 {
			s: i.s,
			n: i.n,
		}
	}
}

func (i *I32) mul(n N32) Z {
	if i.s {
		return - Z(n)*Z(i.n)
	} else {
		return + Z(n)*Z(i.n)
	}
}

func (i *I32) val() Z {
	if i == nil {
		return 0
	} else if i.s {
		return -Z(i.n)
	} else {
		return +Z(i.n)
	}
}

func (i *I32) valPow2() Z {
	if i == nil || i.n == 0 {
		return 0
	} else {
		return Z(i.n)*Z(i.n)
	}
}

func (i *I32) Str(s *Str) {
	if i == nil || i.n == 0 {
		s.Zero()
	} else {
		s.I32(i)
	}
}

func (i *I32) String() string {
	str := &Str{}
	str.I32(i)
	return str.String()
}


type R32s struct {
	primes []N32
}

func NewR32s() *R32s {
	return &R32s{
		primes: NatPrimes(),
	}
}

func (r *R32s) Radical(out, in Z, ext *R32) *R32 {
	if ext == nil {
		if out == 0 || in == 0 {
			return NewR32zero() // zero
		} else if o, i, ok := r.reduce1Z(out, in); ok {
			return NewR32(o, i, nil)
		}
		return nil
	}
	e := out
	f := in
	g := ext.outVal()
	h := ext.inVal()
	if o, a, b, ok := r.reduce2Z(e, f, g); ok {
		return NewR32(o, a, r.Radical(b.val(), h, ext.ext))
	}
	return nil
}


func (r *R32s) reduce1Z(out, inA Z) (*I32, *I32, bool) {
	io := out; if out < 0 { io = -out }
	ia := inA; if inA < 0 { ia = -inA }
	o, a, ok := r.reduce1(N(io), N(ia))
	if !ok {
		return nil, nil, false
	}
	zo := Z(o); if out < 0 { zo = -zo }
	za := Z(a); if inA < 0 { za = -za }
	ro, _ := NewI32(zo)
	ra, _ := NewI32(za)
	return ro, ra, true
}

func (r *R32s) reduce2Z(out, inA, inB Z) (*I32, *I32, *I32, bool) {
	io := out; if out < 0 { io = -out }
	ia := inA; if inA < 0 { ia = -inA }
	ib := inB; if inB < 0 { ib = -inB }
	o, a, b, ok := r.reduce2(N(io), N(ia), N(ib))
	if !ok {
		return nil, nil, nil, false
	}
	zo := Z(o); if out < 0 { zo = -zo }
	za := Z(a); if inA < 0 { za = -za }
	zb := Z(b); if inB < 0 { zb = -zb }
	ro, _ := NewI32(zo)
	ra, _ := NewI32(za)
	rb, _ := NewI32(zb)
	return ro, ra, rb, true
}

// reduce try to decrease in and increase out.
// Example: The input -3√(20) is returned as -6√(5)
// Return ok as false when out or in values are larger than 32 bits (overflow).
func (r *R32s) reduce1(out, in N) (o N32, i N32, ok bool) {
	if out == 0 {
		return 0, 0, true
	}
	if in == 0 {
		return 0, 0, true
	}
	if in > 1 {
		// Try to update out and in
		for _, prime := range r.primes {
			p := N(prime)
			if pp := p*p; in >= pp {
				for {
					if in % pp == 0 {
						// product has a prime factor squared move from in to out
						out *= p
						in  /= pp
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
	if N32overflowN(out) {
		return 0, 0, false
	} else if N32overflowN(in) {
		return 0, 0, false
	}
	return N32(out), N32(in), true
}

// reduce2 try to increase given out by decrease given inA and inB.
// Return reduced out and ins and a flag for 32 bits overflow.
// Example:
//	For given out=5, inA=12=2*2*3, inB=56=2*2*2*7
//  Returns out=10=5*2, inA=3, inB=14=2*7, true
func (r *R32s) reduce2(out, inA, inB N) (N32, N32, N32, bool) {
	if inA > 1 && inB > 1 {
		for _, prime := range r.primes {
			p := N(prime)
			pp := p*p
			for {
				if inA % pp == 0 && inB % pp == 0 {
					out *= p  // multiply by p
					inA /= pp // divide by p squared
					inB /= pp // divide by p squared
					continue
				} else {
					break
				}
			}
		}
	}
	if N32overflowN(out) {
		return 0, 0, 0, false
	} else if N32overflowN(inA) {
		return 0, 0, 0, false
	} else if N32overflowN(inB) {
		return 0, 0, 0, false
	}
	return N32(out), N32(inA), N32(inB), true
}



type R32 struct {
	out *I32 // external integer with sign=true means whole R32 is negative
	in  *I32 // internal integer with sign=true means whole R32 is imaginary
	ext *R32
}

func NewR32zero() *R32 {
	return &R32{
		out: nil,
		in:  nil,
	}
}

func NewR32(out, in *I32, ext *R32) *R32 {
	return &R32{
		out: out,
		in:  in,
		ext: ext,
	}
}

func (r *R32) outVal() Z {
	return r.out.val()
}

func (r *R32) outSet(out *I32) {
	r.out = out.clone()
}

func (r *R32) outValPow2() Z {
	return r.out.valPow2()
}

func (r *R32) inVal() Z {
	return r.in.val()
}

func (r *R32) IsZero() bool {
	if r == nil || r.out == nil || r.in == nil || r.out.n == 0 || r.in.n == 0 {
		return true
	}
	return false
}

// WriteString appends to given buffer very SIMPLE format:
// For nil, out or in zero appends "+0"
// For n > 0 always appends +n or -n including N=1
// For in > 1 appends √ and then in (always positive)
func (r *R32) Str(s *Str) {
	if r == nil {
		s.Infinite()
	} else if r.out == nil || r.out.n == 0 {
		s.Zero()
	} else if r.in == nil || r.in.n == 0 {
		s.Zero()
	} else {
		s.I32(r.out)
		if r.ext == nil {
			s.Radical32(r.in, nil)
		} else {
			s.Radical32(r.in, func(s *Str) {
				r.ext.Str(s)
			})
		}
	}
}

func (r *R32) String() string {
	s := NewStr()
	r.Str(s)
	return s.String()
}


