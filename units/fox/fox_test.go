package fox

import (
	"fmt"
	"testing"

	. "github.com/heptagons/meccano/nest"
)


func TestFoxAll(t *testing.T) {
	max := N32(22)
	m := make(map[string]*ABCD, 0)
	i := 1
	Fox(max, func(a, b, c, d N32, cos *A32) {
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
func TestFoxPentagons(t *testing.T) {
	max := N32(11)
	fmt.Printf("max-lenght=%d a,b,c,d pentagons:\n", max)
	i := 0
	Fox(max, func(a, b, c, d N32, cos *A32) {
		if cos.Equals(4, -1, 1, 5) { //  cos 72°
			i++
			fmt.Printf("% 3d %d,%d,%d,%d\n", i, a, b, c, d)
		}
	})
}

func TestFoxHexagons(t *testing.T) {
	max := N32(40)
	fmt.Printf("max-lenght=%d a,b,c,d efficient hexagons:\n", max)
	i := 0
	Fox(max, func(a, b, c, d N32, cos *A32) {
		if cos.Equals(2,1) { //  cos 60°
			// Efficient hexagons are those when a > b+c
			if a >= b+c {
				i++
				fmt.Printf("% 3d %d,%d,%d,%d\n", i, a, b, c, d)
			}
		}
	})
}
// max=40   found=42 last=[32,1,7,37] time=2.7 seconds
// max=100, found=350 last=[84,1,11,91] time=399 seconds

func TestFoxOctagons(t *testing.T) {
	max := N32(80)
	fmt.Printf("max-lenght=%d a,b,c,d octagons:\n", max)
	i := 0
	Fox(max, func(a, b, c, d N32, cos *A32) {
		if cos.Equals(2,0,1,2) { //  cos 45 degrees sqrt{2}/2
			i++
			fmt.Printf("% 3d %d,%d,%d,%d\n", i, a, b, c, d)
		}
	})
}
// === RUN   TestFoxOctagons
// max-lenght=80 a,b,c,d octagons:
// NO SOLUTIONS

func TestFoxDecagons(t *testing.T) {
	max := N32(80)
	fmt.Printf("max-lenght=%d a,b,c,d decagons:\n", max)
	i := 0
	Fox(max, func(a, b, c, d N32, cos *A32) {
		if cos.Equals(4, 1, 1, 5) { //  cos 36°
			i++
			fmt.Printf("% 3d %d,%d,%d,%d\n", i, a, b, c, d)
		}
	})
}
// max-lenght=80 a,b,c,d decagons:
// --- PASS: TestFoxDecagons (118.12s)
// NO SOLUTIONS

func TestFoxDodecagons(t *testing.T) {
	max := N32(80)
	fmt.Printf("max-lenght=%d a,b,c,d dodecagons:\n", max)
	i := 0
	Fox(max, func(a, b, c, d N32, cos *A32) {
		if cos.Equals(2,0,1,3) { //  cos 30 degrees sqrt{3}/2
			i++
			fmt.Printf("% 3d %d,%d,%d,%d\n", i, a, b, c, d)
		}
	})
}
// max=80 dodecagons=none, time=116 seconds




// max=100, pentagons=1 last=[3,4,8,11] time=399 seconds
func TestFoxSurdPentagons(t *testing.T) {
	max := N32(40)
	fmt.Printf("max-lenght=%d a,b,c,surdD pentagons:\n", max)
	i := 0
	FoxSurd(max, func(a, b, c, surdD N32, cos *A32) {
		if cos.Equals(4, -1, 1, 5) { //  cos 72°
			i++
			fmt.Printf("% 3d) %d,%d,%d,sqrt(%d)\n", i, a, b, c, surdD)
		}
	})
}
/* SOLUTIONS!!!
=== RUN   TestFoxSurdPentagons
max-lenght=40 a,b,c,surdD pentagons:
  1) 2,3,3,sqrt(31)      = sqrt(31)
  2) 3,4,8,sqrt(121)     = sqrt(11*11) Skip
  3) 4,5,15,sqrt(341)    = sqrt(11*31)
  4) 5,6,24,sqrt(781)    = sqrt(11*71)
  5) 6,7,35,sqrt(1555)   = sqrt(5*311)
  6) 6,10,5,sqrt(211)    = sqrt(211)
  7) 10,14,21,sqrt(1031) = sqrt(1031)
  8) 12,21,7,sqrt(781)   = sqrt(11*71)
  9) 15,24,16,sqrt(1441) = sqrt(11*131)
--- PASS: TestFoxSurdPentagons (155.52s)
*/
// One more when max=50
// 10) 20,36,9,sqrt(2101) = sqrt(11*191)
// panic: test timed out after 10m0s



func TestFoxSurdOctagons(t *testing.T) {
	max := N32(80)
	fmt.Printf("max-lenght=%d a,b,c,surdD octagons:\n", max)
	i := 0
	FoxSurd(max, func(a, b, c, d N32, cos *A32) {
		if cos.Equals(2,0,1,2) { //  cos 45 degrees sqrt{2}/2
			i++
			fmt.Printf("% 3d) %d,%d,%d,sqrt(%d)\n", i, a, b, c, d)
		}
	})
}
// === RUN   TestFoxDsurdOctagons
// max-lenght=80 a,b,c,d octagons:
// panic: test timed out after 10m0s
// NO SOLUTIONS



func TestFoxSurdDecagons(t *testing.T) {
	max := N32(40)
	fmt.Printf("max-lenght=%d a,b,c,surdD decagons:\n", max)
	i := 0
	FoxSurd(max, func(a, b, c, surdD N32, cos *A32) {
		if cos.Equals(4, 1, 1, 5) { //  cos 36°
			i++
			fmt.Printf("% 3d) %d,%d,%d,sqrt(%d)\n", i, a, b, c, surdD)
		}
	})
}
// max-lenght=40 a,b,c,surdD decagons:
// --- PASS: TestFoxDsurdDecagons (157.71s)
// NO SOLUTIONS


func TestFoxTrianglesSurdExt(t *testing.T) {
	FoxTrianglesSurdExt(125, 10)
}