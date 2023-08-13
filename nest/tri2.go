package nest

import (
	"fmt"
)

// Tri2 contains references to two trianges Tri which are joined.
// Next example joins two trianges sides 'a' and 'y' and vertices 'C' and 'Z' renamed 'V':
//         B                                     B
//        /\                                    /\
//       /  \ a         X'-_                   /  X'-_
//    c /    \    +      \  '-_ z    =        /    \  '-_
//     /      \         y \    '-_           /      \    '-_
//    /    __--C           Z------Y         /    __--V------Y
//   / __--                    x           / __--
//  A--    b                              A-- 
//
// The joined vertices angles are added to create a new angle

// Up to four new triangles type TriQ can be genterated from each pair.
//
type Tri2 struct {
	tA, tB   *Tri
	pA, pB   int
	sin, cos *A32
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


type Tri2F struct {
	*Tris
	pairs []*Tri2
}

func NewTri2F(tris *Tris) *Tri2F {
	return &Tri2F{
		Tris:  tris,
		pairs: make([]*Tri2, 0),
	}
}

func (ts *Tri2F) tri2NewAll() error {
	return ts.tri2Filter(func (pair *Tri2) {
		ts.pairs = append(ts.pairs, pair)
	})
}

func (ts *Tri2F) tri2NewEqualSin(sin *A32) error {
	return ts.tri2Filter(func (pair *Tri2) {
		if pair.sin.Equal(sin) {
			ts.pairs = append(ts.pairs, pair)
		}
	})
}

func (ts *Tri2F) tri2NewNotEqualSin(sin *A32) error {
	return ts.tri2Filter(func (pair *Tri2) {
		if !pair.sin.Equal(sin) {
			ts.pairs = append(ts.pairs, pair)
		}
	})
}

func (ts *Tri2F) tri2Filter(pairF func(*Tri2)) error {
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

func (ts *Tri2F) tri2New(tA, tB *Tri, pA, pB int) (*Tri2, error) {
	pair, err := newTri2(tA, tB, pA, pB)
	if err != nil {
		return nil, err
	}
	sinA, cosA := tA.sin[pA], tA.cos[pA]
	sinB, cosB := tB.sin[pB], tB.cos[pB]
	// sin(A+B) = sinAcosB + cosAsinB
	if sinAcosB, err := ts.aMul(sinA, cosB); err != nil {
		return nil, err 
	} else if sinBcosA, err := ts.aMul(sinB, cosA); err != nil {
		return nil, err
	} else if sinAB, err := ts.aAdd(sinAcosB, sinBcosA); err != nil {
		return nil, err
	} else {
		pair.sin = sinAB
	}
	// cos(A+B) = cosAcosB - sinAsinB
	if cosAcosB, err := ts.aMul(cosA, cosB); err != nil {
		return nil, err 
	} else if sinAsinB, err := ts.aMul(sinA, sinB); err != nil {
		return nil, err
	} else if cosAB, err := ts.aAdd(cosAcosB, sinAsinB.Neg()); err != nil {
		return nil, err
	} else {
		pair.cos = cosAB
	}
	return pair, nil
}























