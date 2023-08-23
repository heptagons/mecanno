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

	//t.Logf("pows2 %v", factory.pow2s)

	pow2s := [][]N { // first squares
		[]N { 
			0, 1, 4, 9, 16, 25, 36, 49, 
			64, 81, 100, 121, 144, 169, 196, 225,
		},
		[]N {
			256, 289, 324, 361, 400, 441, 484, 529,
			576, 625, 676, 729, 784, 841, 900, 961,
			1024, 1089, 1156, 1225, 1296, 1369, 1444, 1521,
			1600, 1681, 1764, 1849, 1936, 2025, 2116, 2209,
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
		{ 1_000_000, "Overflow" },

		{ 225,  "[225,225]"   }, // first table max, ASAP-1 response
		{ 2209, "[2209,2209]" }, // second table max, ASAP-1 response
		
		// first table
		{ 49, "[49,49]" }, // first midle square (1/2)

		{ 9,   "[9,9]"     }, // second midle square down (1/4)
		{ 121, "[121,121]" }, // second midle square up (3/4)

		{ 133, "[121,144]" }, // non-square


		//{ 197, "[196,225]" }, // higher cell range 197...224
		//{ 224, "[196,225]" }, // higher cell range 197...224

		//{ 144,       "[144,144]"   }, // first table square, ASAP-2 response

		//{ 0,         "[0,0]"       },
	} {
		if floor, ceil, err := factory.nSqrtFloorCeil(r.n); err != nil {
			if got := r.exp; got != err.Error() {
				t.Fatalf("got %s exp %s", got, err.Error())
			}
		} else if got := fmt.Sprintf("[%d,%d]", floor, ceil); got != r.exp {
			t.Fatalf("got %s exp %s", got, r.exp)
		}
	}
}

