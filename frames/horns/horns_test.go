package horns

import (
	"fmt"
	"testing"
	. "github.com/heptagons/meccano/nest"
)

type ABCDE struct {
	a [][]N32
}

func (abcde *ABCDE) append(a, b, c, d, e N32) {
	if abcde.a == nil {
		abcde.a = make([][]N32, 0)
	}
	abcde.a = append(abcde.a, []N32{ a,b,c,d,e })
}

func (abcde *ABCDE) print() {
	for _, a := range abcde.a {
		fmt.Printf("\t%v\n", a)
	}
}

func Horns(max N32, found func(a, b, c, d, e N32, cos *A32)) {
	factory := NewA32s()
	n1 := N32(1)
	for a := n1; a <= max; a++ {
		for b := n1; b <= max; b++ {
			ab := NatGCD(a, b)
			for c := n1; c <= max; c++ {
				if a >= b+c || b >= a+c || c >= a+b {
					continue
				}
				abc := NatGCD(ab, c)
				for d := n1; d <= max; d++ {

					abcd := NatGCD(abc, d)
					f := b + d
					h := (b*d + c*c)*f - a*a*d
					j := a*a*b - b*f*f + h
					na := 4*N(a)*N(b)*N(h)

					for e := a; e <= max; e++ { // e >= a
						if abcde := NatGCD(abcd, e); abcde > 1 {
							continue // skip scale repetitions
						}
						zb := Z(b)*Z(e)*Z(j)
						zd := Z(b)
						zd *= Z(b)*Z(e)*Z(e) - 4*Z(h)
						zd *= Z(j)*Z(j) - 4*Z(a)*Z(a)*Z(b)*Z(h)
						if zd < 0 {
							// skip imaginary numbers invalid fox-face, like d too short
						} else if cos, err := factory.ANew3(na, zb, 1, zd); err != nil {
							// silent overflow
						} else {
							found(a, b, c, d, e, cos)
						}
					}
				}
			}
		}
	}
}

// for e := min; e <= max; e++ {
//	for a := min; a <= e; a++ {
//	 for b := min; b <= max; b++ {
//    for c := min; c <= max; c++ {
//     for d := min; d <= max; d++ {
//      ...
//      found
//     }	
//    }	
//   }
//  }
// }
func HornsE(min, max N32, found func(a, b, c, d, e N32), den N32, num ...Z32) {
	factory := NewA32s()
	for e := min; e <= max; e++ {
		for a := min; a <= e; a++ {
			ea := NatGCD(e, a)
			for b := min; b <= e; b++ {
				eab := NatGCD(ea, b)
				for c:= min; c <= max; c++ {
					if a >= b+c || b >= a+c || c >= a+b {
						continue // impossible triangle abc
					}
					eabc := NatGCD(eab, c)
					for d := min; d <= max; d++ {
						if eabcd := NatGCD(eabc, d); eabcd > 1 {
							continue // scaled repetition
						}
						f := b + d
						h := (b*d + c*c)*f - a*a*d
						j := -(a*a*b - b*f*f + h) // negative for hexagons!
						na := 4*N(a)*N(b)*N(h)

						zb := Z(b)*Z(e)*Z(j)
						zd0 := Z(b)
						zd1 := Z(b)*Z(e)*Z(e) - 4*Z(h)
						zd2 := Z(j)*Z(j) - 4*Z(a)*Z(a)*Z(b)*Z(h)
						if zd1 == 0 || zd2 == 0 { // hexagon arcos=1/2
							if cos, err := factory.ANew1(na, zb); err != nil {
								// silent overflow
							} else if cos.Equals(den, num...) {
								found(a, b, c, d, e)
							}
						} else if zd := zd0*zd1*zd2; zd < 0 {
							// skip imaginary numbers
						} else if cos, err := factory.ANew3(na, zb, 1, zd); err != nil {
							// silent overflow
						} else if cos.Equals(den, num...) {
							found(a, b, c, d, e)
						}
					}
				}
			}
		}
	}
}

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
		"√3/2",      // cos 30°
		"√2/2",      // cos 45°
		"1/2",       // cos 60°
		"(-1+√5)/4", // cos 72°
	} {
		fmt.Printf("arcos(%s)\n", coss)
		fmt.Printf("\t%v\n", m[coss])
	}
}


func TestHornsEPentagons(t *testing.T) {
	max := N32(40)
	fmt.Printf("max-lenght=%d a,b,c,d,e pentagons:\n", max)
	i := 0
	Horns(max, func(a, b, c, d, e N32, cos *A32) {
		if cos.Equals(4, -1, 1, 5) { //  cos 72°
			i++
			fmt.Printf("% 3d) %d,%d,%d,%d,%d\n", i, a, b, c, d, e)
		}
	})
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




