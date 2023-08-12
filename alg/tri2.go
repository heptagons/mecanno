package alg

import (
	"fmt"
)

// Tri2 is a group of two Tri sharing a side and a node.
// The two nodes angles are summed to create another one
// Up to four new triangles type TriQ can be genterated from each pair.
type Tri2 struct {
	tA, tB   *Tri
	pA, pB   int
	sin, cos *Q32
}

func newTri2(tA, tB *Tri, pA, pB int) (*Tri2, error) {
	if tA == nil || tB == nil || pA < 0 || pA > 2 || pB < 0 || pB > 2 {
		return nil, ErrInvalid
	} else {
		return &Tri2{
			tA: tA,
			tB: tB,
			pA: pA,
			pB: pB,
		}, nil
	}
}

func (t *Tri2) String() string {
	return fmt.Sprintf("%v'%d %v'%d sin=%s cos=%s", t.tA.abc, t.pA, t.tB.abc, t.pB, t.sin, t.cos)
}


type Tri2s struct {
	*Tris
	pairs []*Tri2
}

func NewTri2s(tris *Tris) *Tri2s {
	return &Tri2s{
		Tris:  tris,
		pairs: make([]*Tri2, 0),
	}
}

func (ts *Tri2s) tri2NewAll() error {
	return ts.tri2Filter(func (pair *Tri2) {
		ts.pairs = append(ts.pairs, pair)
	})
}

func (ts *Tri2s) tri2NewEqualSin(sin *Q32) error {
	return ts.tri2Filter(func (pair *Tri2) {
		if pair.sin.Equal(sin) {
			ts.pairs = append(ts.pairs, pair)
		}
	})
}

func (ts *Tri2s) tri2NewNotEqualSin(sin *Q32) error {
	return ts.tri2Filter(func (pair *Tri2) {
		if !pair.sin.Equal(sin) {
			ts.pairs = append(ts.pairs, pair)
		}
	})
}

func (ts *Tri2s) tri2Filter(pairF func(*Tri2)) error {
	n := len(ts.tris)
	for p1 := 0; p1 < n; p1++ {
		t1 := ts.tris[p1]
		a1s := make(map[N32]bool, 0)
		for a1, s1 := range t1.abc {
			if _, repeated := a1s[s1]; repeated {
				continue
			}
			a1s[s1] = true
			for p2 := p1; p2 < n; p2++ {
				t2 := ts.tris[p2]
				a2s := make(map[N32]bool, 0)
				for a2, s2 := range t2.abc {
					if _, repeated := a2s[s2]; repeated {
						continue
					}
					a2s[s2] = true
					if p1 == p2 && a1 < a2 {
						continue
					}
					if pair, err := ts.tri2New(t1, t2, a1, a2); err != nil {
						//return err
					} else if pair != nil {
						pairF(pair)
					}
				}
			}
		}
	}
	return nil
}

func (ts *Tri2s) tri2New(tA, tB *Tri, pA, pB int) (*Tri2, error) {
	pair, err := newTri2(tA, tB, pA, pB)
	if err != nil {
		return nil, err
	}
	sinA, cosA := tA.sin[pA], tA.cos[pA]
	sinB, cosB := tB.sin[pB], tB.cos[pB]
	// sin(A+B) = sinAcosB + cosAsinB
	if sinAcosB, err := ts.qMul(sinA, cosB); err != nil {
		return nil, err 
	} else if sinBcosA, err := ts.qMul(sinB, cosA); err != nil {
		return nil, err
	} else if sinAB, err := ts.qAdd(sinAcosB, sinBcosA); err != nil {
		return nil, err
	} else {
		pair.sin = sinAB
	}
	// cos(A+B) = cosAcosB - sinAsinB
	if cosAcosB, err := ts.qMul(cosA, cosB); err != nil {
		return nil, err 
	} else if sinAsinB, err := ts.qMul(sinA, sinB); err != nil {
		return nil, err
	} else if cosAB, err := ts.qAdd(cosAcosB, sinAsinB.Neg()); err != nil {
		return nil, err
	} else {
		pair.cos = cosAB
	}
	return pair, nil
}























