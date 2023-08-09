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
	abc []N32   // Three natural sides
	cos []*Q32
	sin []*Q32
}

// otherSides return the two sides not pointed by pos
func (t *Tri32) otherSides(pos int) []N32 {
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

func (t *Tri32) String() string {
	return fmt.Sprintf("abc:%v cos:%v sin:%v", t.abc, t.cos, t.sin)
}

type Tri32Q struct {
	// Triangle
	a N32  // first natural side
	b N32  // second natural side
	c *Q32 // third rational algebraic side
}

func (t *Tri32Q) String() string {
	return fmt.Sprintf("[%d %d %s]", t.a, t.b, t.c.String())
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

// cosLaw return the third side (squared) cc. Squared to keep simple the Q32 returned.
// Uses the law of cosines to determine the rational algebraic side cc = aa + bb - 2abcosC
func (ts *Tris32) cosLaw(a, b N32, cosC *Q32) (*Q32, error) {
	if aa_bb, err := ts.newQ32(1, Z(a)*Z(a) + Z(b)*Z(b)); err != nil { // a*a + b*b
		return nil, err
	} else if ab, err := ts.newQ32(1, -2*Z(a)*Z(b)); err != nil { // -2a*b
		return nil, err
	} else if abCosC, err := ts.MulQ(ab, cosC); err != nil { // -2a*b*cosC
		return nil, err
	} else if cc, err := ts.AddQ(aa_bb, abCosC); err != nil { // a*a + b*b - 2*a*b*cosC
		return nil, err
	} else {
		return ts.sqrtQ(cc)
	}
}

func (ts *Tris32) addPair(tA, tB *Tri32, pA, pB int) (*TriPair32, error) {
	pair, err := newTriPair32(tA, tB, pA, pB)
	if err != nil {
		return nil, err
	}
	sinA, cosA := tA.sin[pA], tA.cos[pA]
	sinB, cosB := tB.sin[pB], tB.cos[pB]
	// sin(A+B) = sinAcosB + cosAsinB
	if sinAcosB, err := ts.MulQ(sinA, cosB); err != nil {
		return nil, err 
	} else if sinBcosA, err := ts.MulQ(sinB, cosA); err != nil {
		return nil, err
	} else if sinAB, err := ts.AddQ(sinAcosB, sinBcosA); err != nil {
		return nil, err
	} else {
		pair.sin = sinAB
	}
	// cos(A+B) = cosAcosB - sinAsinB
	if cosAcosB, err := ts.MulQ(cosA, cosB); err != nil {
		return nil, err 
	} else if sinAsinB, err := ts.MulQ(sinA, sinB); err != nil {
		return nil, err
	} else if cosAB, err := ts.AddQ(cosAcosB, sinAsinB.Neg()); err != nil {
		return nil, err
	} else {
		pair.cos = cosAB
		// build tris, triangles with two sides natural and one side Q
		triqs := make([]*Tri32Q, 0)
		for _, a := range tA.otherSides(pA) {
			for _, b := range tB.otherSides(pB) {
				max, min := a, b
				if max < min {
					max, min = b, a
				}
				repeated := false
				for _, triq := range triqs {
					if max == triq.a && min == triq.b {
						repeated = true
					}
				}
				if !repeated {
					if c, err := ts.cosLaw(max, min, cosAB); err != nil {
						return nil, err
					} else {
						triqs = append(triqs, &Tri32Q{
							a: max,
							b: min,
							c: c,
						})
					}
				}
			}
		}
		pair.triqs = triqs
	}
	return pair, nil
}

func (ts *Tris32) AddPairs(results func(pair *TriPair32, err error)) {
	n := len(ts.list)
	for p1 := 0; p1 < n; p1++ {
		t1 := ts.list[p1]
		a1s := make(map[N32]bool, 0)
		for a1, s1 := range t1.abc {
			if _, repeated := a1s[s1]; repeated {
				continue
			}
			a1s[s1] = true
			for p2 := p1; p2 < n; p2++ {
				t2 := ts.list[p2]
				a2s := make(map[N32]bool, 0)
				for a2, s2 := range t2.abc {
					if _, repeated := a2s[s2]; repeated {
						continue
					}
					a2s[s2] = true
					if p1 == p2 && a1 < a2 {
						continue
					}
					results(ts.addPair(t1, t2, a1, a2))
				}
			}
		}
	}
}

type TriPair32 struct {
	tA, tB   *Tri32
	pA, pB   int
	sin, cos *Q32
	triqs    []*Tri32Q
}

func newTriPair32(tA, tB *Tri32, pA, pB int) (*TriPair32, error) {
	if tA == nil || tB == nil || pA < 0 || pA > 2 || pB < 0 || pB > 2 {
		return nil, ErrInvalid
	} else {
		return &TriPair32{
			tA: tA,
			tB: tB,
			pA: pA,
			pB: pB,
		}, nil
	}
}

func (t *TriPair32) String() string {
	return fmt.Sprintf("%v'%d %v'%d sin=%s cos=%s tris=%v", t.tA.abc, t.pA, t.tB.abc, t.pB, t.sin, t.cos, t.triqs)
}
