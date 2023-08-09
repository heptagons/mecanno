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

// otherSides return the two sides not pointed by pos
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

type TriQ struct {
	max N32    // max natural side
	min N32    // min natural side
	abc []*Q32 // at leat one side rational algebraic
}

func newTri32Q(max, min N32, cc *Q32, c *Q32) (t *TriQ, e error) {
	t = &TriQ{
		max: max,
		min: min,
	}
	a := newQ32(1, Z32(max))
	b := newQ32(1, Z32(min))
	if cab, err := cc.GreaterThanZ(Z(max)*Z(max)); err != nil {
		e = err
	} else if acb, err := cc.GreaterThanZ(Z(min)*Z(min)); err != nil {
		e = err
	} else if cab {
		t.abc = []*Q32{ c, a, b }
	} else if acb {
		t.abc = []*Q32{ a, c, b }
	} else {
		t.abc = []*Q32{ a, b, c }
	}
	return
}

func (t *TriQ) String() string {
	return fmt.Sprintf("%v", t.abc)
}


type TriPair struct {
	tA, tB   *Tri
	pA, pB   int
	sin, cos *Q32
	triqs    []*TriQ
}

func newTriPair(tA, tB *Tri, pA, pB int) (*TriPair, error) {
	if tA == nil || tB == nil || pA < 0 || pA > 2 || pB < 0 || pB > 2 {
		return nil, ErrInvalid
	} else {
		return &TriPair{
			tA: tA,
			tB: tB,
			pA: pA,
			pB: pB,
		}, nil
	}
}

func (t *TriPair) setTriqs(ts *Tris) error {
	// build tris, triangles with two sides natural and one side Q
	triqs := make([]*TriQ, 0)
	for _, a := range t.tA.otherSides(t.pA) {
		for _, b := range t.tB.otherSides(t.pB) {
			max, min := a, b
			if max < min {
				max, min = b, a
			}
			repeated := false
			for _, triq := range triqs {
				if max == triq.max && min == triq.min {
					repeated = true
				}
			}
			if !repeated {
				if cc, err := ts.cosLaw2(max, min, t.cos); err != nil {
					return err
				} else if c, err := ts.sqrtQ(cc); err != nil {
					return err
				} else if len(c.num) <= 1 && c.den == 1 {
					// prevent a triq with all naturals like below [4 3 2]
						// 4 [2 2 1]'0 [4 3 2]'2 sin=√15/4 cos=-1/4
					// tris=[[2√6 4 2] [4 3 2] [√19 4 1] [√46/2 3 1]]
					continue
				} else if triq, err := newTri32Q(max, min, cc, c); err != nil {
					return err
				} else {
					triqs = append(triqs, triq)
				}
			}
		}
	}
	t.triqs = triqs
	return nil
}

func (t *TriPair) String() string {
	return fmt.Sprintf("%v'%d %v'%d sin=%s cos=%s\n\ttris=%v", t.tA.abc, t.pA, t.tB.abc, t.pB, t.sin, t.cos, t.triqs)
}
