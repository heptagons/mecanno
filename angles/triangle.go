package angles

import (
	"fmt"

	"github.com/heptagons/meccano"
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
	SideA uint
	SideB uint
	SideC uint
	CosA  *Rat
	CosB  *Rat
	CosC  *Rat
	Sin2A *Rat
	Sin2B *Rat
	Sin2C *Rat
}

func NewTriangle(a, b, c int) *Triangle {
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
		Sin2A: NewRatSin2C(b, c, a),
		Sin2B: NewRatSin2C(c, a, b),
		Sin2C: NewRatSin2C(a, b, c),
	}
}

func (t *Triangle) String() string {
	return fmt.Sprintf("a=%d,b=%d,c=%d,cosA=%s,cosB=%s,cosC=%s,sin2A=%s,sin2B=%s,sin2C=%s",
		t.SideA, t.SideB, t.SideC,
		t.CosA,  t.CosB,  t.CosC,
		t.Sin2A, t.Sin2B, t.Sin2C)
}

// Triangles is an array of different triangles 
// Scaled triangles are not repeated
type Triangles []*Triangle

func NewTriangles() Triangles {
	t := make([]*Triangle, 0)
	return Triangles(t)
}

func (t *Triangles) Add(a, b, c int) *Triangle {
	next := NewTriangle(a, b, c)
	if next == nil {
		return nil
	}
	gcd := meccano.Gcd(a, meccano.Gcd(b, c))
	ga, gb, gc := uint(a / gcd), uint(b / gcd), uint(c / gcd)
	for _, prev := range *t {
		if prev.SideA == ga && prev.SideB == gb && prev.SideC == gc {
			// scaled version already stored return nothing
			return nil
		}
	}
	*t = append(*t, next)
	return next
}

