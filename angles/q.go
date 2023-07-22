package angles

import (
	"fmt"
)

type Q struct {
	Neg bool // true=negative
	Num uint // numerator
	Den uint // denominator
}

// NewQ creates a new quotient when n and d are numerator and denominator.
// Greatest common divisor set N and D at minimum.
func NewQ(n, d int) (q *Q) {
	// return nil as NaN or +/- infinity
	if d == 0 {
		return
	}
	q = &Q{}
	// set negative sign and convert to uint
	num := uint(0)
	den := uint(0)
	if n < 0 {
		if d > 0 {
			q.Neg = true
		}
		num = uint(-n)
	} else {
		num = uint(n)
	}
	if d < 0 {
		if n > 0 {
			q.Neg = true
		}
		den = uint(-d)
	} else {
		den = uint(d)
	}
	// return zero
	if num == 0 {
		q.Den = den
		return
	}
	// greatest common divisor (GCD) via Euclidean algorithm
	a, b := num, den
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	// set reduced num/dem
	q.Num = num / a
	q.Den = den / a
	return
}

// Times returns a new Q with the addition of this quotient q with given quotient p.
func (q *Q) Add(p *Q) *Q {
	if q == nil || p == nil {
		return nil
	}
	n1 := int(q.Num * p.Den)
	if q.Neg {
		n1 *= -1
	}
	n2 := int(p.Num * q.Den)
	if p.Neg {
		n2 *= -1
	}
	return NewQ(n1 + n2, int(q.Den * p.Den))
}

// Times returns new Q with the multiplication of this quotient q with given quotient p.
func (q *Q) Times(p *Q) *Q {
	if q == nil || p == nil {
		return nil
	}
	if q.Num == 0 || p.Num == 0 {
		return &Q{ Den:1 }
	}
	n1 := int(q.Num)
	if q.Neg {
		n1 *= -1
	}
	n2 := int(p.Num)
	if p.Neg {
		n2 *= -1
	}
	return NewQ(n1 * n2, int(q.Den * p.Den))
}

// Negate returns new Q with the sign changed.
func (q *Q) Negate() *Q {
	if q == nil {
		return nil
	}
	return &Q{
		Neg: !q.Neg,
		Num: q.Num,
		Den: q.Den,
	}
}

// Inverse returns new Q with the numerator and denominator reversed
func (q *Q) Inverse() *Q {
	if q == nil {
		return nil
	}
	if q.Num == 0 {
		return nil
	}
	return &Q{
		Neg: q.Neg,
		Num: q.Den,
		Den: q.Num,
	}
}

func (q *Q) String() string {
	if q == nil {
		return ""
	}
	if q.Num == 0 {
		return "0"
	}
	neg := ""
	if q.Neg {
		neg = "-"
	}
	if q.Den == 1 {
		return fmt.Sprintf("%s%d", neg, q.Num)	
	}
	return fmt.Sprintf("%s%d/%d", neg, q.Num, q.Den)
}
