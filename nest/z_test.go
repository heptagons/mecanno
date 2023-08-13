package nest

import (
	"fmt"
	"strings"
	"testing"
)

func TestZ32s(t *testing.T) {

	factory := NewZ32s()

	f := func(p ...Z) []Z { return p }

	// zFracN
	for _, s := range []struct { nums []Z ; den N; exp string; } {
		{ f( 0, 0, 0, 0), 1, "(+0+0+0+0)/1" },
		{ f( 1, 2, 3, 4), 5, "(+1+2+3+4)/5" },
		{ f(-8, 4, 2, 2), 2, "(-4+2+1+1)/1" },

		{ f(3*5*7*11),              77, "(+15)/1" },
		{ f(128, 64, 32, 16, 8, 4),  8, "(+32+16+8+4+2+1)/2" },
		{ f(-35, 0, -5, -15),       10, "(-7+0-1-3)/2" },

		{ f(3*3*3*2, 3*3*2*2, 3*2*2*2), 3*3*3*2*2*2, "(+9+6+4)/36" },
		// with zeros
		{ f(0,2),       2, "(+0+1)/1"     },
		{ f(0,4,0,2,0), 4, "(+0+2+0+1+0)/2" },
	} {
		if den, nums, err := factory.zFracN(s.den, s.nums...); err != nil {
			if s.exp != "∞" {
				t.Fatalf("FracN unexpected overflow for %d %v", s.den, s.nums)
			}
		} else {
			var sb strings.Builder
			sb.WriteString("(")
			for _, num := range nums {
				sb.WriteString(fmt.Sprintf("%+d", num))
			}
			sb.WriteString(fmt.Sprintf(")/%d", den))
			if got := sb.String(); got != s.exp {
				t.Fatalf("FracN got %s exp %s", got, s.exp)
			}
		}
	}

	// zSqrt
	for _, s := range []struct{ c, d Z; exp string } {
		{ 0, 10, "0√0"  },
		{ 0,  0, "0√0"  },
		{ 1,  0, "0√0"  },
		{ 1, -4, "2√-1" },
		{ 1, -2, "1√-2" },
		{ 1, -1, "1√-1" },
		{ 1,  1, "1√1"  },
		{ 1,  2, "1√2"  },
		{ 1,  3, "1√3"  },
		{ 1,  4, "2√1"  },

		{ 2,  3*3*3*7*7, "42√3"},

		{ -Z32_MAX,   1, "-2147483647√1" }, // min int32
		{ +Z32_MAX,   1, "2147483647√1"  }, // max int32
		{ +Z32_MAX+1, 0, "0√0"           }, // 0 has precedence over infinite
		{ +Z32_MAX+1, 1, "∞"             },
		{ +Z32_MAX,   4, "∞"             },

		{  0, Z32_MAX*Z32_MAX,   "0√0"  },
		{ +1, Z32_MAX*Z32_MAX/1, "∞"  },
		{ +1, Z32_MAX*Z32_MAX/2,       "98304√238609294" },
		{ +1, Z32_MAX*Z32_MAX/4,       "98304√119304647" },
		{ +1, (Z32_MAX-1)*(Z32_MAX-1), "2147483646√1"    },
	} {
		if o, i, err := factory.zSqrt(s.c, s.d); err != nil {
			if s.exp != "∞" {
				t.Fatalf("Sqrt unexpected overflow for %d√%d", s.c, s.d)
			}
		} else if got := fmt.Sprintf("%d√%d", o, i); got != s.exp {
			t.Fatalf("Sqrt get %s exp %s", got, s.exp)
		}
	}

	// zSqrtN
	for _, s := range []struct{ e, f, g, h Z; exp string } {
		{-1,-1,-1,-1, "-1√(-1-1√-1)" },
		{ 1, 1, 1, 0, "1√1"       }, // +1√(1+1√0) = +1√(1+0) = +1√1 = +1
		{ 1, 1, 0, 1, "1√1"       }, // +1√(1+0√1) = +1√(1+0) = +1√1 = +1
		{ 1, 0, 1, 1, "1√1"       }, // +1√(0+1√1) = +1√(0+1) = +1√1 = +1
		{ 1, 0, 9, 0, "0√0"       }, // +1√(0+9√0) = +1√(0+0) = +1√0 = +0
		{ 1, 0, 0, 9, "0√0"       }, // +1√(0+0√9) = +1√(0+0) = +1√0 = +0
		{ 0, 3, 6, 9, "0√0"       }, // +0√(3+6√9) = +0√(3+12) = +0√15 = +0
		{ 1, 1,-1, 1, "0√0"       }, // +1√(1-1√1) = +1√(1-1) = +1√0 = +0
		{ 1, 1,-2, 1, "1√-1"      }, // +1√(1-2√1) = +1√(1-2) = +1√-1 = +1i
		{ 1, 1, 1, 1, "1√2"       },
		{ 1, 1, 1, 2, "1√(1+1√2)" },
		{ 1, 1, 1, 3, "1√(1+1√3)" },
		{ 1, 1, 1, 4, "1√3"       },
		{ 1, 1, 2, 4, "1√5"       }, // +1√(1+2√4) = +1√(1+4) = +1√5
		{ 1, 1, 3, 4, "1√7"       }, // +1√(1+3√4) = +1√(1+6) = +1√7
		{ 1, 3, 3, 4, "3√1"       }, // +1√(3+3√4) = +1√(3+6) = +1√9 = +3
		{ 3, 3, 3, 4, "9√1"       }, // +3√(3+3√4) = +3√(3+6) = +3√9 = +9
		{ 1, 1, 1, 9, "2√1"       }, // +1√(1+1√9) = +1√(1+3) = +1√4 = +2
		{ 1, 4, 2, 8, "2√(1+1√2)" }, // +1√(4+2√8) = +1√(4+4√2) = +2√(1+1√2)
		{ 1, 8, 8, 2, "2√(2+2√2)" }, // +1√(8+8√2) = +2√(2+2√2)
		{ 1, 8, 4, 8, "2√(2+2√2)" }, // +1√(8+4√8) = +2√(8+8√2) = +2√(2+2√2)
		{ 1,27, 6,27, "3√(3+2√3)" }, // +1√(27+6√27) = +1√(27+18√3) = +3√(3+2√3)
		{ 2, 2, 2, 2, "2√(2+2√2)" },
		{ 3, 3, 3, 3, "3√(3+3√3)" },
		{ 4, 4, 4, 4, "8√3"       }, // +4√(4+4√4) = +4√(4+8) = +4√12 = + 8√3

		{ 8, 8, 8, 8, "16√(2+4√2)" }, // +8√(8+8√8) = +8√(8+16√2) = +16√(2+4√2)
		{ 9, 9, 9, 9, "54√1"         }, // +9√(9+9√9) = +9√(9+27) = +9√36 = +54
		{12,12,12,12, "24√(3+6√3)" }, // +12√(12+12√12) = +12√(12+24√3) = +24√(3+6√3)

	} {
		if g, h, err := factory.zSqrt(s.g, s.h); err != nil {
			t.Fatalf("%v", err)
		} else if g == 0 || h == 0 {
			if e, f, err := factory.zSqrt(s.e, s.f); err != nil {
				t.Fatalf("%v", err)
			} else if got := fmt.Sprintf("%d√%d", e, f); got != s.exp {
				t.Fatalf("zSqrtN get %s exp %s", got, s.exp)
			}
		} else if h == 1 {
			if e, f, err := factory.zSqrt(s.e, s.f + Z(g)); err != nil {
				t.Fatalf("%v", err)
			} else if got := fmt.Sprintf("%d√%d", e, f); got != s.exp {
				t.Fatalf("zSqrtN get %s exp %s", got, s.exp)
			}
		} else if e, fg, err := factory.zSqrtN(s.e, s.f, Z(g)); err != nil {
			t.Fatalf("%v", err)
		} else if got := fmt.Sprintf("%d√(%d%+d√%d)", e, fg[0], fg[1], h); got != s.exp {
			t.Fatalf("zSqrtN get %s exp %s", got, s.exp)
		}
	}

	// zSqrtN general
	for _, s := range []struct{ exp string; o Z; is []Z } {
		// with zeroes
		{ "+0√()",       0, f(1, 1, 1) },
		{ "+0√()",       4, f(0, 0, 0) },
		{ "+1√(+0+1+2)", 1, f(0, 1, 2) },
		{ "+2√(+0+1+1)", 1, f(0, 4, 4) },
		{ "+9√(+0+0+1)", 3, f(0, 0, 9) },
		// negatives
		{ "+20√(-1+3-6)", +10, f(-4, 12,-24) },
		{ "-20√(+1-3+6)", -10, f(+4,-12,+24) },
		{ "+5√(+5+7+11)",   5, f( 5,  7, 11) },
		{ "+8√(+1+1+1)",    1, f(64, 64, 64) },
		// big
		{ "+1√(+1+2+3+4+5+6+7+8+9+10+11+12)",     1, f(1,2,3,4,5,6,7,8,9,10,11,12) },
		{ "+3√(+238609294+238609294)",            1, f(Z32_MAX-1, Z32_MAX-1) },
		{ "+1√(+2147483645+2147483645)",          1, f(Z32_MAX-2, Z32_MAX-2) },
		{ "+2√(+536870911+536870911)",            1, f(Z32_MAX-3, Z32_MAX-3) },
		{ "+1√(+2147483643+2147483643)",          1, f(Z32_MAX-4, Z32_MAX-4) },
		{ "+1√(+2147483642+2147483642)",          1, f(Z32_MAX-5, Z32_MAX-5) },
		{ "+2√(+536870910+536870910)",            1, f(Z32_MAX-7, Z32_MAX-7) },
		{ "+2147483647√(+2147483647-2147483647)", Z32_MAX, f(Z32_MAX, -Z32_MAX) }, // primes
		// overflows
		{ "∞",  1,       f(Z32_MAX*Z32_MAX) },
		{ "∞",  Z32_MAX, f(4,4,4)           },

	} {
		o, is, err := factory.zSqrtN(s.o, s.is...)
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("%+d√(", o))
		for _, i := range is {
			sb.WriteString(fmt.Sprintf("%+d", i))
		}
		sb.WriteString(")")
		got := sb.String()
		if err != nil {
			if s.exp != "∞" {
				t.Fatalf("zSqrtN unexpected overflow for %d %v", s.o, s.is)
			}
		} else if got != s.exp {
			t.Fatalf("zSqrtN got %s exp %s", got, s.exp)
		}
	}
}
