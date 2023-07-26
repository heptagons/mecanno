package alg

import (
	"fmt"
)

// Alg is the simplest algebraic number of the form
// rat.Num
// ------- * sqrt(In)
// rat.Den
type Alg struct {
	*Rat
	In N32
}

func NewAlg(rat *Rat, in N32) *Alg {
	if rat == nil {
		return nil
	}
	return &Alg{
		Rat: rat,
		In: in,
	}
}

func (s *Alg) String() string {
	if s == nil || s.Rat == nil {
		return ""
	} else if s.In == 0 {
		return "0"
	} else if s.In == 1 {
		return s.Rat.String()
	} else {
		return fmt.Sprintf("(%s)√(%d)", s.Rat, s.In)
	}
}

type Algs struct {
	*N32s
}

func NewAlgs(nats *N32s) *Algs {
	return &Algs{
		N32s: nats,
	}
}

// CosC returns the rational cosine of the angle C using the law of cosines:
//	       a² + b² - c²
//	cosC = ------------
//	           2ab
func (algs *Algs) CosC(a, b, c N32) *Rat {
	num := int(a*a + b*b - c*c)
	den := int(2*a*b)
	return NewRat(num, den)
}

// SinC return the algebraic sine of the angle C using the law of sines:
//	       math.Sqrt(4a²b² - (a²+b²-c²)²)
//	sinC = ------------------------------
//	                  2ab 
func (algs *Algs) SinC(a, b, c N32) *Alg {
	p := int(4*a*a*b*b)
	q := int((a*a + b*b - c*c))
	d := int(2*a*b)
	if rat := NewRat(p - q*q, d*d); rat == nil {
		return nil
	} else {
		return algs.Sqrt(rat)
	}
}








