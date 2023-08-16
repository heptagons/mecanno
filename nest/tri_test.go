package nest

import (
	"fmt"
	"testing"
)


func TestTri1(t *testing.T) {
	max := 20
	ts := NewTriF(max)
	ts.triSinCos()
	got, exp := len(ts.t1s), 825
	if got != exp {
		t.Fatalf("Tris32 max:%d got:%d exp:%d", max, got, exp)
	}
	t.Logf(" Tris: %d", exp)
	t.Logf("First: %v", ts.t1s[0])
	t.Logf(" Last: %v", ts.t1s[exp-1])

	for pos, exp := range []string {
		"[1 1 1] cos:[1/2 1/2 1/2] sin:[√3/2 √3/2 √3/2]",                // √3
		"[2 2 1] cos:[1/4 1/4 7/8] sin:[√15/4 √15/4 √15/8]",             // √15
		"[2 2 2] pri:[1 1 1]",               
		"[3 2 2] cos:[-1/8 3/4 3/4] sin:[3√7/8 √7/4 √7/4]",              // √7
		"[3 3 1] cos:[1/6 1/6 17/18] sin:[√35/6 √35/6 √35/18]",          // √35
		"[3 3 2] cos:[1/3 1/3 7/9] sin:[2√2/3 2√2/3 4√2/9]",             // √2
		"[3 3 3] pri:[1 1 1]",           
		"[4 3 2] cos:[-1/4 11/16 7/8] sin:[√15/4 3√15/16 √15/8]",        // √15
		"[4 3 3] cos:[1/9 2/3 2/3] sin:[4√5/9 √5/3 √5/3]",               // √5
		"[4 4 1] cos:[1/8 1/8 31/32] sin:[3√7/8 3√7/8 3√7/32]",          // √7
		"[4 4 2] pri:[2 2 1]",
		"[4 4 3] cos:[3/8 3/8 23/32] sin:[√55/8 √55/8 3√55/32]",         // √55
		"[4 4 4] pri:[1 1 1]",
		"[5 3 3] cos:[-7/18 5/6 5/6] sin:[5√11/18 √11/6 √11/6]",         // √11
		"[5 4 2] cos:[-5/16 13/20 37/40] sin:[√231/16 √231/20 √231/40]", // √231
		"[5 4 3] cos:[0 3/5 4/5] sin:[1 4/5 3/5]",                       // √1
		"[5 4 4] cos:[7/32 5/8 5/8] sin:[5√39/32 √39/8 √39/8]",          // √39
		"[5 5 1] cos:[1/10 1/10 49/50] sin:[3√11/10 3√11/10 3√11/50]",   // √11
		"[5 5 2] cos:[1/5 1/5 23/25] sin:[2√6/5 2√6/5 4√6/25]",          // √6
		"[5 5 3] cos:[3/10 3/10 41/50] sin:[√91/10 √91/10 3√91/50]",     // √91
		"[5 5 4] cos:[2/5 2/5 17/25] sin:[√21/5 √21/5 4√21/25]",         // √21
	} {
		if got := ts.t1s[pos].String(); got != exp {
			t.Fatalf("Tris got %s exp %s", got, exp)		
		} else {
			t.Log(got)
		}
	}
}

