package alg

import (
	"testing"
)

func TestAZ32(t *testing.T) {

	factory := &AZ32s{}

	f0 := factory.F0(-1)
	f1 := factory.F1(1,17)
	f2 := factory.F2(1,34,-2,17)
	f3 := factory.F3(2,17,3,17,-1,170,38,17)
	for _, s := range []struct{ az *AZ32; exp string } {
		{ az: f0, exp:"" },
		{ az: f1, exp:"" },
		{ az: f2, exp:"" },
		{ az: f3, exp:"" },
	} {
		if got := s.az.String(); got != s.exp {
			t.Fatalf("AZ32 got=%s exp=%s", got, s.exp)
		}
	}

}
