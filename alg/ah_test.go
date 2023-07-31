package alg

import (
	"fmt"
	"testing"
)

func TestZ(t *testing.T) {
	for _, u := range []struct{ a N; b, c Z; exp string } {
		{ a:0,    b:0,    c:0,    exp:"0: 0,0,0" },
		{ a:1,    b:0,    c:0,    exp:"1: 1,0,0" },
		{ a:1,    b:1,    c:0,    exp:"1: 1,1,0" },
		{ a:1,    b:1,    c:1,    exp:"1: 1,1,1" },

		{ a:3,    b:4,    c:6,    exp:"1: 3,4,6"  },
		{ a:1,    b:10,   c:20,   exp:"1: 1,10,20"},
		
		{ a:19,   b:23,   c:29,   exp:"1: 19,23,29" },
		{ a:1001, b:1001, c:1001, exp:"1001: 1,1,1" },
		{ a:3003, b:1001, c:7007, exp:"1001: 3,1,7" },
		{ a:1000, b:1001, c:1002, exp:"1: 1000,1001,1002" },
		{ a:30,   b:-45,  c:-60,  exp:"15: 2,-3,-4" },
		{ a:2,    b:0,    c:1,    exp:"1: 2,0,1" },
	} {
		g := (&u.a).Reduce3(&u.b, &u.c)
		if got := fmt.Sprintf("%d: %d,%d,%d", g, u.a, u.b, u.c); got != u.exp {
			t.Fatalf("Z.Reduce3 got %s exp %s", got, u.exp)
		}
	}
}

func TestB(t *testing.T) {
	none     := NewB( 0,0)
	zero     := NewB( 0,1)
	minus1   := NewB(-1,1)
	half     := NewB( 2,4)
	ten      := NewB( 10,1)
	one_17th := NewB( 2*3*5*7*11*13, 2*3*5*7*11*13*17)
	for _, b := range []struct { e string; b *B } {
		{ e:"∞",     b:none },
		{ e:"+0",    b:zero },
		{ e:"-1",    b:minus1 },
		{ e:"+1/2",  b:half },
		{ e:"-7/3",  b:NewB(-2*5*7*11, 2*3*5*11) },
		{ e:"+1/17", b:one_17th },
		{ e:"+10",   b:ten },

		{ e:"+1",            b:NewB(+Z(N32_MAX),   N32_MAX  ) },
		{ e:"+4294967295",   b:NewB(+Z(N32_MAX),   1        ) }, // max positive
		{ e:"+1/4294967295", b:NewB(1,             N32_MAX  ) }, // min positive

		{ e:"∞", b:NewB(1,             N32_MAX+1) }, // overflow
		{ e:"∞", b:NewB(+Z(N32_MAX+1), 1        ) }, // overflow
		{ e:"∞", b:NewB(-Z(N32_MAX+1), 1        ) }, // overflow
	} {
		if got := b.b.String(); got != b.e {
			t.Fatalf("B got %s exp %s", got, b.e)
		}
	}
	// AddB
	for exp, b := range map[string]*B {
		"∞":      zero.AddB(none),
		"+0":     zero.AddB(zero),
		"-1":     zero.AddB(minus1),
		"-2":     minus1.AddB(minus1),
		"+1":     half.AddB(half),
		"+2/17":  one_17th.AddB(one_17th),
		"+19/34": half.AddB(one_17th),
		"+21/34": half.AddB(one_17th).AddB(one_17th),
		"+19/17": half.AddB(one_17th).AddB(one_17th).AddB(half),
		"+21/2":  ten.AddB(half),
	} {
		if got := b.String(); got != exp {
			t.Fatalf("B-add got %s exp %s", got, exp)
		}
	}
	for _, b := range []struct { e string; b *B } {


		// InvB
		{ e:"∞",     b: none.Inv() },
		{ e:"∞",     b: zero.Inv() },
		{ e:"-1",    b: minus1.Inv() },
		{ e:"-1",    b: minus1.Inv().Inv().Inv() },
		{ e:"+2",    b: half.Inv() },
		{ e:"-3/7",  b: NewB(-2*5*7*11, 2*3*5*11).Inv() },
		{ e:"+17",   b: one_17th.Inv() },
		{ e:"+1/10", b: ten.Inv() },
		{ e:"+10",   b: ten.Inv().Inv() },
	} {
		if got := b.b.String(); got != b.e {
			t.Fatalf("B got %s exp %s", got, b.e)
		}
	}

	// MulB
	for _, b := range []struct { e string; b *B } {
		{ e:"∞",        b: zero.MulB(none) },
		{ e:"+0",       b: zero.MulB(zero) },
		{ e:"+0",       b: zero.MulB(minus1) },
		{ e:"+1",       b: minus1.MulB(minus1) },
		{ e:"+1/4",     b: half.MulB(half) },
		{ e:"+1/289",   b: one_17th.MulB(one_17th) },
		{ e:"+1/34",    b: half.MulB(one_17th) },
		{ e:"+1/578",   b: half.MulB(one_17th).MulB(one_17th) },
		{ e:"+1/1156",  b: half.MulB(one_17th).MulB(one_17th).MulB(half) },
		{ e:"+5",       b: ten.MulB(half) },
		{ e:"-12/5",    b: NewB(-144,130).MulB(NewB(26,12)) },
	} {
		if got := b.b.String(); got != b.e {
			t.Fatalf("B-mul got %s exp %s", got, b.e)
		}
	}
}

