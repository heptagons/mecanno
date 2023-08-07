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
type Tri32 struct { // Triangle
	abc []N32
	cos []*Q32
	sin []*Q32
}

func (t *Tri32) String() string {
	return fmt.Sprintf("abc:%v cos:%v sin:%v", t.abc, t.cos, t.sin)
}



type Tris32 struct {
	*Q32s

	list []*Tri32
}

func NewA32Tris(max int, factory *N32s) *Tris32 {
	ts := &Tris32 {
		Q32s: &Q32s{
			N32s: factory,
		},
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

func (ts *Tris32) add(a, b, c N32) {
	gcd := NatGCD(a, NatGCD(b, c))
	ga, gb, gc := a / gcd, b / gcd, c / gcd
	for _, t := range ts.list {
		if t.abc[0] == ga && t.abc[1] == gb && t.abc[2] == gc {
			// scaled version already stored don't append
			return
		}
	}
	ts.list = append(ts.list, &Tri32{
		abc: []N32{ a, b, c },
	})
}

func (ts *Tris32) setSinCos() error {
	for _, t := range ts.list {
		a, b, c := t.abc[0], t.abc[1], t.abc[2]
		if cosA, err := ts.cosC(b, c, a); err != nil {
			return err
		} else if cosB, err := ts.cosC(c, a, b); err != nil {
			return err
		} else if cosC, err := ts.cosC(a, b, c); err != nil {
			return err
		} else {
			t.cos = []*Q32 { cosA, cosB, cosC }
		}
		if sinA, err := ts.sinC(b, c, a); err != nil {
			return err
		} else if sinB, err := ts.sinC(c, a, b); err != nil {
			return err
		} else if sinC, err := ts.sinC(a, b, c); err != nil {
			return err
		} else {
			t.sin = []*Q32 { sinA, sinB, sinC }
		}
	}
	return nil
}

// cosC returns the rational cosine of the angle C using the law of cosines:
//	       a² + b² - c²
//	cosC = ------------
//	           2ab
func (ts *Tris32) cosC(a, b, c N32) (*Q32, error) {
	den64 := 2*N(a)*N(b)
	num64 := Z(a)*Z(a) + Z(b)*Z(b) - Z(c)*Z(c)
	return ts.newQ32(den64, num64)
}

// sinC return the algebraic sine of the angle C using the law of sines:
//	       math.Sqrt(4a²b² - (a²+b²-c²)²)
//	sinC = ------------------------------
//	                  2ab 
func (ts *Tris32) sinC(a, b, c N32) (*Q32, error) {
	//	stable := N(a+(b+c)) * N(c-(a-b)) * N(c+(a-b)) * N(a+(b-c))
	//	if out, in, ok := t.algs.roiN(1, stable); ok {
	//		next.area = NewAlg(NewRat(int(out), 4), in)
	//	}
	ab := 2*N(a)*N(b)
	abc := Z(a)*Z(a) + Z(b)*Z(b) - Z(c)*Z(c)
	return ts.newQ32(ab, 0, 1, Z(ab)*Z(ab) - abc*abc)
}

func (ts *Tris32) sinsAdd(triAngs [][]uint) error {
	// sin(A+B) = sinAcosB + cosAsinB
	if len(triAngs) < 2 {
		return fmt.Errorf("Need at least two triangles")
	}
	if sinA, cosA, err := ts.sinCos(triAngs[0]); err != nil {
		return err
	} else if sinB, cosB, err := ts.sinCos(triAngs[1]); err != nil {
		return err
	} else if sinAcosB, err := ts.MulQ(sinA, cosB); err != nil {
		return err
	} else if sinBcosA, err := ts.MulQ(sinB, cosA); err != nil {
		return err
	} else {
		fmt.Println(sinA, cosA, sinB, cosB)
		fmt.Println(sinAcosB.String() + " + " + sinBcosA.String())
		return nil
	}
}

func (ts *Tris32) sinCos(posAngle []uint) (*Q32, *Q32, error) {
	if len(posAngle) != 2 {
		return nil, nil, fmt.Errorf("Format error: needs two uints for triangle pos and angle")
	} else if pos := int(posAngle[0]); pos > len(ts.list) {
		return nil, nil, fmt.Errorf("No triangle for pos:%d", pos)
	} else if angle := int(posAngle[1]); angle > 2 {
		return nil, nil, fmt.Errorf("Invalid angle:%d", angle)
	} else {
		t := ts.list[pos]
		return t.sin[angle], t.cos[angle], nil
	}
}
