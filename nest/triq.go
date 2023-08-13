package nest

import (
	"fmt"
)

type TriQ struct {
	pair int
	max  N32      // max natural side
	min  N32      // min natural side
	cc   *A32     // c rational algebraic side
	abc  []*A32   // at leat one side rational algebraic
}

func newTriQ(pair int, max, min N32, cc *A32, c *A32) (t *TriQ, e error) {
	t = &TriQ{
		pair: pair,
		max:  max,
		min:  min,
		cc:   cc,
	}
	a := newA32(1, Z32(max))
	b := newA32(1, Z32(min))
	if cab, err := cc.GreaterThanZ(Z(max)*Z(max)); err != nil {
		e = err
	} else if cab { // c >= a >= b
		t.abc = []*A32{ c, a, b }
		return
	} else if acb, err := cc.GreaterThanZ(Z(min)*Z(min)); err != nil {
		e = err
	} else if acb { // a >= c >= c
		t.abc = []*A32{ a, c, b }
	} else { // a >= b >= c
		t.abc = []*A32{ a, b, c }
	}
	return
}

func (t *TriQ) Equal(u *TriQ) bool {
	if t == nil || u == nil {
		return true
	}
	if t.max != u.max {
		return false
	}
	if t.min != u.min {
		return false
	}
	return t.cc.Equal(u.cc)
}

func (t *TriQ) String() string {
	return fmt.Sprintf("pair=%d %v", t.pair, t.abc)
}


type TriQs struct {
	*Tri2s
	triqs []*TriQ
	errs  []error
}

func NewTriQs(tri2s *Tri2s) *TriQs {
	return &TriQs{
		Tri2s: tri2s,
		triqs: make([]*TriQ, 0),
		errs:  make([]error, 0),
	}
}

func (t *TriQs) triqsAll() {
	for pair := range t.pairs {
		if triqs, err := t.triqsNew(pair); err != nil {
			t.errs = append(t.errs, err)
		} else {
			for _, triq := range triqs {
				t.appendUnique(triq)
			}
		}
	}
}

func (t *TriQs) appendUnique(s *TriQ) {
	for _, triq := range t.triqs {
		if triq.Equal(s) {
			return
		}
	}
	t.triqs = append(t.triqs, s)
}

func (t *TriQs) triqsNew(pair int) ([]*TriQ, error) {
	triqs := make([]*TriQ, 0)
	// build tris, triangles with two sides natural and one side Q
	p := t.pairs[pair]
	for _, a := range p.tA.otherSides(p.pA) {
		for _, b := range p.tB.otherSides(p.pB) {
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
			if repeated {
				continue
			}
			if cc, err := t.triqsCosLaw2(max, min, p.cos); err != nil {
				return nil, err
			} else if c, err := t.qSqrt(cc); err != nil {
				return nil, err
			} else if len(c.num) <= 1 && c.den == 1 {
				// reject triq with three sides natural (c.num=1, c.den=1)
				// since already is a regular triangle tri
				// prevent a triq with all naturals like below [4 3 2]
					// 4 [2 2 1]'0 [4 3 2]'2 sin=√15/4 cos=-1/4
				// tris=[[2√6 4 2] [4 3 2] [√19 4 1] [√46/2 3 1]]
				continue
			} else if triq, err := newTriQ(pair, max, min, cc, c); err != nil {
				//return nil, err
			} else {
				triqs = append(triqs, triq)
			}
		}
	}
	return triqs, nil
}

// triCosLaw2 return the third side (squared) cc. Squared to keep simple the A32 returned.
// Uses the law of cosines to determine the rational algebraic side cc = aa + bb - 2abcosC
func (t *TriQs) triqsCosLaw2(a, b N32, cosC *A32) (*A32, error) {
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





