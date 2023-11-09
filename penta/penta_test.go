package penta

import (
	"fmt"
	"math"
	"testing"

	"github.com/heptagons/meccano"
)

func Test_Type_1(t *testing.T) {
	sols := Type_1(100)
	exp := [][]int{
		[]int{ 12,3,4,11 },
	}
	if err := sols.Compare(exp); err != nil {
		t.Fatal(err)
	}
}

func Test_Type_2(t *testing.T) {
	sols := Type_2(100)
	exp := [][]int{ // a, b, c, d, e
		[]int{ 12,  2,  9,  6, 11 },
		[]int{ 12,  3,  0,  4, 11 },
		[]int{ 12,  6,  3, 10, 11 },
		[]int{ 31,  4, 28, 16, 31 },
		[]int{ 31, 15,  3, 27, 31 },
		[]int{ 38, 12, 18, 21, 31 },
		[]int{ 38, 17, 20, 26, 31 },
		[]int{ 48,  8, 24, 21, 41 },
		[]int{ 48, 12,  9, 20, 41 },
		[]int{ 48, 27, 24, 40, 41 },
		[]int{ 48, 28, 39, 36, 41 },
		[]int{ 72, 21, 48, 40, 61 },
		[]int{ 72, 24, 16, 39, 61 },
		[]int{ 72, 32, 24, 51, 61 },
		[]int{ 72, 33, 56, 48, 61 },
		[]int{ 78, 27,  4, 42, 71 },
		[]int{ 78, 36, 74, 51, 71 },
		[]int{ 87, 28, 36, 48, 71 },
		[]int{ 87, 39, 51, 59, 71 },
	}
	if err := sols.Compare(exp); err != nil {
		t.Fatal(err)
	}
}

func Test_Type_2_Half(t *testing.T) {
	sols := Type_2_Half(100)
	exp := [][]int { // a,b,c,d,e
		[]int{ 12,  2,  9,  6, 11 },
		[]int{ 12,  3,  0,  4, 11 },
		[]int{ 31,  4, 28, 16, 31 },
		[]int{ 38, 12, 18, 21, 31 },
		[]int{ 48,  8, 24, 21, 41 },
		[]int{ 48, 12,  9, 20, 41 },
		[]int{ 72, 21, 48, 40, 61 },
		[]int{ 72, 24, 16, 39, 61 },
		[]int{ 78, 27,  4, 42, 71 },
		[]int{ 87, 28, 36, 48, 71 },
	}
	if err := sols.Compare(exp); err != nil {
		t.Fatal(err)
	}
}


