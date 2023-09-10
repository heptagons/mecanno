package face

import (
	"fmt"
	"testing"

	"github.com/heptagons/meccano/nest"
)

func TestFoxFace(t *testing.T) {

	factory := nest.NewA32s()

	m := make(map[string]bool, 0)
	max := nest.Z(11)
	i := 1
	for a := nest.Z(1); a <= max; a++ {
		for b := nest.Z(1); b <= max; b++ {
			for c := nest.Z(1); c <= max; c++ {

				_a := nest.N(4)*nest.N(b)*(nest.N(b) + nest.N(c))
				_b := -a*(2*b + c)
				_c := nest.Z(1)

				for d := c; d <= max; d++ {

					_d := a*a*c*c + 4*b*(b+c)*(d*d - c*c)
					if _d < 0 {
						// skip imaginary numbers invalid fox-face, like d too short
						continue
					}
					if cos, err := factory.ANew3(_a, _b, _c, _d); err != nil {

					} else if coss := cos.String(); coss != "0" {
						m[coss] = true
						fmt.Printf("% 5d a=%d b=%d c=%d d=%d cos=%s\n", i, a, b, c, d, coss)
						i++
					}

				}
			}
		}
	}
	fmt.Printf("diff cosines %d\n", len(m))
}