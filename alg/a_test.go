package alg

import (
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





