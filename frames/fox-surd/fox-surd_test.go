package surd

import (
	"fmt"
	"testing"
	"github.com/heptagons/meccano"
	. "github.com/heptagons/meccano/nest"
)

func TestFoxSurdPentaSimple(t *testing.T) {
	sols := meccano.Sols{}
	sols.Chars("nij")
	foxSurdPenta(1, 1000, false, func(n, i, j Z, s*A32) {
		sols.Add2(fmt.Sprintf("s=%s", s), int(n), int(i), int(j))
	})
}

func TestFoxSurdPentaNest(t *testing.T) {
	sols := meccano.Sols{}
	sols.Chars("nij")
	foxSurdPenta(1, 13, true, func(n, i, j Z, s*A32) {
		sols.Add2(fmt.Sprintf("s=%s", s), int(n), int(i), int(j))
	})
}

func TestFoxSurdPentaZero(t *testing.T) {
	c := 0
	foxSurdPenta(0, 12, false, func(n, i, j Z, s*A32) {
		c++
		fmt.Printf("% 3d) n=% 3d i=% 3d j=% 3d %s\n", c, n, i, j, s)
	})
}

func foxSurdPenta(min, max Z, nest bool, sol func(n,i,j Z, s *A32)) {

	factory := NewA32s()

	for n := Z(1); n <= max; n++ {
		nn := n*n
		for i := min; i <= n; i++ {
			ni := n*i
			ii := i*i
			for j := min; j <= n; j++ {
				if n == 1 && i == 0 && j == 0 {
					continue
				}
				nj := n*j
				jj := j*j
				ij := i*j

				u := 6*nn + 4*ii + 4*jj - 4*ni - 4*nj - 2*ij
				v := 2*nn - 4*ni + 2*ij

				if v == 0 {
					//	((b + c√d) / a) = (0 + 1√u)/2
					if s, err := factory.ANew3(2, 0, 1, u); err != nil {
						// silent 
					} else if min == 0 {
						if i == 0 || j == 0 {
							sol(n, i, j, s)
						}
					} else if nest == false {
						sol(n, i, j, s)
					}
				} else {
					//	((b + c√d + e√(f+g√h)) / a) = (√(u + v√5)) / 2
					if s, err := factory.ANew7(2, 0, 0, 0, 1, u, v, 5); err != nil {
						// silent
					} else if min == 0 {
						if i==0 || j==0 {
							sol(n, i, j, s)	
						}
					} else if nest == true {
						sol(n, i, j, s)
					}
				}
			}
		}
	}
}

/*
=== RUN   TestFoxSurdPentaSimple
   1)  n=  1 i=  1 j=  1 s=1
   2)  n=  6 i=  4 j=  3 s=√31
   3)  n= 12 i=  9 j=  8 s=11
   4)  n= 15 i=  9 j=  5 s=√211
   5)  n= 20 i= 16 j= 15 s=√341
   6)  n= 28 i= 16 j=  7 s=√781
   7)  n= 30 i= 25 j= 24 s=√781
   8)  n= 35 i= 25 j= 21 s=√1031
   9)  n= 40 i= 25 j= 16 s=√1441
  10)  n= 42 i= 36 j= 35 s=√1555
  11)  n= 45 i= 25 j=  9 s=√2101
  12)  n= 56 i= 49 j= 48 s=√2801
  13)  n= 63 i= 49 j= 45 s=√3355
  14)  n= 66 i= 36 j= 11 s=√4651
  15)  n= 70 i= 49 j= 40 s=√4141
  16)  n= 72 i= 64 j= 63 s=√4681
  17)  n= 77 i= 49 j= 33 s=√5261
  18)  n= 84 i= 49 j= 24 s=√6841
  19)  n= 88 i= 64 j= 55 s=√6505
  20)  n= 90 i= 81 j= 80 s=11√61
  21)  n= 91 i= 49 j= 13 s=√9031
  22)  n= 99 i= 81 j= 77 s=√8431
  23)  n=104 i= 64 j= 39 s=√9881
  24)  n=110 i=100 j= 99 s=√11111
  25)  n=117 i= 81 j= 65 s=√11605
  26)  n=120 i= 64 j= 15 s=√15961
  27)  n=126 i= 81 j= 56 s=√13981
  28)  n=130 i=100 j= 91 s=√14251
  29)  n=132 i=121 j=120 s=√16105
  30)  n=143 i=121 j=117 s=√17891
  31)  n=144 i= 81 j= 32 s=√21121
  32)  n=153 i= 81 j= 17 s=√26281
  33)  n=154 i=121 j=112 s=√20101
  34)  n=156 i=144 j=143 s=√22621
  35)  n=165 i=121 j=105 s=√22861
  36)  n=170 i=100 j= 51 s=√27731
  37)  n=176 i=121 j= 96 s=√26321
  38)  n=182 i=169 j=168 s=√30941
  39)  n=187 i=121 j= 85 s=√30655
  40)  n=190 i=100 j= 19 s=√40951
  41)  n=195 i=169 j=165 s=√33751
  42)  n=198 i=121 j= 72 s=√36061
  43)  n=204 i=144 j=119 s=√35101
  44)  n=208 i=169 j=160 s=√37105
  45)  n=209 i=121 j= 57 s=√42761
  46)  n=210 i=196 j=195 s=√41371
  47)  n=220 i=121 j= 40 s=√51001
  48)  n=221 i=169 j=153 s=√41141
  49)  n=228 i=144 j= 95 s=√46405
  50)  n=231 i=121 j= 21 s=√61051
  51)  n=234 i=169 j=144 s=√46021
  52)  n=238 i=196 j=187 s=√48871
  53)  n=240 i=225 j=224 s=√54241
  54)  n=247 i=169 j=133 s=√51931
  55)  n=255 i=225 j=221 s=√58411
  56)  n=260 i=169 j=120 s=√59081
  57)  n=266 i=196 j=171 s=11√491
  58)  n=272 i=256 j=255 s=√69905
  59)  n=273 i=169 j=105 s=√67705
  60)  n=276 i=144 j= 23 s=√87781
  61)  n=285 i=225 j=209 s=√68941
  62)  n=286 i=169 j= 88 s=√78061
  63)  n=299 i=169 j= 69 s=√90431
  64)  n=304 i=256 j=247 s=√80641
  65)  n=306 i=289 j=288 s=√88741
  66)  n=312 i=169 j= 48 s=√105121
  67)  n=322 i=196 j=115 s=√95755
  68)  n=323 i=289 j=285 s=√94655
  69)  n=325 i=169 j= 25 s=√122461
panic: test timed out after 10m0s
*/