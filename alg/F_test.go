package alg

import (
	"testing"
)

func TestB(t *testing.T) {

	none     := NewB( 0,0)
	zero     := NewB( 0,1)
	minus1   := NewB(-1,1)
	half     := NewB( 2,4)
	ten      := NewB( 10,1)
	one_17th := NewB( 2*3*5*7*11*13, 2*3*5*7*11*13*17)
	for exp, b := range map[string]*B {
		"":     none,
		"0":    zero,
		"-1":   minus1,
		"1/2":  half,
		"-7/3": NewB(-2*5*7*11, 2*3*5*11),
		"1/17":	one_17th,
		"10":   ten,
	} {
		if got := b.StringB(); got != exp {
			t.Fatalf("B got %s exp %s", got, exp)
		}
	}
	// AddB
	for exp, b := range map[string]*B {
		"":      zero.AddB(none),
		"0":     zero.AddB(zero),
		"-1":    zero.AddB(minus1),
		"-2":    minus1.AddB(minus1),
		"1":     half.AddB(half),
		"2/17":  one_17th.AddB(one_17th),
		"19/34": half.AddB(one_17th),
		"21/34": half.AddB(one_17th).AddB(one_17th),
		"19/17": half.AddB(one_17th).AddB(one_17th).AddB(half),
		"21/2":  ten.AddB(half),
	} {
		if got := b.StringB(); got != exp {
			t.Fatalf("B-add got %s exp %s", got, exp)
		}
	}
	// MulB

}