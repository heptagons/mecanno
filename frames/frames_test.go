package frames

import (
	"fmt"
	"os"
	"testing"

	. "github.com/heptagons/meccano/nest"
)

func TestTrianglesSurdExt(t *testing.T) {
	tri := NewTriangles()
	surd := Z(7)
	max := N32(10)
	frames := tri.SurdExt(surd, max)
	fmt.Printf("√%d max=%d qty=%d:\n", surd, max, len(frames))
	for f, frame := range frames {
		fmt.Fprintf(os.Stdout, "% 3d) ", f+1)
		frame.WriteString(os.Stdout)
		fmt.Println()
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
func TestFrame4AABB(t *testing.T) {
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


func TestFrame4AAB(t *testing.T) {
	factory := NewA32s()
	min := N32(1)
	max := N32(5)
	for a := min; a <= max; a++ {
		do := true
		for b := min; do; b++ {
			ab := 4*Z(a)*Z(a) - Z(b)
			if ab <= 0 {
				do = false
			} else if o, i, err := factory.ZSqrt(Z(1), ab); err != nil {

			} else if i != 1 {
				fmt.Printf("a=% 3d b=√%d c=%d√%d\n", a, b, o, i)
			}
		}
	}
}









//    C-_                    a^2 + b^2 - c^2
//    |  -_           cosC = ----------------
//  a |    -_ b                   2*a*b
//    |      -_
//    B---___  -_
//    |    c ---_A
//    |       _/
// √n |     _/
//    |   _/  x
//    | _/
//    |/
//    N  
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
func TestFrameX(t *testing.T) {

	var f, g, h, n Z32

	f, g, h, n = 5, 2, 5, 2*2*5  // √(5 + 2√5)            pentagon height (decagon r) -> a=9,b=5,c=5, d=2√5
	//f, g, h, n = 25, 10, 5, 5*5*5  // √(25 + 10√5),  5√5  pentagon radius -> a=18,b=10,c=10, d=5√5
	//f, g, h, n = 50, 10, 5, 20*20*5  // √(50 + 10√5), 3√5  pentagon Radius -> none

	factory := NewA32s()
	min := Z32(1)
	max := Z32(100)
	for a := min; a <= max; a++ {
		for b := min; b <= max; b++ {
			for c := min; c <= max; c++ {
				if a + b <= c || b + c <= a || c + a <= b {
					continue
				}
				B := Z(a*n) + Z(a*c*c)
				C := Z(a*a) - Z(b*b) + Z(c*c)
				D := Z(n)
				if xx, err := factory.ANew3(N(a), B, C, D); err != nil {

				} else if x, err := factory.ASqrt(xx); err == nil {
					if H, ok := x.Num(6); ok && H == h { // √(F + G√5)
						F, _ := x.Num(4)
						G, _ := x.Num(5)
						if F == f && G == g {
							fmt.Printf("a=% 3d b=% 3d c=% 3d d=√%d x= %v\n", a, b, c, n, x.String())
							return
						}
					}
				}
			}
		}
		fmt.Println(a)
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
func TestTri2(t *testing.T) {
	factory := NewA32s()
	min := N32(1)
	max := N32(15)
	for a := min; a <= max; a++ {
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
							if gg, err := factory.ANew3(A, B, C, D); err != nil {

							} else if g, err := factory.ASqrt(gg); err == nil {

								if H, ok := g.Num(6); ok && H == 5 { // √(F + G√5)
									F, _ := g.Num(4)
									G, _ := g.Num(5)
									if F == 5 && G == 2 {
										fmt.Printf("a=% 3d b=% 3d c=% 3d d=% 3d e=% 3d f= %3d  g=%v\n", a, b, c, d, e, f, g)
									}
								}
							}
						}
					}
				}
			}
		}
		fmt.Println(a)
	}
}
