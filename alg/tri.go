package alg

import (
	"fmt"
)

type Tri32 struct { // Triangle
	sides []N32
	cos   []*Q32
	sin   []*Q32
}

func (t *Tri32) String() string {
	return fmt.Sprintf("abc:%v cos:%v sin:%v", t.sides, t.cos, t.sin)
}

func (t *Tri32) setSinCos(n *N32s) (overflow bool) {
	a, b, c := t.sides[0], t.sides[1], t.sides[2]
	if nA, dA, overflow := n.CosC(b, c, a); overflow {
		return true
	} else if nB, dB, overflow := n.CosC(c, a, b); overflow {
		return true
	} else if nC, dC, overflow := n.CosC(a, b, c); overflow {
		return true
	} else {
		t.cos = []*Q32 {
			newQ32(nA, dA),
			newQ32(nB, dB),
 			newQ32(nC, dC),
 		}
	}
	if oA, iA, dA, overflow := n.SinC(b, c, a); overflow {
		return true
	} else if oB, iB, dB, overflow := n.SinC(c, a, b); overflow {
		return true
	} else if oC, iC, dC, overflow := n.SinC(a, b, c); overflow {
		return true
	} else {
		t.sin = []*Q32 {
			newQ32Root(oA, iA, dA),
			newQ32Root(oB, iB, dB),
 			newQ32Root(oC, iC, dC),
 		}
	}
	return false
}






type Tris32 struct {
	list []*Tri32
}

func NewA32Tris(max int) *Tris32 {
	ts := &Tris32 {
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

func (ts *Tris32) SetSinCos(factory *N32s) (overflow bool) {
	for _, t := range ts.list {
		if overflow = t.setSinCos(factory); overflow {
			return true
		}
	}
	return false
}

func (ts *Tris32) SinsAdd(tris [][]int) {

}



func (ts *Tris32) add(a, b, c N32) {
	gcd := NatGCD(a, NatGCD(b, c))
	ga, gb, gc := a / gcd, b / gcd, c / gcd
	for _, t := range ts.list {
		if t.sides[0] == ga && t.sides[1] == gb && t.sides[2] == gc {
			// scaled version already stored don't append
			return
		}
	}
	ts.list = append(ts.list, &Tri32{
		sides: []N32{ a, b, c },
	})
}


