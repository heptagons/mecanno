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

// FoxFaceCosines takes max an iterate a,b,c,d and calls
// cosFunc with a algebraic found cosine value.
func FoxFaceCosines(max N32, cosFunc func(a, b, c, d N32, cos *A32)) {
	factory := NewA32s()
	n1 := N32(1)
	for a := n1; a <= max; a++ {
		for b := n1; b <= max; b++ {
			ab := NatGCD(a, b)
			for c := n1; c <= max; c++ {
				abc := NatGCD(ab, c)

				na := N32(4)*b*(b+c)
				zb := -Z(a)*(2*Z(b) + Z(c))
				zc := Z(1)
				a2c2 := Z(a*a)*Z(c*c)

				for d := c; d <= max; d++ {
					if g := NatGCD(abc, d); g > 1 {
						continue // skip scale repetitions, eg. 1,2,3,4 = 2,4,6,8
					}
					if zd := a2c2 + 4*Z(b)*Z(b+c)*(Z(d*d) - Z(c*c)); zd < 0 {
						// skip imaginary numbers invalid fox-face, like d too short
					} else if cos, err := factory.ANew3(N(na), zb, zc, zd); err != nil {
						// silent overflow
					} else {
						cosFunc(a, b, c, d, cos)
					}
				}
			}
		}
	}
}

func TestFoxFace(t *testing.T) {

	max := N32(22)
	m := make(map[string]*ABCD, 0)
	i := 1
	FoxFaceCosines(max, func(a, b, c, d N32, cos *A32) {
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


	/*
	factory := NewA32s()
	n1 := N32(1)
	for a := n1; a <= max; a++ {
		for b := n1; b <= max; b++ {
			ab := NatGCD(a, b)
			for c := n1; c <= max; c++ {
				abc := NatGCD(ab, c)

				na := N32(4)*b*(b+c)
				zb := -Z(a)*(2*Z(b) + Z(c))
				zc := Z(1)
				a2c2 := Z(a*a)*Z(c*c)

				for d := c; d <= max; d++ {
					if g := NatGCD(abc, d); g > 1 {
						continue // skip scale repetitions, eg. 1,2,3,4 = 2,4,6,8
					}
					if zd := a2c2 + 4*Z(b)*Z(b+c)*(Z(d*d) - Z(c*c)); zd < 0 {
						// skip imaginary numbers invalid fox-face, like d too short
					} else if cos, err := factory.ANew3(N(na), zb, zc, zd); err != nil {
						// silent overflow
					} else if coss := cos.String(); coss != "0" && coss != "-1" && coss != "1" {
						// reject uninteresting cos=0,-1,+1
						if _, ok := m[coss]; !ok {
							m[coss] = &ABCD{}
						}
						m[coss].append(a, b, c, d)
						i++
					}

				}
			}
		}
	}
	*/
}