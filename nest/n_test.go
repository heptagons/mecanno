package nest

import (
	"fmt"
	"testing"
)

func TestN32s(t *testing.T) {
	a := NatGCD(8,12)
	if a != 4 {
		t.Fatalf("a != 4")
	}
	factory := NewN32s() // with primes for reductions
	// primes
	if got, exp := len(factory.primes), 6542; got != exp {
		t.Fatalf("primes32 got %d exp:%d", got, exp)
	}
	for _, s := range []struct { pos int; prime N32 } {
		{ pos:     0, prime:      2 },
		{ pos:     1, prime:      3 },
		{ pos:     2, prime:      5 },
		{ pos:     3, prime:      7 },
		{ pos:    10, prime:     31 },
		{ pos:   100, prime:    547 },
		{ pos: 1_000, prime:  7_927 },
		{ pos: 6_541, prime: 65_521 },
	} {
		if got, exp := factory.primes[s.pos], s.prime; got != exp {
			t.Fatalf("primes32 pos=%d got=%d exp=%d", s.pos, got, exp)
		}
	}

	pow2s := [][]N { // first pow2s                         // floor-sqrs
		[]N { 
			0, 1, 4, 9, 16, 25, 36, 49,                     //  0, 1, 2, 3, 4, 5, 6, 7
			64, 81, 100, 121, 144, 169, 196, 225,           //  8, 9,10,11,12,13,14,15
		},
		[]N {
			256, 289, 324, 361, 400, 441, 484, 529,         // 16,17,18,19,20,21,22,23
			576, 625, 676, 729, 784, 841, 900, 961,         // 24,25,26,27,28,29,30,31
			1024, 1089, 1156, 1225, 1296, 1369, 1444, 1521, // 32,33,34,35,36,37,38,39
			1600, 1681, 1764, 1849, 1936, 2025, 2116, 2209, // 40,41,42,43,44,45,46,47
		},
	}
	for i, pow2 := range pow2s {
		for j, exp := range pow2 {
			if got := factory.pow2s[i][j]; got != exp {
				t.Fatalf("got %d exp %d i=%d j=%d", got, exp, i, j)
			}
		}
	}

	for _, r := range []struct { n N; exp string } {
		// first table
		{ 0,   "[0,0,0]"      },
		{ 1,   "[1,1,1]"      },
		{ 2,   "[1,1,4]"      },
		{ 3,   "[1,1,4]"      },
		{ 9,   "[3,9,9]"      }, // second midle square down (1/4)
		{ 13,  "[3,9,16]"     },
		{ 49,  "[7,49,49]"    }, // first midle square (1/2)
		{ 121, "[11,121,121]" }, // second midle square up (3/4)
		{ 133, "[11,121,144]" }, // non-square
		{ 134, "[11,121,144]" }, // non-square
		{ 145, "[12,144,169]" }, // non-square
		{ 197, "[14,196,225]" }, // non-square
		{ 224, "[14,196,225]" }, // non-square
		{ 225, "[15,225,225]" }, // non-square max

		// second table
		{ 226,  "[15,225,256]"   }, // floor from table-1, ceil from table-2
		{ 255,  "[15,225,256]"   }, // floor from table-1, ceil from table-2
		{ 289,  "[17,289,289]"   }, // square
		{ 300,  "[17,289,324]"   }, // non-square
		{ 2209, "[47,2209,2209]" }, // second table max, ASAP response

		{      10_000, "[100,10000,10000]"       },
		{     100_000, "[316,99856,100489]"      },
		{   1_000_000, "[1000,1000000,1000000]"  },
		{  10_000_000, "[3162,9998244,10004569]" },
		{ 100_000_000, "Overflow" },
	} {
		if sqrt, floor, ceil, err := factory.nSqrtFloor(r.n); err != nil {
			if got := r.exp; got != err.Error() {
				t.Fatalf("nSqrtFloor num %d got %s exp %s", r.n, got, err.Error())
			}
		} else if got := fmt.Sprintf("[%d,%d,%d]", sqrt, floor, ceil); got != r.exp {
			t.Fatalf("nSqrtFloor num %d got %s exp %s", r.n, got, r.exp)
		} else {
			t.Logf("num %d got %s", r.n, got)
		}
	}
}