func Test_Type_2_HalfWithConjecture(t *testing.T) {
	sols := Type_2_HalfWithConjecture(100)
	exp := [][]int { // a,b,c,d,e
		[]int{ 12,  2,  9,  6, 11 },
		[]int{ 12,  3,  0,  4, 11 },
		[]int{ 31,  4, 28, 16, 31 },
		[]int{ 38, 12, 18, 21, 31 },
		[]int{ 48,  8, 24, 21, 41 },
		[]int{ 48, 12,  9, 20, 41 },
		[]int{ 72, 21, 48, 40, 61 },
		[]int{ 72, 24, 16, 39, 61 },
		[]int{ 78, 27,  4, 42, 71 },
		[]int{ 87, 28, 36, 48, 71 },
	}
	if err := sols.Compare(exp); err != nil {
		t.Fatal(err)
	}
}
/*
  1  a= 12 b=  2 c=  9 d=  6 e= 11
  2  a= 12 b=  3 c=  0 d=  4 e= 11
  3  a= 31 b=  4 c= 28 d= 16 e= 31
  4  a= 38 b= 12 c= 18 d= 21 e= 31
  5  a= 48 b=  8 c= 24 d= 21 e= 41
  6  a= 48 b= 12 c=  9 d= 20 e= 41
  7  a= 72 b= 21 c= 48 d= 40 e= 61
  8  a= 72 b= 24 c= 16 d= 39 e= 61
  9  a= 78 b= 27 c=  4 d= 42 e= 71
 10  a= 87 b= 28 c= 36 d= 48 e= 71
 11  a=111 b= 39 c= 99 d= 67 e=101
 12  a=121 b= 33 c= 33 d= 57 e=101
 13  a=128 b=  8 c= 89 d= 56 e=121
 14  a=138 b= 12 c= 54 d= 47 e=121
 15  a=145 b= 45 c= 39 d= 75 e=121
 16  a=147 b= 43 c= 51 d= 75 e=121
 17  a=151 b= 19 c= 73 d= 61 e=131
 18  a=156 b= 43 c= 96 d= 84 e=131
 19  a=165 b= 36 c=132 d= 88 e=151
 20  a=179 b= 15 c=177 d= 93 e=191
 21  a=183 b= 66 c= 62 d=108 e=151
 22  a=201 b=  9 c= 13 d= 21 e=191
 23  a=204 b= 21 c=112 d= 84 e=181
 24  a=216 b= 48 c=111 d=104 e=181
 25  a=236 b= 80 c= 20 d=125 e=211
 26  a=249 b= 45 c= 75 d= 95 e=211
 27  a=264 b= 76 c=  3 d=108 e=241
 28  a=285 b= 73 c= 27 d=111 e=251
 29  a=296 b=104 c=128 d=173 e=241
 30  a=303 b= 51 c= 29 d= 81 e=271
 31  a=304 b= 76 c=133 d=148 e=251
 32  a=312 b= 36 c= 93 d=100 e=271
 33  a=315 b= 24 c=160 d=120 e=281
 34  a=324 b= 64 c=204 d=159 e=281
 35  a=343 b=  7 c=115 d= 91 e=311
 36  a=352 b=  3 c=240 d=144 e=341
 37  a=354 b= 53 c= 60 d=102 e=311
 38  a=368 b= 36 c=219 d=156 e=331
 39  a=369 b= 37 c= 27 d= 63 e=341
 40  a=370 b=  1 c=172 d=118 e=341
 41  a=375 b= 15 c=191 d=135 e=341
 42  a=378 b= 21 c= 84 d= 86 e=341
 43  a=384 b=120 c=312 d=223 e=341
 44  a=390 b= 84 c= 50 d=135 e=341
 45  a=390 b= 87 c=228 d=194 e=331
 46  a=392 b=119 c=296 d=224 e=341
 47  a=392 b=128 c= 56 d=203 e=341
 48  a=393 b= 98 c= 54 d=156 e=341
 49  a=396 b=138 c= 73 d=222 e=341
 50  a=399 b= 70 c=210 d=180 e=341
 51  a=403 b= 78 c=114 d=156 e=341
 52  a=404 b= 89 c=104 d=164 e=341
 53  a=408 b= 16 c=312 d=183 e=401
 54  a=408 b= 84 c=167 d=180 e=341
 55  a=411 b=123 c=243 d=227 e=341
 56  a=435 b= 96 c=400 d=240 e=421
 57  a=450 b= 92 c=438 d=249 e=451
 58  a=468 b=173 c= 24 d=276 e=431
 59  a=480 b= 80 c= 75 d=144 e=421
 60  a=486 b=180 c= 18 d=287 e=451
 61  a=488 b= 72 c= 15 d= 96 e=451
 62  a=488 b=132 c=423 d=276 e=451
 63  a=488 b=152 c=269 d=272 e=401
 64  a=495 b=135 c=415 d=279 e=451
 65  a=502 b= 93 c= 36 d=138 e=451
 66  a=507 b= 18 c=366 d=220 e=491
 67  a=507 b= 60 c= 84 d=128 e=451
 68  a=509 b=150 c= 42 d=228 e=451
 69  a=516 b=114 c=169 d=222 e=431
 70  a=520 b= 36 c=225 d=180 e=461
 71  a=525 b=185 c=399 d=315 e=451
 72  a=525 b=189 c=105 d=305 e=451
 73  a=528 b= 80 c=171 d=192 e=451
 74  a=540 b=150 c=321 d=290 e=451
 75  a=543 b=123 c=221 d=249 e=451
 76  a=546 b=135 c=228 d=262 e=451
 77  a=552 b=179 c=288 d=312 e=451
 78  a=553 b=180 c=276 d=312 e=451
 79  a=560 b=200 c=344 d=335 e=461
 80  a=565 b= 69 c=153 d=177 e=491
 81  a=588 b=104 c= 12 d=135 e=541
 82  a=600 b= 65 c=240 d=216 e=521
 83  a=600 b=120 c= 96 d=205 e=521
 84  a=617 b= 89 c=533 d=317 e=601
 85  a=632 b=113 c=152 d=224 e=541
 86  a=652 b= 58 c=235 d=214 e=571
 87  a=661 b=109 c= 37 d=157 e=601
 88  a=684 b=237 c=192 d=388 e=571
 89  a=699 b= 84 c=564 d=344 e=671
 90  a=701 b=254 c=698 d=428 e=671
 91  a=713 b=234 c=582 d=420 e=631
 92  a=715 b=211 c=655 d=415 e=671
 93  a=720 b=216 c=712 d=423 e=701
 94  a=724 b=147 c= 72 d=228 e=641
 95  a=728 b= 21 c=192 d=168 e=661
 96  a=729 b= 36 c=428 d=288 e=671
 97  a=732 b= 18 c=681 d=358 e=781
 98  a=732 b= 42 c=111 d=134 e=671
 99  a=744 b=228 c=155 d=372 e=631
100  a=746 b=164 c= 38 d=233 e=671
101  a=755 b=123 c=267 d=291 e=641
102  a=756 b= 69 c=168 d=196 e=671
103  a=762 b= 73 c=372 d=294 e=671
104  a=765 b= 30 c=354 d=260 e=691
105  a=777 b=234 c=118 d=372 e=671
106  a=781 b=108 c=348 d=312 e=671
107  a=784 b=192 c=189 d=336 e=661
108  a=800 b=164 c=263 d=332 e=671
109  a=804 b=177 c=272 d=348 e=671
110  a=805 b=202 c=238 d=364 e=671
111  a=810 b=276 c=510 d=475 e=671
112  a=819 b=136 c=216 d=288 e=701
113  a=824 b=276 c=363 d=468 e=671
114  a=826 b=315 c=420 d=510 e=671
115  a=840 b=196 c=777 d=468 e=811
116  a=845 b=285 c=465 d=489 e=691
117  a=859 b=130 c=502 d=388 e=751
118  a=861 b=126 c= 66 d=196 e=781
119  a=863 b=303 c=711 d=519 e=761
120  a=864 b= 24 c=349 d=264 e=781
121  a=873 b=137 c=453 d=381 e=751
122  a=879 b=231 c= 63 d=343 e=781
123  a=885 b=206 c=642 d=468 e=781
124  a=885 b=309 c= 13 d=477 e=821
125  a=892 b=112 c=196 d=259 e=781
126  a=896 b=144 c=528 d=411 e=781
127  a=896 b=332 c=725 d=548 e=781
128  a=904 b=328 c=640 d=547 e=761
129  a=905 b=161 c=185 d=305 e=781
130  a=912 b=168 c=507 d=424 e=781
131  a=915 b=135 c=345 d=349 e=781
132  a=928 b=319 c=232 d=520 e=781
133  a=938 b=252 c=270 d=441 e=781
134  a=947 b=306 c=558 d=540 e=781
135  a=948 b=342 c=589 d=570 e=781
136  a=949 b=273 c=495 d=507 e=781
137  a=960 b=195 c=760 d=504 e=881
138  a=961 b=249 c=633 d=513 e=821
139  a=987 b=350 c=594 d=588 e=811
*/



