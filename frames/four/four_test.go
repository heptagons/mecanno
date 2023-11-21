package four

import (
	"testing"
	. "github.com/heptagons/meccano/nest"
)

type Four struct{
	a, b, c, m N32
}

func NewFour(a, b, c N) *Four {
	// TODO return error for invalid triangles?
	// see "m" formula in github.com/heptagons/meccano/frames/four/four.pdf
	return &Four{
		a: N32(a),
		b: N32(b),
		c: N32(c),
		m: N32(c)*N32(a + b + c)*N32(a - b + c),
	}
}

type Fours struct {
	*A32s
}

func NewFours() *Fours {
	return &Fours{
		A32s: NewA32s(),
	}
}

func (s *Fours) F0(f *Four) (*A32, error) {
	// see "f0" formula in github.com/heptagons/meccano/frames/four/four.pdf
	//	(b + c√d) / a
	A := N(f.c)
	B := Z(0)
	C := Z(1)
	D := Z(f.a * f.m)
	return s.ANew3(A, B, C, D)
}

func (s *Fours) F1(f *Four, d N32) (*A32, error) {
	// see "f1" formula in github.com/heptagons/meccano/frames/four/four.pdf
	//	(b + c√d) / a
	a := Z(f.a)
	b := Z(f.b)
	c := Z(f.c)
	m := Z(f.m)
	A := N(f.b)*N(f.c)
	B := Z(0)
	C := Z(1)
	D := b*b*c*c*Z(d)*Z(d) + b*m*(a*b + Z(d)*(a-c))
	return s.ANew3(A, B, C, D)
}

func TestFour(t *testing.T) {
	factory := NewFours()
	four := NewFour(4, 2, 3)
	t.Logf("abc=(%d,%d,%d) m=%d", four.a, four.b, four.c, four.m)
	if f0, err := factory.F0(four); err == nil {
		t.Logf("f0=%v", f0)
	}
	if f1, err := factory.F1(four, 2); err == nil {
		t.Logf("d=%d f1=%v", 2, f1)
	}
}