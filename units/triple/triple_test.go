package triple

import (
    "fmt"
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

func (t *Triples) Pentagons(max int) {
    for a := 1; a <= max; a++ {
        for b := 1; b < a; b++ {
            for c := 0; c < a; c++ {
                ab_ac_bc := a*b + a*c - b*c // ab + ac - bc
                aa_bb_cc := a*a + b*b + c*c // aa + bb + cc
                for d := 1; d < a; d++ {
                    if ab_ac_bc != (a-b+c)*d { // ab + c - bc != (a-b+c)d
                        // condition to reject sqrt{5} from ee equation
                        continue
                    }
                    // e = sqrt{ aa + bb + cc + dd - bc - (a+c)d }
                    if e, ok := t.squareRoot(aa_bb_cc + d*d - b*c - (a+c)*d); !ok {
                        // radical negative or not square
                    } else {
                        t.Add(a, b, c, d, e)
                    }
                }
            }
        }
    }
}


func (t *Triples) HexagonsNice(max int) {
    for a := 1; a < max; a++ {
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

func (t *Triples) Octagons(max int) {
    for a := 1; a <= max; a++ {
        for b := 1; b < a; b++ {
            for c := 0; c < a; c++ {
                a_b_c := a*(b+c)
                aa_bb_cc := a*a + b*b + c*c
                for d := 1; d < a; d++ {
                    if a_b_c - (c-b)*d != 0 {
                        continue // condition to reject sqrt{2} from e equation
                    }
                    fmt.Println(a,b,c,d)
                    if e, ok := t.squareRoot(aa_bb_cc + d*d); !ok {
                        // radical negative or not square
                    } else {
                        t.Add(a, b, c, d, e)
                    }
                }
            }
        }
        fmt.Println("a",a)
    }
}


func (t *Triples) Decagons(max int) {
    for a := 1; a <= max; a++ {
        for b := 1; b < a; b++ {
            for c := 0; c < a; c++ {
                ab_ac_bc := a*b + a*c + b*c
                aa_bb_cc := a*a + b*b + c*c
                for d := 1; d < a; d++ {
                    if ab_ac_bc != (c-a-b)*d {
                        continue // condition to reject sqrt{5} from e equation
                    }
                    if e, ok := t.squareRoot(aa_bb_cc + d*d - b*c -(a-c)*d); !ok {
                        // radical negative or not square
                    } else {
                        t.Add(a, b, c, d, e)
                    }
                }
            }
        }
        fmt.Println("a", a)
    }
}

func TestPentagons(t *testing.T) {
    tri := NewTriples()
    tri.Pentagons(500)
}


func TestHexagonsNice(t *testing.T) {
	tri := NewTriples()
	tri.HexagonsNice(60)
}

func TestOctagons(t *testing.T) {
    tri := NewTriples()
    tri.Octagons(600)
    // Conjecture: No possible octagons for triple unit!
}

func TestDecagons(t *testing.T) {
    tri := NewTriples()
    tri.Decagons(500)
    // Conjecture: No possible decagons for triple unit!
}
