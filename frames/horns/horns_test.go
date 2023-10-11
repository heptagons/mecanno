package horns

import (
	"fmt"
	"testing"
	. "github.com/heptagons/meccano/nest"
)


func TestHorns(t *testing.T) {
	max := N32(7)
	Horns(max, func(a, b, c, d, e N32, cos *A32) {
		fmt.Printf("(%d,%d,%d,%d,%d)=%s\n", a, b, c, d, e, cos.String())
	})
}

func TestHornsAll(t *testing.T) {
	max := N32(11)
	m := make(map[string]*ABCDE, 0)
	i := 1
	Horns(max, func(a, b, c, d, e N32, cos *A32) {
		if coss := cos.String(); coss != "0" && coss != "-1" && coss != "1" {
			// reject uninteresting cos=0,-1,+1
			if _, ok := m[coss]; !ok {
				m[coss] = &ABCDE{}
			}
			m[coss].append(a, b, c, d, e)
			i++
		}
	})
	fmt.Printf("max=%d abcd's=%d diff cosines=%d\n", max, i, len(m))
	for _, coss := range []string {
		"√3/2",      // cos 30° dodecagon
		"(1+√5)/4",  // cos 36° decagon
		"√2/2",      // cos 45° octagon
		"1/2",       // cos 60° hexagon
		"(-1+√5)/4", // cos 72° pentagon
	} {
		fmt.Printf("arcos(%s)\n", coss)
		fmt.Printf("\t%v\n", m[coss])
	}
}


func TestHornsEPentagons(t *testing.T) {
	min,max := N32(1), N32(40)
	fmt.Printf("max-lenght=%d a,b,c,d,e pentagons:\n", max)
	i := 0
	HornsE(min, max, func(a, b, c, d, e N32) {
		i++
		fmt.Printf("% 3d) %d,%d,%d,%d,%d\n", i, a, b, c, d, e)
	},4,-1,1,5)
}
/*
max-lenght=40 a,b,c,d,e pentagons: NONE
panic: test timed out after 10m0s
*/

func TestHornsEHexagons(t *testing.T) {
	min, max := N32(0), N32(20)
	fmt.Printf("segments min=%d max=%d a,b,c,d,e:\n", min, max)
	i := 0
	HornsE(min, max, func(a, b, c, d, e N32) {
		// Interesting hexagons are those when abc triangle is not equilateral
		//if a != b && a !=c {
			i++
			fmt.Printf("% 3d) %d,%d,%d,%d,%d\n", i, a, b, c, d, e)
		//}
	},2,1) //  cos 60°
}

func TestHornsEOctagons(t *testing.T) {
	min, max := N32(1), N32(40)
	fmt.Printf("segments min=%d max=%d a,b,c,d,e:\n", min, max)
	i := 0
	HornsE(min, max, func(a, b, c, d, e N32) {
		i++
		fmt.Printf("% 3d) %d,%d,%d,%d,%d\n", i, a, b, c, d, e)
	},2,0,1,2) //  cos 45 degrees sqrt{2}/2
}

func TestHornsEDecagons(t *testing.T) {
	min,max := N32(1), N32(40)
	fmt.Printf("max-lenght=%d a,b,c,d,e decagons:\n", max)
	i := 0
	HornsE(min, max, func(a, b, c, d, e N32) {
		i++
		fmt.Printf("% 3d) %d,%d,%d,%d,%d\n", i, a, b, c, d, e)
	},4,1,1,5)
}

func TestHornsEDodecagons(t *testing.T) {
	min, max := N32(1), N32(40)
	fmt.Printf("segments min=%d max=%d a,b,c,d,e:\n", min, max)
	i := 0
	HornsE(min, max, func(a, b, c, d, e N32) {
		i++
		fmt.Printf("% 3d) %d,%d,%d,%d,%d\n", i, a, b, c, d, e)
	}, 2,0,1,3) //  cos 30 degrees sqrt{3}/2
}
/*
segments min=1 max=40 a,b,c,d,e:
  1) 6,5,5,5,8
  2) 15,4,13,21,20
  3) 15,9,12,16,20
  4) 15,14,13,11,20
  5) 15,18,15,7,20
  6) 10,13,13,13,24
  7) 21,10,17,25,28
  8) 16,17,17,17,30
  9) 30,11,25,39,40
--- PASS: TestHornsEDodecagons (566.19s)
*/




