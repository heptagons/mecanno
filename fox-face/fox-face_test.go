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
	w, x, y, z := a,b,c,d
	/*w.Reduce4(&x, &y, &z)
	for _, e := range e.a {
		if e[0]==w && e[1]==x && e[2]==y && e[3]==z {
			return
		}
	}*/
	e.a = append(e.a, []N32{ w,x,y,z })
}

func (abcd *ABCD) print() {
	for _, a := range abcd.a {
		fmt.Printf("\t%v\n", a)
	}
}

func TestFoxFace(t *testing.T) {

	factory := NewA32s()

	m := make(map[string]*ABCD, 0)
	max := N32(40)
	i := 1
	n1 := N32(1)
	for a := n1; a <= max; a++ {
		for b := n1; b <= max; b++ {
			for c := n1; c <= max; c++ {

				na := N32(4)*b*(b+c)
				zb := -Z(a)*(2*Z(b) + Z(c))
				zc := Z(c)
				a2c2 := Z(a*a)*Z(c*c)

				for d := c; d <= max; d++ {

					if zd := a2c2 + 4*Z(b)*Z(b+c)*(Z(d*d) - Z(c*c)); zd < 0 {
						// skip imaginary numbers invalid fox-face, like d too short
					} else if cos, err := factory.ANew3(N(na), zb, zc, zd); err != nil {
						// silent overflow
					} else if coss := cos.String(); coss != "0" {
						if _, ok := m[coss]; !ok {
							m[coss] = &ABCD{}
						}
						m[coss].append(a, b, c, d)
						//fmt.Printf("% 5d a=%d b=%d c=%d d=%d cos=%s\n", i, a, b, c, d, coss)
						i++
					}

				}
			}
		}
	}
	fmt.Printf("diff abcd %d\n", i)
	fmt.Printf("diff coss %d\n", len(m))
	for _, coss := range []string {
		"1/2",
		"(-1+√5)/4",
		"√3/4",
		"√2/4",
	} {
		fmt.Printf("coss %s\n", coss)
		if abcd, ok := m[coss]; ok {
			abcd.print()
		}
	}
}