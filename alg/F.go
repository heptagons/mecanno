package alg

import (
	"fmt"
)

// N is a natural number of 32 bits
type N uint32

// gcd returns the greatest common divisor of two numbers
func (a N) gcd(b N) N {
	if b == 0 {
		return a
	}
	return b.gcd(a % b)
}

type Z int64

func gcd(a, b Z) Z {
	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}


// I is an integer of 32 bits
type I struct {
	s bool
	n N
}

func NewI(s bool, n N) *I {
	return &I{
		s: s,
		n: n,
	}
}

func (x *I) sgn(y *I) bool {
	if x.s {
		return !y.s
	} else {
		return y.s
	}
}

func (x *I) mul(n N) Z {
	if x.s {
		return Z(n) * -Z(x.n)
	} else {
		return Z(n) * Z(x.n)
	}
}


func (x *I) stringI() string {
	if x == nil {
		return ""
	} else if x.n == 0 {
		return "0"
	} else if x.s {
		return fmt.Sprintf("-%d", x.n)
	} else {
		return fmt.Sprintf("%d", x.n)
	}
}

type B struct { // Rational
	a *I
	b N
}

func NewB(s bool, a, b N) *B {
	if b == 0 {
		return nil // infinity
	}
	if a == 0 {
		return &B{
			a: NewI(s, 0),
			b: b,
		}
	}
	num, den := a, b
	g := num.gcd(den)
	return &B{
		a: NewI(s, num / g),
		b: den / g,
	}
}

func (x *B) AddB(y *B) *B {
	if x == nil || y == nil {
		return nil
	} else if x.a == nil || y.a == nil {
		return nil
	} else if x.b == 0 || y.b == 0 {
		return nil
	}
	num := x.a.mul(y.b) + y.a.mul(x.b)	
	den := Z(x.b) * Z(y.b)
	g := gcd(num, den)
	num /= g
	den /= g
	b := N(den) // TODO check overflow
	var a *I
	if num < 0 {
		// TODO check overflow
		a = NewI(true, N(-num))
	} else {
		// TODO check overflow
		a = NewI(false, N(num))
	}
	return &B{
		a: a,
		b: b,
	}
}

func (x *B) MulB(y *B) *B {
	if x == nil || y == nil {
		return nil
	}
	if x.a == nil || y.a == nil {
		return nil
	}
	if x.b == 0 || y.b == 0 {
		return nil
	}
	num, den := x.a.n * y.a.n, x.b * y.b
	g := num.gcd(den)
	return NewB(x.a.sgn(y.a), num / g, den / g)
}

func (x *B) StringB() string {
	if x == nil {
		return ""
	} else if x.a == nil {
		return "" // undefined
	} else if x.b == 0 {
		return "" // infinity
	} else if x.a.n == 0 {
		return "0"
	} else if x.b == 1 {
		return x.a.stringI()
	} else {
		return fmt.Sprintf("%s/%d", x.a.stringI(), x.b)
	}
}









type C struct { // Algebraic C
	*B
	c N
}

type D struct { // Algebraic D
	*C
	d    I
}

type F struct { // Algebraic F
	*D
	e I
	f N
}



