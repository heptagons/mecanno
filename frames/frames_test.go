package frames

import (
	"fmt"
	"testing"

	. "github.com/heptagons/meccano/nest"
)

// FramesTrianglesSurdExt creates a virtual strip ED of type surd outside an integer triangle.
// Integral triangle ABC has integral extentionss D from A and E from B:
//                              a*a + b*b - c*c 
//       C-_            cosC = ----------------
//      /   -_                      2*a*b
//     /   __ A_        __   __ __   __ __     __ __
//    B__--     -_      ED = CD*CD + CE*CE - 2*CD*CE*cosC
//   /            -_
//  E---------------D
//
func FramesTrianglesSurdExt(surd Z, max N32) {
	factory := NewA32s()
	min := N32(1)
	for a := min; a <= max; a++ {
		for b := min; b <= max; b++ {
			_2ab := N(2)*N(a)*N(b)
			aa_bb := Z(a*a) + Z(b*b)
			for c := min; c <= max; c++ {

				if a + b <= c || b + c <= a || c + a <= b {
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
									if e == 0 {
										fmt.Printf("a=%d b=%d", a, b)
									} else {
										// triangle extension with only 4 bolts not 5
										fmt.Printf("a=%d b=%d+%d", a, b,e)
									}
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

func TestFramesTrianglesSurdExt(t *testing.T) {
	FramesTrianglesSurdExt(25*7, 10)
}



//       A
//      /|
//     / |
//    B  |
//   / \ |
//  /   \| 
// C     D------E
// 
// Let strips AB = BC = DB = a
// Let bar AD = b
// Let bar DE = c
// Fix angle DAE to be right
// Then 
// CD = sqrt((2a)^2 - b^2) = sqrt(d) when 2a > b
// and CE = c + sqrt(d)
func TestFrameY(t *testing.T) {
	factory := NewA32s()
	min := N32(1)
	max := N32(10)
	for a := min; a <= max; a++ {
		for b := min; b < 2*a; b++ {
			ab := 4*Z(a)*Z(a) - Z(b)*Z(b)
			if o, i, err := factory.ZSqrt(Z(1), ab); err != nil {

			} else if i != 1 {
				fmt.Printf("a=% 3d b=% 3d c=%d√%d\n", a, b, o, i)
			}
		}
	}
}

//    C-_                    a^2 + b^2 - c^2
//    |  -_           cosC = ----------------
//  a |    -_ b                   2*a*b
//    |      -_                     _
//    B---___  -_      x^2 = ( a + √n )^2 + b^2 - 2*(a + √n)*b*cosC
//    |    c ---_A                      /      a^2 + b^2 - c^2 \  _
//    |       _/           = n + c^2 + ( 2*a - ---------------  )√n
// √n |     _/                          \            a         /
//    |   _/  x                         a^2 - b^2 + c^2 _
//    | _/                 = n + c^2 + ----------------√n
//    |/                                       a
//    N
//
//
func TestFrameX(t *testing.T) {
	min := 1
	max := 500
	n := 100*5 // 10√5
	for a := min; a <= max; a++ {
		for b := min; b <= max; b++ {
			for c := min; c <= max; c++ {
				if a + b <= c || b + c <= a || c + a <= b {
					continue
				}
				if d := a*a - b*b + c*c; d > 0 && d % a == 0 {
					x1 := n + c*c
					x2 := d/a
					if x1 % x2 == 0 && x1 / x2 == 25 {
					fmt.Printf("a=% 3d b=% 3d c=% 3d x^2= %d%+d√%d\n", a, b, c, x1, x2, n)
						fmt.Println("XXX")
					}
				}
			}
		}
	}	
}