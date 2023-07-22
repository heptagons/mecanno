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
