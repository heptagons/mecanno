package angles

import (
	"fmt"
	"testing"
)

func TestQ(t *testing.T) {

	for _, r := range []struct { q *Q; exp string } {
		{ q:NewQ( 1, 0), exp:   ""     },
		{ q:NewQ( 0, 1), exp:   "0"    },
		{ q:NewQ(44,33), exp:   "4/3"  },
		{ q:NewQ( 5, 7), exp:   "5/7"  },
		{ q:NewQ(99, 6), exp:  "33/2"  },
		{ q:NewQ(66,33), exp:   "2"    },
		{ q:NewQ(14,21), exp:   "2/3"  },
		{ q:NewQ(20,23), exp:  "20/23" },

		{ q:NewQ(-1,10), exp:  "-1/10" },
		{ q:NewQ(-10,1), exp: "-10"    },
		{ q:NewQ(-1,-1), exp:   "1"    },

		{ q:NewQ(2*3*5*7*11*13,2*3*5*7*11*13*17), exp:"1/17" },
	} {
		if got := r.q.String(); got != r.exp {
			t.Fatalf("got: %s exp: %s", got, r.exp)	
		}
	}
}

func TestTriangle(t *testing.T) {
	// test errors
	for _, triangle := range []*Triangle{
		NewTriangle(0,1,1), // a = 0
		NewTriangle(1,0,1), // b = 0
		NewTriangle(1,1,0), // c = 0
		NewTriangle(1,2,1), // b > a
		NewTriangle(1,1,2), // c > b
		NewTriangle(2,1,1), // a >= c+b (area zero)
	} {
		if triangle != nil {
			t.Fatalf("exp nil got %v", t)
		}
	}
}

func TestTriangles(t *testing.T) {

	triangles := NewTriangles()

	for a := 1; a <= 5; a++ {
		for b := 1; b <= a; b++ {
			for c := 1; c <= b; c++ {
				if next := triangles.Add(a, b, c); next != nil {
					fmt.Println(len(triangles), next)
				}
			}
		}
	}
	exp := []string {
		"a=1,b=1,c=1,cosA=1/2,cosB=1/2,cosC=1/2",       // equilateral
		"a=2,b=2,c=1,cosA=1/4,cosB=1/4,cosC=7/8",       // isosceles
		"a=3,b=2,c=2,cosA=-1/8,cosB=3/4,cosC=3/4",      // isosceles
		"a=3,b=3,c=1,cosA=1/6,cosB=1/6,cosC=17/18",     // isosceles
		"a=3,b=3,c=2,cosA=1/3,cosB=1/3,cosC=7/9",       // isosceles
		"a=4,b=3,c=2,cosA=-1/4,cosB=11/16,cosC=7/8",    // scalene
		"a=4,b=3,c=3,cosA=1/9,cosB=2/3,cosC=2/3",
		"a=4,b=4,c=1,cosA=1/8,cosB=1/8,cosC=31/32",
		"a=4,b=4,c=3,cosA=3/8,cosB=3/8,cosC=23/32",
		"a=5,b=3,c=3,cosA=-7/18,cosB=5/6,cosC=5/6",
		"a=5,b=4,c=2,cosA=-5/16,cosB=13/20,cosC=37/40", // scalene
		"a=5,b=4,c=3,cosA=0,cosB=3/5,cosC=4/5",         // pythagoras
		"a=5,b=4,c=4,cosA=7/32,cosB=5/8,cosC=5/8",
		"a=5,b=5,c=1,cosA=1/10,cosB=1/10,cosC=49/50",
		"a=5,b=5,c=2,cosA=1/5,cosB=1/5,cosC=23/25",
		"a=5,b=5,c=3,cosA=3/10,cosB=3/10,cosC=41/50",
		"a=5,b=5,c=4,cosA=2/5,cosB=2/5,cosC=17/25",
	}
	if got := len(triangles); got != len(exp) {
		t.Fatalf("triangles got:%d exp:%d", got, len(exp))
	}
	for i, exp := range exp {
		if got := triangles[i].String(); got != exp {
			t.Fatalf("got %s exp:%s", got, exp)
		}
	}
}