func TestD(t *testing.T) {
	rs := NewAI32s()
	ds := NewDs(rs)
	a1 := N(1)
	a2 := N(2)
	infinite := ds.NewD( 0, 0,0,  0)
	zero     := ds.NewD( 0, 0,0, a1)
	minus1   := ds.NewD(-1, 0,0, a1)
	half     := ds.NewD( 1, 0,0, a2)
	ten      := ds.NewD(10, 0,0, a1)
	one_17th := ds.NewD(2*3*5*7*11*13, 0,0, 2*3*5*7*11*13*17)

	for _, r := range []struct { d *D; exp string } {

		{ d: infinite, exp:"∞" },
		{ d: zero,     exp:"+0" },
		{ d: minus1,   exp:"-1" },
		{ d: half,     exp:"+1/2" },
		{ d: ten,      exp:"+10" },
		{ d: one_17th, exp:"+1/17" },
		
		{ d: ds.NewD(  0,  1,-1, a1),   exp: "+1i"        }, // imag

		{ d: ds.NewD(  0,  1,+2, a1),   exp: "+1√2"       },
		{ d: ds.NewD(  0,  1,+2, a2),   exp: "+1√2/2"     },
		{ d: ds.NewD(  1,  1,+2, a2),   exp: "(+1+1√2)/2" },
		{ d: ds.NewD( -2, -2,+2, a2),   exp: "-1-1√2"     },
		{ d: ds.NewD( -2, -1,+8, a2),   exp: "-1-1√2"     },
		{ d: ds.NewD(  0, +1,+2, a2),   exp: "+1√2/2"     }, // sin(pi/4)
		{ d: ds.NewD( -1, +1,+5,  4),   exp: "(-1+1√5)/4" }, // sin(pi/10)

		{ d: ds.NewD(   3,   4, 1,   6), exp: "+3/6" },
		{ d: ds.NewD( 158, 632, 5, 316), exp: "(+1+4√5)/2" }, 


	} {
		if got := r.d.String(); got != r.exp {
			t.Fatalf("D-new got %s exp %s", got, r.exp)	
		}
	}

	// B sqrts
	for _, r := range []struct { b Z; a N; exp string } {
		// reals
		{ b: 0, a:0, exp:"∞"       },
		{ b: 0, a:1, exp:"+0"      },
		{ b:+1, a:1, exp:"+1"      },
		{ b: 1, a:5, exp:"+1√5/5"  },
		// imaginaries
		{ b:-1, a:1, exp:"+1i"     },
		{ b:-1, a:2, exp:"+1i√2/2" },
	} {
		if got := ds.NewDsqrtB(r.b, r.a).String(); got != r.exp {
			t.Fatalf("D-sqrtB got %s exp %s", got, r.exp)		
		}
	}

	// D inversions
	for _, r := range []struct { d *D; exp string } {
		// reals
		{ d: ds.NewD(  1, 2, 3,  4),  exp: "(-4+8√3)/11"    },
		{ d: ds.NewD( -4, 8, 3, 11),  exp: "(+1+2√3)/4"     },

		{ d: ds.NewD(  2, 3,  4, 5),  exp: "(-5+15)/16"     }, // 5/8
		{ d: ds.NewD( -5, 15, 1, 16), exp: "(+2+6)/5"       }, // 8/5

		{ d: ds.NewD(  3,  4, 5,  6), exp: "(-18+24√5)/71"  }, // 
		{ d: ds.NewD(-18, 24, 5, 71), exp: "(+3+4√5)/6"     },

		{ d: ds.NewD(  4,  5, 6,  7), exp: "(-28+35√6)/134" },
		{ d: ds.NewD(-28, 35, 6,134), exp: "(+4+5√6)/7"     },

		{ d: ds.NewD(  5,  6, 7,  8), exp: "(-40+48√7)/227" },
		{ d: ds.NewD(-40, 48, 7,227), exp: "(+5+6√7)/8"     },

		{ d: ds.NewD(  6,  7, 8,  9), exp: "(-27+63√2)/178" },
		{ d: ds.NewD(-27, 63, 2,178), exp: "(+6+14√2)/9"    },

		{ d: ds.NewD(  7,  8, 9, 10), exp: "(-70+240)/527"  },
		{ d: ds.NewD(-70,240, 1,527), exp: "(+7+24)/10"     },

	} {
		if got := ds.NewInvD(r.d).String(); got != r.exp {
			t.Fatalf("D-inv got %s exp %s", got, r.exp)		
		}
	}
}

func TestH(t *testing.T) {

	rs := NewAI32s()
	hs := NewHs(rs)

	for _, r := range []struct { b, c, d, e, f, g, h Z; a N; exp string } {
		// reals
		{ b:1, c:2, d:3, e:4, f:5, g:6, h:7, a:9, exp: "+6√7" },
	} {
		if got := hs.NewH(r.b, r.c, r.d, r.e, r.f, r.g, r.h, r.a).String(); got != r.exp {
			t.Fatalf("H got %s exp %s", got, r.exp)		
		}
	}
}