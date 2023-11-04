package meccano

import (
	"fmt"
	"testing"

	. "github.com/heptagons/meccano/nest"
)

func TestPythagorasInt(t *testing.T) {
	factory := NewA32s()
	a := N(1)
	b := Z(0)
	c := Z(1)
	max := 25
	// y = 1,2,3...,max
	// x = 1,...,y
	for y := 1; y <= max; y++ {
		fmt.Printf("y=%2d", y)
		for x := 1; x <= y-1; x++ {
			d := Z(y-x) * Z(y+x)
			if surd, err := factory.ANew3(a, b, c, d); err == nil {
				fmt.Printf("  %5s", surd.String())
			}
		}
		fmt.Println()
	}
	fmt.Printf("     ")
	for n := 1; n < max; n++ {
		fmt.Printf("  x=%2d ", n)
	}
	fmt.Println()
}

func TestPythagorasHalf(t *testing.T) {
	factory := NewA32s()
	a := N(2)
	b := Z(0)
	c := Z(1)
	max := 25
	// xs = 1/2, 3/2, 5/2...
	// surd = sqrt((2y-x)(2y+x)) / 2
	for y := 1; y <= max; y++ {
		fmt.Printf("y=%2d", y)
		for x := 1; x < 2*y; x += 2 {
			d := Z(2*y-x)*Z(2*y+x)
			if surd, err := factory.ANew3(a, b, c, d); err == nil {
				fmt.Printf("  %7s", surd.String())
			}
		}
		fmt.Println()
	}
	fmt.Printf("     ")
	for n := 1; n < 2*max; n +=2 {
		fmt.Printf("  x=%2d/2 ", n)
	}
	fmt.Println()
}


//              (D)     x = b^2 - r^2s    r^2s < b^2 < r^2s^2
//               |
//               | y    AB = r√s
//          √x   |      BC = √x
//      (B)-----(C)     AC = b
//       |     _/       CD = y
//   r√s |   _/         AD = √( (√x)^2 + (r√s + y)^2 )
//       | _/   b          = √( x + r^2s + y^2 + 2yr√s )
//      (A)                = √(u + v√s)
func TestNest5Simple(t *testing.T) {
	factory := NewA32s()
	A := N(1)
	B := Z(0)
	C := Z(0)
	D := Z(1)
	E := Z(1)
	H := Z(5) // for pentagons
	for r := Z(1); r <= 20; r++ {
		for b := Z(1); b < 10*r; b++ {
			if x := b*b - r*r*H; x > 0 {
				for y := Z(1); y <= 20; y++ {
					F := Z(x) + Z(r*r)*H + Z(y*y)
					G := Z(2*r*y)
					//	(B + C√D + E√(F+G√H)) / A
					if s, err := factory.ANew7(A, B, C, D, E, F, G, H); err == nil {
						if s.IsNest(25, 10) {
							fmt.Printf("r=%d√%d b=%d x=√%d y=%d s=%v\n", r, H, b, x, y, s)
						}
					}
				}
			}
		}
	}
}
// Can't solve: √(5+2√5), √(10+3√5)

/* Solves pentagon's r
r=5√5 b=12 x=√19 y=9 s=3√(25+10√5)
r=5√5 b=12 x=√19 y=16 s=4√(25+10√5)
r=10√5 b=24 x=√76 y=18 s=6√(25+10√5)
*/

/* Solves pentagon's R
r=1√5 b=5 x=√20 y=5 s=√(50+10√5)
r=2√5 b=10 x=√80 y=10 s=2√(50+10√5)
r=3√5 b=15 x=√180 y=15 s=3√(50+10√5)
r=4√5 b=20 x=√320 y=20 s=4√(50+10√5)
*/



//                 (D)       AB = r√s
//                  |        BC = √x
//                  | y      CD = y
//                  |              ____________________
//      (B)--------(C)       AD = √ (√x)^2 + (r√s+y)^2
//       |    √x                   ________________________
//  r√s  |                      = √ x + r^2s + y^2 + 2ry√s
//       |
//      (A)
//
func TestNest5Complex(t *testing.T) {
	factory := NewA32s()
	A := N(1)
	B := Z(0)
	C := Z(0)
	D := Z(1)
	E := Z(1)
	H := Z(5) // √s for pentagons
	for r := 1; r <= 20; r++ { // √5 - 20√5
		for x := 1; x < 200; x++ { // horizontal
			for y := 1; y < 100; y++ { // vertical
				F := Z(x) + Z(r*r)*H + Z(y*y)
				G := Z(2*r*y)
				//	(B + C√D + E√(F+G√H)) / A
				if s, err := factory.ANew7(A, B, C, D, E, F, G, H); err == nil {
					if s.IsNest(10,3) {
						fmt.Printf("r=%d√%d x=√%d y=%d s=%v\n", r, H, x, y, s)
					}
				}
			}
		}
	}
}
/*
=== RUN   TestNest5Complex   --> Pentagon H -> height (r+R)
a= 4 x= 19 c= 9 s=6√(5+2√5)
a= 8 x= 76 c= 18 s=12√(5+2√5)
a= 9 x= 59 c= 16 s=12√(5+2√5)
a= 9 x= 95 c= 25 s=15√(5+2√5)
a= 12 x= 171 c= 27 s=18√(5+2√5)
a= 16 x= 95 c= 25 s=20√(5+2√5)
*/

/*
=== RUN   TestNest5Complex    ---> Pentagon r -> inradius
a= 5 x= 19 c= 9 s=3√(25+10√5)
a= 5 x= 19 c= 16 s=4√(25+10√5)
a= 9 x= 95 c= 20 s=6√(25+10√5)
a= 10 x= 76 c= 18 s=6√(25+10√5)
a= 10 x= 76 c= 32 s=8√(25+10√5)
*/

/*
=== RUN   TestNest5Complex  ---> Pentagon R -> circumradius
a= 1 x= 20 c= 5 s=√(50+10√5)
a= 2 x= 80 c= 10 s=2√(50+10√5)
a= 4 x= 95 c= 5 s=2√(50+10√5)
a= 5 x= 59 c= 4 s=2√(50+10√5)
a= 9 x= 20 c= 5 s=3√(50+10√5)
*/




//                      (F)        x  = AB
//                       |        r√s = BD
//          (C)_       _(E)         y = DF
//           |  -_   _-  |               __________________________
//           |   _(D)_   |         AF = √ x^2 + r^2s + y^2 + 2rx√s 
//           | _-     -_ |
//  (A)-----(B)         (D)
//
func TestNestSix(t *testing.T) {
	factory := NewA32s()
	A := N(1)
	B := Z(0)
	C := Z(0)
	D := Z(1)
	E := Z(1)
	H := Z(5) // √s for pentagons
	for r := 1; r <= 50; r++ {
		for x := 1; x <= 50; x++ {
			G := Z(2*r*x) // 2rx
			for y := 1; y <= 50; y++ {
				F := Z(x*x) + Z(r)*Z(r)*H + Z(y*y) // x^2 + r^2s + y^2
				//	(B + C√D + E√(F+G√H)) / A
				if s, err := factory.ANew7(A, B, C, D, E, F, G, H); err == nil {
					if s.IsNest(10,2) {
						fmt.Printf("% 2d√%d x=% 2d y=% 2d s=%v\n", r, H, x, y, s)
					}
				}
			}
		}
	}
}
// Simpler TestNestSix: Can't solve √(5+2√5), √(25+10√5), √(50+10√5)
// Can't solve: √(10+3√5), √(46+18√5)
// even with r <= 50, x <=50, y <= 50
// But solves several: √(10+2√5)