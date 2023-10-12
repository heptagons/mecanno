package frames

import (
	"fmt"
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


// Surds returns FrameSurds with a+d, b+d and c <= max and Frame ED distance equals √surd.
// FrameSurd consist of ABC with extentions (lenght 0 to max) D from A and E from B:
//
//                                     a^2 + b^2 - c^2
//        C-_  b               cosC = -----------------
//    a  /   -_                            2*a*b
//      /   __ A_                __            __
//     B__--     -_              CE = a + d,   CD = b + e
// d  /     c       -_ e
//   /                -_         __   
//  E                   D        ED = (a+d)^2 + (b+e)^2 - 2(a+d)(b+e)cosC
//
func (t *Frames) SurdsInt(surd Z, max N32, frame func(a *FrameSurd)) {
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
							// reject rational surds example:
							// a=2 b=1 c=2 d=1 e=0 surd= √34/2
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

func (t *Frames) SurdsRat(max N32, frame func(n []N32, s *A32)) {
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
						if e == 0 && d == 0 {
							continue
						}
						b_e := Z(b+e)
						if m, err := t.ANew1(N(1), Z(-2)*a_d*b_e); err != nil {
							// silent
						} else if prod, err := t.AMulN(cos, m); err != nil {
							// silent
						} else if s2, err := prod.AddInt(a_d*a_d + b_e*b_e); err != nil {
							// silent
						} else if s, err := t.ASqrt(s2); err != nil {

						} else if s.IsRational() {
							frame([]N32{ a, b, c, d, e }, s)
						}
					}
				}
			}
		}
	}
}


func (f *Frames) AlgsNotRight(surd Z, max N32) {
	fmt.Printf("surd=%d, max=%d\n", surd, max)
	f.SurdsInt(surd, max, func(fs *FrameSurd) {
		fmt.Println(fs)
		for d := fs.a; d <= max; d++ {
			for e := N32(1); e <= max; e++ {
				for f := N32(1); f <= max; f++ {
					if surd != Z(d*d) + Z(e*e) - Z(f*f) {
						continue
					}
					//fmt.Println(">>>")
					if (Z(fs.a*fs.a) + Z(fs.b*fs.b) + surd)*Z(d) == 2*Z(fs.a)*surd {
						fmt.Printf("\td=%d e=%d f=%d\n", d, e, f)
					}
				}
			}
		}
	})
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




