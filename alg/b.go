package alg

type B struct { // Rational
	b *I32 // numerator integer optional
	a N32  // denominator natural >= 1
}

func NewB0() *B {
	return &B{
		a:1,
	}
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
	////// TODO replace with (&num.Reduce2(&den))
	g := gcd(num, den)
	num /= g
	den /= g
	////// end replace
	if num > MaxN {
		return nil // numerator overflow
	} else if den > MaxN {
		return nil // denominator overflow
	}
	var b *I32
	if s { // fast negative
		b = newI32minus(N32(num))
	} else { // fast positive
		b = newI32plus(N32(num))
	}
	return &B{
		b: b,
		a: N32(den),
	}
}

// NewBcosC returns the rational cosine of the angle opposed to segment c
// using the law of cosines:
//	       a² + b² - c²
//	cosC = ------------
//	           2ab
func NewBcosC(a, b, c N32) *B {
	num := Z(a*a) + Z(b*b) - Z(c*c)
	den := 2*Z(a*b)
	return NewB(num, den)
}

func NewBplus(num, den N32) *B {
	return newB(false, num, den)
}

func NewBminus(num, den N32) *B {
	return newB(true, num, den)	
}

func newB(s bool, num, den N32) *B {
	if den == 0 {
		return nil // infinity
	}
	if num == 0 {
		return NewB0() // zero
	}
	n, d := num, den
	g := n.gcd(d)
	return &B{
		b: &I32{ s:s, n: n / g },
		a: d / g,
	}
}

func (x *B) IsZero() bool {
	if x == nil || x.b == nil || x.b.n == 0 {
		return true
	}
	return false
}

func (x *B) clone() *B {
	if x == nil || x.a == 0 {
		return nil // infinite
	} else if x.b == nil || x.b.n == 0 {
		return NewB0()
	} else {
		return newB(x.b.s, x.b.n, x.a)
	}
}

func (x *B) AddB(y *B) *B {
	if x == nil || y == nil || x.a == 0 || y.a == 0 {
		return nil // infinite
	} else if x.b == nil || x.b.n == 0 {
		return y.clone() // y
	} else if y.b == nil || y.b.n == 0 {
		return x.clone() // x
	}
	num := x.b.mul(y.a) + y.b.mul(x.a)
	den := Z(x.a) * Z(y.a)
	return NewB(num, den)
}

func (x *B) Inv() *B {
	if x == nil || x.a == 0 || x.b == nil || x.b.n == 0 {
		return nil
	} 
	return &B{
		b: &I32{ s:x.b.s, n: x.a },
		a: x.b.n,
	}
}

func (x *B) MulB(y *B) *B {
	if x == nil || y == nil {
		return nil
	} else if x.b == nil || x.b.n == 0 {
		return NewB0()
	} else if y.b == nil || y.b.n == 0 {
		return NewB0()
	}
	num := Z(x.b.n) * Z(y.b.n)
	if x.b.s != y.b.s {
		num = -num
	}
	den := Z(x.a) * Z(y.a)
	return NewB(num, den)
}

func (x *B) Str(s *Str) {
	if x == nil || x.a == 0 {
		s.Infinite()
	} else if x.b == nil || x.b.n == 0 {
		s.Zero()
	} else if x.a == 1 {
		x.b.Str(s)
	} else {
		x.b.Str(s)
		s.Divisor()
		s.N32(x.a)
	}
}

func (x *B) String() string {
	s := NewStr()
	x.Str(s)
	return s.String()
}

func (x *B) Reduce3(third *I32) {
	if third == nil {
		return
	}
	if x.b == nil {
		// this B is zero just reduce denominator and given not nil third
		(&(x.a)).Reduce2(&(third.n))
	} else {
		(&(x.a)).Reduce3(&(x.b.n), &(third.n))
	}
}
