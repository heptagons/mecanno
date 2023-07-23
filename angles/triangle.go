package angles

import (
	"fmt"

	"github.com/heptagons/meccano"
)

// Triangles is an array of different triangles 
// Scaled triangles are not repeated
type Triangles struct {
	nats *Nats
	list []*Triangle
}

func NewTriangles() *Triangles {
	return &Triangles{
		nats: NewNats(), // to use Sqrt to calculate sines
		list: make([]*Triangle, 0),
	}	
}

func (t *Triangles) Add(a, b, c int) *Triangle {
	next := NewTriangle(t.nats, a, b, c)
	if next == nil {
		return nil
	}
	gcd := meccano.Gcd(a, meccano.Gcd(b, c))
	ga, gb, gc := uint(a / gcd), uint(b / gcd), uint(c / gcd)
	for _, prev := range t.list {
		if prev.SideA == ga && prev.SideB == gb && prev.SideC == gc {
			// scaled version already stored return nothing
			return nil
		}
	}
	t.list = append(t.list, next)
	return next
}



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
	SideA uint
	SideB uint
	SideC uint
	CosA  *Rat
	CosB  *Rat
	CosC  *Rat
	SinA  *Alg
	SinB  *Alg
	SinC  *Alg
}

func NewTriangle(nats *Nats, a, b, c int) *Triangle {
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
		SideA: uint(a),
		SideB: uint(b),
		SideC: uint(c),
		CosA:  NewRatCosC(b, c, a),
		CosB:  NewRatCosC(c, a, b),
		CosC:  NewRatCosC(a, b, c),
		SinA:  nats.SinC(b, c, a),
		SinB:  nats.SinC(c, a, b),
		SinC:  nats.SinC(a, b, c),
	}
}

func (t *Triangle) String() string {
	return fmt.Sprintf("a=%d,b=%d,c=%d,cosA=%s,cosB=%s,cosC=%s,sinA=%s,sinB=%s,sinC=%s",
		t.SideA, t.SideB, t.SideC,
		t.CosA,  t.CosB,  t.CosC,
		t.SinA,  t.SinB,  t.SinC)
}

