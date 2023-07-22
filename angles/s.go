package angles

import (
	"fmt"
)

const MAX32 = uint64(0xffffffff)

type Squares struct {
	primes []uint
}


func NewSquares32() *Squares {
	return &Squares{
		primes: PrimesList(0xffff),
	}
}

// Returns the square root of the product of two integers

func (s *Squares) productRoot(a, b uint) (uint, uint, bool) {
	in := uint64(a) * uint64(b)
	out   := uint64(1)
	if in == 0 {
		return 0, 0, true
	}
	if in == 1 {
		return 1, 1, true
	}
	for _, prime := range s.primes {
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
	if in > MAX32 {
		return 0, 0, false
	} else if out > MAX32 {
		return 0, 0, false
	} else {
		return uint(in), uint(out), true
	}
}

// Root returns the square root of the given q in format: s.Neg * (s.Num/s.Den)sqrt(s.R)
func (s *Squares) Root(q *Q) *S {
	if q == nil {
		return nil // NaN
	}
	if q.Neg {
		return nil // Imaginary
	}
	//
	//	   num * den     sqrt(num*den)                   n1
	// Q = ----------- -> ------------ -> sqrt(radical)*----
	//     den * den         den                        den
	if in, out, ok := s.productRoot(q.Num, q.Den); !ok {
		return nil // overflow
	} else {
		return NewS(in, int(out), int(q.Den))
	}
}

func (s *Squares) Times(a, b *S) *S {
	if a == nil || b == nil {
		return nil
	} else if in, out, ok := s.productRoot(a.In, b.In); !ok {
		return nil
	} else {
		return &S{
			In: in,
			Q:  a.Times(NewQ(int(out), 1)).
					Times(b.Q),
		}
	}
}

type S struct {
	In uint
	*Q
}

func NewS(in uint, num, den int) *S {
	if q := NewQ(num, den); q == nil {
		return nil
	} else {
		return &S{
			In: in,
			Q:  q,
		}
	}
}

func (s *S) String() string {
	if s == nil || s.Q == nil {
		return ""
	} else if s.In == 0 {
		return "0"
	} else if s.In == 1 {
		return s.Q.String()
	} else {
		return fmt.Sprintf("sqrt(%d)(%s)", s.In, s.Q)
	}
}




