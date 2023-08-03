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

	// roie
	for _, s := range []struct{ o int64; is []int64; exp string } {
		
		{ +10, []int64{ -4, 12,-24 }, "+20√(-1+3-6)" },
		{ -10, []int64{ +4,-12,+24 }, "-20√(+1-3+6)" },
		{   5, []int64{  5,  7, 11 }, "+5√(+5+7+11)" },
		{   1, []int64{ 64, 64, 64 }, "+8√(+1+1+1)" },


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
				t.Fatalf("roie unexpected overflow for %s", got)
			}
		} else if got != s.exp {
			t.Fatalf("roie got %s exp %s", got, s.exp)
		}
	}

	// F1 -> roi
	for _, s := range []struct{ o, i int64; exp string } {
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
		if z, overflow := factory.F1(s.o, s.i); overflow {
			if s.exp != "∞" {
				t.Fatalf("F0 unexpected overflow for %d√%d", s.o, s.i)
			}
		} else if got := z.String(); got != s.exp {
			t.Fatalf("F0 get %s exp %s", got, s.exp)
		}
	}
}
