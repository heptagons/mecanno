package frames

import (
	"fmt"
	"os"
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


// Triangles returns Triangles with a+d, b+d and c <= max and Frame ED distance equals √surd.
// Triangle consist of ABC with extentions (lenght 0 to max) D from A and E from B:
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
func (t *Frames) Triangles(surd Z, max N32, frame func(a *Triangle)) {
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
							// reject here rational surds. Example:
							// a=2 b=1 c=2 d=1 e=0 surd= √34/2
							// silent
						} else if g := a_d*a_d + b_e*b_e + Z(f); g == surd {
							frame(&Triangle{
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

func (f *Frames) AlgsNoPythagoras(surd Z, max N32) {
	fmt.Printf("surd=%d, max=%d\n", surd, max)
	f.Triangles(surd, max, func(t *Triangle) {
		if d, e := t.RightAngles(); d || e {
			return
		}
		t.WriteString(os.Stdout)
		fmt.Println()
		ad := Z(t.a + t.d)
		be := Z(t.b + t.e)
		for g := t.a + t.d; g <= max; g++ {
			for h := N32(1); h <= max; h++ {
				for i := N32(1); i <= max; i++ {
					if surd != Z(g*g) + Z(h*h) - Z(i*i) {
						continue
					}
					if (ad*ad - be*be + surd)*Z(g) == 2*ad*surd {
						fmt.Printf("\tg=%d h=%d i=%d\n", g, h, i)
					}
				}
			}
		}
	})
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

//                                                                                      _________     ___________
//                                   a^2 + b^2 - c^2    m            ____________      /     m^2     √(n^2 - m^2)
//        B                   cosX = --------------- = ---   sinX = √ 1 - cos^2X   =  / 1 - ----- = -------------
//       / \_                              2ab          n                            √       n^2          n
//    a /    \_ c                                                                       _________     ___________
//     /       \_                    d^2 + e^2 - f^2    o            ____________      /     o^2     √(p^2 - o^2)
//    / X   b    \            cosY = --------------- = ---   sinY = √ 1 - cos^2Y   =  / 1 - ----- = -------------
//   C------------A-------D                2de          p                            √       p^2          p
//    \_ Y     e       _-'
//     '\_           _-'            __
//       '\_       _-' f      g^2 = BE = a^2 + d^2 - 2adcos(X+Y)
//      d   '\_  _-'
//            'E-'            
//
// cos(X+Y) = cosXcosY - sinXsinY
//                        ___________    ___________
//             m   o     √(n^2 - m^2)   √(p^2 - o^2)
//          = ---x--- - ------------- x ------------
//             n   p           n              p
//                   _______________________
//             mo - √(n^2 - m^2)(p^2 - o^2))
//          =  -----------------------------
//                           np
//                               ______________________
//                         mo - √(n^2 - m^2)(p^2 - o^2)
// g^2 = a^2 + d^2 - (2ad)-----------------------------
//                                   4abde
//                               ______________________
//       2a^2be + 2bd^2e - mo + √(n^2 - m^2)(p^2 - o^2)
// g^2 = ------------------------------------------------
//                        2be
// 
func (fr *Frames) TrianglePairsOld(max N32, fgh []int) {
	min := N32(1)
	for a := min; a <= max; a++ {
		fmt.Printf("a=%d...\n", a)
		for b := min; b <= max; b++ {
			for c := min; c <= max; c++ {
				if a + b <= c || b + c <= a || c + a <= b {
					continue // invalid triangle
				}
				m, n := Z(a*a) + Z(b*b) - Z(c*c), Z(2*a*b)
				for d := min; d <= max; d++ {
					for e := a; e <= max; e++ {
						for f := min; f <= max; f++ {
							if d + e <= f || e + f <= d || f + d <= e {
								continue // invalid triangle
							}
							o, p := Z(d*d) + Z(e*e) - Z(f*f), Z(2*d*e)
							//	(B + C√D) / A
							A := N(2*b*e)
							B := Z(2*a*a*b*e) + Z(2*b*d*d*e) - m*o
							C := Z(1)
							D := (n-m)*(n+m)*(p-o)*(p+o)//(n*n - m*m) * (p*p - o*o)
							if gg, err := fr.ANew3(A, B, C, D); err != nil {
								// silent error
							} else if g, err := fr.ASqrt(gg); err != nil {  // √(F + G√H)
								// silent error
							} else if F, ok := g.Num(4); !ok || F != Z32(fgh[0]) {
								// doesn't match f
							} else if G, ok := g.Num(5); !ok || G != Z32(fgh[1]) {
								//fmt.Println("g error", fgh[1], g)
								// doesn't match g
							} else if H, ok := g.Num(6); !ok || H != Z32(fgh[2]) {
								// doesn't match h
							} else {
								fmt.Printf("a=%3d b=%3d c=%3d | d=%3d e=%3d f=%3d | g=%v\n", a, b, c, d, e, f, g)
							}
						}
					}
				}
			}
		}
	}
}

// TrianglePairsTex uses factory.ANew7 algebraic number to be simplified.
// Prints Tex rows to be pasted in latex documents.
func (fr *Frames) TrianglePairsTex(max N32, fgh []int) {
	fmt.Println("%%----------- start")
	fmt.Println("\\begin{align*}")
	fmt.Println("Folder &: \\texttt{github.com/heptagons/meccano/frames}\\\\")
	fmt.Printf("Call &: \\texttt{NewFrames().TrianglePairsTex(%d, %v)}", max, fgh)
	fmt.Println("\\end{align*}")
	fmt.Println("\\begin{align*}")
	fmt.Printf("(a,b,c) \\oplus (d,e,f) &\\mapsto g\\\\\n")
	fmt.Println("\\hline")
	B := Z(0)
	C := Z(0)
	D := Z(1)
	E := Z(1)
	min := N32(1)
	for a := min; a <= max; a++ {
		fmt.Printf("%%a=%d...\\\\ \n", a) // just to notify console user for large a's
		for b := min; b <= max; b++ {
			for c := min; c <= max; c++ {
				if a + b <= c || b + c <= a || c + a <= b {
					continue // invalid triangle
				}
				m, n := Z(a*a) + Z(b*b) - Z(c*c), Z(2*a*b)
				nn_mm := (n-m)*(n+m)
				for d := min; d <= max; d++ {
					for e := b; e <= max; e++ { // e should be at least equal to b to reject duplications
						for f := min; f <= max; f++ {
							if d + e <= f || e + f <= d || f + d <= e {
								continue // invalid triangle
							}
							o, p := Z(d*d) + Z(e*e) - Z(f*f), Z(2*d*e)
							pp_oo := (p-o)*(p+o)
							//	(B + C√D + E√(F + G√H)) / A
							A := Z(2*b*e)
							F := A*A*Z(a*a + d*d) - A*m*o
							G := A
							H := nn_mm * pp_oo
							if g, err := fr.ANew7(N(A), B, C, D, E, F, G, H); err != nil {
								// silent error
							} else if fgh == nil { // no filter
								fmt.Printf("%d,%d,%d | %d,%d,%d | %v\n", a, b, c, d, e, f, g)
							} else if F, ok := g.Num(4); !ok || F != Z32(fgh[0]) {
								// f doesn't match
							} else if G, ok := g.Num(5); !ok || G != Z32(fgh[1]) {
								// g doesn't match
							} else if H, ok := g.Num(6); !ok || H != Z32(fgh[2]) {
								// h doesn't match
							} else {
								fmt.Printf("(%d,%d,%d) \\oplus (%d,%d,%d) &\\mapsto %v \\\\\n", a, b, c, d, e, f, g.Tex())
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("\\end{align*}")
	fmt.Println("%%----------- end")
}




