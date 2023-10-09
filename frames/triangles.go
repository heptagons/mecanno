package frames

import (
	. "github.com/heptagons/meccano/nest"
)

type Triangles struct {
	*A32s
}

func NewTriangles() *Triangles {
	return &Triangles{
		A32s: NewA32s(),
	}
}

// SurdExt create virtual strips ED of type surd outside an integer triangle.
// Integral triangle ABC has integral extentionss D from A and E from B:
//
//                              a*a + b*b - c*c 
//       C-_            cosC = ----------------
//      /   -_                      2*a*b
//     /   __ B_        __   __ __   __ __     __ __
//    A__--     -_      ED = CD*CD + CE*CE - 2*CD*CE*cosC
//   /            -_
//  D---------------E
//
// results limits are a+d <= max, b+e <= max and c <= max.
func (t *Triangles) SurdExt(surd Z, max N32) (frames []*FrameA) {
	min := N32(1)
	frames = make([]*FrameA, 0)
	for a := min; a <= max; a++ {
		// a greater-equal b for symmetric redundancy
		for b := min; b <= a; b++ {
			_2ab := N(2)*N(a)*N(b)
			aa_bb := Z(a*a) + Z(b*b)
			for c := min; c <= max; c++ {
				if a + b <= c || b + c <= a || c + a <= b {
					continue // invalid triangle
				}
				cos, err := t.ANew1(_2ab, aa_bb - Z(c*c))
				if err != nil {
					continue // silent
				}
				for d := N32(0); d <= (max-a); d++ {
					a_d := Z(a+d)
					for e := N32(0); e <= (max-b); e++ {
						b_e := Z(b+e)
						if m, err := t.ANew1(N(1), Z(-2)*a_d*b_e); err != nil {
							// silent
						} else if prod, err := t.AMulN(cos, m); err != nil {
							// silent
						} else if f, ok := prod.IsInteger(); !ok {
							// silent
						} else if g := a_d*a_d + b_e*b_e + Z(f); g == surd {
							frames = append(frames, &FrameA{
								a:   a,
								b:   b,
								c:   c,
								d:   d,
								e:   e,
								cos: cos,
							})
						}
					}
				}
			}
		}
	}
	return
}
