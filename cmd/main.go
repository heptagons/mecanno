package main

import (
	"fmt"
	"math"
)

func main() {
	pentagons_type_2_half_with_conjecture(1000)
	//pentagons_type_2_half(1000)
	//octagons_2(60)
	//triangle_diagonals(200)
}

type Sols struct {
	sols [][]int
}

func (s *Sols) Add(rods ...int) {
	if len(rods) < 0 {
		return
	}
	const RODS = "abcdefhijkl"
	for _, s := range s.sols {
		a := rods[0]
		if a % s[0] != 0 { 
			continue
		}
		// new a is a factor of previous a
		f := a / s[0]
		cont := false
		for r := 1; r < len(rods); r++ {
			b := rods[r]
			if t := b % s[r] == 0 && b / s[r] == f; !t {
				cont = true
				break
			}
		}
		if cont {
			continue // scaled solution already found (reject)
		}
		return
	}
	// solution!
	if s.sols == nil {
		s.sols = make([][]int, 0)
	}
	s.sols = append(s.sols, rods)
	fmt.Printf("%3d ", len(s.sols))
	for i, r := range rods {
		fmt.Printf(" %c=%3d", RODS[i], r)
	}
	fmt.Println()
}

func pentagons_type_1(max int) {

	sols := &Sols{}

	check := func(a, b, c int) {
		f := float64(a*a + b*b + c*c - a*c)
		if f < 0 {
			return
		}
		if d := int(math.Sqrt(f)); math.Pow(float64(d), 2) == f {
			sols.Add(a, b, c, d)
		}
	}

	for a := 1; a < max; a++ {
		for b := 1; b <= a; b++ {
			for c := 0; c <= a; c++ {
				if a*c == (a + c)*b {
					check(a, b, c)
				}
			}
		}
	}
}
/*
a=12 b=3 c=4 d=11
*/

func pentagons_type_2(max int) {

	sols := &Sols{}

	check := func(a, b, c, d int) {
		//f := float64(a*a + b*b + c*c + d*d - a*d - b*c - c*d)
		f := float64(a*a + b*b + c*c + d*d - a*c - b*d - a*b)
	    if f < 0 {
	    	return
	    }
		if e := int(math.Sqrt(f)); math.Pow(float64(e), 2) == f {
			sols.Add(a, b, c, d, e)
		}
	}

    for a := 1 ; a < max; a++ {
    	for b := 1; b < a; b++ {
        	for c := 1; c < a; c++ {
          		for d := 1; d < a; d++ {
            		if ((a - b)*(c - d) + a*b == c*d) {
              			check(a, b, c, d)
              		}
              	}
            }
        }
    }
}/*
  1 a= 12 b=  2 c=  9 d=  6 e= 11
  2 a= 12 b=  6 c=  3 d= 10 e= 11
  3 a= 31 b=  4 c= 28 d= 16 e= 31
  4 a= 31 b= 15 c=  3 d= 27 e= 31
  5 a= 38 b= 12 c= 18 d= 21 e= 31
  6 a= 38 b= 17 c= 20 d= 26 e= 31
  7 a= 48 b=  8 c= 24 d= 21 e= 41
  8 a= 48 b= 12 c=  9 d= 20 e= 41
  9 a= 48 b= 27 c= 24 d= 40 e= 41
 10 a= 48 b= 28 c= 39 d= 36 e= 41
 11 a= 72 b= 21 c= 48 d= 40 e= 61
 12 a= 72 b= 24 c= 16 d= 39 e= 61
 13 a= 72 b= 32 c= 24 d= 51 e= 61
 14 a= 72 b= 33 c= 56 d= 48 e= 61
 15 a= 78 b= 27 c=  4 d= 42 e= 71
 16 a= 78 b= 36 c= 74 d= 51 e= 71
 17 a= 87 b= 28 c= 36 d= 48 e= 71
 18 a= 87 b= 39 c= 51 d= 59 e= 71
*/

// this function fails to find what pentagons_type_2 find (roudoff errors)
func pentagons_type_2b(max int) {
	s := &Sols{}
	cosA, sinA := math.Cos(2*math.Pi/5), math.Sin(2*math.Pi/5)
	cosB, sinB := math.Cos(  math.Pi/5), math.Sin(  math.Pi/5)
	for a := 1; a <= max; a++ {
		ax, ay := float64(a)*cosA, float64(a)*sinA
		for b := 1; b < a; b++ {
			bx, by := float64(b)*cosB, float64(b)*sinB
			for d := 1; d < (a-b); d++ {
				dx, dy := float64(d)*cosB, float64(d)*sinB
				for c := 1; c < a; c++ {
					cx := float64(c)
					ex := cx + ax - bx - dx
					ey := ay + by - dy
					f := ex*ex + ey*ey
					e := int(math.Sqrt(f))
					if (math.Pow(float64(e), 2) - f) == 0 {
						s.Add(a, b, c, d, e)
					}
				}
			}
		}
	}
}

