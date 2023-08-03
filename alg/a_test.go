package alg

import (
	"fmt"
	"strings"
	"testing"
)

func TestA32(t *testing.T) {

	factory := &A32s{} // without primes 

	const cospi_15 = "(1+1√5+1√(30-6√5))/8"
	d0 := factory.f0(1)
	d1 := factory.f1(1, 5)
	d2 := factory.f2(1, 30, -6, 5)

	const cospi_24  = "(1√(2+0+1√(2+1√3)))/2"
	e3 := factory.f3(1,2,0,0,1,2,1,3)

	const cos2pi_17 = "(-1+1√17+1√(34-2√17)+2√(17+3√17-1√(170+38√17)))/16"
	f0 := factory.f0(-1)
	f1 := factory.f1(1,17)
	f2 := factory.f2(1,34,-2,17)
	f3 := factory.f3(2,17,3,17,-1,170,38,17)
	// AZ32
	for _, s := range []struct{ az *AZ32; exp string } {
		{ az: d0, exp:"+1" },
		{ az: d1, exp:"+1√5" },
		{ az: d2, exp:"+1√(30-6√5)" },
		{ az: e3, exp:"+1√(2+0+1√(2+1√3))" },
		{ az: f0, exp:"-1" },
		{ az: f1, exp:"+1√17" },
		{ az: f2, exp:"+1√(34-2√17)" },
		{ az: f3, exp:"+2√(17+3√17-1√(170+38√17))" },
	} {
		if got := s.az.String(); got != s.exp {
			t.Fatalf("AZ32 got=%s exp=%s", got, s.exp)
		}
	}

	// AQ32
	for _, s := range []struct{ exp string; num []*AZ32; den uint32 } {
		{ cospi_15,  []*AZ32{ d0,d1,d2    },  8 },
		{ cospi_24,  []*AZ32{          e3 },  2 },
		{ cos2pi_17, []*AZ32{ f0,f1,f2,f3 }, 16 },
	} {
		q := factory.Q(s.num, s.den)
		if got := q.String(); got != s.exp {
			t.Fatalf("AQ32 got=%s exp=%s", got, s.exp)
		}
	}
}

