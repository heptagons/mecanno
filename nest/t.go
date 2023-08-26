package nest

// T is a triangle
type T struct {
	a N32
	b N32
	c N32
	d Z32 // area = √d/4
}

func newT(a, b, c N32) *T {
	if a >= b && b >= c && b+c > a {
		return &T{
			a: a,
			b: b,
			c: c,
			// https://en.wikipedia.org/wiki/Heron%27s_formula Numerical stability
			d: Z32(a+(b+c)) * Z32(c-(a-b)) * Z32(c+(a-b)) * Z32(a+(b-c)),
		}
	}
	return nil
}

func (t *T) cosZ(x, y, z N32) (num, den Z) {
	num = Z(x)*Z(x) + Z(y)*Z(y) - Z(z)*Z(z)
	den = 2*Z(x)*Z(y)
	return
}

func (t *T) cosA() (num, den Z) {
	return t.cosZ(t.b, t.c, t.a)
}

func (t *T) cosB() (num, den Z) {
	return t.cosZ(t.c, t.a, t.b)
}

func (t *T) cosC() (num, den Z) {
	return t.cosZ(t.a, t.b, t.c)
}

func (t *T) sinA() (surd, den Z32) {
	return t.d, 2*Z32(t.b * t.c)
}

func (t *T) sinB() (surd, den Z32) {
	return t.d, 2*Z32(t.c * t.a)
}

func (t *T) sinC() (surd, den Z32) {
	return t.d, 2*Z32(t.a * t.b)
}

type T32s struct {
	*Z32s
}

func NewT32s() *T32s {
	return &T32s{
		Z32s:  NewZ32s(),
	}
}

func (ts *T32s) aDiags(t *T) ([][]N, N) {
	num, den := t.cosA()
	return ts.tDiags(num, den, t.b, t.c)
}

func (ts *T32s) bDiags(t *T) ([][]N, N) {
	num, den := t.cosB()
	return ts.tDiags(num, den, t.a, t.c)
}

func (ts *T32s) cDiags(t *T) ([][]N, N) {
	num, den := t.cosC()
	return ts.tDiags(num, den, t.a, t.b)
}

// Example for b=6, c=5:
//
//	a0   a1   a2   a3   a4   a5   a6   a7
//   0    1    2    3    4    5    6    7
//	   +----+----+----+----+----+----+----+
//	   | A0 | B0 | C0 | D0 | E0 | F0 | G0 |  b1
//	   +----+----+----+----+----+----+----+
//	        | A1 | B1 | C1 | D1 | E1 | F1 |  b2
//	        +----+----+----+----+----+----+
//	             | A2 | B2 | C2 | D2 | E2 |  b3
//	             +----+----+----+----+----+
//	                  | A3 | B3 | C3 | D3 |  b4
//	                  +----+----+----+----+
//	                       | A4 | B4 | C4 |  b5
//	                       +----+----+----+
//                              | A5 |  6 |  b6
//                              +----+----+
//
// diagsBC return and array of diagonal factors size = b + c - 1
func (ts *T32s) tDiags(num, den Z, s1, s2 N32) ([][]N, N) {
	diags := make([][]N, s1)
	for d := range diags {
		diags[d] = make([]N, 0)
	}
	for x := N(1); x <= N(s1); x++ {
		for y := N(1); y <= x; y++ {
			pos := x - y
			d := (x*x + y*y)*N(den) - 2*x*y*N(num)
			diags[pos] = append(diags[pos], d)
		}
	}
	return diags, N(den)
}
