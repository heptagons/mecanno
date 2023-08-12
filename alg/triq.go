package alg

import (
	"fmt"
)

type TriQ struct {
	pair int
	max  N32      // max natural side
	min  N32      // min natural side
	c    *Q32     // c rational algebraic side
	abc  []*Q32   // at leat one side rational algebraic
}

func newTriQ(pair int, max, min N32, cc *Q32, c *Q32) (t *TriQ, e error) {
	t = &TriQ{
		pair: pair,
		max:  max,
		min:  min,
		c:    c,
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
	return t.c.Equal(u.c)
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
			if !repeated {
				if cc, err := t.triCosLaw2(max, min, p.cos); err != nil {
					return nil, err
				} else if c, err := t.qSqrt(cc); err != nil {
					return nil, err
				} else if len(c.num) <= 1 && c.den == 1 {
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
	}
	return triqs, nil
}



