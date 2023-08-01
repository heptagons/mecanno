package alg

import (
	"math"
	//"fmt"
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

// newI32 returns a 32 bit integer
// for z = 0 returns nil, true
// for overflow return nil, false
// for 0 < z <= N32_MAX return positive i32, true
// for 0 < -z <= N32_MAX return negative i32, true
func newI32(z Z) (*I32, bool) {
	if z == 0 {
		return nil, true // zero
	}
	if z > 0 {
		if N(z) > N32_MAX {
			return nil, false // overflow
		}
		return &I32{ s:false, n:N32(z) }, true // positive
	}
	if N(-z) > N32_MAX {
		return nil, false // overflow
	}
	return &I32{ s:true, n:N32(-z) }, true // negative
}

func (i *I32) clone() *I32 {
	if i == nil {
		return nil
	}
	return &I32 {
		s: i.s,
		n: i.n,
	}
}

func (i *I32) mul(n N32) Z {
	if i == nil || i.n == 0 || n == 0 {
		return 0
	}
	if i.s {
		return - Z(n)*Z(i.n)
	}
	return + Z(n)*Z(i.n)
}

func (i *I32) val() Z {
	if i == nil || i.n == 0 {
		return 0
	}
	if i.s {
		return -Z(i.n)
	}
	return +Z(i.n)
}

func (i *I32) valPow2() Z {
	if i == nil || i.n == 0 {
		return 0
	}
	return Z(i.n)*Z(i.n) // always positive
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

func (r *AI32s) AI(out, in Z, ext *AI32) (ai *AI32, overflow bool) {
	if ext == nil {
		// out√in is c√d or g√h or m√n ...
		return r.roi(out, in)
	}
	eo1 := ext.outVal()
	if eo1 == 0 {
		// out√in ext.o = 0
		// e√f g=0 or i√j k=0 or k√l m=0 or ...
		return r.roi(out, in)
	}
	oi, ii, eo2, ok := r.roie(out, in, eo1)
	if !ok {
		// c, dd, g, hh, m, nn overflow
		return nil, true
	
	}
	if ext.e == nil {
		if ext.inVal() == +1 {
			// e√(f+g√h) or k√(l+m√n) ... where h == +1, n == +1, ...
			// return reduced e√(f+g) or k√(l+m) ...
			return r.roi(oi.val(), ii.val() + eo2)
		}
	}
	// i√(j+k√l+(...)) or ...
	ext, over2 := r.AI( // extension
		eo2,         // equals to eoi or reduced (decreased)
		ext.inVal(), // extra-in original expected irreducible
		ext.e,       // extra-extra original expected irreducible
	)
	return &AI32{
		o: oi, // equals to out or out increased
		i: ii, // equals to in or in decreased
		e: ext,
	}, over2
}

// roi returns a AI32 with a maximum o, minimum i and without e.
// Looks value inA can be expressed as a product of p*p*inB where inB is square-free.
// Then returns a new AI32 with values o=out*p and i=inA/p
// Case 1: For out = inA = 0 returns empty AI32 which means is zero.
// Case 2: When iB = +1, return AI32 with only o = out*p*inA
// Case 3: When p = 1, returns AI32 with o=out, i=inA
// Case 4: When p > 1, returns AI32 with o=out*p, i=inB/p
// Returns nil for overflows
func (r *AI32s) roi(out, inA Z) (ai *AI32, overflow bool) {
	if out == 0 || inA == 0 { // case 1
		return // zero
	}
	io := out; if out < 0 { io = -out }
	ia := inA; if inA < 0 { ia = -inA }
	if no, na, ok := r.reduce1(N(io), N(ia)); !ok {
		overflow = true
	} else if na == 1 && inA > 0 { // case 2
		ai = &AI32{
			o: &I32{ n:no, s:out < 0 },
			i: &I32{ n:na, s:inA < 0 },
		}
	} else { // cases 3 and 4
		ai = &AI32{
			o: &I32{ n:no, s:out < 0 },
			i: &I32{ n:na, s:inA < 0 },
		}
	}
	return
}

func (r *AI32s) roie(out, inA, inB Z) (*I32, *I32, Z, bool) {
	if out == 0 { // case 1
		return nil, nil, 0, true
	}
	io := out; if out < 0 { io = -out }
	ia := inA; if inA < 0 { ia = -inA }
	ib := inB; if inB < 0 { ib = -inB }
	no, na, nb, ok := r.reduce2(N(io), N(ia), N(ib))
	if !ok {
		return nil, nil, 0, false // overflow
	}
	zo := Z(no); if out < 0 { zo = -zo }
	za := Z(na); if inA < 0 { za = -za }
	zb := Z(nb); if inB < 0 { zb = -zb }
	ro, _ := newI32(zo)
	ra, _ := newI32(za)
	return ro, ra, zb, true
}



// reduce1N try to decrease in and increase out.
// Example: The input -3√(20) is returned as -6√(5)
// Return ok as false when out or in values are larger than 32 bits (overflow).
func (r *AI32s) reduce1(out, in N) (o N32, i N32, ok bool) {
	if out == 0 || in == 0 {
		return 0, 0, true
	} else if in > 1 {
		for _, prime := range r.primes {
			p := N(prime)
			if pp := p*p; in >= pp {
				for {
					if in % pp == 0 {
						// reduce ok: increase out, decrease in.
						out *= p
						in  /= pp
						// look for repeated squares in reduced in
						continue
					}
					break // check next prime
				}
			} else {
				// in has no more factors to check
				break
			}
		}
	}
	if out > N32_MAX || in > N32_MAX {
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
	if out > N32_MAX || inA > N32_MAX || inB > N32_MAX {
		return 0, 0, 0, false
	}
	return N32(out), N32(inA), N32(inB), true
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
	if r == nil || r.o == nil || r.o.n == 0 {
		s.Zero() // return +0
	} else {
		s.I32(r.o) // append ±out
		if r.i == nil || r.i.n == 0 {
			return // return only printed out
		}
		if r.e == nil {
			// append √±in and return
			s.Radical32(r.i, nil)
		} else {
			s.Radical32(r.i, func(s *Str) {
				// append √±in( ... ) and return
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


