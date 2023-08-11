package alg

import (
	"fmt"
)

// TriPair is a group of two Tri sharing a side and a node.
// The two nodes angles are summed to create another one
// Up to four new triangles type TriQ can be genterated from each pair.
type TriPair struct {
	tA, tB   *Tri
	pA, pB   int
	sin, cos *Q32
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

func (t *TriPair) String() string {
	return fmt.Sprintf("%v'%d %v'%d sin=%s cos=%s", t.tA.abc, t.pA, t.tB.abc, t.pB, t.sin, t.cos)
}


type TriPairs struct {
	*Tris
	pairs []*TriPair
}

func NewTriPairs(tris *Tris) *TriPairs {
	return &TriPairs{
		Tris:  tris,
		pairs: make([]*TriPair, 0),
	}
}

func (ts *TriPairs) NewPairs() error {
	return ts.newPairs(func (pair *TriPair) {
		ts.pairs = append(ts.pairs, pair)
	})
}

func (ts *TriPairs) NewPairsSin(sin *Q32) error {
	return ts.newPairs(func (pair *TriPair) {
		if pair.sin.Equal(sin) {
			ts.pairs = append(ts.pairs, pair)
		}
	})
}

func (ts *TriPairs) NewPairsNoSin(sin *Q32) error {
	return ts.newPairs(func (pair *TriPair) {
		if !pair.sin.Equal(sin) {
			ts.pairs = append(ts.pairs, pair)
		}
	})
}

func (ts *TriPairs) newPairs(pairF func(*TriPair)) error {
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
					if pair, err := ts.newPair(t1, t2, a1, a2); err != nil {
//fmt.Println("addPairs err", err)
//						return err
					} else if pair != nil {
						pairF(pair)
					}
				}
			}
		}
	}
	return nil
}

func (ts *TriPairs) newPair(tA, tB *Tri, pA, pB int) (*TriPair, error) {
	pair, err := newTriPair(tA, tB, pA, pB)
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
	}
	return pair, nil
}























