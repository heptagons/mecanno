package four

import (
	"testing"
	. "github.com/heptagons/meccano/nest"
)

func TestFour(t *testing.T) {
	factory := NewFours()
	a,b,c := N(1),N(1),N(1)
	d,e   := N32(1),N32(1)
	four := NewFour(a, b, c)
	t.Logf("abc=(%d,%d,%d)", four.a, four.b, four.c)
	if f, err := factory.F(four, d); err == nil {
		t.Logf("d=%d f=%v", d, f)
	} else {
		t.Fatal(err)
	}

	if cos, err := factory.CosTheta(four, d); err == nil {
		t.Logf("d=%d cosTheta=%v", d, cos)
	}
	if g, err := factory.G1(four, d, e); err == nil {
		t.Logf("de=%d,%d g1=%v", d, e, g)
	} else {
		t.Fatal(err)
	}
}