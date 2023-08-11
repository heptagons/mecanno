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
type Tri struct { // Triangle
	abc []N32   // Three natural sides
	cos []*Q32
	sin []*Q32
}

// otherSides return the two sides not pointed by pos (0,1,2)
func (t *Tri) otherSides(pos int) []N32 {
	switch pos {
	case 0:
		return []N32{ t.abc[1], t.abc[2] }
	case 1:
		return []N32{ t.abc[0], t.abc[2] }
	case 2:
		return []N32{ t.abc[0], t.abc[1] }
	}
	return nil
}

func (t *Tri) String() string {
	return fmt.Sprintf("abc:%v cos:%v sin:%v", t.abc, t.cos, t.sin)
}


// Tris holds a set of Tri with no sides repeated
type Tris struct {
	*Q32s
	tris  []*Tri
}

// NewTris build a Tris set with ordered Tri's with sides a,b,c, where:
//	1 <= a <= max
//	a >= b >= c
//	b+c > a
func NewTris(max int, factory *N32s) *Tris {
	ts := &Tris {
		Q32s: &Q32s{
			N32s: factory,
		},
		tris: make([]*Tri, 0),
	}
	for a := N32(1); a <= N32(max); a++ {
		for b := N32(1); b <= a; b++ {
			for c := N32(1); c <= b; c++ {
				if b+c > a {
					ts.addTri(a, b, c)
				}
			}
		}
	}
	return ts
}

// addTri appends a new Tri to the list rejecting
// repeated Tri sides by scaling, for instance Tri a=6, b=4, c=2
// rejected if already exists a=3, b=2, c=1
func (ts *Tris) addTri(a, b, c N32) {
	gcd := NatGCD(a, NatGCD(b, c))
	ga, gb, gc := a / gcd, b / gcd, c / gcd
	for _, t := range ts.tris {
		if t.abc[0] == ga && t.abc[1] == gb && t.abc[2] == gc {
			// scaled version already stored don't append
			return
		}
	}
	ts.tris = append(ts.tris, &Tri{
		abc: []N32{ a, b, c },
	})
}

// SetSinCos calculate for every Tri in the set its sines and cosines
// for each of the three angles.
func (ts *Tris) SetSinCos() error {
	for _, t := range ts.tris {
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
func (ts *Tris) cosC(a, b, c N32) (*Q32, error) {
	den64 := 2*N(a)*N(b)
	num64 := Z(a)*Z(a) + Z(b)*Z(b) - Z(c)*Z(c)
	return ts.newQ32(den64, num64)
}

