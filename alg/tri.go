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
// A,B, and C the angles to opposite sides a,b and c.
type Tri32 struct { // Triangle
	sides []N32
	cos   []*Q32
	sin   []*Q32
}

func (t *Tri32) String() string {
	return fmt.Sprintf("abc:%v cos:%v sin:%v", t.sides, t.cos, t.sin)
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
		if t.sides[0] == ga && t.sides[1] == gb && t.sides[2] == gc {
			// scaled version already stored don't append
			return
		}
	}
	ts.list = append(ts.list, &Tri32{
		sides: []N32{ a, b, c },
	})
}

func (ts *Tris32) setSinCos() (overflow bool) {
	for _, t := range ts.list {
		a, b, c := t.sides[0], t.sides[1], t.sides[2]
		if nA, dA, overflow := ts.cosC(b, c, a); overflow {
			return true
		} else if nB, dB, overflow := ts.cosC(c, a, b); overflow {
			return true
		} else if nC, dC, overflow := ts.cosC(a, b, c); overflow {
			return true
		} else {
			t.cos = []*Q32 {
				newQ32(nA, dA),
				newQ32(nB, dB),
	 			newQ32(nC, dC),
	 		}
		}
		if oA, iA, dA, overflow := ts.sinC(b, c, a); overflow {
			return true
		} else if oB, iB, dB, overflow := ts.sinC(c, a, b); overflow {
			return true
		} else if oC, iC, dC, overflow := ts.sinC(a, b, c); overflow {
			return true
		} else {
			t.sin = []*Q32 {
				newQ32Root(oA, iA, dA),
				newQ32Root(oB, iB, dB),
	 			newQ32Root(oC, iC, dC),
	 		}
		}
	}
	return false
}

func (ts *Tris32) sinsAdd(triAngs [][]uint) error {
	// sin(A+B) = sinAcosB + cosAsinB
	if len(triAngs) < 2 {
		return fmt.Errorf("Need at least two triangles")
	}
	if sinA, cosA, err := ts.triSinCos(triAngs[0]); err != nil {
		return err
	} else if sinB, cosB, err := ts.triSinCos(triAngs[1]); err != nil {
		return err
	} else if sinAcosB, overflow := ts.reduceQMul(sinA, cosB); overflow {
		return fmt.Errorf("sinAcosB overflow")
	} else if sinBcosA, overflow := ts.reduceQMul(sinB, cosA); overflow {
		return fmt.Errorf("sinBcosA overflow")
	} else {
		fmt.Println(sinA, cosA, sinB, cosB)
		fmt.Println(sinAcosB + " + " + sinBcosA)
		return nil
	}
}


func (ts *Tris32) triSinCos(posAngle []uint) (*Q32, *Q32, error) {
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


// cosC returns the rational cosine of the angle C using the law of cosines:
//	       a² + b² - c²
//	cosC = ------------
//	           2ab
func (ts *Tris32) cosC(a, b, c N32) (num Z32, den N32, overflow bool) {
	num64 := Z(a)*Z(a) + Z(b)*Z(b) - Z(c)*Z(c)
	den64 := 2*N(a)*N(b)
	if den32, nums32, overflow := ts.reduceQ(den64, num64); overflow {
		return 0, 0, true
	} else {
		return nums32[0], den32, false
	}
}

// sinC return the algebraic sine of the angle C using the law of sines:
//	       math.Sqrt(4a²b² - (a²+b²-c²)²)
//	sinC = ------------------------------
//	                  2ab 
func (ts *Tris32) sinC(a, b, c N32) (out, in Z32, den N32, overflow bool) {
	//	stable := N(a+(b+c)) * N(c-(a-b)) * N(c+(a-b)) * N(a+(b-c))
	//	if out, in, ok := t.algs.roiN(1, stable); ok {
	//		next.area = NewAlg(NewRat(int(out), 4), in)
	//	}
	ab := 2*N(a)*N(b)
	aa := Z(a)*Z(a)
	bb := Z(b)*Z(b)
	cc := Z(c)*Z(c)
	i2 := aa + bb - cc
	out, in, overflow = ts.reduceRoot(1, Z(ab)*Z(ab) - i2*i2)
	if overflow {
		return
	}
	var nums32 []Z32
	den, nums32, overflow = ts.reduceQ(ab, Z(out))
	if overflow {
		return
	}
	out = nums32[0]
	return
}






