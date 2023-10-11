package frames

import (
	. "github.com/heptagons/meccano/nest"
)

type Frames struct {
	*A32s
}

func NewFrames() *Frames {
	return &Frames{
		A32s: NewA32s(),
	}
}

// As create virtual strips ED of type surd outside an integer triangle.
// Integral triangle ABC has integral extentionss D from A and E from B:
//
//                              a*a + b*b - c*c 
//       C-_            cosC = ----------------
//      /   -_                      2*a*b
//     /   __ A_        __   __ __   __ __     __ __
//    B__--     -_      ED = CD*CD + CE*CE - 2*CD*CE*cosC
//   /            -_
//  E---------------D
//
// results limits are a+d <= max, b+e <= max and c <= max.
func (t *Frames) Surds(surd Z, max N32, frame func(a *FrameSurd)) {
	min := N32(1)
	for a := min; a <= max; a++ {
		// a > b for symmetric redundancy
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
							frame(&FrameSurd{
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
}

//            _A
//        a _- |
//        _-   |
//       B     | b
//   a _- -_   |
//   _-   a -_ | 
// C-         -D-------E
// 
// Let strips AB = BC = DB = a
// Let bar AD = b
// Let bar DE = c
// Fix angle DAE to be right
// Then 
// CD = sqrt((2a)^2 - b^2) = sqrt(d) when 2a > b
// and CE = c + sqrt(d)
func (f *Frames) Algs(max N32, frame func(*FrameAlg)) {
	min := N32(1)
	for a := min; a <= max; a++ {
		for b := min; b < 2*a; b++ {
			ab := 4*Z(a)*Z(a) - Z(b)*Z(b)
			if o, i, err := f.ZSqrt(Z(1), ab); err != nil {

			} else if i != 1 {
				frame(&FrameAlg{
					a:a,
					b:b,
					o:o,
					i:i,
				})
			}
		}
	}
}



//    C-_                           C-_ 
//    |  -_                         *  -_
//    |a   -_ b                    /|a   -_
//    |      -_                   / |      -_
//    B---___  -_             *--*--B---___  -_
//    .    c ----A             \   /       ----A
//    .       _.                \ /         _.
// √s .     _.                   x  √s    _.  _____
//    .   _.  nest                \     _.   √x+y√z
//    . _.                         \  _.
//    N                             N
//      
//               a^2 + b^2 - c^2
//       cosC = -----------------
//                   2*a*b
//                                   _
//   x^2 = (a + √n)^2 + b^2 - 2(a + √n)(b)cosC
//                                       _     a^2 + b^2 - c^2
//       = a^2 + 2a√n + n + b^2 - 2(a + √n)(b)-----------------
//                                                  2ab
//                                      _  a^2 + b^2 - c^2
//       = a^2 + 2a√n + n + b^2 - (a + √n)-----------------
//                                                a
//                                                 a^2 + b^2 - c^2
//       = a^2 + n + b^2 - a^2 - b^2 + c^2 + (2a - ---------------)√n
//                                                      a
//                     2a^2 - a^2 - b^2 + c^2  _
//       = n + c^2 + (-----------------------)√n
//                                a       _
//          an + ac^2 + (a^2 - b^2 + c^2)√n
//       = --------------------------------
//                         a
//
func (t *Frames) Nests(max, surd N32, frame func(n *FrameNest)) {
	factory := NewA32s()
	min := N32(1)
	for a := min; a <= max; a++ {
		for b := min; b <= max; b++ {
			for c := min; c <= max; c++ {
				if a + b <= c || b + c <= a || c + a <= b {
					continue
				}
				B := Z(a*surd) + Z(a*c*c)
				C := Z(a*a) - Z(b*b) + Z(c*c)
				D := Z(surd)
				if xx, err := factory.ANew3(N(a), B, C, D); err != nil {

				} else if nest, err := factory.ASqrt(xx); err == nil {
					frame(&FrameNest{
						a:    a,
						b:    b,
						c:    c,
						surd: surd,
						nest: nest,
					})
				}
			}
		}
	}	
}




