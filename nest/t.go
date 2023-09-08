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

func (t *T) String() string {
	return fmt.Sprintf("%d,%d,%d", t.a, t.b, t.c)
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


// newTalphas return all triangles with sides:
//	 _____
//	√alpha >= b >= c
func newTalphas(alpha N32) (tris []*T) {
	tris = make([]*T, 0)
	b := N32(1)
	for {
		if b*b > alpha {
			break
		}
		for c := N32(1); c <= b; c++ {
			if alpha < (b+c)*(b+c) {
				tris = append(tris, &T{ a:alpha, b:b, c:c })
			}
		}
		b++
	}
	return
}

// newTbetas return all triangles with sides:
//	      ____
//	a >= √beta >= c
func newTbetas(beta N32, max N32) (tris []*T) {
	tris = make([]*T, 0)
	maxC, minA := N32(1),N32(1)
	for { // naive way to get i*i > beta
		minA = maxC+1
		if minA*minA > beta {
			break
		}
		maxC = minA
	}
	for a := minA; a <= max; a++ {
		for c := N32(1); c <= maxC; c++ {
			if (a - c)*(a - c) < beta {
				tris = append(tris, &T{ a:a, b:beta, c:c })
			}
		}
	}
	return
}

// newTgammas return some triangles with sides:
//	          _____
//	a >= b > √gamma
func newTgammas(gamma N32, max N32) (tris []*T) {
	tris = make([]*T, 0)
	min := N32(1)
	for { // naive way to get i*i > gamma
		if min*min > gamma {
			break
		}
		min++
	}
	for a := min; a <= max; a++ {
		for b := min; b <= a; b++ {
			if (a - b)*(a - b) < gamma {
				tris = append(tris, &T{ a:a, b:b, c:gamma })
			}
		}
	}
	return
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

func (ts *T32s) tRatCosinesAll(t *T) (a, b, c *TRat) {
	var n, d Z
	n, d = ts.cos(t, TangA); a = &TRat{ num:n, den:d }
	n, d = ts.cos(t, TangB); b = &TRat{ num:n, den:d }
	n, d = ts.cos(t, TangC); c = &TRat{ num:n, den:d }
	return
}

func (ts *T32s) tRatSinesAll(t *T) (a, b, c *A32) {
	var s, d Z32
	s, d = ts.sin(t, TangA); a,_ = ts.aNew3(N(d), 0, 1, Z(s))
	s, d = ts.sin(t, TangB); b,_ = ts.aNew3(N(d), 0, 1, Z(s))
	s, d = ts.sin(t, TangC); c,_ = ts.aNew3(N(d), 0, 1, Z(s))
	return
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
// a0   a1   a2   a3   a4   a5   a6   a7
//  0    1    2    3    4    5    6    7
//	   +---,----,----,----,----,----,----,
//	   |A0 \ B0 \ C0 \ D0 \ E0 \ F0 \ G0 \  b1
//	   '--, '--, '--, '--, '--, '--, '--, '
//	       \ A1 \ B1 \ C1 \ D1 \ E1 \ F1 \  b2
//	        '--, '--, '--, '--, '--, '--, '
//	            \ A2 \ B2 \ C2 \ D2 \ E2 \  b3
//	             '--, '--, '--, '--, '--, '
//	                 \ A3 \ B3 \ C3 \ D3 \  b4
//	                  '--, '--, '--, '--, '
//	                      \ A4 \ B4 \ C4 \  b5
//	                       '--, '--, '--, '
//                             \ A5 \  5 \  b6
//                              '----'----'
//
// diagsBC return and array of diagonal factors size = b + c - 1
func (ts *T32s) tDiags(num, den Z, s1, s2 N32) ([][]N, N) {
	diags := make([][]N, s1)
	for d := range diags {
		diags[d] = make([]N, 0)
	}
	for x := N(1); x <= N(s1); x++ {
		for y := N(1); y <= x; y++ {
			if y > N(s2) {
				continue
			}
			pos := int(x - y)
			// cos = n/d
			// z²  = x² + y² - 2xycos
			// z²  = x² + y² - 2xyn/d
			// z²d = (x² + y²)d - 2xyn
			z := (x*x + y*y)*N(den) - 2*x*y*N(num)
			diags[pos] = append(diags[pos], z*N(den))
		}
	}
	return diags, N(den)
}

// tCosAplusB returns cosAcosB - sinAsinB.
// cosA and cosB are rationals: cosA = nA/dA, cosB= nB/dB
// cosAcosB - sinAsinB simplifies to:
//	         ____________________________________
//	nA*nB - √(dA + nA)(dA - nA)(dB + nB)(dB - nB)
//	---------------------------------------------
//	                     dA*dB
//	where are rationals: 
func (ts *T32s) tCosAplusB(tA *T, aA Tang, tB *T, aB Tang) (*A32, error) {
	nA,dA := ts.cos(tA, aA)
	nB,dB := ts.cos(tB, aB)
	a := N(dA)*N(dB)
	b := nA*nB
	c := Z(-1)
	d := (dA + nA)*(dA - nA) * (dB + nB)*(dB - nB)
	return ts.aNew3(a,b,c,d)
}

// tLawOfCos returns
//    ___________________
//	 √ y² + z² - 2xycosX
func (ts *T32s) tLawOfCos(y, z N32, cosX *A32) (*A32, error) {
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

func (ts *T32s) tAlphaCosines(t *T) (cosA, cosB, cosC *A32, err error) {
	alpha := Z(t.a) // slur
	b := Z(t.b)
	c := Z(t.c)
	if cosA, err = ts.aNew1(2*N(b)*N(c), b*b + c*c - alpha); err != nil {
		return
	} else if cosB, err = ts.aNew3(2*N(alpha)*N(c), 0, alpha + c*c - b*b, alpha); err != nil {
		return
	} else if cosC, err = ts.aNew3(2*N(alpha)*N(b), 0, alpha + b*b - c*c, alpha); err != nil {
		return
	}
	return
}

func (ts *T32s) tBetaCosines(t *T) (cosA, cosB, cosC *A32, err error) {
	a := Z(t.a)
	beta := Z(t.b) // slur
	c := Z(t.c)
	if cosA, err = ts.aNew3(2*N(beta)*N(c), 0, beta + c*c - a*a, beta); err != nil {
		return
	} else if cosB, err = ts.aNew1(2*N(a)*N(c), a*a + c*c - beta); err != nil {
		return
	} else if cosC, err = ts.aNew3(2*N(a)*N(beta), 0, a*a + beta - c*c, beta); err != nil {
		return
	}
	return
}

func (ts *T32s) tGammaCosines(t *T) (cosA, cosB, cosC *A32, err error) {
	a := Z(t.a)
	b := Z(t.b)
	gamma := Z(t.c)
	if cosA, err = ts.aNew3(2*N(b)*N(gamma), 0, b*b + gamma - a*a, gamma); err != nil {
		return
	} else if cosB, err = ts.aNew3(2*N(a)*N(gamma), 0, a*a + gamma - b*b, gamma); err != nil {
		return
	} else if cosC, err = ts.aNew1(2*N(a)*N(b), a*a + b*b - gamma); err != nil {
		return
	}
	return
}

func (ts *T32s) tRatCosines(tri *T) (tRats *TRats) {
	tRats = &TRats{}
	for _, ang := range []Tang{ TangA, TangB, TangC } {
		num, den := ts.cos(tri, ang)
		tRats.addRat(ang, num, den)
	}
	return
}

func (ts *T32s) tRatCosXY(x, y *TRat) (*A32, error) {
	dA, nA := x.den, x.num
	dB, nB := y.den, y.num
	a := N(dA)*N(dB)
	b := nA*nB
	c := Z(-1)
	d := (dA + nA)*(dA - nA) * (dB + nB)*(dB - nB)
	return ts.aNew3(a,b,c,d)
}

func (ts *T32s) tRatCos2XY(a, d *TRat) (*A32, error) {
	an2 := a.num*a.num
	ad2 := a.den*a.den
	return ts.aNew3(
		N(ad2)*N(d.den),     // a
		(2*an2 - ad2)*d.num, // b
		-2*a.num,            // c
		a.S()*d.S(),         // d
	)
}

func (ts *T32s) tRatCosXYZ(aRat, bRat, cRat *TRat) (*A32, error) {
	return nil, nil
}




type Ts struct {
	tris []*T
}

func (t *Ts) AddTris(max N32) {
	t.tris = make([]*T, 0)
	for a:=N32(1); a <= max; a++ {
		for b:=N32(1); b <= a; b++ {
			for c:=N32(1); c <= b; c++ {
				if a < b+c {
					t.triNew(a, b, c)
				}
			}
		}
	}
}

func (t *Ts) triNew(a, b, c N32) {
	gcd := NatGCD(a, NatGCD(b, c))
	ga, gb, gc := a / gcd, b / gcd, c / gcd
	for _, t1 := range t.tris {
		if t1.a == ga && t1.b == gb && t1.c == gc {
			return
		}
	}
	t.tris = append(t.tris, &T{ a:a, b:b, c:c })
}



type TRat struct {
	angle Tang
	num   Z
	den   Z
}

// S returns den*den - num*num
func (t *TRat) S() Z {
	return t.den*t.den - t.num*t.num
}

func (t *TRat) Tex() string {
	n, d := t.num, t.den
	if d < 0 {
		n = -n
		d = -d
	} 
	if n == 0 {
		return "0"
	} else if d == 1 {
		return fmt.Sprintf("%d", n)
	} else if n > 0 {
		return fmt.Sprintf("\\frac{%d}{%d}", n, d)
	} else {
		return fmt.Sprintf("-\\frac{%d}{%d}", -n, d)
	}
}

type TRats struct {
	rats []*TRat
}

func (t *TRats) addRat(angle Tang, num, den Z) {
	if t.rats == nil {
		t.rats = make([]*TRat, 0)
	} else {
		for _, rat := range t.rats {
			if rat.num == num && rat.den == den {
				return
			}
		}
	}
	t.rats = append(t.rats, &TRat{angle, num, den})
}




