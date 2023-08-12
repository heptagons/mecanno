package alg

import (
	"fmt"
	"testing"
)


func TestTri1s(t *testing.T) {
	max := 20
	ts := NewTris(max)
	ts.triSinCos()
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
}

func TestTri2s(t *testing.T) {
	max := 12
	ts := NewTris(max)
	ts.triSinCos()
	exp := len(ts.tris)
	fmt.Printf("  Tris: %d\n", exp)
	fmt.Printf(" First: %v\n", ts.tris[0])
	fmt.Printf("  Last: %v\n", ts.tris[exp-1])

	comp180, _ := ts.qNew(1, 0)             // sin(0)         = 180°     430 pairs (6 sec aprox)
	comp90,  _ := ts.qNew(1, 1)             // sin(1)         =  90°      25
	comp60,  _ := ts.qNew(2, 0, 1, 3)       // sin((√3)/2)    =  60°      74
	comp54,  _ := ts.qNew(4, 1, 1, 5)       // sin((1+√5)/4)  =  54°       0
	comp45,  _ := ts.qNew(2, 0, 1, 2)       // sin(√2/2)      =  45°       0
	comp30,  _ := ts.qNew(2, 1)             // sin(1/2)       =  30°      15
	comp18,  _ := ts.qNew(4,-1, 1, 5)       // sin((-1+√5)/4) =  18°       0
	comp15,  _ := ts.qNew(4, 0,-1, 2, 1, 6) // sin((-√2+√6)4) =  15°       0

	// add two triangle angles pairs sines to get new angle and print matching above sines'
	for _, s := range []struct { sin *Q32; angle string } {
		{ comp180, "180" },
		{ comp90,   "90" },
		{ comp60,   "60" },
		{ comp54,   "54" },
		{ comp45,   "45" },
		{ comp30,   "30" },
		{ comp18,   "18" },
		{ comp15,   "15" },
	} {
		ps := NewTri2s(ts)
		ps.tri2NewEqualSin(s.sin)
		fmt.Printf("-----\n")
		fmt.Printf("Pairs: %d filtered by sin=%v (%s°)\n", len(ps.pairs), s.sin, s.angle)
		if len(ps.pairs) == 0 {
			continue
		}
		fmt.Printf("First: %v\n", ps.pairs[0])
		fmt.Printf(" Last: %v\n", ps.pairs[len(ps.pairs) - 1])
	}
}

func TestTri123(t *testing.T) {
	max := 4
	t1s := NewTris(max)
	t1s.triSinCos()
	for i, tri := range t1s.tris {
		fmt.Printf("Tri % 3d %v\n", i, tri)
	}
	t2s := NewTri2s(t1s)
	t2s.tri2NewAll()
	for i, pair := range t2s.pairs {
		fmt.Printf("Pair % 3d %v\n", i, pair)
	}
	t3s := NewTriQs(t2s)
	t3s.triqsAll()
	for i, triq := range t3s.triqs {
		fmt.Printf("Triq % 3d %v\n", i, triq)
	}
}

// max  Tris  Pairs  TriQs  errs
// ---  ----  -----  -----------
//   1     1      1      1     0
//   2     2      4      7     0
//   3     5     13     25     0
//   4     9     36     77     0
//   5    17    168    158    99 can't sqrt of b+c√d
//   6    24    370    304   245  "" 
//   7    39    672    592   425  ""
//   8    53   1232    922   833  ""
//   9    74   1875   1512  1211  ""
//  10    94   2572   2377  1589  ""
//  11   129   3783   3842  2219  ""

//  15   294  18772  12663 13600  ""

//  20   658  71780  39573 56000  ""
func TestTri3s(t *testing.T) {
	max := 3
	t1s := NewTris(max)
	t1s.triSinCos()
	n1 := len(t1s.tris)
	fmt.Printf("      Tris: %d\n", n1)
	fmt.Printf(" First tri: %v\n", t1s.tris[0])
	fmt.Printf("  Last tri: %v\n", t1s.tris[n1 - 1])

	t2s := NewTri2s(t1s)
	sin0, _ := t1s.qNew(1, 0) // sin(0)= 180°
	t2s.tri2NewNotEqualSin(sin0)
	n2 := len(t2s.pairs)
	fmt.Printf("     Pairs: %d no sin0\n", n2)
	fmt.Printf("First pair: %v\n", t2s.pairs[0])
	fmt.Printf(" Last pair: %v\n", t2s.pairs[n2 - 1])

	t3s := NewTriQs(t2s)
	t3s.triqsAll()
	//for i, triq := range qs.triqs {
	//	fmt.Printf("% 3d %v\n", i+1, triq)
	//}
	n3 := len(t3s.triqs)
	fmt.Printf("     TriQs: %d all errs:%d\n", n3, len(t3s.errs))
	fmt.Printf("First triq: %v\n", t3s.triqs[0])
	fmt.Printf(" Last triq: %v\n", t3s.triqs[n3 - 1])
}