// TestPentaAsymmDiagonalSlow looks integers solutions m,n of equation:
// for integers a,b:
//
//  a*b = 4*m*n
//	a*a + b*b = m*m + 2*m*n + 5*n*n, 
//
func TestPentaAsymmDiagonalSlow(t *testing.T) {
	max := 1000
	for a := 1; a < max; a++ {
		for b := 1; b < a; b++ { // reject already diagonals complete (a==b)
			ab := a*b
			if ab % 4 != 0 {
				continue
			}
			for m := 1; m < 2*a; m++ {
				for n :=1; n < 2*a; n++ {
					if ab == 4*m*n {
						if a*a + b*b == m*m + 2*m*n + 5*n*n {
							eleven := a % 11 == 0
							t.Logf("a=%d b=%d -> m=%d n=%d eleven=%t\n", a, b, m, n, eleven)
						}
					}
				}
			}
		}
	}
}

func TestPentaAsymmDiagonalFast(t *testing.T) {
	max := 4000
	sols := &meccano.Sols{}
	for a := 1; a < max; a++ {
		for b := 1; b < a; b++ { // reject already diagonals complete (a==b)
			ab := a*b
			if ab % 4 != 0 {
				continue
			}
			mn := ab >> 2 // ab/4 as integer
			for m := 1; m <= mn; m++ {
				if mn % m == 0 {
					n := mn / m
					if a*a + b*b == m*m + 2*m*n + 5*n*n {
						sols.Add(a, b, m, n)
					}
				}
			}
		}
		if a % 100 == 0 {
			fmt.Printf("a=%d/%d\n", a, max)
		}
	}
}
/*
=== RUN   TestPentaAsymmDiagonalFast
  1  a= 11 b=  8 c= 11 d=  2
a=100/4000
a=200/4000
  2  a=246 b= 70 c= 41 d=105
a=300/4000
a=400/4000
a=500/4000
a=600/4000
a=700/4000
a=800/4000
a=900/4000
a=1000/4000
a=1100/4000
a=1200/4000
a=1300/4000
a=1400/4000
a=1500/4000
a=1600/4000
a=1700/4000
a=1800/4000
panic: test timed out after 10m0s
*/

