package penta

import (
	"fmt"
	"math"
	"testing"

	"github.com/heptagons/meccano"
	"github.com/heptagons/meccano/nest"
	//"github.com/heptagons/meccano/frames"
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
 ...
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


	// test to get e = sqrt(46 + 18sqrt(5))/2
	min, max := 12, 12
	fmt.Printf("NewDiagonals min=%d max=%d\n", min, max)
	// we print rows to be copied to penta-diagonals table for pentagon size=3
	NewDiagonals().Get(min, max, func(a, b, c int, surd *nest.A32) {
		fmt.Printf("a=%d b=%d c=%d %s\n", a, b, c, surd.String())
	})

	//frames := frames.New()
	//fmt.Println("TwoTriangles")
	//frames.TwoTriangles(15, []int{ 46, 18, 5})
}

