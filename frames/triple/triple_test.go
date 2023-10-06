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
		return radical, false
	} else {
		return e, true
	}
}

func (t *Triples) Pentagons(max int, surds bool) {
    cos_3pi_5 := (1.0 - math.Sqrt(5)) / 4.0 // cos(3pi/5) = -cos(2pi/5)
    for a := 1; a <= max; a++ {
        for b := 1; b < a; b++ {
            maxD := a - int(2.0*float64(a-b)*cos_3pi_5)
            for c := 0; c < a; c++ {
                ab_ac_bc := a*b + a*c - b*c // ab + ac - bc
                aa_bb_cc := a*a + b*b + c*c // aa + bb + cc
                for d := 1; d <= maxD; d++ {
                    if ab_ac_bc != (a-b+c)*d { // ab + c - bc != (a-b+c)d
                        // condition to reject sqrt{5} from ee equation
                        continue
                    }
                    // e = sqrt{ aa + bb + cc + dd - bc - (a+c)d }
                    if e, ok := t.squareRoot(aa_bb_cc + d*d - b*c - (a+c)*d); !ok {
                        // radical negative or not square
                        if surds {
                            last := fmt.Sprintf("e=sqrt(%d)", e)
                            t.Add2(last, a, b, c, d)
                        }
                    } else {
                        if !surds {
                            t.Add(a, b, c, d, e)
                        }
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

func (t *Triples) Octagons(max int, surds bool) {
    cos_3pi_4 := -math.Sqrt(2) / 2.0 // cos(6pi/8) = -cos(2pi/8)
    for a := 1; a <= max; a++ {
        for b := 1; b <= a; b++ {
            maxD := a - int(2.0*float64(a-b)*cos_3pi_4)
            for c := 0; c <= a; c++ {
                a_b_c := a*(b+c)
                aa_bb_cc := a*a + b*b + c*c
                for d := 1; d <= maxD; d++ {
                    if a_b_c - (c-b)*d != 0 {
                        continue // condition to reject sqrt{2} from e equation
                    }
                    if e, ok := t.squareRoot(aa_bb_cc + d*d); !ok {
                        // radical negative or not square
                        if surds {
                            last := fmt.Sprintf("e=sqrt(%d)", e)
                            t.Add2(last, a, b, c, d)
                        }
                    } else {
                        if !surds {
                            t.Add(a, b, c, d, e)
                        }
                    }
                }
            }
        }
    }
}


func (t *Triples) DecagonsCBA(max int) {
	for c := 1; c <= max; c++ {
		for b := 1; b <= c; b++ {
			for a := 1; a <= c; a++ {
				ab_ac_bc := a*b + a*c + b*c
				aa_bb_cc := a*a + b*b + c*c
				for d := 1; d <= max; d++ {
					if ab_ac_bc != (c-a-b)*d {
						continue // condition to reject sqrt{5} from e equation
					}
					if e, ok := t.squareRoot(aa_bb_cc + d*d - b*c -(a-c)*d); ok {
						t.Add(a, b, c, d, e)
					}
				}
			}
		}
	}
}

func TestPentagonsIntE(t *testing.T) {
    tri := NewTriples()
    tri.Pentagons(100, false) // no surds
}
/*
=== RUN   TestPentagons
  1  a= 12 b=  2 c=  9 d=  6 e= 11
  2  a= 12 b=  3 c=  0 d=  4 e= 11
  3  a= 12 b=  6 c=  3 d= 10 e= 11
  4  a= 12 b=  9 c=  8 d= 12 e= 11 // d = a
  5  a= 28 b= 16 c=  4 d= 31 e= 31 // d > a
  6  a= 31 b=  4 c= 28 d= 16 e= 31
  7  a= 31 b= 15 c=  3 d= 27 e= 31
  8  a= 38 b= 12 c= 18 d= 21 e= 31
  9  a= 38 b= 17 c= 20 d= 26 e= 31
 10  a= 48 b=  8 c= 24 d= 21 e= 41
 11  a= 48 b= 12 c=  9 d= 20 e= 41
 12  a= 48 b= 27 c= 24 d= 40 e= 41
 13  a= 48 b= 28 c= 39 d= 36 e= 41
 14  a= 72 b= 21 c= 48 d= 40 e= 61
 15  a= 72 b= 24 c= 16 d= 39 e= 61
 16  a= 72 b= 32 c= 24 d= 51 e= 61
 17  a= 72 b= 33 c= 56 d= 48 e= 61
 18  a= 74 b= 51 c= 36 d= 78 e= 71 // d > a
 19  a= 78 b= 27 c=  4 d= 42 e= 71
 20  a= 78 b= 36 c= 74 d= 51 e= 71
 21  a= 87 b= 28 c= 36 d= 48 e= 71
 22  a= 87 b= 39 c= 51 d= 59 e= 71
 23  a= 99 b= 67 c= 39 d=111 e=101 // d > a
--- PASS: TestPentagons (0.07s)
*/

func TestPentagonsSurdE(t *testing.T) {
    tri := NewTriples()
    tri.Pentagons(12, true)
}
/*
=== RUN   TestPentagonsSurdE
  1  a=  2 b=  1 c=  0 d=  2 e=sqrt(5)
  2  a=  4 b=  1 c=  2 d=  2 e=sqrt(11)
  3  a=  4 b=  2 c=  2 d=  3 e=sqrt(11)
  4  a=  6 b=  2 c=  0 d=  3 e=sqrt(31)
  5  a=  6 b=  4 c=  2 d=  7 e=sqrt(41)
  6  a=  6 b=  4 c=  3 d=  6 e=sqrt(31)
  7  a=  7 b=  2 c=  6 d=  4 e=sqrt(41)
  8  a=  7 b=  3 c=  1 d=  5 e=sqrt(41)
  9  a=  9 b=  1 c=  3 d=  3 e=sqrt(61)
 10  a=  9 b=  3 c=  3 d=  5 e=sqrt(55)
 11  a=  9 b=  4 c=  6 d=  6 e=sqrt(55)
 12  a=  9 b=  6 c=  6 d=  8 e=sqrt(61)
 13  a= 10 b=  2 c=  3 d=  4 e=sqrt(71)
 14  a= 10 b=  6 c=  7 d=  8 e=sqrt(71)
 15  a= 12 b=  4 c=  8 d=  7 e=sqrt(101)
 16  a= 12 b=  5 c=  4 d=  8 e=sqrt(101)
*/




func TestHexagonsNice(t *testing.T) {
	tri := NewTriples()
	tri.HexagonsNice(60)
}
/*
=== RUN   TestHexagonsNice
  1  a=  7 b=  1 c=  2 d=  6 e=  7
  2  a=  7 b=  1 c=  4 d=  6 e=  7
  3  a= 13 b=  2 c=  5 d= 11 e= 13
  4  a= 13 b=  2 c=  6 d= 11 e= 13
  5  a= 14 b=  1 c=  6 d= 13 e= 13
  6  a= 14 b=  1 c=  7 d= 13 e= 13
  7  a= 15 b=  1 c=  5 d= 14 e= 14
  8  a= 15 b=  1 c=  9 d= 14 e= 14
  9  a= 19 b=  2 c=  3 d= 17 e= 19
 10  a= 19 b=  2 c= 14 d= 17 e= 19
 11  a= 20 b=  1 c=  4 d= 19 e= 19
 12  a= 20 b=  1 c= 15 d= 19 e= 19
 13  a= 22 b=  2 c=  7 d= 20 e= 21
 14  a= 22 b=  2 c= 13 d= 20 e= 21
 15  a= 23 b=  1 c=  8 d= 22 e= 21
 16  a= 23 b=  1 c= 14 d= 22 e= 21
 17  a= 27 b=  3 c= 11 d= 24 e= 26
 18  a= 27 b=  3 c= 13 d= 24 e= 26
 19  a= 29 b=  1 c= 13 d= 28 e= 26
 20  a= 29 b=  1 c= 15 d= 28 e= 26
 21  a= 29 b=  3 c=  9 d= 26 e= 28
 ...
106  a= 58 b=  5 c= 43 d= 53 e= 57
107  a= 59 b=  1 c= 27 d= 58 e= 52
108  a= 59 b=  1 c= 31 d= 58 e= 52
109  a= 59 b=  4 c= 11 d= 55 e= 57
110  a= 59 b=  4 c= 44 d= 55 e= 57
111  a= 59 b=  5 c= 19 d= 54 e= 56
112  a= 59 b=  5 c= 35 d= 54 e= 56
--- PASS: TestHexagonsNice (0.01s)
*/

func TestOctagonsIntE(t *testing.T) {
    tri := NewTriples()
    tri.Octagons(100, false)
    // Conjecture: No possible octagons for triple frame where a > d
}
/*
=== RUN   TestOctagonsIntE
  1  a= 24 b=  5 c= 20 d= 40 e= 51
  2  a= 30 b=  3 c=  8 d= 66 e= 73
  3  a= 34 b=  7 c= 24 d= 62 e= 75
  4  a= 45 b=  6 c= 26 d= 72 e= 89
  5  a= 46 b= 12 c= 35 d= 94 e=111
  6  a= 51 b=  6 c= 42 d= 68 e= 95
  7  a= 62 b=  9 c= 40 d= 98 e=123
  8  a= 63 b=  4 c= 60 d= 72 e=113
  9  a= 76 b=  2 c= 18 d= 95 e=123
 10  a= 78 b=  5 c= 20 d=130 e=153
 11  a= 93 b=  8 c= 56 d=124 e=165
 12  a= 94 b= 16 c= 63 d=158 e=195
 13  a= 96 b=  3 c= 12 d=160 e=187
 14  a= 98 b= 11 c= 60 d=142 e=183
--- PASS: TestOctagonsIntE (0.09s)
*/


func TestOctagonsSurdE(t *testing.T) {
    tri := NewTriples()
    tri.Octagons(10, true)
}
/*
=== RUN   TestOctagonsSurdE
  1  a=  4 b=  1 c=  3 d=  8 e=sqrt(90)
  2  a=  5 b=  1 c=  3 d= 10 e=sqrt(135)
  3  a=  6 b=  1 c=  3 d= 12 e=sqrt(190)
  4  a=  6 b=  1 c=  4 d= 10 e=sqrt(153)
  5  a=  6 b=  1 c=  5 d=  9 e=sqrt(143)
  6  a=  7 b=  1 c=  3 d= 14 e=sqrt(255)
  7  a=  7 b=  2 c=  6 d= 14 e=sqrt(285)
  8  a=  8 b=  1 c=  3 d= 16 e=sqrt(330)
  9  a=  8 b=  1 c=  5 d= 12 e=sqrt(234)
 10  a=  9 b=  1 c=  3 d= 18 e=sqrt(415)
 11  a=  9 b=  1 c=  4 d= 15 e=sqrt(323)
 12  a=  9 b=  1 c=  7 d= 12 e=sqrt(275)
 13  a=  9 b=  2 c=  6 d= 18 e=sqrt(445)
 14  a=  9 b=  2 c=  8 d= 15 e=sqrt(374)
 15  a= 10 b=  1 c=  3 d= 20 e=sqrt(510)
 16  a= 10 b=  1 c=  5 d= 15 e=sqrt(351)
 17  a= 10 b=  1 c=  6 d= 14 e=sqrt(333)
 18  a= 10 b=  2 c=  7 d= 18 e=sqrt(477)
 19  a= 10 b=  2 c= 10 d= 15 e=sqrt(429)
--- PASS: TestOctagonsSurdE (0.00s)
*/



func TestDecagonsCBA(t *testing.T) {
    tri := NewTriples()
    tri.DecagonsCBA(500)
}
/*
=== RUN   TestDecagonsCBA
  1  a=  8 b=  4 c= 13 d=188 e=191
  2  a=  3 b=  6 c= 18 d= 20 e= 31
  3  a=  6 b=  3 c= 20 d= 18 e= 31
  4  a= 12 b=  8 c= 36 d= 51 e= 71
  5  a= 24 b=  8 c= 51 d= 96 e=121
  6  a=  8 b= 12 c= 51 d= 36 e= 71
  7  a= 42 b=  7 c= 60 d=294 e=311
  8  a= 20 b= 30 c= 75 d=174 e=211
  9  a= 44 b= 24 c= 84 d=423 e=451
 10  a=  2 b= 63 c= 84 d=294 e=341
 11  a=  7 b= 57 c= 93 d=219 e=271
 12  a=  8 b= 24 c= 96 d= 51 e=121
 13  a= 60 b= 15 c=104 d=300 e=341
 14  a= 42 b= 36 c=114 d=289 e=341
 15  a= 45 b= 24 c=128 d=168 e=241
 16  a= 15 b= 57 c=133 d=171 e=251
 17  a= 72 b= 39 c=152 d=480 e=541
 18  a= 24 b= 84 c=153 d=412 e=491
 19  a= 13 b= 83 c=167 d=241 e=341
 20  a= 24 b= 45 c=168 d=128 e=241
 21  a= 53 b= 55 c=169 d=347 e=431
 22  a= 57 b= 15 c=171 d=133 e=251
 23  a= 21 b= 91 c=171 d=357 e=451
 24  a= 30 b= 20 c=174 d= 75 e=211
 25  a=  4 b=  8 c=188 d= 13 e=191
 26  a=117 b=  3 c=219 d=269 e=401
 27  a= 57 b=  7 c=219 d= 93 e=271
 28  a= 28 b= 98 c=221 d=322 e=451
 29  a= 34 b= 93 c=228 d=318 e=451
 30  a= 83 b= 13 c=241 d=167 e=341
 31  a=109 b= 24 c=264 d=288 e=451
 32  a= 24 b=144 c=267 d=488 e=641
 33  a=  3 b=117 c=269 d=219 e=401
 34  a= 36 b= 96 c=276 d=277 e=451
 35  a= 96 b= 36 c=277 d=276 e=451
 36  a= 24 b=109 c=288 d=264 e=451
 37  a= 36 b= 42 c=289 d=114 e=341
 38  a= 63 b=  2 c=294 d= 84 e=341
 39  a=  7 b= 42 c=294 d= 60 e=311
 40  a= 15 b= 60 c=300 d=104 e=341
 41  a= 93 b= 34 c=318 d=228 e=451
 42  a= 98 b= 28 c=322 d=221 e=451
 43  a= 55 b= 53 c=347 d=169 e=431
 44  a= 91 b= 21 c=357 d=171 e=451
 45  a=105 b= 87 c=363 d=461 e=671
 46  a=180 b= 24 c=380 d=465 e=691
 47  a=105 b= 90 c=406 d=420 e=671
 48  a= 84 b= 24 c=412 d=153 e=491
 49  a= 90 b=105 c=420 d=406 e=671
 50  a= 24 b= 44 c=423 d= 84 e=451
 51  a=222 b= 12 c=454 d=495 e=781
 52  a= 87 b=105 c=461 d=363 e=671
 53  a= 24 b=180 c=465 d=380 e=691
 54  a= 39 b= 72 c=480 d=152 e=541
 55  a=144 b= 24 c=488 d=267 e=641
 56  a= 12 b=222 c=495 d=454 e=781
--- PASS: TestDecagonsCBA (42.60s)
*/

