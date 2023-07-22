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
	CosA  *Q
	CosB  *Q
	CosC  *Q
	Sin2A *Q
	Sin2B *Q
	Sin2C *Q
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
	cosZ := func(x, y, z int) *Q {
		// cosines law
		// cosC = (a² + b² - c²) / 2ab
		return NewQ(x*x + y*y - z*z, 2*x*y)
	}
	sin2Z := func(x, y, z int) *Q {
		// sines law
		// sinC = math.Sqrt(4a²b² - (a²+b²-c²)²) / 2ab
		a := 4*x*x*y*y
		b := (x*x + y*y - z*z)
		return NewQ(a - b*b, 4*x*x*y*y)
	}
	return &Triangle {
		SideA: uint(a),
		SideB: uint(b),
		SideC: uint(c),
		CosA:  cosZ(b, c, a),
		CosB:  cosZ(c, a, b),
		CosC:  cosZ(a, b, c),
		Sin2A: sin2Z(b, c, a),
		Sin2B: sin2Z(c, a, b),
		Sin2C: sin2Z(a, b, c),
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

