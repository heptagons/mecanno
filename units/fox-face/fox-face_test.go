package face

import (
	"fmt"
	"testing"

	. "github.com/heptagons/meccano/nest"
)

type ABCD struct {
	a [][]N32
}

func (e *ABCD) append(a, b, c, d N32) {
	if e.a == nil {
		e.a = make([][]N32, 0)
	}
	e.a = append(e.a, []N32{ a,b,c,d })
}

func (abcd *ABCD) print() {
	for _, a := range abcd.a {
		fmt.Printf("\t%v\n", a)
	}
}

// FoxFace iterate a,b,c,d until max an reports a,b,c,d and
//	                    ,---------------------
//	       -a(2b+c) +- √ a²c² + 4b(b+c)(d²-c²)
//	cos = ------------------------------------
//	                    4b(b+c)
func FoxFace(max N32, found func(a, b, c, d N32, cos *A32)) {
	factory := NewA32s()
	n1 := N32(1)
	for a := n1; a <= max; a++ {
		for b := n1; b <= max; b++ {
			ab := NatGCD(a, b)
			for c := n1; c <= max; c++ {
				abc := NatGCD(ab, c)
				na := N32(4)*b*(b+c)        // 4b(b+c)
				zb := -Z(a)*(2*Z(b) + Z(c)) // -a(2b+c)
				zc := Z(1)                  // 1
				a2c2 := Z(a*a)*Z(c*c)       // a
				for d := c; d <= max; d++ { // d >= c always
					if g := NatGCD(abc, d); g > 1 {
						continue // skip scale repetitions, eg. [1,2,3,4] = [2,4,6,8]
					}
					if zd := a2c2 + 4*Z(b)*Z(b+c)*(Z(d*d) - Z(c*c)); zd < 0 {
						// skip imaginary numbers invalid fox-face, like d too short
					} else if cos, err := factory.ANew3(N(na), zb, zc, zd); err != nil {
						// silent overflow
					} else {
						found(a, b, c, d, cos)
					}
				}
			}
		}
	}
}

func TestFoxFaceAll(t *testing.T) {
	max := N32(22)
	m := make(map[string]*ABCD, 0)
	i := 1
	FoxFace(max, func(a, b, c, d N32, cos *A32) {
		if coss := cos.String(); coss != "0" && coss != "-1" && coss != "1" {
			// reject uninteresting cos=0,-1,+1
			if _, ok := m[coss]; !ok {
				m[coss] = &ABCD{}
			}
			m[coss].append(a, b, c, d)
			i++
		}
	})
	fmt.Printf("max=%d abcd's=%d diff cosines=%d\n", max, i, len(m))
	for _, coss := range []string {
		"√3/2",      // cos 30°
		"√2/2",      // cos 45°
		"1/2",       // cos 60°
		"(-1+√5)/4", // cos 72°
	} {
		fmt.Printf("arcos(%s)\n", coss)
		fmt.Printf("\t%v\n", m[coss])
	}
}

// max=100, pentagons=1 last=[3,4,8,11] time=399 seconds
func TestFoxFacePentagons(t *testing.T) {
	max := N32(11)
	fmt.Printf("max-lenght=%d a,b,c,d pentagons:\n", max)
	i := 0
	FoxFace(max, func(a, b, c, d N32, cos *A32) {
		if cos.Equals(4, -1, 1, 5) { //  cos 72°
			i++
			fmt.Printf("% 3d %d,%d,%d,%d\n", i, a, b, c, d)
		}
	})
}

// max=40   found=42 last=[32,1,7,37] time=2.7 seconds
// max=100, found=350 last=[84,1,11,91] time=399 seconds
func TestFoxFaceHexagons(t *testing.T) {
	max := N32(40)
	fmt.Printf("max-lenght=%d a,b,c,d efficient hexagons:\n", max)
	i := 0
	FoxFace(max, func(a, b, c, d N32, cos *A32) {
		if cos.Equals(2,1) { //  cos 60°
			// Efficient hexagons are those when a > b+c
			if a >= b+c {
				i++
				fmt.Printf("% 3d %d,%d,%d,%d\n", i, a, b, c, d)
			}
		}
	})
}

func TestFoxFaceOctagons(t *testing.T) {
	max := N32(80)
	fmt.Printf("max-lenght=%d a,b,c,d octagons:\n", max)
	i := 0
	FoxFace(max, func(a, b, c, d N32, cos *A32) {
		if cos.Equals(2,0,1,2) { //  cos 45 degrees sqrt{2}/2
			i++
			fmt.Printf("% 3d %d,%d,%d,%d\n", i, a, b, c, d)
		}
	})
}

func TestFoxFaceDecagons(t *testing.T) {
	max := N32(80)
	fmt.Printf("max-lenght=%d a,b,c,d decagons:\n", max)
	i := 0
	FoxFace(max, func(a, b, c, d N32, cos *A32) {
		if cos.Equals(4, 1, 1, 5) { //  cos 36°
			i++
			fmt.Printf("% 3d %d,%d,%d,%d\n", i, a, b, c, d)
		}
	})
}
// max-lenght=80 a,b,c,d decagons:
// --- PASS: TestFoxFaceDecagons (118.12s)
// NO SOLUTIONS