func pentagons_type_2_half(max int) {
	sols := &Sols{}
	aa, a_b, ab, bb, dd, ad, bc, c_d, cd, cc := 0,0,0,0,0,0,0,0,0,0
	for a := 1; a <= max; a++ {
		aa = a*a
		for b := 1; b < a; b++ {
			a_b, ab, bb = a - b, a*b, b*b
			for d := 1; d < (a-b); d++ {
				dd, ad = d*d, a*d
				for c := 1; c < a; c++ {
					bc, c_d, cd, cc = b*c, c - d, c*d, c*c
					if a_b * c_d + ab == cd {
						if f := float64(aa + bb + cc + dd - ad - bc - cd); f > 0 {
							if e := int(math.Sqrt(f)); math.Pow(float64(e), 2) == f {
								sols.Add(a, b, c, d, e)
							}
						}
					}
				}
			}
		}
	}
}
/*
  1 a= 12 b=  2 c=  9 d=  6 e= 11
  2 a= 31 b=  4 c= 28 d= 16 e= 31
  3 a= 38 b= 12 c= 18 d= 21 e= 31
  4 a= 48 b=  8 c= 24 d= 21 e= 41
  5 a= 48 b= 12 c=  9 d= 20 e= 41
  6 a= 72 b= 21 c= 48 d= 40 e= 61
  7 a= 72 b= 24 c= 16 d= 39 e= 61
  8 a= 78 b= 27 c=  4 d= 42 e= 71
  9 a= 87 b= 28 c= 36 d= 48 e= 71
 10 a=111 b= 39 c= 99 d= 67 e=101
 11 a=121 b= 33 c= 33 d= 57 e=101
 12 a=128 b=  8 c= 89 d= 56 e=121
 13 a=138 b= 12 c= 54 d= 47 e=121
 14 a=145 b= 45 c= 39 d= 75 e=121
 15 a=147 b= 43 c= 51 d= 75 e=121
 16 a=151 b= 19 c= 73 d= 61 e=131
 17 a=156 b= 43 c= 96 d= 84 e=131
 18 a=165 b= 36 c=132 d= 88 e=151
 19 a=179 b= 15 c=177 d= 93 e=191
 20 a=183 b= 66 c= 62 d=108 e=151
 21 a=201 b=  9 c= 13 d= 21 e=191
 22 a=204 b= 21 c=112 d= 84 e=181
 23 a=216 b= 48 c=111 d=104 e=181
 24 a=236 b= 80 c= 20 d=125 e=211
 25 a=249 b= 45 c= 75 d= 95 e=211
 26 a=264 b= 76 c=  3 d=108 e=241
 27 a=285 b= 73 c= 27 d=111 e=251
 28 a=296 b=104 c=128 d=173 e=241
 29 a=303 b= 51 c= 29 d= 81 e=271
 30 a=304 b= 76 c=133 d=148 e=251
 31 a=312 b= 36 c= 93 d=100 e=271
 32 a=315 b= 24 c=160 d=120 e=281
 33 a=324 b= 64 c=204 d=159 e=281
 34 a=343 b=  7 c=115 d= 91 e=311
 35 a=352 b=  3 c=240 d=144 e=341
 36 a=354 b= 53 c= 60 d=102 e=311
 37 a=368 b= 36 c=219 d=156 e=331
 38 a=369 b= 37 c= 27 d= 63 e=341
 39 a=370 b=  1 c=172 d=118 e=341
 40 a=375 b= 15 c=191 d=135 e=341
 41 a=378 b= 21 c= 84 d= 86 e=341
 42 a=384 b=120 c=312 d=223 e=341
 43 a=390 b= 84 c= 50 d=135 e=341
 44 a=390 b= 87 c=228 d=194 e=331
 45 a=392 b=119 c=296 d=224 e=341
 46 a=392 b=128 c= 56 d=203 e=341
 47 a=393 b= 98 c= 54 d=156 e=341
 48 a=396 b=138 c= 73 d=222 e=341
 49 a=399 b= 70 c=210 d=180 e=341
 50 a=403 b= 78 c=114 d=156 e=341
 51 a=404 b= 89 c=104 d=164 e=341
 52 a=408 b= 16 c=312 d=183 e=401
 53 a=408 b= 84 c=167 d=180 e=341
 54 a=411 b=123 c=243 d=227 e=341
 55 a=435 b= 96 c=400 d=240 e=421
 56 a=450 b= 92 c=438 d=249 e=451
 57 a=468 b=173 c= 24 d=276 e=431
 58 a=480 b= 80 c= 75 d=144 e=421
 59 a=486 b=180 c= 18 d=287 e=451
 60 a=488 b= 72 c= 15 d= 96 e=451
 61 a=488 b=132 c=423 d=276 e=451
 62 a=488 b=152 c=269 d=272 e=401
 63 a=495 b=135 c=415 d=279 e=451
 64 a=502 b= 93 c= 36 d=138 e=451
 65 a=507 b= 18 c=366 d=220 e=491
 66 a=507 b= 60 c= 84 d=128 e=451
 67 a=509 b=150 c= 42 d=228 e=451
 68 a=516 b=114 c=169 d=222 e=431
 69 a=520 b= 36 c=225 d=180 e=461
 70 a=525 b=185 c=399 d=315 e=451
 71 a=525 b=189 c=105 d=305 e=451
 72 a=528 b= 80 c=171 d=192 e=451
 73 a=540 b=150 c=321 d=290 e=451
 74 a=543 b=123 c=221 d=249 e=451
 75 a=546 b=135 c=228 d=262 e=451
 76 a=552 b=179 c=288 d=312 e=451
 77 a=553 b=180 c=276 d=312 e=451
 78 a=560 b=200 c=344 d=335 e=461
 79 a=565 b= 69 c=153 d=177 e=491
 80 a=588 b=104 c= 12 d=135 e=541
 81 a=600 b= 65 c=240 d=216 e=521
 82 a=600 b=120 c= 96 d=205 e=521
 83 a=617 b= 89 c=533 d=317 e=601
 84 a=632 b=113 c=152 d=224 e=541
 85 a=652 b= 58 c=235 d=214 e=571
 86 a=661 b=109 c= 37 d=157 e=601
 87 a=684 b=237 c=192 d=388 e=571
 88 a=699 b= 84 c=564 d=344 e=671
 89 a=701 b=254 c=698 d=428 e=671
 90 a=713 b=234 c=582 d=420 e=631
 91 a=715 b=211 c=655 d=415 e=671
 92 a=720 b=216 c=712 d=423 e=701
 93 a=724 b=147 c= 72 d=228 e=641
 94 a=728 b= 21 c=192 d=168 e=661
 95 a=729 b= 36 c=428 d=288 e=671
 96 a=732 b= 18 c=681 d=358 e=781
 97 a=732 b= 42 c=111 d=134 e=671
 98 a=744 b=228 c=155 d=372 e=631
 99 a=746 b=164 c= 38 d=233 e=671
100 a=755 b=123 c=267 d=291 e=641
101 a=756 b= 69 c=168 d=196 e=671
102 a=762 b= 73 c=372 d=294 e=671
103 a=765 b= 30 c=354 d=260 e=691
104 a=777 b=234 c=118 d=372 e=671
105 a=781 b=108 c=348 d=312 e=671
106 a=784 b=192 c=189 d=336 e=661
107 a=800 b=164 c=263 d=332 e=671
108 a=804 b=177 c=272 d=348 e=671
109 a=805 b=202 c=238 d=364 e=671
110 a=810 b=276 c=510 d=475 e=671
111 a=819 b=136 c=216 d=288 e=701
112 a=824 b=276 c=363 d=468 e=671
113 a=826 b=315 c=420 d=510 e=671
114 a=840 b=196 c=777 d=468 e=811
115 a=845 b=285 c=465 d=489 e=691
116 a=859 b=130 c=502 d=388 e=751
117 a=861 b=126 c= 66 d=196 e=781
118 a=863 b=303 c=711 d=519 e=761
119 a=864 b= 24 c=349 d=264 e=781
120 a=873 b=137 c=453 d=381 e=751
121 a=879 b=231 c= 63 d=343 e=781
122 a=885 b=206 c=642 d=468 e=781
123 a=885 b=309 c= 13 d=477 e=821
124 a=892 b=112 c=196 d=259 e=781
125 a=896 b=144 c=528 d=411 e=781
126 a=896 b=332 c=725 d=548 e=781
127 a=904 b=328 c=640 d=547 e=761
128 a=905 b=161 c=185 d=305 e=781
129 a=912 b=168 c=507 d=424 e=781
130 a=915 b=135 c=345 d=349 e=781
131 a=928 b=319 c=232 d=520 e=781
132 a=938 b=252 c=270 d=441 e=781
133 a=947 b=306 c=558 d=540 e=781
134 a=948 b=342 c=589 d=570 e=781
135 a=949 b=273 c=495 d=507 e=781
136 a=960 b=195 c=760 d=504 e=881
137 a=961 b=249 c=633 d=513 e=821
138 a=987 b=350 c=594 d=588 e=811
*/

