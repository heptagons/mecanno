package nest

import (
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
}
