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

// Example for b=6, c=5:
//
//	     b1   b2   b3   b4   b5   b6
//	   +----+----+----+----+----+----+
//	c1 | A0 | B0 | C0 | D0 | E0 | F0 |
//	   +----+----+----+----+----+----+
//	c2      | C1 | D1 | E1 | F1 | G0 |
//	        +----+----+----+----+----+
//	c3           | E2 | F2 | G1 | H0 |
//	             +----+----+----+----+
//	c4                | G2 | H1 | I0 |
//	                  +----+----+----+
//	c5                     | I1 | J1 |
//	                       +----+----+
// Where: A < B < C < D < E < F < G < H < I < J
//
// diagsBC return and array of diagonal factors size = b + c - 1
func (t *T) diagsBC() ([][]N, N) {
	num, den := t.cosA()
	size := int(t.b + t.c - 1)
	diags := make([][]N, size)
	for d := range diags {
		diags[d] = make([]N, 0)
	}
	for x := N(1); x <= N(t.b); x++ {
		for y := N(1); y <= x; y++ {
			r := int(x + y - 2)
			if r < size {
				d := (x*x + y*y)*N(den) - 2*x*y*N(num)
				diags[r] = append(diags[r], d)
			}
		}
	}
	return diags, N(den)
}
