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

func TestTriangles(t *testing.T) {
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

	// test valid
	for exp, triangle := range map[string]*Triangle {
		"a=1,b=1,c=1,cosA=1/2,cosB=1/2,cosC=1/2":    NewTriangle(1,1,1), // equilateral
		"a=3,b=2,c=2,cosA=-1/8,cosB=3/4,cosC=3/4":   NewTriangle(3,2,2), // isosceles 
		"a=3,b=3,c=2,cosA=1/3,cosB=1/3,cosC=7/9":    NewTriangle(3,3,2), // isosceles
		"a=4,b=3,c=2,cosA=-1/4,cosB=11/16,cosC=7/8": NewTriangle(4,3,2), // scalene
		"a=5,b=4,c=3,cosA=0,cosB=3/5,cosC=4/5":      NewTriangle(5,4,3), // pythagoras
	} {
		if got := triangle.String(); exp != got {
			t.Fatalf("got: %s exp: %s", got, exp)
		}
	}
}

func TestTrianglesGroup(t *testing.T) {

	triangles := make([]*Triangle, 0)

	add := func(next *Triangle) {
		if next == nil {
			return
		}
		gcd := uint(next.Gcd())
		a, b, c := next.SideA / gcd, next.SideB / gcd, next.SideC / gcd
		for _, prev := range triangles {
			if prev.SideA == a && prev.SideB == b && prev.SideC == c {
				return
			}
		}
		triangles = append(triangles, next)
		fmt.Println(len(triangles), next)
	}

	for a := 1; a <= 10; a++ {
		for b := 1; b <= a; b++ {
			for c := 1; c <= b; c++ {
				add(NewTriangle(a, b, c))
			}
		}
	}
}