package alg

import (
	"fmt"
)

// Triangle is a valid triangle with positive sides:
//	a >= b >= c > 0
//  a > b+c
//
//           _ -C
//     a _ -   /
//   _ -      /
// B_        / b
//   -_     /
//  c  -_  /  
//       -A
//
// A,B, and C the angles to opposite sides a,b and c.
type Triangle struct {
	// sides are naturals > 0
	a N32
	b N32
	c N32
	// cosines are rational
	cosA *Rat
	cosB *Rat
	cosC *Rat

	// sines are algebraic
	sinA *Alg
	sinB *Alg
	sinC *Alg
	area *Alg
}

func NewTriangle(a, b, c N32) *Triangle {
	if a <= 0 || b <= 0 || c <= 0 {
		return nil // fmt.Errorf("side equals or less than 0")
	} else if a < b {
		return nil // fmt.Errorf("a < b")
	} else if b < c {
		return nil // fmt.Errorf("b < c")	
	} else if a >= b+c {
		return nil // fmt.Errorf("a > b+c")	
	}
	return &Triangle {
		a: a,
		b: b,
		c: c,
	}
}

func (t *Triangle) String() string {
	return fmt.Sprintf("a=%d,b=%d,c=%d,cosA=%s,cosB=%s,cosC=%s,sinA=%s,sinB=%s,sinC=%s",
		t.a, t.b, t.c, t.cosA, t.cosB, t.cosC, t.sinA, t.sinB, t.sinC)
}


// Triangles is an array of different triangles 
// Scaled triangles are not repeated
type Triangles struct {
	algs *Algs
	list []*Triangle
}

func NewTriangles(algs *Algs) *Triangles {
	return &Triangles{
		algs: algs, // to use Sqrt to calculate sines
		list: make([]*Triangle, 0),
	}	
}

func (t *Triangles) Find(max N32) {
	// first valid triangles
	for a := N32(1); a <= max; a++ {
		for b := N32(1); b <= a; b++ {
			for c := N32(1); c <= b; c++ {
				t.Add(a, b, c)
			}
		}
	}
}

func (t *Triangles) Add(a, b, c N32) *Triangle {
	next := NewTriangle(a, b, c)
	if next == nil {
		return nil
	}
	gcd := NatGCD(a, NatGCD(b, c))
	ga, gb, gc := a / gcd, b / gcd, c / gcd
	for _, prev := range t.list {
		if prev.a == ga && prev.b == gb && prev.c == gc {
			// scaled version already stored return nothing
			return nil
		}
	}
	next.cosA = t.algs.CosC(b, c, a)
	next.cosB = t.algs.CosC(c, a, b)
	next.cosC = t.algs.CosC(a, b, c)
	next.sinA = t.algs.SinC(b, c, a)
	next.sinB = t.algs.SinC(c, a, b)
	next.sinC = t.algs.SinC(a, b, c)

	stable := N(a+(b+c)) * N(c-(a-b)) * N(c+(a-b)) * N(a+(b-c))
	if out, in, ok := t.algs.roiN(1, stable); ok {
		next.area = NewAlg(NewRat(int(out), 4), in)
	}

	t.list = append(t.list, next)
	return next
}



