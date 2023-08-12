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
func NewTris(max int) *Tris {
	t := &Tris {
		Q32s: NewQ32s(),
		tris: make([]*Tri, 0),
	}
	for a := N32(1); a <= N32(max); a++ {
		for b := N32(1); b <= a; b++ {
			for c := N32(1); c <= b; c++ {
				if b+c > a {
					t.triNew(a, b, c)
				}
			}
		}
	}
	return t
}

// triNew appends a new Tri to the list rejecting
// repeated Tri sides by scaling, for instance Tri a=6, b=4, c=2
// rejected if already exists a=3, b=2, c=1
func (t *Tris) triNew(a, b, c N32) {
	gcd := NatGCD(a, NatGCD(b, c))
	ga, gb, gc := a / gcd, b / gcd, c / gcd
	for _, tri := range t.tris {
		if tri.abc[0] == ga && tri.abc[1] == gb && tri.abc[2] == gc {
			// scaled version already stored don't append
			return
		}
	}
	t.tris = append(t.tris, &Tri{
		abc: []N32{ a, b, c },
	})
}

// triSinCos calculate for every Tri in the set its sines and cosines
// for each of the three angles.
func (t *Tris) triSinCos() error {
	for _, tri := range t.tris {
		a, b, c := tri.abc[0], tri.abc[1], tri.abc[2]

		if cosA, err := t.triCosC(b, c, a); err != nil {
			return err
		} else
		if cosB, err := t.triCosC(c, a, b); err != nil {
			return err
		} else
		if cosC, err := t.triCosC(a, b, c); err != nil {
			return err
		} else {
			tri.cos = []*Q32 { cosA, cosB, cosC }
		}
		// https://en.wikipedia.org/wiki/Heron%27s_formula Numerical stability
		area2_4 := Z(a+(b+c)) * Z(c-(a-b)) * Z(c+(a-b)) * Z(a+(b-c))
		// area = √(area2_4)/4
		// https://en.wikipedia.org/wiki/Law_of_sines
		// Area = (absinC)/2 -> sinC = 2A/(a*b)
		if sinA, err := t.qNew(2*N(b*c), 0, 1, area2_4); err != nil {
			return err
		} else 
		if sinB, err := t.qNew(2*N(c*a), 0, 1, area2_4); err != nil {
			return err
		} else
		if sinC, err := t.qNew(2*N(a*b), 0, 1, area2_4); err != nil {
			return err
		} else {
			tri.sin = []*Q32 { sinA, sinB, sinC }
		}
	}
	return nil
}

// triCosC returns the rational cosine of the angle C using the law of cosines:
//	       a² + b² - c²
//	cosC = ------------
//	           2ab
func (t *Tris) triCosC(a, b, c N32) (*Q32, error) {
	den64 := 2*N(a)*N(b)
	num64 := Z(a)*Z(a) + Z(b)*Z(b) - Z(c)*Z(c)
	return t.qNew(den64, num64)
}

// triCosLaw2 return the third side (squared) cc. Squared to keep simple the Q32 returned.
// Uses the law of cosines to determine the rational algebraic side cc = aa + bb - 2abcosC
func (t *Tris) triCosLaw2(a, b N32, cosC *Q32) (*Q32, error) {
	if aa_bb, err := t.qNew(1, Z(a)*Z(a) + Z(b)*Z(b)); err != nil { // a*a + b*b
		return nil, err
	} else if ab, err := t.qNew(1, -2*Z(a)*Z(b)); err != nil { // -2a*b
		return nil, err
	} else if abCosC, err := t.qMul(ab, cosC); err != nil { // -2a*b*cosC
		return nil, err
	} else {
		return t.qAdd(aa_bb, abCosC) // a*a + b*b - 2*a*b*cosC
	}
}



