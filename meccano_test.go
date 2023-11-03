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


//              (D)     x = b^2 - 5a^2    5a^2 < b^2 < 25a^2
//               |
//               | c    AB = a√5
//          √x   |      BC = √x
//      (B)-----(C)     AC = b
//       |     _/       CD = c
//   a√5 |   _/         AD = √( (√x)^2 + (a√5 + c)^2 )
//       | _/   b          = √( x + 5a^2 + c^2 + 2ac√5 )
//      (A)                = √(u + v√5)
func TestNest5Simple(t *testing.T) {
	factory := NewA32s()
	A := N(1)
	B := Z(0)
	C := Z(0)
	D := Z(1)
	E := Z(1)
	for a := 1; a <= 20; a++ {
		for b := 1; b < 10*a; b++ {
			if x := b*b - 5*a*a; x > 0 {
				for c := 1; c <= 20; c++ {
					F := Z(x) + Z(5*a*a) + Z(c*c)
					G := Z(2*a*c)
					H := Z(5)
					//	(B + C√D + E√(F+G√H)) / A
					if s, err := factory.ANew7(A, B, C, D, E, F, G, H); err == nil {
						if s.IsNest(25, 10) {
							fmt.Printf("a=% 2d b=% 2d x=% 2d c=% 2d s=%v\n", a, b, x, c, s)
						}
					}
				}
			}
		}
	}
}
/* Resolves: s.IsNest(25,10):
a= 5 b= 12 x= 19 c= 9 s=3√(25+10√5)
a= 5 b= 12 x= 19 c= 16 s=4√(25+10√5)
a= 10 b= 24 x= 76 c= 18 s=6√(25+10√5)
*/

func TestNest5Complex(t *testing.T) {
	factory := NewA32s()
	A := N(1)
	B := Z(0)
	C := Z(0)
	D := Z(1)
	E := Z(1)
	for a := 1; a <= 20; a++ { // √5 - 20√5
		for x := 1; x < 200; x++ { // horizontal
			for c := 1; c < 100; c++ { // vertical
				F := Z(x) + Z(5*a*a) + Z(c*c)
				G := Z(2*a*c)
				H := Z(5)
				//	(B + C√D + E√(F+G√H)) / A
				if s, err := factory.ANew7(A, B, C, D, E, F, G, H); err == nil {
					if s.IsNest(50, 10) {
						fmt.Printf("a=% 2d x=% 2d c=% 2d s=%v\n", a, x, c, s)
					}
				}
			}
		}
	}
}
/*
=== RUN   TestNest5Complex   --> Pentagon height
a= 4 x= 19 c= 9 s=6√(5+2√5)
a= 8 x= 76 c= 18 s=12√(5+2√5)
a= 9 x= 59 c= 16 s=12√(5+2√5)
a= 9 x= 95 c= 25 s=15√(5+2√5)
*/

/*
=== RUN   TestNest5Complex    ---> Pentagon r
a= 5 x= 19 c= 9 s=3√(25+10√5)
a= 5 x= 19 c= 16 s=4√(25+10√5)
a= 9 x= 95 c= 20 s=6√(25+10√5)
a= 10 x= 76 c= 18 s=6√(25+10√5)
a= 10 x= 76 c= 32 s=8√(25+10√5)
*/

/*
=== RUN   TestNest5Complex  ---> Pentagon R!!!
a= 1 x= 20 c= 5 s=√(50+10√5)
a= 2 x= 80 c= 10 s=2√(50+10√5)
a= 4 x= 95 c= 5 s=2√(50+10√5)
a= 5 x= 59 c= 4 s=2√(50+10√5)
a= 9 x= 20 c= 5 s=3√(50+10√5)
*/