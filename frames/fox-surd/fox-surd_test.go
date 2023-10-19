package surd

import (
	"fmt"
	"testing"
	. "github.com/heptagons/meccano/nest"
)

func TestFoxSurdPenta(t *testing.T) {

	factory := NewA32s()

	max := Z(6)
	for n := Z(1); n <= max; n++ {
		nn := n*n
		for i := Z(0); i <= n; i++ {
			ni := n*i
			ii := i*i
			for j := Z(0); j <= n; j++ {
				if n == i && n == j {
					continue
				}
				nj := n*j
				jj := j*j
				ij := i*j

				u := 6*nn + 4*ii + 4*jj - 4*ni - 4*nj - 2*ij
				v := 2*nn - 4*ni + 2*ij

				if v == 0 {
					//	((b + c√d) / a) = (0 + 1√u)/2
					if s, err := factory.ANew3(2, 0, Z(1), u); err == nil {
						fmt.Printf("n=%d i=%d j=%d s=%v\n", n, i, j, s)
					}
				} else {
					//	((b + c√d + e√(f+g√h)) / a) = (√(u + v√5)) / 2
					f := u
					g := v
					h := Z(5)
					if s, err := factory.ANew7(2, 0, 0, 0, Z(1), f, g, h); err == nil {
						fmt.Printf("n=%d i=%d j=%d s=%v\n", n, i, j, s.String())
					}
				}
			}
		}
	}
}