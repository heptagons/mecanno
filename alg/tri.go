package alg

import (
	"fmt"
)

// Triangle is a valid triangle with positive sides:
//	a >= b >= c > 0
//  a > b+c
//
//           _ -C
//     a _ -   /
//   _ -      /
// B_        / b
//   -_     /
//  c  -_  /  
//       -A
//
// A,B, and C the angles to opposite abc a,b and c.
type Tri32 struct { // Triangle
	abc []N32
	cos []*Q32
	sin []*Q32
}

func (t *Tri32) String() string {
	return fmt.Sprintf("abc:%v cos:%v sin:%v", t.abc, t.cos, t.sin)
}



type Tris32 struct {
	*Q32s

	list []*Tri32
}

func NewA32Tris(max int, factory *N32s) *Tris32 {
	ts := &Tris32 {
		Q32s: &Q32s{
			N32s: factory,
		},
		list: make([]*Tri32, 0),
	}
	for a := N32(1); a <= N32(max); a++ {
		for b := N32(1); b <= a; b++ {
			for c := N32(1); c <= b; c++ {
				if b+c > a {
					ts.add(a, b, c)
				}
			}
		}
	}
	return ts
}

func (ts *Tris32) add(a, b, c N32) {
	gcd := NatGCD(a, NatGCD(b, c))
	ga, gb, gc := a / gcd, b / gcd, c / gcd
	for _, t := range ts.list {
		if t.abc[0] == ga && t.abc[1] == gb && t.abc[2] == gc {
			// scaled version already stored don't append
			return
		}
	}
	ts.list = append(ts.list, &Tri32{
		abc: []N32{ a, b, c },
	})
}

func (ts *Tris32) setSinCos() error {
	for _, t := range ts.list {
		a, b, c := t.abc[0], t.abc[1], t.abc[2]

		if cosA, err := ts.cosC(b, c, a); err != nil {
			return err
		} else
		if cosB, err := ts.cosC(c, a, b); err != nil {
			return err
		} else
		if cosC, err := ts.cosC(a, b, c); err != nil {
			return err
		} else {
			t.cos = []*Q32 { cosA, cosB, cosC }
		}
		// https://en.wikipedia.org/wiki/Heron%27s_formula Numerical stability
		area2_4 := Z(a+(b+c)) * Z(c-(a-b)) * Z(c+(a-b)) * Z(a+(b-c))
		// area = √(area2_4)/4
		// https://en.wikipedia.org/wiki/Law_of_sines
		// Area = (absinC)/2 -> sinC = 2A/(a*b)
		if sinA, err := ts.newQ32(2*N(b*c), 0, 1, area2_4); err != nil {
			return err
		} else 
		if sinB, err := ts.newQ32(2*N(c*a), 0, 1, area2_4); err != nil {
			return err
		} else
		if sinC, err := ts.newQ32(2*N(a*b), 0, 1, area2_4); err != nil {
			return err
		} else {
			t.sin = []*Q32 { sinA, sinB, sinC }
		}
	}
	return nil
}

// cosC returns the rational cosine of the angle C using the law of cosines:
//	       a² + b² - c²
//	cosC = ------------
//	           2ab
func (ts *Tris32) cosC(a, b, c N32) (*Q32, error) {
	den64 := 2*N(a)*N(b)
	num64 := Z(a)*Z(a) + Z(b)*Z(b) - Z(c)*Z(c)
	return ts.newQ32(den64, num64)
}

// sin(A+B) = sinAcosB + cosAsinB
func (ts *Tris32) SinAdd(tA, tB *Tri32, pA, pB int) (*Q32, error) {
	if tA == nil || tB == nil || pA < 0 || pA > 2 || pB < 0 || pB > 2 {
		return nil, ErrInvalid
	}
	sinA, cosA := tA.sin[pA], tA.cos[pA]
	sinB, cosB := tB.sin[pB], tB.cos[pB]
	if sinAcosB, err := ts.MulQ(sinA, cosB); err != nil {
		return nil, err
	} else if sinBcosA, err := ts.MulQ(sinB, cosA); err != nil {
		return nil, err
	} else {
		return ts.AddQ(sinAcosB, sinBcosA)
	}
}
