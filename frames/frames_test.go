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








func TestFramesTrianglesSurdExt(t *testing.T) {
	FramesTrianglesSurdExt(90, 20)
}