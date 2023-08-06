package alg

import (
	"testing"
)

func TestQ32(t *testing.T) {

	factory := NewN32s() // with primes for reductions

	f := func(p ...Z) []Z { return p }

	qs := NewQ32s(factory)

	for _, s := range []struct { nums []Z ; den N; exp string; } {
		{ f(),  1, "err" },
		{ f(0), 0, "err" },
		{ f(0), 1, "0"   },
		{ f(4), 2, "2"   },
		{ f(1), 2, "1/2" },

		{ f(1,2),   2, "err" },
		{ f(0,1,1), 1, "1"   }, // (0+1√1)/1 = (1)/1 = 1
		{ f(1,1,1), 1, "2"   }, // (1+1√1)/1 = (1+1)/1 = 2
		{ f(1,1,2), 1, "1+1√2" }, // (1+1√2)/1 = (1+√2)/1 = 2
	} {
		if q, err := qs.newQ32(s.den, s.nums...); err != nil {
			if s.exp != "err" {
				t.Fatalf("reduceQ unexpected overflow for %d %v", s.den, s.nums)
			}
		} else if got := q.String(); got != s.exp {
			t.Fatalf("qs.new got %s exp %s", got, s.exp)
		}
	}
}