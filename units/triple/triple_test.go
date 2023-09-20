package triple

import (
	"math"
	"testing"

	"github.com/heptagons/meccano"
)

type Triples struct {
	*meccano.Sols
}

func NewTriples() *Triples {
	return &Triples{
		Sols: &meccano.Sols{},
	}
}

func (t *Triples) squareRoot(radical int) (int, bool) {
	if f := float64(radical); f <= 0 {
		// radical is negative
		return 0, false
	} else if e := int(math.Sqrt(f)); math.Pow(float64(e), 2) != f {
		// e is not square
		return 0, false
	} else {
		return e, true
	}
}

func (t *Triples) HexagonsNice(max int) {
    for a := 1 ; a < max; a++ {
    	for b := 1; b < a; b++ {
    		ab2, a_b, ab := (a+b)*(a+b), a - b, a*b
        	for c := 0; c < a; c++ {
          		for d := 1; d < a; d++ {
          			if c == b+d {
          				continue; // reject trivial e y-coordinate = 0 (e horizontal)
          			}
          			if a - b != d {
          				continue; // reject displaced copies
          			}
          			c_d := c - d
           			if e, ok := t.squareRoot(ab2 + (c_d)*(a_b+c_d) - ab); !ok {
           				// radical negative or not square
           			} else if e <= a {
           				// just "nice" hexagons without e larger than a
           				t.Add(a, b, c, d, e)
           			}
              	}
            }
        }
    }
}

func TestHexagonsNice(t *testing.T) {
	tri := NewTriples()
	tri.HexagonsNice(60)
}
