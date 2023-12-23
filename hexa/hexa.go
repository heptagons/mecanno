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
func HexagonTrianglesTex(max int) {

	fmt.Println("\\begin{tabular}{| c | c c c |}")
	fmt.Println("\\hline")
 	fmt.Println("$a$ & $c$ & $p$ & $b$ \\\\ [0.5ex]")
 	fmt.Println("\\hline\\hline")
	for a := 1; a <= max; a++ {
		for b := 1; b <= a/2; b++ {
			if meccano.Gcd(a, b) == 1 {
				diag := (a-b)*(a-b) + a*b
				cd := math.Sqrt(float64(diag))
				d := int(cd)
				if cd == float64(d) {
					//num := float64(diag + a*a - b*b)
					//den := 2.0 * cd * float64(a)
					//angle := 180*math.Acos(num/den)/math.Pi
					//fmt.Printf("a=%3d b=%3d d=%3d angle=%8.4f\n", a, b, d, angle)
					//  8 &  7 &  3 &  7 &  5 \\ \hline
					fmt.Printf("%d & %d & %d & %d ", a, d, a-b, b)
					fmt.Print("\\\\ \\hline\n")

				}
			}
		}
	}
	fmt.Println("\\end{tabular}")
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

