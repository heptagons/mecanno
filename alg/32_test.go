package alg

/*
import (
	"testing"
)

func Test32(t *testing.T) {

	i0,  _ := newI32(0)
	i1,  _ := newI32(1)
	i2,  _ := newI32(2)
	i4,  _ := newI32(4)
	i5_, _ := newI32(-5)
	// I32 
	for _, s := range[]struct{ exp string; i *I32 } {
		{ "+0",  i0 },
		{ "+1",  i1 },
		{ "+2",  i2 },
		{ "+4",  i4 },
		{ "-5",  i5_ },
	} {
		if got := s.i.String(); got != s.exp {
			t.Fatalf("I32 got=%s exp=%s", got, s.exp)
		}
	}

	for _, s := range[] struct{ a, b Z; exp string; overflow bool } {
		{ a:-7,   b:-8,   exp:"-15"   },
		{ a:1000, b:3000, exp:"+4000" },

		{ a:-Z(N32_MAX), b: 0,          exp:"-4294967295" },
		{ a: 0,          b: Z(N32_MAX), exp:"+4294967295" },
		{ a: Z(N32_MAX), b:-Z(N32_MAX), exp:"+0" },

		{ a:Z(N32_MAX), b:1,          overflow:true },
		{ a:1,          b:Z(N32_MAX), overflow:true },
		{ a:Z(N32_MAX), b:Z(N32_MAX), overflow:true },
	} {
		a, _ := newI32(s.a)
		b, _ := newI32(s.b)
		if sum, overflow := a.add(b); overflow {
			if !s.overflow {
				t.Fatalf("add Expected overflow for %d,%d", s.a, s.b)
			}
		} else if got := sum.String(); got != s.exp {
			t.Fatalf("I32.add got %s exp %s", got, s.exp)
		}
	}

	for _, s := range[] struct{ a, b Z; exp string; overflow bool } {
		{ a:-7,   b:-8,   exp:"+56"      },
		{ a:1000, b:3000, exp:"+3000000" },

		{ a:-Z(N32_MAX), b: 0,          exp:"+0" },
		{ a: 0,          b: Z(N32_MAX), exp:"+0" },
		{ a: Z(N32_MAX), b:-Z(N32_MAX), overflow:true },

		{ a:-Z(N32_MAX), b: 1,          exp:"-4294967295" },
		{ a: 1,          b: Z(N32_MAX), exp:"+4294967295" },
		{ a: Z(N32_MAX), b:-Z(N32_MAX), overflow:true },
	} {
		a, _ := newI32(s.a)
		b, _ := newI32(s.b)
		if mul, overflow := a.mul(b); overflow {
			if !s.overflow {
				t.Fatalf("I32.mul Expected overflow for %d,%d", s.a, s.b)
			}
		} else if got := mul.String(); got != s.exp {
			t.Fatalf("I32.mul got %s exp %s", got, s.exp)
		}
	}

	// AI32 (not irreducible)
	for _, s := range []struct{ exp string; ai *AI32 } {
		{ "+0", &AI32{} },
		{ "+0", &AI32{ o:i0 }},
		{ "+0", &AI32{ o:i1 }},
		{ "+1", &AI32{ o:i1, i:i1 }}, // 1√1
		{ "+2", &AI32{ o:i2, i:i1 }}, // 2√1
		
		{ "+1√2", &AI32{ o:i1, i:i2 }}, // 1√2
		{ "+1√4", &AI32{ o:i1, i:i4 }}, // 1√4

		{ "+1√4", &AI32{o:i1, i:i4, e:&AI32{} }},
		{ "+1√4", &AI32{o:i1, i:i4, e:&AI32{ o:i0 } }}, // 1√(4+0)
		{ "+1√4", &AI32{o:i1, i:i4, e:&AI32{ o:i1 } }}, // 1√(4+1√0)
		
		{ "+1√(4+1)",   &AI32{o:i1, i:i4, e:&AI32{ o:i1, i:i1 } }}, // 1√(4+1√1)
		{ "+1√(4+1√2)", &AI32{o:i1, i:i4, e:&AI32{ o:i1, i:i2 } }}, // 1√(4+1√2)
		{ "+1√(4+1√4)", &AI32{o:i1, i:i4, e:&AI32{ o:i1, i:i4 } }}, // 1√(4+1√4)
	} {
		if got := s.ai.String(); got != s.exp {
			t.Fatalf("AI32 got=%s exp=%s", got, s.exp)
		}
	}
}
*/