func TestFoxFaceDodecagons(t *testing.T) {
	max := N32(80)
	fmt.Printf("max-lenght=%d a,b,c,d dodecagons:\n", max)
	i := 0
	FoxFace(max, func(a, b, c, d N32, cos *A32) {
		if cos.Equals(2,0,1,3) { //  cos 30 degrees sqrt{3}/2
			i++
			fmt.Printf("% 3d %d,%d,%d,%d\n", i, a, b, c, d)
		}
	})
}
// max=80 dodecagons=none, time=116 seconds





func FoxTriangle() {
	min := N32(1)
	max := N32(10)
	for a := min; a <= max; a++ {
		for b := min; b <= max; b++ {
			for c := min; c <= max; c++ {
				if cos, err := factory.ANew1(2*a*b, c*c - a*a - b*b); err != nil {
					// silent
				} else {
					for d := min; d <= max; d++ {
						for e := min; e <= max; e++ {

						}
					}
				}
			}
		}
	}
}



// FoxFace iterate a,b,c,√d until max an reports a,b,c,√d and
//	                    ,---------------------
//	       -a(2b+c) +- √ a²c² + 4b(b+c)(d-c²)
//	cos = ------------------------------------
//	                    4b(b+c)
func FoxFaceDsurd(max N32, found func(a, b, c, d N32, cos *A32)) {
	factory := NewA32s()
	n1 := N32(1)
	for a := n1; a <= max; a++ {
		for b := n1; b <= max; b++ {
			ab := NatGCD(a, b)
			for c := n1; c <= max; c++ {
				abc := NatGCD(ab, c)
				na := N32(4)*b*(b+c)        // 4b(b+c)
				zb := -Z(a)*(2*Z(b) + Z(c)) // -a(2b+c)
				zc := Z(1)                  // 1
				a2c2 := Z(a*a)*Z(c*c)       // a
				for surdD := c*c; surdD <= max*max; surdD++ { // surdD >= c*c always
					if g := NatGCD(abc, surdD); g > 1 {
						continue // skip scale repetitions, eg. [1,2,3,4] = [2,4,6,8]
					}
					if zd := a2c2 + 4*Z(b)*Z(b+c)*(Z(surdD) - Z(c*c)); zd < 0 {
						// skip imaginary numbers invalid fox-face, like d too short
					} else if cos, err := factory.ANew3(N(na), zb, zc, zd); err != nil {
						// silent overflow
					} else {
						found(a, b, c, surdD, cos)
					}
				}
			}
		}
	}
}

// max=100, pentagons=1 last=[3,4,8,11] time=399 seconds
func TestFoxFaceDsurdPentagons(t *testing.T) {
	max := N32(50)
	fmt.Printf("max-lenght=%d a,b,c,surdD pentagons:\n", max)
	i := 0
	FoxFaceDsurd(max, func(a, b, c, surdD N32, cos *A32) {
		if cos.Equals(4, -1, 1, 5) { //  cos 72°
			i++
			fmt.Printf("% 3d) %d,%d,%d,sqrt(%d)\n", i, a, b, c, surdD)
		}
	})
}
/* SOLUTIONS!!!
=== RUN   TestFoxFaceDsurdPentagons
max-lenght=40 a,b,c,surdD pentagons:
  1) 2,3,3,sqrt(31)
  2) 3,4,8,sqrt(121) <- skip this is d*d=11 (original solution)
  3) 4,5,15,sqrt(341)
  4) 5,6,24,sqrt(781)
  5) 6,7,35,sqrt(1555)
  6) 6,10,5,sqrt(211)
  7) 10,14,21,sqrt(1031)
  8) 12,21,7,sqrt(781)
  9) 15,24,16,sqrt(1441)
--- PASS: TestFoxFaceDsurdPentagons (155.52s)
*/
// One more when max=50
// 10) 20,36,9,sqrt(2101)
// panic: test timed out after 10m0s



func TestFoxFaceDsurdOctagons(t *testing.T) {
	max := N32(80)
	fmt.Printf("max-lenght=%d a,b,c,surdD octagons:\n", max)
	i := 0
	FoxFaceDsurd(max, func(a, b, c, d N32, cos *A32) {
		if cos.Equals(2,0,1,2) { //  cos 45 degrees sqrt{2}/2
			i++
			fmt.Printf("% 3d) %d,%d,%d,sqrt(%d)\n", i, a, b, c, d)
		}
	})
}
// === RUN   TestFoxFaceDsurdOctagons
// max-lenght=80 a,b,c,d octagons:
// panic: test timed out after 10m0s
// NO SOLUTIONS



func TestFoxFaceDsurdDecagons(t *testing.T) {
	max := N32(40)
	fmt.Printf("max-lenght=%d a,b,c,surdD decagons:\n", max)
	i := 0
	FoxFaceDsurd(max, func(a, b, c, surdD N32, cos *A32) {
		if cos.Equals(4, 1, 1, 5) { //  cos 36°
			i++
			fmt.Printf("% 3d) %d,%d,%d,sqrt(%d)\n", i, a, b, c, surdD)
		}
	})
}
// max-lenght=40 a,b,c,surdD decagons:
// --- PASS: TestFoxFaceDsurdDecagons (157.71s)
// NO SOLUTIONS

