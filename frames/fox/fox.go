package fox

import (
	"fmt"
	. "github.com/heptagons/meccano/nest"
)

type ABCD struct {
	a [][]N32
}

func (e *ABCD) append(a, b, c, d N32) {
	if e.a == nil {
		e.a = make([][]N32, 0)
	}
	e.a = append(e.a, []N32{ a,b,c,d })
}

func (abcd *ABCD) print() {
	for _, a := range abcd.a {
		fmt.Printf("\t%v\n", a)
	}
}

// Fox iterate a,b,c,d until max an reports a,b,c,d and
//	                    ,---------------------
//	       -a(2b+c) +- √ a²c² + 4b(b+c)(d²-c²)
//	cos = ------------------------------------
//	                    4b(b+c)
func Fox(max N32, found func(a, b, c, d N32, cos *A32)) {
	factory := NewA32s()
	n1 := N32(1)
	for a := n1; a <= max; a++ {
		for b := n1; b <= max; b++ {
			ab := NatGCD(a, b)
			for c := n1; c <= max; c++ {
				abc := NatGCD(ab, c)
				na := N32(4)*b*(b+c)        // 4b(b+c)
				zb := -Z(a)*(2*Z(b) + Z(c)) // -a(2b+c)
				zc := Z(1)                  // 1
				a2c2 := Z(a*a)*Z(c*c)       // a
				for d := c; d <= max; d++ { // d >= c always
					if g := NatGCD(abc, d); g > 1 {
						continue // skip scale repetitions, eg. [1,2,3,4] = [2,4,6,8]
					}
					if zd := a2c2 + 4*Z(b)*Z(b+c)*(Z(d*d) - Z(c*c)); zd < 0 {
						// skip imaginary numbers invalid fox-face, like d too short
					} else if cos, err := factory.ANew3(N(na), zb, zc, zd); err != nil {
						// silent overflow
					} else {
						found(a, b, c, d, cos)
					}
				}
			}
		}
	}
}

// Fox iterate a,b,c,√d until max an reports a,b,c,√d and
//	                    ,---------------------
//	       -a(2b+c) +- √ a²c² + 4b(b+c)(d-c²)
//	cos = ------------------------------------
//	                    4b(b+c)
func FoxSurd(max N32, found func(a, b, c, d N32, cos *A32)) {
	factory := NewA32s()
	n1 := N32(1)
	for a := n1; a <= max; a++ {
		for b := n1; b <= max; b++ {
			ab := NatGCD(a, b)
			for c := n1; c <= max; c++ {
				abc := NatGCD(ab, c)
				na := N32(4)*b*(b+c)        // 4b(b+c)
				zb := -Z(a)*(2*Z(b) + Z(c)) // -a(2b+c)
				zc := Z(1)                  // 1
				a2c2 := Z(a*a)*Z(c*c)       // a
				for surdD := c*c; surdD <= max*max; surdD++ { // surdD >= c*c always

					if o, _, err := factory.ZSqrt(1, Z(surdD)); err != nil {
						// 
					} else if g := NatGCD(abc, N32(o)); g > 1 {
						// skip scale repetitions, eg. [1,2,3,4] = [2,4,6,8]
					} else if zd := a2c2 + 4*Z(b)*Z(b+c)*(Z(surdD) - Z(c*c)); zd < 0 {
						// skip imaginary numbers invalid fox-face, like d too short
					} else if cos, err := factory.ANew3(N(na), zb, zc, zd); err != nil {
						// silent overflow
					} else {
						found(a, b, c, surdD, cos)
					}
				}
			}
		}
	}
}

// FoxTrianglesSurdExt creates a virtual strip ED of type surd outside an integer triangle.
// Integral triangle ABC has integral extentionss D from A and E from B:
//                              a*a + b*b - c*c 
//       C-_            cosC = ----------------
//      /   -_                      2*a*b
//     /   __ A_        __   __ __   __ __     __ __
//    B__--     -_      ED = CD*CD + CE*CE - 2*CD*CE*cosC
//   /            -_
//  E---------------D
//
func FoxTrianglesSurdExt(surd Z, max N32) {
	factory := NewA32s()
	min := N32(1)
	for a := min; a <= max; a++ {
		for b := min; b <= max; b++ {
			_2ab := N(2)*N(a)*N(b)
			aa_bb := Z(a*a) + Z(b*b)
			for c := min; c <= max; c++ {

				if a + b < c || b + c < a || c + a < b {
					continue
				}
				if cos, err := factory.ANew1(_2ab, aa_bb - Z(c*c)); err != nil {
					// silent
				} else {
					for d := N32(0); d <= max; d++ {
						for e := N32(0); e <= max; e++ {
							if m, err := factory.ANew1(N(1), Z(-2)*Z(a+d)*Z(b+e)); err != nil {
								// silent
							} else if prod, err := factory.AMulN(cos, m); err != nil {
								// silent
							} else if f, ok := prod.IsInteger(); !ok {
								// silent
							} else if g := Z(a+d)*Z(a+d) + Z(b+e)*Z(b+e) + Z(f); g == surd {
									fmt.Printf("√%d ", g)
								if d == 0 {
									// triangle extension with only 4 bolts not 5
									fmt.Printf("a=%d b=%d+%d", a, b,e)
								} else if e == 0 {
									fmt.Printf("a+d=%d+%d b=%d", a,d, b)
								} else {
									fmt.Printf("a+d=%d+%d b=%d+%d", a,d, b,e)
								}
								fmt.Printf(" c=%d cos=%s\n", c, cos)
							}
						}
					}
				}
			}
		}
	}
}






