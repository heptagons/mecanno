package hexa

import (
	"fmt"
	"math"
	"github.com/heptagons/meccano"
)

func TriangleInsideHexagon(max int) {

	for s := 1; s < max; s++ {
		for o := 1; o < s/2; o++ {
			if meccano.Gcd(s,o) == 1 {
				diag := 3*s*s - 3*s*o + 3*o*o
				cd := math.Sqrt(float64(diag))
				c := int(cd)
				if cd == float64(c) {
					fmt.Printf("s=%d o=%d c=%d p=%d b=%d\n", s, o, c, s+o, s-2*o)
				}
			}
		}
	}
	// s=13 o=2 c=21 p=15 b=9
	// s=23 o=1 c=39 p=24 b=21
	// s=37 o=11 c=57 p=48 b=15
	// s=59 o=13 c=93 p=72 b=33
	// s=73 o=26 c=111 p=99 b=21
	// s=83 o=22 c=129 p=105 b=39
	// s=94 o=23 c=147 p=117 b=48
}

// triangle_diagonals finds the integer diagonals inside equilateral meccano
// triangles:
//       C      sizes AB = BD = CA, angle ABC = 60°
//      / \     diagonal AD=Math.sqrt((a-b)*(a-b) + ab)
//     /   \        where a = AB, b = BD
//    /   _ D
//   /_ -    \
//  A---------B
func TriangleDiagonals(max int) {
	for a := 1; a < max; a++ {
		for b := 1; b <= a/2; b++ {
			if meccano.Gcd(a, b) == 1 {
				diag := (a-b)*(a-b) + a*b
				cd := math.Sqrt(float64(diag))
				d := int(cd)
				if cd == float64(d) {
					num := float64(diag + a*a - b*b)
					den := 2.0 * cd * float64(a)
					angle := 180*math.Acos(num/den)/math.Pi
					fmt.Printf("a=%3d b=%3d d=%3d angle=%8.4f\n", a, b, d, angle)
				}
			}
		}
	}
	// a=  8 b=  3 d=  7 angle= 21.7868
	// a= 15 b=  7 d= 13 angle= 27.7958
	// a= 21 b=  5 d= 19 angle= 13.1736
	// a= 35 b= 11 d= 31 angle= 17.8966
	// a= 40 b=  7 d= 37 angle=  9.4300
	// a= 48 b= 13 d= 43 angle= 15.1782
	// a= 55 b= 16 d= 49 angle= 16.4264
	// a= 65 b=  9 d= 61 angle=  7.3410
	// a= 77 b= 32 d= 67 angle= 24.4327
	// a= 80 b= 17 d= 73 angle= 11.6351
	// a= 91 b= 40 d= 79 angle= 26.0078
	// a= 96 b= 11 d= 91 angle=  6.0090
	// a= 99 b= 19 d= 91 angle= 10.4174
}

func Diagonals(max int) {
	for a := 1; a < max; a++ {
		for b := 1; b <= a/2; b++ {
			if meccano.Gcd(a, b) == 1 {
				diag := (a-b)*(a-b) + a*b
				cd := math.Sqrt(float64(diag))
				d := int(cd)
				if cd == float64(d) {
					fmt.Printf("s=%3d b=%3d d=%3d\n", a-b, b, d)
				}
			}
		}
	}
}

