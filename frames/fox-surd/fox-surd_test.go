package surd

import (
	"fmt"
	"testing"
	. "github.com/heptagons/meccano/nest"
)

func TestFoxSurdPenta(t *testing.T) {

	max := Z(36)
	for n := Z(1); n <= max; n++ {
		for i := Z(0); i <= n; i++ {
			xi := n - 2*i
			for j := Z(0); j <= n; j++ {

				u := 4*j*j + xi*xi + 5*n*n + xi*j - 5*n*j
				v := 2*n*xi - xi*j + n*j

				if v == 0 {
					if n == i && n == j {

					} else {
						fmt.Printf("n=%d i=%d j=%d u=%d v=%d\n", n, i, j, u, v)
					}
				}
			}
		}
	}
}