func TestA32Red(t *testing.T) {

	factory := NewA32s() // with primes for reductions

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

	// roie
	for _, s := range []struct{ o int64; is []int64; exp string } {
		// with zeroes
		{   0, []int64{  1,  1,  1 }, "+0√()"       },
		{   4, []int64{  0,  0,  0 }, "+0√()"       },
		{   1, []int64{  0,  1,  2 }, "+1√(+0+1+2)" },
		{   1, []int64{  0,  4,  4 }, "+2√(+0+1+1)" },
		{   3, []int64{  0,  0,  9 }, "+9√(+0+0+1)" },
		// negatives
		{ +10, []int64{ -4, 12,-24 }, "+20√(-1+3-6)" },
		{ -10, []int64{ +4,-12,+24 }, "-20√(+1-3+6)" },
		{   5, []int64{  5,  7, 11 }, "+5√(+5+7+11)" },
		{   1, []int64{ 64, 64, 64 }, "+8√(+1+1+1)"  },
		// big
		{   1,      []int64{ 1,2,3,4,5,6,7,8,9,10,11,12 }, "+1√(+1+2+3+4+5+6+7+8+9+10+11+12)" },
		{   1,      []int64{ AZ_MAX-1, AZ_MAX-1 },         "+3√(+238609294+238609294)"   },
		{   1,      []int64{ AZ_MAX-2, AZ_MAX-2 },         "+1√(+2147483645+2147483645)" },
		{   1,      []int64{ AZ_MAX-3, AZ_MAX-3 },         "+2√(+536870911+536870911)" },
		{   1,      []int64{ AZ_MAX-4, AZ_MAX-4 },         "+1√(+2147483643+2147483643)" },
		{   1,      []int64{ AZ_MAX-5, AZ_MAX-5 },         "+1√(+2147483642+2147483642)" },
		{   1,      []int64{ AZ_MAX-7, AZ_MAX-7 },         "+2√(+536870910+536870910)" },
		// overflows
		{   1,      []int64{ AZ_MAX*AZ_MAX },   "∞" },
		{   AZ_MAX, []int64{ AZ_MAX, -AZ_MAX }, "+2147483647√(+2147483647-2147483647)" }, // primes
		{   AZ_MAX, []int64{ 4,4,4 },           "∞" },

	} {
		o, is, overflow := factory.roie(s.o, s.is)
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("%+d√(", o))
		for _, i := range is {
			sb.WriteString(fmt.Sprintf("%+d", i))
		}
		sb.WriteString(")")
		got := sb.String()
		if overflow {
			if s.exp != "∞" {
				t.Fatalf("roie unexpected overflow for %d %v", s.o, s.is)
			}
		} else if got != s.exp {
			t.Fatalf("roie got %s exp %s", got, s.exp)
		}
	}

	// F1 -> roi
	for _, s := range []struct{ c, d int64; exp string } {
		{ 0, 10, "+0"    },
		{ 0,  0, "+0"    },
		{ 1,  0, "+0"    },
		{ 1, -4, "+2i"   },
		{ 1, -2, "+1i√2" },
		{ 1, -1, "+1i"   },
		{ 1,  1, "+1"    },
		{ 1,  2, "+1√2"  },
		{ 1,  3, "+1√3"  },
		{ 1,  4, "+2"    },

		{ 2,  3*3*3*7*7,  "+42√3"},

		{ -AZ_MAX,   1, "-2147483647" }, // min int32
		{ +AZ_MAX,   1, "+2147483647" }, // max int32
		{ +AZ_MAX+1, 0, "+0"          }, // 0 has precedence over infinite
		{ +AZ_MAX+1, 1, "∞"           },
		{ +AZ_MAX,   4, "∞"           },

		{  0, AZ_MAX*AZ_MAX,   "+0" },
		{ +1, AZ_MAX*AZ_MAX/1, "∞"  },
		{ +1, AZ_MAX*AZ_MAX/2,       "+98304√238609294" },
		{ +1, AZ_MAX*AZ_MAX/4,       "+98304√119304647" },
		{ +1, (AZ_MAX-1)*(AZ_MAX-1), "+2147483646"      },
	} {
		if z, overflow := factory.F1(s.c, s.d); overflow {
			if s.exp != "∞" {
				t.Fatalf("F1 unexpected overflow for %d√%d", s.c, s.d)
			}
		} else if got := z.String(); got != s.exp {
			t.Fatalf("F1 get %s exp %s", got, s.exp)
		}
	}

	// F2 -> roi -> roie
	for _, s := range []struct{ e, f, g, h int64; exp string } {
		{-1,-1,-1,-1, "-1√(-1-1i)" },
		{ 1, 1, 1, 0, "+1"         }, // +1√(1+1√0) = +1√(1+0) = +1√1 = +1
		{ 1, 1, 0, 1, "+1"         }, // +1√(1+0√1) = +1√(1+0) = +1√1 = +1
		{ 1, 0, 1, 1, "+1"         }, // +1√(0+1√1) = +1√(0+1) = +1√1 = +1
		{ 1, 0, 9, 0, "+0"         }, // +1√(0+9√0) = +1√(0+0) = +1√0 = +0
		{ 1, 0, 0, 9, "+0"         }, // +1√(0+0√9) = +1√(0+0) = +1√0 = +0
		{ 0, 3, 6, 9, "+0"         }, // +0√(3+6√9) = +0√(3+12) = +0√15 = +0
		{ 1, 1,-1, 1, "+0"         }, // +1√(1-1√1) = +1√(1-1) = +1√0 = +0
		{ 1, 1,-2, 1, "+1i"        }, // +1√(1-2√1) = +1√(1-2) = +1√-1 = +1i
		{ 1, 1, 1, 1, "+1√2"       },
		{ 1, 1, 1, 2, "+1√(1+1√2)" },
		{ 1, 1, 1, 3, "+1√(1+1√3)" },
		{ 1, 1, 1, 4, "+1√3"       },
		{ 1, 1, 2, 4, "+1√5"       }, // +1√(1+2√4) = +1√(1+4) = +1√5
		{ 1, 1, 3, 4, "+1√7"       }, // +1√(1+3√4) = +1√(1+6) = +1√7
		{ 1, 3, 3, 4, "+3"         }, // +1√(3+3√4) = +1√(3+6) = +1√9 = +3
		{ 3, 3, 3, 4, "+9"         }, // +3√(3+3√4) = +3√(3+6) = +3√9 = +9
		{ 1, 1, 1, 9, "+2"         }, // +1√(1+1√9) = +1√(1+3) = +1√4 = +2
		{ 1, 4, 2, 8, "+2√(1+1√2)" }, // +1√(4+2√8) = +1√(4+4√2) = +2√(1+1√2)
		{ 1, 8, 8, 2, "+2√(2+2√2)" }, // +1√(8+8√2) = +2√(2+2√2)
		{ 1, 8, 4, 8, "+2√(2+2√2)" }, // +1√(8+4√8) = +2√(8+8√2) = +2√(2+2√2)
		{ 1,27, 6,27, "+3√(3+2√3)" }, // +1√(27+6√27) = +1√(27+18√3) = +3√(3+2√3)
		{ 2, 2, 2, 2, "+2√(2+2√2)" },
		{ 3, 3, 3, 3, "+3√(3+3√3)" },
		{ 4, 4, 4, 4, "+8√3"       }, // +4√(4+4√4) = +4√(4+8) = +4√12 = + 8√3

		{ 8, 8, 8, 8, "+16√(2+4√2)" }, // +8√(8+8√8) = +8√(8+16√2) = +16√(2+4√2)
		{ 9, 9, 9, 9, "+54"         }, // +9√(9+9√9) = +9√(9+27) = +9√36 = +54
		{12,12,12,12, "+24√(3+6√3)" }, // +12√(12+12√12) = +12√(12+24√3) = +24√(3+6√3)

	} {
		if z, overflow := factory.F2(s.e, s.f, s.g, s.h); overflow {
			if s.exp != "∞" {
				t.Fatalf("F2 unexpected overflow for %+d√(%d%+d√%d)", s.e, s.f, s.g, s.h)
			}
		} else if got := z.String(); got != s.exp {
			t.Fatalf("F2 get %s exp %s", got, s.exp)
		}
	}
}
