package main

import (
	"fmt"
	"math"

	"github.com/heptagons/meccano/penta"
)

func main() {
	//penta.Type_1(5000)
	//penta.Type_2(100)
	//penta.Type_2_half(1000)
	penta.Type_2_half_with_conjecture(1000)
	//octagons_2(60)
	//triangle_diagonals(200)
}






// triangle_diagonals finds the integer diagonals inside equilateral meccano
// triangles:
//       C      sizes AB = BD = CA, angle ABC = 60°
//      / \     diagonal AD=Math.sqrt((a-b)*(a-b) + ab)
//     /   \        where a = AB, b = BD
//    /   _ D
//   /_ -    \
//  A---------B
func triangle_diagonals(max int) {
	for a := 1; a < max; a++ {
		for b := 1; b <= a/2; b++ {
			if gcd(a, b) == 1 {
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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}

// Find bars forming octagon internal angle 135°.
//  
//  D . _ 
//  | \   ` - _ C
//  |   \       /
//  |     \   /
//  A-------B

//  Angles: ABD=45°, DBC=90°, ABC=135°

//  i = A-D = A-D, B-D = i*sqrt(2)
//  we iterate bc to find DC integer
//  
func octagons_1() {
	for ab := 1; ab <= 30; ab++ {
		h_2 := 2*ab*ab
		for bc := 1; bc <= 4*ab; bc++ {
			h_3 :=  float64(bc*bc + h_2)
			cd := math.Sqrt(h_3)
			_cd := int(cd)
			if cd == float64(_cd) {
				if gcd(_cd, gcd(ab, bc)) == 1 {
					fmt.Printf("✔ ab=%2d bc=%2d cd=%2d\n", ab, bc, _cd)
				}
			}
		}
	}
	// Expected
	// ✔ ab= 2 bc= 1 cd= 3
	// ✔ ab= 4 bc= 7 cd= 9
	// ✔ ab= 6 bc= 7 cd=11
	// ✔ ab= 6 bc=17 cd=19
	// ✔ ab=10 bc=23 cd=27
	// ✔ ab=12 bc= 1 cd=17
	// ✔ ab=14 bc=47 cd=51
	// ✔ ab=20 bc=17 cd=33
	// ✔ ab=24 bc=23 cd=41
	// ✔ ab=28 bc=41 cd=57
	// ✔ ab=30 bc= 7 cd=43
	// ✔ ab=30 bc=41 cd=59
}

// Octagons bars second method
// 1) i iterate cd=1,2,3,...
// 2) j iterate bc=1,2,3 < cd
// 3) calculate bd*bd = cd**cd - bc*bc
// 4) accept when bd*db = 2 * square
func octagons_2(max int) {
	for a := 1; a < max; a++ {
		for b := 1; b < a; b++ {
			cc := a*a - b*b
			if cc % 2 == 0 {
				f := math.Sqrt(float64(cc/2))
				c := int(f)
				if f == float64(c) {
					if gcd(c, gcd(b, a)) == 1 {
						s := int(math.Max(float64(b), f))
						fmt.Printf("a=%2d b=%2d c=%2d s=%2d\n", a, b, c, s)
					}
				}
			}
		}
	}
}