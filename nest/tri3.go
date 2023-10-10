package nest

import (
	"fmt"
)

type Tri3 struct {
	tri2 int
	max  N32    // max natural side
	min  N32    // min natural side
	cc   *A32   // c rational algebraic side
	abc  []*A32 // at leat one side rational algebraic
}

func newTri3(tri2 int, max, min N32, cc *A32, c *A32) (t *Tri3, e error) {
	t = &Tri3{
		tri2: tri2,
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

func (t *Tri3) Equal(u *Tri3) bool {
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

func (t *Tri3) String() string {
	return fmt.Sprintf("tri2=%d %v", t.tri2, t.abc)
}


type Tri3F struct {
	*Tri2F
	tri3s []*Tri3
	errs  []error
}

func NewTri3F(tri2F *Tri2F) *Tri3F {
	return &Tri3F{
		Tri2F: tri2F,
		tri3s: make([]*Tri3, 0),
		errs:  make([]error, 0),
	}
}

func (t *Tri3F) tri3All() {
	for tri2 := range t.tri2s {
		if tri3s, err := t.tri3New(tri2); err != nil {
			t.errs = append(t.errs, err)
		} else {
			for _, triq := range tri3s {
				t.appendUnique(triq)
			}
		}
	}
}

func (t *Tri3F) appendUnique(s *Tri3) {
	for _, triq := range t.tri3s {
		if triq.Equal(s) {
			return
		}
	}
	t.tri3s = append(t.tri3s, s)
}

func (t *Tri3F) tri3New(tri2 int) ([]*Tri3, error) {
	tri3s := make([]*Tri3, 0)
	// build tris, triangles with two sides natural and one side Q
	p := t.tri2s[tri2]
	for _, a := range p.tA.otherSides(p.vA) {
		for _, b := range p.tB.otherSides(p.vB) {
			max, min := a, b
			if max < min {
				max, min = b, a
			}
			repeated := false
			for _, triq := range tri3s {
				if max == triq.max && min == triq.min {
					repeated = true
				}
			}
			if repeated {
				continue
			}
			if cc, err := t.tri3CosLaw2(max, min, p.cos); err != nil {
				return nil, err
			} else if c, err := t.ASqrt(cc); err != nil {
				return nil, err
			} else if len(c.num) <= 1 && c.den == 1 {
				// reject triq with three sides natural (c.num=1, c.den=1)
				// since already is a regular triangle tri
				// prevent a triq with all naturals like below [4 3 2]
					// 4 [2 2 1]'0 [4 3 2]'2 sin=√15/4 cos=-1/4
				// tris=[[2√6 4 2] [4 3 2] [√19 4 1] [√46/2 3 1]]
				continue
			} else if triq, err := newTri3(tri2, max, min, cc, c); err != nil {
				//return nil, err
			} else {
				tri3s = append(tri3s, triq)
			}
		}
	}
	return tri3s, nil
}

// triCosLaw2 return the third side (squared) cc. Squared to keep simple the A32 returned.
// Uses the law of cosines to determine the rational algebraic side cc = aa + bb - 2abcosC
func (t *Tri3F) tri3CosLaw2(a, b N32, cosC *A32) (*A32, error) {
	if aa_bb, err := t.ANew(1, Z(a)*Z(a) + Z(b)*Z(b)); err != nil { // a*a + b*b
		return nil, err
	} else if ab, err := t.ANew(1, -2*Z(a)*Z(b)); err != nil { // -2a*b
		return nil, err
	} else if abCosC, err := t.aMul(ab, cosC); err != nil { // -2a*b*cosC
		return nil, err
	} else {
		return t.aAdd(aa_bb, abCosC) // a*a + b*b - 2*a*b*cosC
	}
}





