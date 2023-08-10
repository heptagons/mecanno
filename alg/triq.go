package alg

import (
	"fmt"
)

type TriQ struct {
	pair *TriPair
	max  N32      // max natural side
	min  N32      // min natural side
	abc  []*Q32   // at leat one side rational algebraic
}

func newTriQ(pair *TriPair, max, min N32, cc *Q32, c *Q32) (t *TriQ, e error) {
	t = &TriQ{
		pair: pair,
		max:  max,
		min:  min,
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


type TriQs struct {
	*TriPairs
	triqs []*TriQ
}

func NewTriQs(p *TriPairs) *TriQs {
	return &TriQs{
		TriPairs: p,
		triqs:    make([]*TriQ, 0),
	}
}

func (tqs *TriQs) All() error {
	for _, pair := range tqs.pairs {
		if err := tqs.SetTriqs(pair); err != nil {
			//
		}
	}
	return nil
}

func (tqs *TriQs) SetTriqs(t *TriPair) error {
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
				if cc, err := tqs.cosLaw2(max, min, t.cos); err != nil {
					return err
				} else if c, err := tqs.sqrtQ(cc); err != nil {
					return err
				} else if len(c.num) <= 1 && c.den == 1 {
					// prevent a triq with all naturals like below [4 3 2]
						// 4 [2 2 1]'0 [4 3 2]'2 sin=√15/4 cos=-1/4
					// tris=[[2√6 4 2] [4 3 2] [√19 4 1] [√46/2 3 1]]
					continue
				} else if triq, err := newTriQ(t, max, min, cc, c); err != nil {
					return err
				} else {
					triqs = append(triqs, triq)
				}
			}
		}
	}
	tqs.triqs = triqs
	return nil
}

// cosLaw return the third side (squared) cc. Squared to keep simple the Q32 returned.
// Uses the law of cosines to determine the rational algebraic side cc = aa + bb - 2abcosC
func (tqs *TriQs) cosLaw2(a, b N32, cosC *Q32) (*Q32, error) {
	if aa_bb, err := tqs.newQ32(1, Z(a)*Z(a) + Z(b)*Z(b)); err != nil { // a*a + b*b
		return nil, err
	} else if ab, err := tqs.newQ32(1, -2*Z(a)*Z(b)); err != nil { // -2a*b
		return nil, err
	} else if abCosC, err := tqs.MulQ(ab, cosC); err != nil { // -2a*b*cosC
		return nil, err
	} else {
		return tqs.AddQ(aa_bb, abCosC) // a*a + b*b - 2*a*b*cosC
	}
}