func TestPentaAsymmDiagonal2(t *testing.T) {
	max := 2000
	sols := &meccano.Sols{}
	for a := 1; a < max; a++ {
		aa := a*a
		if aa % 4 != 0 {
			continue
		}
		mn := aa >> 2 // a*a/4 as integer
		for b := 1; b <= a/2; b++ {
			for m := 1; m <= mn; m++ {
				if mn % m == 0 {
					n := mn / m
					if b*b - a*b == m*m - 6*m*n + 5*n*n {
						sols.Add(a, b, m, n)
					}
				}
			}
		}
		if a % 100 == 0 {
			fmt.Printf("a=%d/%d\n", a, max)
		}
	}
}
/*
=== RUN   TestPentaAsymmDiagonal2
  1  a=  4 b=  1 c=  4 d=  1
a=100/3000
a=200/3000
a=300/3000
a=400/3000
a=500/3000
a=600/3000
a=700/3000
a=800/3000
a=900/3000
a=1000/3000
a=1100/3000
a=1200/3000
a=1300/3000
a=1400/3000
a=1500/3000
a=1600/3000
a=1700/3000
a=1800/3000
panic: test timed out after 10m0s
*/

// Test complicate e formula in penta-diagonals.pdf
// in section "Regular polygon diagonal e"
// with known pentagon height
// 
func TestDiagonals(t *testing.T) {

	a := 1.0
	b := a
	c := a/2

	// expected
	expD    := a*(1+math.Sqrt(5))/2
	expCosA := math.Cos(math.Pi/5) // 36 degrees
	expSinA := math.Sin(math.Pi/5) // 36 degrees
	expE    := math.Sqrt(5 + 2*math.Sqrt(5))/2 // pentagons side=1 height
	expCosB := math.Cos(2*math.Pi/5) // 72 degrees
	expF    := (a*a + b*b + c*c - expE*expE)/2


	u := (1 - math.Sqrt(5))/4
	u2 := u*u
	//u3 := u*u*u
	//u4 := u*u*u*u

	d := math.Sqrt(a*a + b*b - 2*a*b*u)

	t.Logf("a=%f b=%f c=%f u=%+f", a, b, c, u)
	t.Logf("   exp d = %+f", expD)
	t.Logf("   got d = %+f", d)
	t.Log()

	// cosine from quadratic equation AX^2 + BX + C = 0 where
	// A = 1
	// B = -2*u*cosA
	// C := u^2 - sin^2A
	cosB1 := u*expCosA + math.Sqrt(u2*expCosA*expCosA - u2 + expSinA*expSinA)


	cosA  := (a - b*u)/d
	cosB2 := (a*a + b*b + c*c - expE*expE - 2*a*b*u)/(2*c*d)

	t.Logf("exp cosA = %+f", expCosA)
	t.Logf("    cosA = %+f", cosA)
	t.Logf("exp cosB = %+f", expCosB)
	t.Logf("exp cosB1= %+f", cosB1)
	t.Logf("got cosB = %+f", cosB2)
	t.Log()

	// Section "Regular polygon diagonal e"
	// where we define m,n to have a simpler f
	// We find that for f when we choose positive from plus/minus
	// we match expectedE (?)
	m := a*(b+c)*u - b*c*u2
	n := math.Abs(b*(1-u2))
	f1 := m + c*n
	f2 := m - c*n
	e := math.Sqrt(a*a + b*b + c*c - 2*m - 2*c*n)

	f5 := (b*c + a*(b+c) + (b*c - a*(b+c))*math.Sqrt(5))/4

	t.Logf("   exp f = %+f", expF)
	t.Logf("   got f1= %+f", f1)
	t.Logf("   got f2= %+f", f2)
	t.Logf("   exp e = %+f", expE)
	t.Logf("   got e = %+f", e)
	t.Logf("   got f5= %+f", f5)









	//t.Logf("         m = %+f", m)
	//t.Logf("         n = %+f", n)
	//t.Logf("         f = %+f", f)


}