func pentagons_type_2_half_with_conjecture(max int) {
	sols := &Sols{}
	aa, a_b, ab, bb, dd, ad, bc, c_d, cd, cc := 0,0,0,0,0,0,0,0,0,0
	for a := 1; a <= max; a++ {
		aa = a*a
		for b := 1; b < a; b++ {
			a_b, ab, bb = a - b, a*b, b*b
			for d := 1; d < (a-b); d++ {
				dd, ad = d*d, a*d
				for c := 1; c < a; c++ {
					bc, c_d, cd, cc = b*c, c - d, c*d, c*c
					if a_b * c_d + ab == cd {

						e2 := aa + bb + cc + dd - ad - bc - cd

						x := 1
						for {
							if e := 10*x + 1; e*e == e2 {
								sols.Add(a, b, c, d, e)
								break
							} else if e*e > e2 {
								break
							}
							x++
						}
					}
				}
			}
		}
	}
}



// triangle_diagonals finds the integer diagonals inside equilateral meccano
// triangles:
//       C      sizes AB = BD = CA, angle ABC = 60°
//      / \     diagonal AD=Math.sqrt((a-b)*(a-b) + ab)
//     /   \        where a = AB, b = BD
//    /   _ D
//   /_ -    \
//  A---------B
func triangle_diagonals(max int) {
	for a := 1; a < max; a++ {
		for b := 1; b <= a/2; b++ {
			if gcd(a, b) == 1 {
				diag := (a-b)*(a-b) + a*b
				cd := math.Sqrt(float64(diag))
				d := int(cd)
				if cd == float64(d) {
					num := float64(diag + a*a - b*b)
					den := 2.0 * cd * float64(a)
					angle := 180*math.Acos(num/den)/math.Pi
					fmt.Printf("a=%3d b=%3d d=%3d angle=%8.4f\n", a, b, d, angle)
				}
			}
		}
	}
	// a=  8 b=  3 d=  7 angle= 21.7868
	// a= 15 b=  7 d= 13 angle= 27.7958
	// a= 21 b=  5 d= 19 angle= 13.1736
	// a= 35 b= 11 d= 31 angle= 17.8966
	// a= 40 b=  7 d= 37 angle=  9.4300
	// a= 48 b= 13 d= 43 angle= 15.1782
	// a= 55 b= 16 d= 49 angle= 16.4264
	// a= 65 b=  9 d= 61 angle=  7.3410
	// a= 77 b= 32 d= 67 angle= 24.4327
	// a= 80 b= 17 d= 73 angle= 11.6351
	// a= 91 b= 40 d= 79 angle= 26.0078
	// a= 96 b= 11 d= 91 angle=  6.0090
	// a= 99 b= 19 d= 91 angle= 10.4174
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}

