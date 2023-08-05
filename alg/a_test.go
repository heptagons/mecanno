package alg

import (
	"fmt"
	"strings"
	"testing"
)


func TestA32(t *testing.T) {

	const cospi_15 = "(1+1√5+1√(30-6√5))/8"
	d0 := newAZ32(1)
	d1 := newAZ32(1, 5)
	d2 := newAZ32(1, 30, -6, 5)

	const cospi_24  = "(1√(2+0+1√(2+1√3)))/2"
	e3 := newAZ32(1,2,0,0,1,2,1,3)

	const cos2pi_17 = "(-1+1√17+1√(34-2√17)+2√(17+3√17-1√(170+38√17)))/16"
	f0 := newAZ32(-1)
	f1 := newAZ32(1,17)
	f2 := newAZ32(1,34,-2,17)
	f3 := newAZ32(2,17,3,17,-1,170,38,17)
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
	/*for _, s := range []struct{ exp string; num []*AZ32; den uint32 } {
		{ cospi_15,  []*AZ32{ d0,d1,d2    },  8 },
		{ cospi_24,  []*AZ32{          e3 },  2 },
		{ cos2pi_17, []*AZ32{ f0,f1,f2,f3 }, 16 },
	} {
		q := newQ(s.num, s.den)
		if got := q.String(); got != s.exp {
			t.Fatalf("AQ32 got=%s exp=%s", got, s.exp)
		}
	}*/
}

func TestA32Red(t *testing.T) {

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

	f := func(p ...int64) []int64 { return p }


	// rQ
	for _, s := range []struct { nums []int64 ; den N; exp string; } {
		{ f( 0, 0, 0, 0), 1, "+0+0+0+0/1" },
		{ f( 1, 2, 3, 4), 5, "+1+2+3+4/5" },
		{ f(-8, 4, 2, 2), 2, "-4+2+1+1/1" },

		{ f(3*5*7*11),            77, "+15/1" },
		{ f(128, 64, 32, 16, 8, 4), 8, "+32+16+8+4+2+1/2"},
	} {
		if den, nums, overflow := factory.rQ(s.den, s.nums...); overflow {
			if s.exp != "∞" {
				t.Fatalf("rQ unexpected overflow for %d %v", s.den, s.nums)
			}
		} else {
			var sb strings.Builder
			for _, num := range nums {
				sb.WriteString(fmt.Sprintf("%+d", num))
			}
			sb.WriteString(fmt.Sprintf("/%d", den))
			if got := sb.String(); got != s.exp {
				t.Fatalf("rQ got %s exp %s", got, s.exp)
			}
		}
	}

	// roie
	for _, s := range []struct{ exp string; o int64; is []int64 } {
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
		{ "+3√(+238609294+238609294)",            1, f(AZ_MAX-1, AZ_MAX-1) },
		{ "+1√(+2147483645+2147483645)",          1, f(AZ_MAX-2, AZ_MAX-2) },
		{ "+2√(+536870911+536870911)",            1, f(AZ_MAX-3, AZ_MAX-3) },
		{ "+1√(+2147483643+2147483643)",          1, f(AZ_MAX-4, AZ_MAX-4) },
		{ "+1√(+2147483642+2147483642)",          1, f(AZ_MAX-5, AZ_MAX-5) },
		{ "+2√(+536870910+536870910)",            1, f(AZ_MAX-7, AZ_MAX-7) },
		{ "+2147483647√(+2147483647-2147483647)", AZ_MAX, f(AZ_MAX, -AZ_MAX) }, // primes
		// overflows
		{ "∞",  1,      f(AZ_MAX*AZ_MAX) },
		{ "∞",  AZ_MAX, f(4,4,4)         },

	} {
		o, is, overflow := factory.roie(s.o, s.is...)
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

	// Reduce 1: F1 -> roi
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
		if r, overflow := factory.Reduce(s.c, s.d); overflow {
			if s.exp != "∞" {
				t.Fatalf("Reduce1 unexpected overflow for %d√%d", s.c, s.d)
			}
		} else if got := newAZ32(r...).String(); got != s.exp {
			t.Fatalf("Reduce1 get %s exp %s", got, s.exp)
		}
	}

	// Reduce 2: F2 -> roi -> roie
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
		if r, overflow := factory.Reduce(s.e, s.f, s.g, s.h); overflow {
			if s.exp != "∞" {
				t.Fatalf("Reduce2 unexpected overflow for %+d√(%d%+d√%d)", s.e, s.f, s.g, s.h)
			}
		} else if got := newAZ32(r...).String(); got != s.exp {
			t.Fatalf("Reduce2 get %s exp %s", got, s.exp)
		}
	}

	/*
	// Reduce general
	for _, s := range []struct{ exp string; is []int64 } {
		{ "", f(1) }, // 1 is invalid
		{ "", f(1,1,1) }, // 3 is invalid
		{ "", f(1,1,1,1,1) }, // 5 is invalid
		{ "", f(1,1,1,1,1,1) }, // 6 is invalid
		{ "", f(1,1,1,1,1,1,1) }, // 7 is invalid


		{ "+1√(1+1+1√(2+0))",     f(1,1,1,1,1,1,1,1) }, // TODO
		{ "+2√(2+2√2+2√(2+2√2))", f(2,2,2,2,2,2,2,2) },
		{ "+3√(3+3√3+3√(3+3√3))", f(3,3,3,3,3,3,3,3) },
		{ "+8√(1+2+2√(3+0))",     f(4,4,4,4,4,4,4,4) }, // TODO

		// +1√(1+1√1+1√(1+1√1))
		// +1√(1+1√1+1√(1+1))
		// +1√(1+1√1+1√2)
		// +1√(1+1+1√2)
		// +1√(2+1√2)

		// +4√(4+4√4+4√(4+4√4))
		// +4√(4+4√4+4√(4+8))
		// +4√(4+4√4+4√12)
		// +4√(4+4√4+8√3)
		// +4√(4+8+8√3)
		// +8√(1+2+2√3)
		// +8√(3+2√3)

	} {
		if r, overflow := factory.Reduce(s.is...); overflow {
			if s.exp != "∞" {
				t.Fatalf("Reduce unexpected overflow for %+v", s.is)
			}
		} else if got := newAZ32(r...).String(); got != s.exp {
			t.Fatalf("Reduce get %s exp %s", got, s.exp)
		}
	}
	*/
	//t.Logf("22222222 %v", newAZ32(2,2,2,2,2,2,2,2))
}



func TestA32Tris(t *testing.T) {
	//factory := NewA32s() // with primes for reductions
	max := 5
	ts := NewA32Tris(max)
	//ts.Cosines(factory)
	if got, exp := len(ts.list), 29; got != exp {
		t.Fatalf("A32Tris max:%d got:%d exp:%d", max, got, exp)
	}
	for pos, exp := range []string {
		"sides:[1 1 1]",
		"sides:[2 1 1]",
		"sides:[2 2 1]",
		"sides:[3 1 1]",
	} {
		if got := ts.list[pos].String(); got != exp {
			t.Fatalf("A32Tris got:%s exp:%s", got, exp)		
		}
	}
}

