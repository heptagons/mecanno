package horns

import (
	"fmt"
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
