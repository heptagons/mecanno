package frames

import (
	"fmt"
	"os"
	"testing"

	. "github.com/heptagons/meccano/nest"
)

// ~/github.com/heptagons/meccano$ go test ./frames/. -run TestFramesTriangleSurds -v -count 1

func TestFramesTriangleSurds(t *testing.T) {
	surd := Z(13)
	max := N32(10)
	n := 0
	fmt.Printf("NewFrames().TriangleSurds surd=%d max=%d\n", surd, max)
	NewFrames().TriangleSurds(surd, max, func(frame *Triangle) {
		n++
		fmt.Fprintf(os.Stdout, "% 3d) ", n)
		frame.WriteString(os.Stdout)
		fmt.Println()
	})
}

func TestFramesAlgsNotRight(t *testing.T) {
	surd := Z(11)
	max := N32(30)
	NewFrames().AlgsNotRight(surd, max)
}

/*
√7 max=10
  1) a=1 b+e=1+2 c=1 cos=1/2
  2) a+d=1+1 b+e=1+2 c=1 cos=1/2
  3) a+d=1+2 b=1 c=1 cos=1/2
  4) a+d=1+2 b+e=1+1 c=1 cos=1/2
  5) a=2 b+e=2+1 c=2 cos=1/2
  6) a+d=2+1 b=2 c=2 cos=1/2
  7) a=3 b+e=2+2 c=2 cos=3/4 E=pi/2
  8) a+d=3+1 b+e=2+1 c=2 cos=3/4 D=pi/2
  9) a+d=4+2 b+e=4+4 c=1 cos=31/32
 10) a+d=4+4 b+e=4+2 c=1 cos=31/32
 11) a=7 b+e=5+1 c=3 cos=13/14
 12) a=7 b+e=5+2 c=3 cos=13/14
*/


func TestFramesSurdsRat(t *testing.T) {
	max := N32(10)
	n := 0
	NewFrames().SurdsRat(max, func(d []N32, surd *A32) {
		n++
		fmt.Printf("% 3d) %v %v\n", n, d, surd)
	})
}
/* solutions: [a b c d e] surd
  1) [3 1 3 1 0] √141/3
  2) [3 1 3 1 1] 2√39/3
  3) [3 1 3 1 3] 4√15/3
  4) [3 1 3 1 4] √309/3
  5) [3 1 3 2 0] √219/3
  6) [3 1 3 2 1] √231/3
  7) [3 1 3 2 3] √309/3
  8) [3 1 3 2 4] 5√15/3
  9) [3 2 3 1 0] 2√33/3
 10) [3 2 3 1 2] 8√3/3
 11) [3 2 3 1 3] √249/3
 12) [3 2 3 2 0] √201/3
 13) [3 2 3 2 2] √249/3
 14) [3 2 3 2 3] 10√3/3
 15) [3 3 1 0 1] √21/3
 16) [3 3 1 0 2] √51/3
 17) [3 3 1 1 0] √21/3
 18) [3 3 1 1 1] 4/3
 ...
 57) [4 4 5 1 1] 25/4
 58) [5 3 3 0 1] √69/3
 59) [5 3 3 0 2] 5√3/3
*/





func TestFramesAlgs(t *testing.T) {
	NewFrames().Algs(10, func(frame *FrameAlg) {
		if frame.i == 3 { // x√5
			fmt.Println(frame)
		}
	})
}

