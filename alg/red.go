package alg

import (
	"math"
)

// NewRed32 creates algebraic integers factory
// A fixed internal 32-bit primes list is used to create
// algebraic integers reduced
type Red32 struct {
	primes []N32
}

func NewRed32() *Red32 {
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
	return &Red32{
		primes: primes,
	}
}

func (r *Red32) AI(out, in Z, ext *AI32) (ai *AI32, overflow bool) {
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
	if oi, ii, eo2, overflow := r.roie(out, in, eo1); overflow {
		return nil, true // c, dd, g, hh, m, nn overflow
	} else {
		if ext.e == nil {
			if ext.inVal() == +1 && ii != 0 {
				// e√(f+g√h) or k√(l+m√n) ... where h == +1, n == +1, ...
				// return reduced e√(f+g) or k√(l+m) ...
				return r.roi(oi, ii + eo2)
			}
		}
		// i√(j+k√l+(...)) or ...
		// extension
		// eo2 equals to eoi or reduced (decreased)
		// ext.inVal() extra-in original expected irreducible
		// ext.e extra-extra original expected irreducible
		if ext, overflow := r.AI(eo2, ext.inVal(), ext.e); overflow {
			return nil, true
		} else {
			o, _ := newI32(oi)
			i, _ := newI32(ii)
			return &AI32{
				o: o, // equals to out or out increased
				i: i, // equals to in or in decreased
				e: ext,
			}, false
		}
	}
}

// roi returns a AI32 with a maximum o, minimum i and without e.
// Looks value inA can be expressed as a product of p*p*inB where inB is square-free.
// Then returns a new AI32 with values o=out*p and i=inA/p
// Case 1: For out = inA = 0 returns empty AI32 which means is zero.
// Case 2: When iB = +1, return AI32 with only o = out*p*inA
// Case 3: When p = 1, returns AI32 with o=out, i=inA
// Case 4: When p > 1, returns AI32 with o=out*p, i=inB/p
// Returns nil for overflows
func (r *Red32) roi(out, inA Z) (ai *AI32, overflow bool) {
	if out == 0 || inA == 0 { // case 1a
		return nil, false // zero
	}
	io := out; if out < 0 { io = -out }
	ia := inA; if inA < 0 { ia = -inA }
	if no, na, overflow := r.roiN(N(io), N(ia)); overflow {
		return nil, true
	} else { // cases 1b,2a,2b
		return &AI32{
			o: &I32{ n:no, s:out < 0 },
			i: &I32{ n:na, s:inA < 0 },
		}, false
	}
}

// roiN finds a p such that in = p*p*in2 where in2 is a square-free number.
// For p > 1, returns new values out*p and in/p
// Example: The input -3√(20) is returned as -6√(5)
// Return overflow when out or in values are larger than 32 bits.
func (r *Red32) roiN(out, in N) (o N32, i N32, overflow bool) {
	if out == 0 || in == 0 {
		return 0, 0, false // zero
	}
	for _, prime := range r.primes {
		p := N(prime)
		if pp := p*p; in >= pp {
			for {
				if in % pp == 0 { // reduce ok
					out *= p
					in  /= pp
					continue // look for repeated squares in reduced in
				}
				break // check next prime
			}
		} else {
			break // in has no more factors to check
		}
	}
	if out > N32_MAX || in > N32_MAX {
		return 0, 0, true // overflow
	}
	return N32(out), N32(in), false
}

func (r *Red32) roie(out, inA, inB Z) (o, i, j Z, overflow bool) {
	if out == 0 { // case 1
		return 0, 0, 0, false // zero
	}
	io := out; if out < 0 { io = -out }
	ia := inA; if inA < 0 { ia = -inA }
	ib := inB; if inB < 0 { ib = -inB }
	if no, na, nb, overflow := r.roieN(N(io), N(ia), N(ib)); overflow {
		return 0, 0, 0, true // overflow
	} else {
		zo := Z(no); if out < 0 { zo = -zo }
		za := Z(na); if inA < 0 { za = -za }
		zb := Z(nb); if inB < 0 { zb = -zb }
		return zo, za, zb, false
	}
}

// reduce2 try to increase given out by decrease given inA and inB.
// Return reduced out and ins and a flag for 32 bits overflow.
// Example:
//	For given out=5, inA=12=2*2*3, inB=56=2*2*2*7
//  Returns out=10=5*2, inA=3, inB=14=2*7, true
func (r *Red32) roieN(out, inA, inB N) (o, i, j N32, overflow bool) {
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
		return 0, 0, 0, true // overflow
	}
	return N32(out), N32(inA), N32(inB), false 
}





