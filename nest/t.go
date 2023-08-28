package nest

import (
	"fmt"
)

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

func (t *T) String() string {
	return fmt.Sprintf("%d,%d,%d", t.a, t.b, t.c)
}


type Tang byte

const (
	TangA Tang = 'A'
	TangB Tang = 'B'
	TangC Tang = 'C'
)

type T32s struct {
	*A32s
}

func NewT32s() *T32s {
	return &T32s{
		A32s:  NewA32s(),
	}
}

func (ts *T32s) cosZ(x, y, z N32) (num, den Z) {
	d32, n32, _ := ts.zFrac(
		2*N(x)*N(y),
		Z(x)*Z(x) + Z(y)*Z(y) - Z(z)*Z(z),
	)
	return Z(n32), Z(d32)
}

func (ts *T32s) cos(t *T, a Tang) (num, den Z) {
	switch a {
	case TangA:
		return ts.cosZ(t.b, t.c, t.a)
	case TangB:
		return ts.cosZ(t.c, t.a, t.b)
	case TangC:
		return ts.cosZ(t.a, t.b, t.c)
	default:
		panic("Invalid Tang")
	}
}

func (ts *T32s) sin(t *T, a Tang) (surd, den Z32) {
	switch a {
	case TangA:
		return t.d, 2*Z32(t.b * t.c)
	case TangB:
		return t.d, 2*Z32(t.c * t.a)
	case TangC:
		return t.d, 2*Z32(t.a * t.b)
	default:
		panic("Invalid Tang")
	}
}


// diags return the diagonals:
//	- For TangC return ths diagonals between sides a and b.
//	- For TangA return the diagonals between sides b and c.
// in case a = c return nothing, so would be repetitions
// already obtained by calling abDiags.
//	- For TancB return the diagonals between sides a and c.
// in case b = c return nothing, so would be repetitions
// already obtained by calling abDiags.
func (ts *T32s) tDiagsAng(t *T, a Tang) ([][]N, N) {
	switch a {
	case TangC:
		num, den := ts.cos(t, TangC)
		return ts.tDiags(num, den, t.a, t.b)
	
	case TangA:
		if t.a == t.c {
			return nil, 0
		}
		num, den := ts.cos(t, TangA)
		return ts.tDiags(num, den, t.b, t.c)
	
	case TangB:
		if t.a == t.b || t.b == t.c {
			return nil, 0
		}
		num, den := ts.cos(t, TangB)
		return ts.tDiags(num, den, t.a, t.c)
	
	default:
		panic("Invalid Tang")
	}
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
			pos := int(x - y)
			d := (x*x + y*y)*N(den) - 2*x*y*N(num)
			diags[pos] = append(diags[pos], d)
		}
	}
	denN := N(den)
	ts.nFracN(&denN, diags)
	return diags, denN
}


func (ts *T32s) tCosAplusB(t1 *T, a1 Tang, t2 *T, a2 Tang) (*A32, error) {
	n1,d1 := ts.cos(t1, a1)
	n2,d2 := ts.cos(t2, a2)
	//fmt.Printf("%v cos %d/%d\n", t1.String(), n1, d1)
	//fmt.Printf("%v cos %d/%d\n", t2.String(), n2, d2)
	an := d1*d2
	bz := n1*n2
	cz := Z(-1)
	dz := (d1*d1 - n1*n1) * (d2*d2 - n2*n2)
	if o, in, err := ts.zSqrt(cz, dz); err != nil {
		return nil, err
	} else if den32, n32s, err := ts.zFracN(N(an), bz, Z(o)); err != nil {
		return nil, err
	} else {
		return ts.aNew3(N(den32), Z(n32s[0]), Z(n32s[1]), Z(in))
	}
}

func (ts *T32s) tSide(y, z N32, cosX *A32) (*A32, error) {
	if y2_z2, err := ts.aNew1(1, Z(y)*Z(y) + Z(z)*Z(z)); err != nil {
		return nil, err
	} else if _2zy, err := ts.aNew1(1, -2*Z(y)*Z(z)); err != nil {
 		return nil, err
	} else if _2zycosX, err := ts.aMul(cosX, _2zy); err != nil {
		return nil, err
	} else if y2_z2_2zycosX, err := ts.aAdd(y2_z2, _2zycosX); err != nil {
		return nil, err
	} else {
		return ts.aSqrt(y2_z2_2zycosX)
	}
}


