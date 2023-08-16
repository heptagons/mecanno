package nest

import (
	"fmt"
)

type Tri2 struct {
	tA  *Tri
	tB  *Tri
	vA  int
	vB  int
	sin *A32
	cos *A32
}

func newTri2(tA, tB *Tri, vA, vB int) (*Tri2, error) {
	if tA == nil || tB == nil || vA < 0 || vA > 2 || vB < 0 || vB > 2 {
		return nil, ErrInvalid
	} else {
		return &Tri2{
			tA: tA,
			tB: tB,
			vA: vA,
			vB: vB,
		}, nil
	}
}

func (t *Tri2) String() string {
	return fmt.Sprintf("%v'%d %v'%d sin=%s cos=%s", t.tA.abc, t.vA, t.tB.abc, t.vB, t.sin, t.cos)
}


type Tri2F struct {
	*TriF
	tri2s []*Tri2
}

func NewTri2F(triF *TriF) *Tri2F {
	return &Tri2F{
		TriF:  triF,
		tri2s: make([]*Tri2, 0),
	}
}

func (t *Tri2F) tri2NewAll() error {
	return t.tri2Unique(func (pair *Tri2) {
		t.tri2s = append(t.tri2s, pair)
	})
}

func (t *Tri2F) tri2NewEqualSin(sin *A32) error {
	return t.tri2Unique(func (pair *Tri2) {
		if pair.sin.Equal(sin) {
			t.tri2s = append(t.tri2s, pair)
		}
	})
}

func (t *Tri2F) tri2NewNotEqualSin(sin *A32) error {
	return t.tri2Unique(func (pair *Tri2) {
		if !pair.sin.Equal(sin) {
			t.tri2s = append(t.tri2s, pair)
		}
	})
}

// tri2Unique scan all Tri triangles stored, compare any three vertice with the
// rest Tri triangles three vertices
// to calculate allcreate unique pairs Tri2
// Tri2 contains references to two trianges Tri which are joined.
// Next example joins two trianges sides 'a' and 'y' and vertices 'C' and 'Z' renamed 'V':
//
//         V                                       V
//        /\                                      /\
//       /  \ u           X'-_                   /  X'-_
//    w /    \      +      \  '-_ z     =       /    \  '-_
//     /      \           y \    '-_           /      \    '-_
//    /     _.-W             Z------Y         /     _.-M------Y
//   /  _.-'                      x          /  _.-'
//  U.-'   v                                U.-' 
//
// The joined vertices angles are added to create a new angle C + Z.

// Up to four new triangles type TriQ can be genterated from each pair.
func (t *Tri2F) tri2Unique(pairF func(*Tri2)) error {
	n := len(t.t1s)
	for p1 := 0; p1 < n; p1++ {
		tA := t.t1s[p1]
		a1s := make(map[N32]bool, 0)
		for a1, s1 := range tA.abc {
			if _, repeated := a1s[s1]; repeated {
				continue
			}
			a1s[s1] = true
			for p2 := p1; p2 < n; p2++ {
				tB := t.t1s[p2]
				a2s := make(map[N32]bool, 0)
				for a2, s2 := range tB.abc {
					if _, repeated := a2s[s2]; repeated {
						continue
					}
					a2s[s2] = true
					if p1 == p2 && a1 < a2 {
						continue
					}
					if pair, err := t.tri2New(tA, tB, a1, a2); err != nil {
						return err
					} else if pair != nil {
						pairF(pair)
					}
				}
			}
		}
	}
	return nil
}

func (t *Tri2F) tri2New(tA, tB *Tri, pA, pB int) (*Tri2, error) {
	tri2, err := newTri2(tA, tB, pA, pB)
	if err != nil {
		return nil, err
	}
	sinA, cosA := tA.sin[pA], tA.cos[pA]
	sinB, cosB := tB.sin[pB], tB.cos[pB]
	// sin(A+B) = sinAcosB + cosAsinB
	if sinAcosB, err := t.aMul(sinA, cosB); err != nil {
		return nil, err 
	} else if sinBcosA, err := t.aMul(sinB, cosA); err != nil {
		return nil, err
	} else if sinAB, err := t.aAdd(sinAcosB, sinBcosA); err != nil {
		return nil, err
	} else {
		tri2.sin = sinAB
	}
	// cos(A+B) = cosAcosB - sinAsinB
	if cosAcosB, err := t.aMul(cosA, cosB); err != nil {
		return nil, err 
	} else if sinAsinB, err := t.aMul(sinA, sinB); err != nil {
		return nil, err
	} else if cosAB, err := t.aAdd(cosAcosB, sinAsinB.Neg()); err != nil {
		return nil, err
	} else {
		tri2.cos = cosAB
	}
	return tri2, nil
}























