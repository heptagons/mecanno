package alg

import (
	"testing"
	"fmt"
)

func TestTri20(t *testing.T) {
	max := 20
	factory := NewN32s()
	ts := NewA32Tris(max, factory)
	if got, exp := len(ts.list), 658; got != exp {
		t.Fatalf("Tris32 max:%d got:%d exp:%d", max, got, exp)
	}
	ts.setSinCos()

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
		if got := ts.list[pos].String(); got != exp {
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

	i := 0
	ts.AddPairs(func(pair *TriPair32, err error) {
		if pair == nil {
			return
		}
		t1, t2 := pair.tA, pair.tB
		a1, a2 := pair.pA, pair.pB
		if err != nil {
			t.Fatalf("%v(%d) %v(%d) %v", t1.abc, a1, t2.abc, a2, err)
		} else if pair.sin.Equal(comp30) {
			i++
			fmt.Printf("% 3d %v(%d) %v(%d) {%s + %s} sin=%s cos=%s\n", i, t1.abc, a1, t2.abc, a2, t1.sin[a1], t2.sin[a2], pair.sin, pair.cos)
		}
	})
}


func TestTri5(t *testing.T) {
	max := 4
	factory := NewN32s()
	ts := NewA32Tris(max, factory)
	ts.setSinCos()
	pairs := make([]*TriPair32, 0)
	comp180, _ := ts.newQ32(1, 0)
	ts.AddPairs(func(pair *TriPair32, err error) {
		if pair == nil {
			return
		}
		t1, t2 := pair.tA, pair.tB
		a1, a2 := pair.pA, pair.pB
		if err != nil {
			t.Fatalf("%v(%d) %v(%d) %v", t1.abc, a1, t2.abc, a2, err)
		} else if pair.sin.Equal(comp180) {
			// Dont count natural (again) new sides
		} else {
			pairs = append(pairs, pair)
		}
	})
	for c, pair := range pairs {
		fmt.Printf("% 3d %v\n", c+1, pair)
	}
}
