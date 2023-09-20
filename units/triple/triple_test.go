package triple

import (
	"math"
	"testing"

	"github.com/heptagons/meccano"
)

func Hexagons(max int) *meccano.Sols {
	sols := &meccano.Sols{}
    for a := 1 ; a < max; a++ {
    	for b := 1; b < a; b++ {
    		abab := (a+b)*(a+b)
    		a_b := a - b
    		ab := a*b
        	for c := 0; c < a; c++ {
          		for d := 1; d < a; d++ {
          			if c == b+d {
          				continue; // trivial ey = 0
          			}
          			if a - b != d {
          				continue; // reject displaced copies
          			}
          			c_d := c - d
           			if f := float64(abab + (c_d)*(a_b+c_d) - ab); f <= 0 {

           			} else if e := int(math.Sqrt(f)); math.Pow(float64(e), 2) == f {
	          			if e > a {
	          				continue; // reject 
	          			}
           				sols.Add(a, b, c, d, e)
           			}
              	}
            }
        }
    }
    return sols
}

func TestHexagons(t *testing.T) {
	sols := Hexagons(60)
	_ = sols
}
