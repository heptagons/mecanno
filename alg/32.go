package alg

import (
	"math"
)

type N32 uint32 // range 0 - 0xffffffff

type I32 struct {
  s bool // sign: true means negative
  n N32  // positive value
}

type AI32 struct {
  o *I32  // outside
  i *I32  // inside
  e *AI32 // inside extension
}

type AI32s struct {
	primes []N32
}







const N32_MAX = N(0xffffffff)
//const MAXNAT = uint64(0xffffffff)



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

// NewAI32s creates algebraic integers factory
// A fixed internal 32-bit primes list is used to create
// algebraic integers reduced
func NewAI32s() *AI32s {
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
	return &AI32s{
		primes: primes,
	}
}

func (r *AI32s) AI(out, in Z, ext *AI32) *AI32 {
	if ext == nil {
		if out == 0 {
			return NewAI32zero() // zero
		} else if o, i, ok := r.reduce1Z(out, in); ok {
			return NewAI32(o, i, nil)
		}
		return nil
	}
	eo := ext.outVal()
	ei := ext.inVal()
	if o, ia, ib, ok := r.reduce2Z(out, in, eo); ok {
		return NewAI32(o, ia, r.AI(ib.val(), ei, ext.e))
	}
	return nil
}


func (r *AI32s) reduce1Z(out, inA Z) (*I32, *I32, bool) {
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
	// TODO only here because we have the ra sign
	// if ra = +squared (not integer imaginary)
	// set ro = ro + √(ra)
	// set ra = nil
	return ro, ra, true
}

func (r *AI32s) reduce2Z(out, inA, inB Z) (*I32, *I32, *I32, bool) {
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
func (r *AI32s) reduce1(out, in N) (o N32, i N32, ok bool) {
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
func (r *AI32s) reduce2(out, inA, inB N) (N32, N32, N32, bool) {
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




func NewAI32zero() *AI32 {
	return &AI32{
		o: nil,
		i: nil,
	}
}

func NewAI32(out, in *I32, ext *AI32) *AI32 {
	return &AI32{
		o: out,
		i: in,
		e: ext,
	}
}

func (r *AI32) outVal() Z {
	return r.o.val()
}

func (r *AI32) outSet(out *I32) {
	r.o = out.clone()
}

func (r *AI32) outValPow2() Z {
	return r.o.valPow2()
}

func (r *AI32) inVal() Z {
	return r.i.val()
}

func (r *AI32) IsZero() bool {
	if r == nil || r.o == nil || r.i == nil || r.o.n == 0 || r.i.n == 0 {
		return true
	}
	return false
}

// WriteString appends to given buffer very SIMPLE format:
// For nil, out or in zero appends "+0"
// For n > 0 always appends +n or -n including N=1
// For in > 1 appends √ and then in (always positive)
func (r *AI32) Str(s *Str) {
	if r == nil {
		s.Infinite()
	} else if r.o == nil || r.o.n == 0 {
		s.Zero()
	} else if r.i == nil || r.i.n == 0 {
		s.Zero()
	} else {
		s.I32(r.o)
		if r.e == nil {
			s.Radical32(r.i, nil)
		} else {
			s.Radical32(r.i, func(s *Str) {
				r.e.Str(s)
			})
		}
	}
}

func (r *AI32) String() string {
	s := NewStr()
	r.Str(s)
	return s.String()
}


