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

const MaxN = Z(0xffffffff)

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

func newIplus(n N) *I {
	return &I{ s:false, n:n }
}

func newIminus(n N) *I {
	return &I{ s:true, n:n }	
}

func (x *I) mul(n N) Z {
	if x.s {
		return -Z(n) * Z(x.n)
	} else {
		return +Z(n) * Z(x.n)
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

func NewB(num Z, den Z) *B {
	if den == 0 {
		return nil // infinite
	}
	if num == 0 {
		return NewB0() // zero
	}
	s := false
	if num < 0 {
		num = -num
		s = !s
	}
	if den < 0 {
		den = -den
		s = !s
	}
	g := gcd(num, den)
	num /= g
	den /= g
	if num > MaxN {
		return nil // numerator overflow
	} else if den > MaxN {
		return nil // denominator overflow
	}
	var a *I
	if s { // fast negative
		a = newIminus(N(num))
	} else { // fast positive
		a = newIplus(N(num))
	}
	return &B{
		a: a,
		b: N(den),
	}
}

func NewB0() *B {
	return &B{ b:1 }
}

func NewBplus(a, b N) *B {
	return newB(false, a, b)
}

func NewBminus(a, b N) *B {
	return newB(true, a, b)	
}

func newB(s bool, a, b N) *B {
	if b == 0 {
		return nil // infinity
	}
	if a == 0 {
		return NewB0() // zero
	}
	num, den := a, b
	g := num.gcd(den)
	return &B{
		a: &I{ s:s, n: num / g },
		b: den / g,
	}
}

func (x *B) clone() *B {
	if x == nil {
		return nil // infinite
	} else if x.a == nil || x.a.n == 0 {
		return NewB0()
	} else {
		return newB(x.a.s, x.a.n, x.b)
	}
}

func (x *B) AddB(y *B) *B {
	if x == nil || y == nil {
		return nil // infinite
	} else if x.a == nil || x.a.n == 0 {
		return y.clone() // y
	} else if y.a == nil || y.a.n == 0 {
		return x.clone() // x
	}
	num := x.a.mul(y.b) + y.a.mul(x.b)
	den := Z(x.b) * Z(y.b)
	return NewB(num, den)
}

func (x *B) MulB(y *B) *B {
	if x == nil || y == nil {
		return nil
	} else if x.a == nil || x.a.n == 0 {
		return NewB0()
	} else if y.a == nil || y.a.n == 0 {
		return NewB0()
	}
	num := Z(x.a.n) * Z(y.a.n)
	den := Z(x.b) * Z(y.b)
	return NewB(num, den)
}

func (x *B) StringB() string {
	if x == nil || x.b == 0 {
		return "" // infinity
	} else if x.a == nil || x.a.n == 0 {
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



