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
		m: Z(b*b) + Z(c*c) - Z(a*a),
		n: Z(2*b*c),
		s: Z(a*a) + Z(b*b) - Z(c*c),
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

func (s *Fours) G1(f *Four, d, e N32) (*A32, error) {
	// see "g" formula in github.com/heptagons/meccano/frames/four/four.pdf
	//i1 := f.n * f.t
	a, b, c := Z(f.a), Z(f.b), Z(f.c)
	x, y := f.xy(d)
	ee := Z(e*e)
	//i := Z(e)*f.s*(f.n*x - f.m*y) - f.m*x*y*f.t - Z(e)*y*(f.t-f.s)*(f.s+f.t)
	//i2 := i1*i1*(ee + x*x + y*y) + 2*i1*i
	i1 := 2*a*b*c
	i2 := 2*a*c*( 2*a*b*b*c*(ee + x*x + y*y) - 2*a*b*f.m*x*y  + Z(e)*f.s*(f.n*x-f.m*y) - Z(e)*y*(f.t*f.t - f.s*f.s))
	return s.ANew3(
		N(i1), // A
		0, // B
		1, // C
		i2, // D
	)
}

func (s *Fours) G2(f *Four, d, e N32) (*A32, error) {
	// see "g" formula in github.com/heptagons/meccano/frames/four/four.pdf
	i1 := f.n * f.t
	x, y := f.xy(d)
	ee := Z(e*e)
	i2 := i1*i1*(ee + x*x + y*y) -2*i1*f.m*x*y*f.t + 2*i1*Z(e)*f.s*(f.n*x - f.m*y)
	i3 := -2*i1*Z(e)*y
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