func testFramesNests(max, n N32, f, g, h Z32) {
	surd := n*n*N32(h)
	NewFrames().Nests(max, surd, func(frame *FrameNest) {
		nest := frame.nest
		if H, ok := nest.Num(6); ok && H == h { // √(F + G√5)
			F, _ := nest.Num(4)
			G, _ := nest.Num(5)
			if F == f && G == g {
				fmt.Printf("a+s=%d+%d√%d, b=%d, c=%d : nest=%v\n", 
					frame.a, n, h, frame.b, frame.c, frame.nest.String())
			}
		}
	})
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




func TestFramesNestPentagonsHeight(t *testing.T) {
	// pentagon height (decagon r)
	for n := N32(1); n < N32(50); n++ {
		fmt.Printf("surd=%d√5:\n", n)
		testFramesNests(50, n, 5, 2, 5) // √(5 + 2√5)
	}
}
/* Solutions
a+surd=9+2√5, b=5, c=5 : nest=3√(5+2√5)
a+surd=18+4√5, b=10, c=10 : nest=6√(5+2√5)
a+surd=11+6√5, b=7, c=15 : nest=9√(5+2√5)
a+surd=16+6√5, b=7, c=15 : nest=9√(5+2√5)
a+surd=27+6√5, b=15, c=15 : nest=9√(5+2√5)
a+surd=36+8√5, b=20, c=20 : nest=12√(5+2√5)
a+surd=21+10√5, b=11, c=25 : nest=15√(5+2√5)
a+surd=24+10√5, b=11, c=25 : nest=15√(5+2√5)
a+surd=45+10√5, b=25, c=25 : nest=15√(5+2√5)
a+surd=22+12√5, b=14, c=30 : nest=18√(5+2√5)
a+surd=32+12√5, b=14, c=30 : nest=18√(5+2√5)
a+surd=24+14√5, b=17, c=35 : nest=21√(5+2√5)
a+surd=39+14√5, b=17, c=35 : nest=21√(5+2√5)
a+surd=21+16√5, b=23, c=40 : nest=24√(5+2√5)
a+surd=25+18√5, b=25, c=45 : nest=27√(5+2√5)
a+surd=33+18√5, b=21, c=45 : nest=27√(5+2√5)
a+surd=48+18√5, b=21, c=45 : nest=27√(5+2√5)
a+surd=42+20√5, b=22, c=50 : nest=30√(5+2√5)
a+surd=48+20√5, b=22, c=50 : nest=30√(5+2√5)
*/

func TestFramesNestPentagonsInradius(t *testing.T) {
	for n := N32(1); n < N32(50); n++ {
		fmt.Printf("surd=%d√5:\n", n)
		testFramesNests(50, n, 25, 10, 5) // √(25 + 10√5)
	}
}
/* Solutions:
a+surd=18+5√5, b=10, c=10 : nest=3√(25+10√5)
a+surd=15+8√5, b=7, c=16 : nest=24√(25+10√5)/5
a+surd=10+9√5, b=10, c=18 : nest=27√(25+10√5)/5
a+surd=36+10√5, b=20, c=20 : nest=6√(25+10√5)
a+surd=30+11√5, b=14, c=22 : nest=33√(25+10√5)/5
a+surd=25+12√5, b=11, c=24 : nest=36√(25+10√5)/5
a+surd=35+12√5, b=17, c=24 : nest=36√(25+10√5)/5
a+surd=22+15√5, b=14, c=30 : nest=9√(25+10√5)
a+surd=32+15√5, b=14, c=30 : nest=9√(25+10√5)
a+surd=30+16√5, b=14, c=32 : nest=48√(25+10√5)/5
a+surd=20+18√5, b=20, c=36 : nest=54√(25+10√5)/5
a+surd=21+20√5, b=23, c=40 : nest=12√(25+10√5)
a+surd=50+21√5, b=22, c=42 : nest=63√(25+10√5)/5
a+surd=45+24√5, b=21, c=48 : nest=72√(25+10√5)/5
a+surd=50+24√5, b=22, c=48 : nest=72√(25+10√5)/5
a+surd=42+25√5, b=22, c=50 : nest=15√(25+10√5)
a+surd=48+25√5, b=22, c=50 : nest=15√(25+10√5)
*/

func TestFramesNestPentagonsOutradius(t *testing.T) {
	for n := N32(1); n < N32(50); n++ {
		fmt.Printf("surd=%d√5:\n", n)
		testFramesNests(50, n, 50, 10, 5) // √(50 + 10√5)
	}
}
// No solutions

func TestFramesNestSin2Pi_5(t *testing.T) {
	for n := N32(1); n < N32(50); n++ {
		fmt.Printf("surd=%d√5:\n", n)
		testFramesNests(50, n, 10, 2, 5) // √(10 + 2√5)
	}
}
// No solutions





 
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
// Solutions for:
//   √(10+2√5)
func TestTri2(t *testing.T) {
	factory := NewA32s()
	min := N32(1)
	max := N32(15)
	for a := min; a <= max; a++ {
		fmt.Println(a)
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
								//	F, _ := g.Num(4)
								//	G, _ := g.Num(5)
								//	if F == 10 && G == 2 {
										fmt.Printf("a=% 3d b=% 3d c=% 3d d=% 3d e=% 3d f= %3d  g=%v\n", a, b, c, d, e, f, g)
								//	}
								}
							}
						}
					}
				}
			}
		}
	}
}