func TestTri2(t *testing.T) {
	max := 12
	ts := NewTriF(max)
	ts.triSinCos()
	exp := len(ts.t1s)
	fmt.Printf(" Tris: %d\n", exp)
	fmt.Printf("First: %v\n", ts.t1s[0])
	fmt.Printf(" Last: %v\n", ts.t1s[exp-1])

	comp180, _ := ts.aNew(1, 0)             // sin(0)         = 180°     430 pairs (6 sec aprox)
	comp90,  _ := ts.aNew(1, 1)             // sin(1)         =  90°      25
	comp60,  _ := ts.aNew(2, 0, 1, 3)       // sin((√3)/2)    =  60°      74
	comp54,  _ := ts.aNew(4, 1, 1, 5)       // sin((1+√5)/4)  =  54°       0
	comp45,  _ := ts.aNew(2, 0, 1, 2)       // sin(√2/2)      =  45°       0
	comp30,  _ := ts.aNew(2, 1)             // sin(1/2)       =  30°      15
	comp18,  _ := ts.aNew(4,-1, 1, 5)       // sin((-1+√5)/4) =  18°       0
	comp15,  _ := ts.aNew(4, 0,-1, 2, 1, 6) // sin((-√2+√6)4) =  15°       0

	// add two triangle angles pairs sines to get new angle and print matching above sines'
	for _, s := range []struct { sin *A32; angle string } {
		{ comp180, "180" },
		{ comp90,   "90" },
		{ comp60,   "60" },
		{ comp54,   "54" },
		{ comp45,   "45" },
		{ comp30,   "30" },
		{ comp18,   "18" },
		{ comp15,   "15" },
	} {
		ps := NewTri2F(ts)
		ps.tri2NewEqualSin(s.sin)
		fmt.Printf("-----\n")
		fmt.Printf("Tri2s: %d filtered by sin=%v (%s°)\n", len(ps.tri2s), s.sin, s.angle)
		if len(ps.tri2s) == 0 {
			continue
		}
		fmt.Printf("First: %v\n", ps.tri2s[0])
		fmt.Printf(" Last: %v\n", ps.tri2s[len(ps.tri2s) - 1])
	}
}


// max  Tris  Tri2s  Tri3s  errs
// ---  ----  -----  -----------
//   1     1      1      1     0
//   2     2      6      7     0
//   3     5     45     25     0
//   4     9    169     77     0
//   5    17    663    158     0
//   6    24   1481    304     0  "" 
//   7    39    672    592   425  ""
//   8    53   1232    922   833  ""
//   9    74   1875   1512  1211  ""
//  10    94   2572   2377  1589  ""
//  11   129   3783   3842  2219  ""

//  15   294  18772  12663 13600  ""

//  20   658  71780  39573 56000  ""
func TestTri3(t *testing.T) {
	max := 6
	t1s := NewTriF(max)
	t1s.triSinCos()
	n1 := len(t1s.t1s)
	fmt.Printf("  Tris: %d\n", n1)
	fmt.Printf(" First: %v\n", t1s.t1s[0])
	fmt.Printf("  Last: %v\n", t1s.t1s[n1 - 1])

	t2s := NewTri2F(t1s)
	sin0, _ := t1s.aNew(1, 0) // sin(0)= 180°
	t2s.tri2NewNotEqualSin(sin0)
	n2 := len(t2s.tri2s)
	fmt.Printf(" Tri2: %d no sin0\n", n2)
	fmt.Printf("First: %v\n", t2s.tri2s[0])
	fmt.Printf(" Last: %v\n", t2s.tri2s[n2 - 1])

	t3s := NewTri3F(t2s)
	t3s.tri3All()
	//for i, triq := range qs.triqs {
	//	fmt.Printf("% 3d %v\n", i+1, triq)
	//}
	n3 := len(t3s.tri3s)
	fmt.Printf("Tri3s: %d all errs:%d\n", n3, len(t3s.errs))
	fmt.Printf("First: %v\n", t3s.tri3s[0])
	fmt.Printf(" Last: %v\n", t3s.tri3s[n3 - 1])
}

func TestTriList(t *testing.T) {
	max := 4
	t1s := NewTriF(max)
	t1s.triSinCos()
	for i, tri := range t1s.t1s {
		fmt.Printf("T1 % 3d %v\n", i, tri)
	}
	t2s := NewTri2F(t1s)
	t2s.tri2NewAll()
	for i, tri2 := range t2s.tri2s {
		fmt.Printf("T2 % 3d %v\n", i, tri2)
	}
	t3s := NewTri3F(t2s)
	t3s.tri3All()
	for i, tri3 := range t3s.tri3s {
		fmt.Printf("T3 % 3d %v\n", i, tri3)
	}
}
