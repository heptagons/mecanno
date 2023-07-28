package alg

import (
	"math"
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


/*func gcd(a, b Z) Z {
	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}*/

// I is an integer of 32 bits
type I32 struct {
	s bool
	n N32
}

func NewI32(z Z) (*I32, bool) {
	if N32overflowZ(z) {
		return nil, false
	} else if z > 0 {
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
		s.Integer32(i)
	}
}


type R32s struct {
	primes []N32
}

func NewR32s() *R32s {
	return &R32s{
		primes: NatPrimes(),
	}
}

// sqrt reduces the radical (out)√(in) in in two parts
// Example: -3√(20) returns -6√(5)
// Return ok as false when returned values are larger than 32 bits (overflow).
func (r *R32s) sqrt(out, in N) (o N32, i N32, ok bool) {
	if out == 0 {
		return 0, 0, true
	}
	if in == 0 {
		return 0, 0, true
	}
	if in > 1 {
		// Try to modify out and in
		for _, prime := range r.primes {
			p := N(prime)
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
	if N32overflowN(out) {
		return 0, 0, false
	}
	if N32overflowN(in) {
		return 0, 0, false
	}
	return N32(out), N32(in), true
}

func (rs *R32s) NewR32(out, in Z) *R32 {
	if out == 0 || in == 0 {
		return NewR32zero()
	}
	out64 := out
	if out < 0 { // radical negative
		out64 = -out
	}
	in64 := in
	if in < 0 { // radical imaginary
		in64 = -in
	}
	out32, in32, ok := rs.sqrt(N(out64), N(in64))
	if !ok {
		return nil // reject overflows
	}
	zo := Z(out32); if out < 0 { zo = -zo }
	zi := Z(in32);  if in  < 0 { zi = -zi }


	if out, ok := NewI32(zo); !ok {
		return nil
	} else if in, ok := NewI32(zi); !ok {
		return nil
	} else {
		return &R32{
			out: out,
			in:  in,
		}
	}
}

type R32 struct {
	out *I32 // external integer with sign=true means whole R32 is negative
	in  *I32 // internal integer with sign=true means whole R32 is imaginary
}

func (r *R32) outVal() Z {
	return r.out.val()
}

func (r *R32) outValPow2() Z {
	return r.out.valPow2()
}

func (r *R32) inVal() Z {
	return r.in.val()
}

func NewR32zero() *R32 {
	return &R32{
		out: nil,
		in:  nil,
	}
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
		s.Integer32(r.out)
		s.Radical32(r.in)
	}
}

func (r *R32) String() string {
	s := NewStr()
	r.Str(s)
	return s.String()
}


