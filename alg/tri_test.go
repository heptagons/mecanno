package alg

import (
	"testing"
)

func TestTris20sin60(t *testing.T) {
	max := 20
	factory := NewN32s()
	ts := NewTris(max, factory)
	ts.SetSinCos()
	got, exp := len(ts.tris), 658
	if got != exp {
		t.Fatalf("Tris32 max:%d got:%d exp:%d", max, got, exp)
	}
	t.Logf("      Tris: %d", exp)
	t.Logf(" First tri: %v", ts.tris[0])
	t.Logf("  Last tri: %v", ts.tris[exp-1])

	for pos, exp := range []string {
		"abc:[1 1 1] cos:[1/2 1/2 1/2] sin:[√3/2 √3/2 √3/2]",                // √3
		"abc:[2 2 1] cos:[1/4 1/4 7/8] sin:[√15/4 √15/4 √15/8]",             // √15
		"abc:[3 2 2] cos:[-1/8 3/4 3/4] sin:[3√7/8 √7/4 √7/4]",              // √7
		"abc:[3 3 1] cos:[1/6 1/6 17/18] sin:[√35/6 √35/6 √35/18]",          // √35
		"abc:[3 3 2] cos:[1/3 1/3 7/9] sin:[2√2/3 2√2/3 4√2/9]",             // √2
		"abc:[4 3 2] cos:[-1/4 11/16 7/8] sin:[√15/4 3√15/16 √15/8]",        // √15
		"abc:[4 3 3] cos:[1/9 2/3 2/3] sin:[4√5/9 √5/3 √5/3]",               // √5
		"abc:[4 4 1] cos:[1/8 1/8 31/32] sin:[3√7/8 3√7/8 3√7/32]",          // √7
		"abc:[4 4 3] cos:[3/8 3/8 23/32] sin:[√55/8 √55/8 3√55/32]",         // √55
		"abc:[5 3 3] cos:[-7/18 5/6 5/6] sin:[5√11/18 √11/6 √11/6]",         // √11
		"abc:[5 4 2] cos:[-5/16 13/20 37/40] sin:[√231/16 √231/20 √231/40]", // √231
		"abc:[5 4 3] cos:[0 3/5 4/5] sin:[1 4/5 3/5]",                       // √1
		"abc:[5 4 4] cos:[7/32 5/8 5/8] sin:[5√39/32 √39/8 √39/8]",          // √39
		"abc:[5 5 1] cos:[1/10 1/10 49/50] sin:[3√11/10 3√11/10 3√11/50]",   // √11
		"abc:[5 5 2] cos:[1/5 1/5 23/25] sin:[2√6/5 2√6/5 4√6/25]",          // √6
		"abc:[5 5 3] cos:[3/10 3/10 41/50] sin:[√91/10 √91/10 3√91/50]",     // √91
		"abc:[5 5 4] cos:[2/5 2/5 17/25] sin:[√21/5 √21/5 4√21/25]",         // √21
	} {
		if got := ts.tris[pos].String(); got != exp {
			t.Fatalf("Tris32 got %s exp %s", got, exp)		
		}
	}

	comp180, _ := ts.newQ32(1, 0)             // sin(0)         = 180°     430 pairs (6 sec aprox)
	comp90,  _ := ts.newQ32(1, 1)             // sin(1)         =  90°      25
	comp60,  _ := ts.newQ32(2, 0, 1, 3)       // sin((√3)/2)    =  60°      74
	comp54,  _ := ts.newQ32(4, 1, 1, 5)       // sin((1+√5)/4)  =  54°       0
	comp45,  _ := ts.newQ32(2, 0, 1, 2)       // sin(√2/2)      =  45°       0
	comp30,  _ := ts.newQ32(2, 1)             // sin(1/2)       =  30°      15
	comp18,  _ := ts.newQ32(4,-1, 1, 5)       // sin((-1+√5)/4) =  18°       0
	comp15,  _ := ts.newQ32(4, 0,-1, 2, 1, 6) // sin((-√2+√6)4) =  15°       0

	// add two triangle angles pairs sines to get new angle and print matching above sines'
	_ = comp180
	_ = comp90
	_ = comp60
	_ = comp54
	_ = comp45
	_ = comp30
	_ = comp18
	_ = comp15

	ps := NewTriPairs(ts)
	sin := comp60
	ps.Sin(sin)
	t.Logf("     Pairs: %d filtered by sin=%v", len(ps.pairs), sin)
	t.Logf("First pair: %v", ps.pairs[0])
	t.Logf(" Last pair: %v", ps.pairs[len(ps.pairs) - 1])

	qs := NewTriQs(ps)
	qs.All()
	t.Logf("     TriQs: %d all", len(qs.triqs))
	t.Logf("First triq: %v", qs.triqs[0])
	t.Logf(" Last triq: %v", qs.triqs[len(qs.triqs) - 1])

}

func TestTris10(t *testing.T) {

	max := 3
	factory := NewN32s()
	ts := NewTris(max, factory)
	ts.SetSinCos()
	n1 := len(ts.tris)
	t.Logf("      Tris: %d", n1)
	t.Logf(" First tri: %v", ts.tris[0])
	t.Logf("  Last tri: %v", ts.tris[n1 - 1])

	ps := NewTriPairs(ts)
	ps.All()
	n2 := len(ps.pairs)
	t.Logf("     Pairs: %d all", n2)
	t.Logf("First pair: %v", ps.pairs[0])
	t.Logf(" Last pair: %v", ps.pairs[n2 - 1])

	qs := NewTriQs(ps)
	qs.All()
	n3 := len(qs.triqs)
	t.Logf("     TriQs: %d all", n3)
	t.Logf("First triq: %v", qs.triqs[0])
	t.Logf(" Last triq: %v", qs.triqs[n3 - 1])
}

