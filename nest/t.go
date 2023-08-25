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
