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
		if got := b.String(); got != exp {
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
		if got := b.String(); got != exp {
			t.Fatalf("B-add got %s exp %s", got, exp)
		}
	}
	// InvB
	for _, b := range []struct { e string; b *B } {
		{ e:"",     b: none.Inv() },
		{ e:"",     b: zero.Inv() },
		{ e:"-1",   b: minus1.Inv() },
		{ e:"-1",   b: minus1.Inv().Inv().Inv() },
		{ e:"2",    b: half.Inv() },
		{ e:"-3/7", b: NewB(-2*5*7*11, 2*3*5*11).Inv() },
		{ e:"17",	b: one_17th.Inv() },
		{ e:"1/10", b: ten.Inv() },
		{ e:"10",   b: ten.Inv().Inv() },
	} {
		if got := b.b.String(); got != b.e {
			t.Fatalf("B-inv got %s exp %s", got, b.e)
		}
	}

	// MulB
	for _, b := range []struct { e string; b *B } {
		{ e:"",        b: zero.MulB(none) },
		{ e:"0",       b: zero.MulB(zero) },
		{ e:"0",       b: zero.MulB(minus1) },
		{ e:"1",       b: minus1.MulB(minus1) },
		{ e:"1/4",     b: half.MulB(half) },
		{ e:"1/289",   b: one_17th.MulB(one_17th) },
		{ e:"1/34",    b: half.MulB(one_17th) },
		{ e:"1/578",   b: half.MulB(one_17th).MulB(one_17th) },
		{ e:"1/1156",  b: half.MulB(one_17th).MulB(one_17th).MulB(half) },
		{ e:"5",       b: ten.MulB(half) },
		{ e:"-12/5",   b: NewB(-144,130).MulB(NewB(26,12)) },
	} {
		if got := b.b.String(); got != b.e {
			t.Fatalf("B-mul got %s exp %s", got, b.e)
		}
	}
}