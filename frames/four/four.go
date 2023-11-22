package four

import (
	. "github.com/heptagons/meccano/nest"
)

type Four struct{
	a, b, c    N32
	m, n, t, s Z
}

func NewFour(a, b, c N) *Four {
	// TODO return error for invalid triangles?
	// see "m" formula in github.com/heptagons/meccano/frames/four/four.pdf
	return &Four{
		a: N32(a),
		b: N32(b),
		c: N32(c),
		m: Z(b*b + c*c - a*a),
		n: Z(2*b*c),
		s: Z(a*a + b*b - c*c),
		t: Z(2*a*b),
	}
}

func (s *Four) xy(d N32) (Z, Z) {
	return Z(s.a + s.c), Z(s.b + d)
}

type Fours struct {
	*A32s
}

func NewFours() *Fours {
	return &Fours{
		A32s: NewA32s(),
	}
}

func (s *Fours) F(f *Four, d N32) (*A32, error) {
	// see "f" formula in github.com/heptagons/meccano/frames/four/four.pdf
	x, y := f.xy(d)
	return s.ANew3(
		N(f.n), // A
		0, // B
		1, // C
		f.n*f.n*(x*x + y*y) - 2*f.m*f.n*x*y, // D
	)
}

func (s *Fours) CosTheta(f *Four, d N32) (*A32, error) {
	x, y := f.xy(d)
	z := f.n*f.n*(x*x + y*y) - 2*f.m*f.n*x*y
	// cos = (nx - my)sqrt(z)/z
	return s.ANew3(
		N(z), // A
		0, // B
		f.n*x - f.m*y, // C
		z, //D
	)
}

func (s *Fours) G(f *Four, d, e N32) (*A32, error) {
	// see "g" formula in github.com/heptagons/meccano/frames/four/four.pdf
	i1 := f.n * f.t
	x, y := f.xy(d)
	ee := Z(e*e)
	_2ei1 := 2*Z(e)*i1
	i2 := (ee + x*x + y*y)*i1*i1 -2*f.m*x*y*f.t*i1 + _2ei1*(f.n*x - f.m*y)
	i3 := _2ei1*y
	i4 := (f.n - f.m)*(f.n + f.m)*(f.t - f.s)*(f.t + f.s)
	A := N(i1)
	B := Z(0)
	C := Z(0)
	D := Z(0)
	E := Z(1)
	F := i2
	G := i3
	H := i4
	return s.ANew7(A, B, C, D, E, F, G, H)
}

