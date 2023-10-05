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
                        fmt.Println(a, b, c, d, "sqrt", e)
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
        for b := 1; b <= max; b++ {
            for c := 0; c <= max; c++ {
                a_b_c := a*(b+c)
                aa_bb_cc := a*a + b*b + c*c
                for d := 1; d <= max; d++ {
                    if a_b_c - (c-b)*d != 0 {
                        continue // condition to reject sqrt{2} from e equation
                    }
                    fmt.Println(a,b,c,d, math.Sqrt(float64(aa_bb_cc + d*d)))
                    if e, ok := t.squareRoot(aa_bb_cc + d*d); !ok {
                        // radical negative or not square
                    } else {
                        t.Add(a, b, c, d, e)
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

func TestPentagons(t *testing.T) {
    tri := NewTriples()
    tri.Pentagons(20)
}
/*
=== RUN   TestPentagons
  1  a= 12 b=  2 c=  9 d=  6 e= 11
  2  a= 12 b=  3 c=  0 d=  4 e= 11
  3  a= 12 b=  6 c=  3 d= 10 e= 11
  4  a= 31 b=  4 c= 28 d= 16 e= 31
  5  a= 31 b= 15 c=  3 d= 27 e= 31
  6  a= 38 b= 12 c= 18 d= 21 e= 31
  7  a= 38 b= 17 c= 20 d= 26 e= 31
  8  a= 48 b=  8 c= 24 d= 21 e= 41
  9  a= 48 b= 12 c=  9 d= 20 e= 41
 10  a= 48 b= 27 c= 24 d= 40 e= 41
 11  a= 48 b= 28 c= 39 d= 36 e= 41
 12  a= 72 b= 21 c= 48 d= 40 e= 61
 13  a= 72 b= 24 c= 16 d= 39 e= 61
 14  a= 72 b= 32 c= 24 d= 51 e= 61
 15  a= 72 b= 33 c= 56 d= 48 e= 61
 16  a= 78 b= 27 c=  4 d= 42 e= 71
 17  a= 78 b= 36 c= 74 d= 51 e= 71
 18  a= 87 b= 28 c= 36 d= 48 e= 71
 19  a= 87 b= 39 c= 51 d= 59 e= 71
 20  a=111 b= 39 c= 99 d= 67 e=101
 21  a=111 b= 44 c= 12 d= 72 e=101
 22  a=121 b= 33 c= 33 d= 57 e=101
 23  a=121 b= 64 c= 88 d= 88 e=101
 24  a=128 b=  8 c= 89 d= 56 e=121
 25  a=128 b= 72 c= 39 d=120 e=121
 26  a=138 b= 12 c= 54 d= 47 e=121
 27  a=138 b= 91 c= 84 d=126 e=121
 28  a=145 b= 45 c= 39 d= 75 e=121
 29  a=145 b= 70 c=106 d=100 e=121
 30  a=147 b= 43 c= 51 d= 75 e=121
 31  a=147 b= 72 c= 96 d=104 e=121
 32  a=151 b= 19 c= 73 d= 61 e=131
 33  a=151 b= 90 c= 78 d=132 e=131
 34  a=156 b= 43 c= 96 d= 84 e=131
 35  a=156 b= 72 c= 60 d=113 e=131
 36  a=165 b= 36 c=132 d= 88 e=151
 37  a=165 b= 77 c= 33 d=129 e=151
 38  a=179 b= 15 c=177 d= 93 e=191
 39  a=179 b= 86 c=  2 d=164 e=191
 40  a=183 b= 66 c= 62 d=108 e=151
 41  a=183 b= 75 c=121 d=117 e=151
 42  a=201 b=  9 c= 13 d= 21 e=191
 43  a=201 b=180 c=188 d=192 e=191
 44  a=204 b= 21 c=112 d= 84 e=181
 45  a=204 b=120 c= 92 d=183 e=181
 46  a=216 b= 48 c=111 d=104 e=181
 47  a=216 b=112 c=105 d=168 e=181
 48  a=236 b= 80 c= 20 d=125 e=211
 49  a=236 b=111 c=216 d=156 e=211
 50  a=249 b= 45 c= 75 d= 95 e=211
 51  a=249 b=154 c=174 d=204 e=211
 52  a=264 b= 76 c=  3 d=108 e=241
 53  a=264 b=156 c=261 d=188 e=241
 54  a=285 b= 73 c= 27 d=111 e=251
 55  a=285 b=174 c=258 d=212 e=251
 56  a=296 b=104 c=128 d=173 e=241
 57  a=296 b=123 c=168 d=192 e=241
 58  a=303 b= 51 c= 29 d= 81 e=271
 59  a=303 b=222 c=274 d=252 e=271
 60  a=304 b= 76 c=133 d=148 e=251
 61  a=304 b=156 c=171 d=228 e=251
 62  a=312 b= 36 c= 93 d=100 e=271
 63  a=312 b=212 c=219 d=276 e=271
 64  a=315 b= 24 c=160 d=120 e=281
 65  a=315 b=195 c=155 d=291 e=281
 66  a=324 b= 64 c=204 d=159 e=281
 67  a=324 b=165 c=120 d=260 e=281
 68  a=343 b=  7 c=115 d= 91 e=311
 69  a=343 b=252 c=228 d=336 e=311
 70  a=352 b=  3 c=240 d=144 e=341
 71  a=352 b=208 c=112 d=349 e=341
 72  a=354 b= 53 c= 60 d=102 e=311
 73  a=354 b=252 c=294 d=301 e=311
 74  a=368 b= 36 c=219 d=156 e=331
 75  a=368 b=212 c=149 d=332 e=331
 76  a=369 b= 37 c= 27 d= 63 e=341
 77  a=369 b=306 c=342 d=332 e=341
 78  a=370 b=  1 c=172 d=118 e=341
 79  a=370 b=252 c=198 d=369 e=341
 80  a=375 b= 15 c=191 d=135 e=341
 81  a=375 b=240 c=184 d=360 e=341
 82  a=378 b= 21 c= 84 d= 86 e=341
 83  a=378 b=292 c=294 d=357 e=341
 84  a=384 b=120 c=312 d=223 e=341
 85  a=384 b=161 c= 72 d=264 e=341
 86  a=390 b= 84 c= 50 d=135 e=341
 87  a=390 b= 87 c=228 d=194 e=331
 88  a=390 b=196 c=162 d=303 e=331
 89  a=390 b=255 c=340 d=306 e=341
 90  a=392 b=119 c=296 d=224 e=341
 91  a=392 b=128 c= 56 d=203 e=341
 92  a=392 b=168 c= 96 d=273 e=341
 93  a=392 b=189 c=336 d=264 e=341
 94  a=393 b= 98 c= 54 d=156 e=341
 95  a=393 b=237 c=339 d=295 e=341
 96  a=396 b=138 c= 73 d=222 e=341
 97  a=396 b=174 c=323 d=258 e=341
 98  a=399 b= 70 c=210 d=180 e=341
 99  a=399 b=219 c=189 d=329 e=341
100  a=403 b= 78 c=114 d=156 e=341
101  a=403 b=247 c=289 d=325 e=341
102  a=404 b= 89 c=104 d=164 e=341
103  a=404 b=240 c=300 d=315 e=341
104  a=408 b= 16 c=312 d=183 e=401
105  a=408 b= 84 c=167 d=180 e=341
106  a=408 b=225 c= 96 d=392 e=401
107  a=408 b=228 c=241 d=324 e=341
108  a=411 b=123 c=243 d=227 e=341
109  a=411 b=184 c=168 d=288 e=341
110  a=435 b= 96 c=400 d=240 e=421
111  a=435 b=195 c= 35 d=339 e=421
112  a=450 b= 92 c=438 d=249 e=451
113  a=450 b=201 c= 12 d=358 e=451
114  a=468 b=173 c= 24 d=276 e=431
115  a=468 b=192 c=444 d=295 e=431
116  a=480 b= 80 c= 75 d=144 e=421
117  a=480 b=336 c=405 d=400 e=421
118  a=486 b=180 c= 18 d=287 e=451
119  a=486 b=199 c=468 d=306 e=451
120  a=488 b= 72 c= 15 d= 96 e=451
121  a=488 b=132 c=423 d=276 e=451
122  a=488 b=152 c=269 d=272 e=401
123  a=488 b=212 c= 65 d=356 e=451
124  a=488 b=216 c=219 d=336 e=401
125  a=488 b=392 c=473 d=416 e=451
126  a=495 b=135 c=415 d=279 e=451
127  a=495 b=216 c= 80 d=360 e=451
--- PASS: TestPentagons (31.58s)
PASS
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
 22  a= 29 b=  3 c= 17 d= 26 e= 28
 23  a= 31 b=  1 c= 11 d= 30 e= 28
 24  a= 31 b=  1 c= 19 d= 30 e= 28
 25  a= 31 b=  4 c=  7 d= 27 e= 31
 26  a= 31 b=  4 c= 20 d= 27 e= 31
 27  a= 32 b=  3 c=  8 d= 29 e= 31
 28  a= 32 b=  3 c= 21 d= 29 e= 31
 29  a= 33 b=  2 c=  9 d= 31 e= 31
 30  a= 33 b=  2 c= 22 d= 31 e= 31
 31  a= 34 b=  1 c= 10 d= 33 e= 31
 32  a= 34 b=  1 c= 23 d= 33 e= 31
 33  a= 36 b=  4 c= 11 d= 32 e= 35
 34  a= 36 b=  4 c= 21 d= 32 e= 35
 35  a= 37 b=  3 c=  4 d= 34 e= 37
 36  a= 37 b=  3 c= 12 d= 34 e= 35
 37  a= 37 b=  3 c= 22 d= 34 e= 35
 38  a= 37 b=  3 c= 30 d= 34 e= 37
 39  a= 38 b=  2 c=  5 d= 36 e= 37
 40  a= 38 b=  2 c= 13 d= 36 e= 35
 41  a= 38 b=  2 c= 23 d= 36 e= 35
 42  a= 38 b=  2 c= 31 d= 36 e= 37
 43  a= 39 b=  1 c=  6 d= 38 e= 37
 44  a= 39 b=  1 c= 14 d= 38 e= 35
 45  a= 39 b=  1 c= 24 d= 38 e= 35
 46  a= 39 b=  1 c= 32 d= 38 e= 37
 47  a= 39 b=  3 c=  7 d= 36 e= 38
 48  a= 39 b=  3 c= 29 d= 36 e= 38
 49  a= 40 b=  5 c= 16 d= 35 e= 39
 50  a= 40 b=  5 c= 19 d= 35 e= 39
 51  a= 41 b=  1 c=  9 d= 40 e= 38
 52  a= 41 b=  1 c= 31 d= 40 e= 38
 53  a= 41 b=  4 c= 17 d= 37 e= 39
 54  a= 41 b=  4 c= 20 d= 37 e= 39
 55  a= 43 b=  2 c= 19 d= 41 e= 39
 56  a= 43 b=  2 c= 22 d= 41 e= 39
 57  a= 43 b=  5 c=  8 d= 38 e= 43
 58  a= 43 b=  5 c= 13 d= 38 e= 42
 59  a= 43 b=  5 c= 25 d= 38 e= 42
 60  a= 43 b=  5 c= 30 d= 38 e= 43
 61  a= 44 b=  1 c= 20 d= 43 e= 39
 62  a= 44 b=  1 c= 23 d= 43 e= 39
 63  a= 44 b=  4 c=  9 d= 40 e= 43
 64  a= 44 b=  4 c= 31 d= 40 e= 43
 65  a= 45 b=  3 c= 10 d= 42 e= 43
 66  a= 45 b=  3 c= 32 d= 42 e= 43
 67  a= 46 b=  2 c= 11 d= 44 e= 43
 68  a= 46 b=  2 c= 33 d= 44 e= 43
 69  a= 47 b=  1 c= 12 d= 46 e= 43
 70  a= 47 b=  1 c= 17 d= 46 e= 42
 71  a= 47 b=  1 c= 29 d= 46 e= 42
 72  a= 47 b=  1 c= 34 d= 46 e= 43
 73  a= 49 b=  6 c= 10 d= 43 e= 49
 74  a= 49 b=  6 c= 33 d= 43 e= 49
 75  a= 50 b=  5 c= 11 d= 45 e= 49
 76  a= 50 b=  5 c= 34 d= 45 e= 49
 77  a= 50 b=  6 c= 15 d= 44 e= 49
 78  a= 50 b=  6 c= 29 d= 44 e= 49
 79  a= 51 b=  4 c= 12 d= 47 e= 49
 80  a= 51 b=  4 c= 35 d= 47 e= 49
 81  a= 51 b=  5 c= 16 d= 46 e= 49
 82  a= 51 b=  5 c= 30 d= 46 e= 49
 83  a= 52 b=  3 c= 13 d= 49 e= 49
 84  a= 52 b=  3 c= 36 d= 49 e= 49
 85  a= 52 b=  4 c= 17 d= 48 e= 49
 86  a= 52 b=  4 c= 31 d= 48 e= 49
 87  a= 53 b=  2 c= 14 d= 51 e= 49
 88  a= 53 b=  2 c= 37 d= 51 e= 49
 89  a= 53 b=  3 c= 18 d= 50 e= 49
 90  a= 53 b=  3 c= 32 d= 50 e= 49
 91  a= 53 b=  7 c= 21 d= 46 e= 52
 92  a= 53 b=  7 c= 25 d= 46 e= 52
 93  a= 54 b=  1 c= 15 d= 53 e= 49
 94  a= 54 b=  1 c= 38 d= 53 e= 49
 95  a= 54 b=  2 c= 19 d= 52 e= 49
 96  a= 54 b=  2 c= 33 d= 52 e= 49
 97  a= 55 b=  1 c= 20 d= 54 e= 49
 98  a= 55 b=  1 c= 34 d= 54 e= 49
 99  a= 55 b=  5 c= 23 d= 50 e= 52
100  a= 55 b=  5 c= 27 d= 50 e= 52
101  a= 57 b=  3 c= 25 d= 54 e= 52
102  a= 57 b=  3 c= 29 d= 54 e= 52
103  a= 57 b=  7 c= 17 d= 50 e= 56
104  a= 57 b=  7 c= 33 d= 50 e= 56
105  a= 58 b=  5 c= 10 d= 53 e= 57
106  a= 58 b=  5 c= 43 d= 53 e= 57
107  a= 59 b=  1 c= 27 d= 58 e= 52
108  a= 59 b=  1 c= 31 d= 58 e= 52
109  a= 59 b=  4 c= 11 d= 55 e= 57
110  a= 59 b=  4 c= 44 d= 55 e= 57
111  a= 59 b=  5 c= 19 d= 54 e= 56
112  a= 59 b=  5 c= 35 d= 54 e= 56
--- PASS: TestHexagonsNice (0.01s)
*/

func TestOctagons(t *testing.T) {
    tri := NewTriples()
    tri.Octagons(100)
    // Conjecture: No possible octagons for triple frame!
}

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

