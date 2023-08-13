package nest

import (
	"fmt"
)

// Tri represents a triangle with sides as natural numbers a, b, c.
// To reduce the lists we use the condition: a >= b >= c > 0.
// and the condition a > b+c to avoid "open" trianges with (complex numbers) angles.
// Contains also the three sines and three cosines of the angles, but they are
// calculated by a factory to simplify and reduce only valid algebraic numbers values.
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
	cos []*A32
	sin []*A32
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

// String represent returns a string representation
// including the three sides a,b,c and cosines and sines.
func (t *Tri) String() string {
	return fmt.Sprintf("abc:%v cos:%v sin:%v", t.abc, t.cos, t.sin)
}


// TriF is a factory of triangles and repository.
// Uses the A32s factory to calculate angles for each triangle.
type TriF struct {
	*A32s
	tris  []*Tri
}

// NewTriF build a TriF set with ordered Tri's with sides a,b,c, where:
//	1 <= a <= max
//	a >= b >= c
//	b+c > a
func NewTriF(max int) *TriF {
	t := &TriF {
		A32s: NewA32s(),
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

// triNew appends a new triangle Tri to the list preventing repetitions
// by scaling, for instance triangle a=6, b=4, c=2 is rejected if already
// exists triangle a=3, b=2, c=1.
func (t *TriF) triNew(a, b, c N32) {
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

// triSinCos calculate for every triangle Trie already int he list,
// the triangles three angles sines and cosines stored in each triangle.
func (t *TriF) triSinCos() error {
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
			tri.cos = []*A32 { cosA, cosB, cosC }
		}
		// https://en.wikipedia.org/wiki/Heron%27s_formula Numerical stability
		area2_4 := Z(a+(b+c)) * Z(c-(a-b)) * Z(c+(a-b)) * Z(a+(b-c))
		// area = √(area2_4)/4
		// https://en.wikipedia.org/wiki/Law_of_sines
		// Area = (absinC)/2 -> sinC = 2A/(a*b)
		if sinA, err := t.aNew(2*N(b*c), 0, 1, area2_4); err != nil {
			return err
		} else 
		if sinB, err := t.aNew(2*N(c*a), 0, 1, area2_4); err != nil {
			return err
		} else
		if sinC, err := t.aNew(2*N(a*b), 0, 1, area2_4); err != nil {
			return err
		} else {
			tri.sin = []*A32 { sinA, sinB, sinC }
		}
	}
	return nil
}

// triCosC returns the rational cosine (A32 of size 1) of the angle C using the law of cosines:
//	       a² + b² - c²
//	cosC = ------------
//	           2ab
func (t *TriF) triCosC(a, b, c N32) (*A32, error) {
	den64 := 2*N(a)*N(b)
	num64 := Z(a)*Z(a) + Z(b)*Z(b) - Z(c)*Z(c)
	return t.aNew(den64, num64)
}

// triCosLaw2 return the third side (squared) cc to help later comparisons.
// Uses the law of cosines to determine the rational algebraic side cc = aa + bb - 2abcosC
func (t *TriF) triCosLaw2(a, b N32, cosC *A32) (*A32, error) {
	A := 1
	B := Z(a)*Z(a) + Z(b)*Z(b)
	if aa_bb, err := t.aNew(1, B); err != nil { // a*a + b*b
		return nil, err
	} else if ab, err := t.aNew(1, -2*Z(a)*Z(b)); err != nil { // -2a*b
		return nil, err
	} else if abCosC, err := t.aMul(ab, cosC); err != nil { // -2a*b*cosC
		return nil, err
	} else {
		return t.aAdd(aa_bb, abCosC) // a*a + b*b - 2*a*b*cosC
	}
}



