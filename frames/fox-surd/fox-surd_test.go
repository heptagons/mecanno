package surd

import (
	"fmt"
	"testing"
	. "github.com/heptagons/meccano/nest"
)

func TestFoxSurdPenta(t *testing.T) {

	max := Z(6)
	for n := Z(6); n <= max; n++ {
		for i := Z(0); i <= n; i++ {
			ii := i*i
			n5i := n - 5*i
			n5i2 := n5i*n5i
			for j := Z(0); j <= n; j++ {
				jj := j*j
				_16jjii :=  16*(jj - ii)
				u := n5i2 + _16jjii + 5 -ii

				v := 2*((2*i-1)*(2*i-1) - (n-1)*(i-1))

				//if v == 0 {
					fmt.Printf("n=%d i=%d j=%d u=%d v=%d\n", n, i, j, u, v)
				//}
			}
		}
	}
}