// Find bars forming octagon internal angle 135°.
//  
//  D . _ 
//  | \   ` - _ C
//  |   \       /
//  |     \   /
//  A-------B

//  Angles: ABD=45°, DBC=90°, ABC=135°

//  i = A-D = A-D, B-D = i*sqrt(2)
//  we iterate bc to find DC integer
//  
func octagons_1() {
	for ab := 1; ab <= 30; ab++ {
		h_2 := 2*ab*ab
		for bc := 1; bc <= 4*ab; bc++ {
			h_3 :=  float64(bc*bc + h_2)
			cd := math.Sqrt(h_3)
			_cd := int(cd)
			if cd == float64(_cd) {
				if gcd(_cd, gcd(ab, bc)) == 1 {
					fmt.Printf("✔ ab=%2d bc=%2d cd=%2d\n", ab, bc, _cd)
				}
			}
		}
	}
	// Expected
	// ✔ ab= 2 bc= 1 cd= 3
	// ✔ ab= 4 bc= 7 cd= 9
	// ✔ ab= 6 bc= 7 cd=11
	// ✔ ab= 6 bc=17 cd=19
	// ✔ ab=10 bc=23 cd=27
	// ✔ ab=12 bc= 1 cd=17
	// ✔ ab=14 bc=47 cd=51
	// ✔ ab=20 bc=17 cd=33
	// ✔ ab=24 bc=23 cd=41
	// ✔ ab=28 bc=41 cd=57
	// ✔ ab=30 bc= 7 cd=43
	// ✔ ab=30 bc=41 cd=59
}

// Octagons bars second method
// 1) i iterate cd=1,2,3,...
// 2) j iterate bc=1,2,3 < cd
// 3) calculate bd*bd = cd**cd - bc*bc
// 4) accept when bd*db = 2 * square
func octagons_2(max int) {
	for a := 1; a < max; a++ {
		for b := 1; b < a; b++ {
			cc := a*a - b*b
			if cc % 2 == 0 {
				f := math.Sqrt(float64(cc/2))
				c := int(f)
				if f == float64(c) {
					if gcd(c, gcd(b, a)) == 1 {
						s := int(math.Max(float64(b), f))
						fmt.Printf("a=%2d b=%2d c=%2d s=%2d\n", a, b, c, s)
					}
				}
			}
		}
	}
}