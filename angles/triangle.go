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
		return NewQ(x*x + y*y - z*z, 2*x*y)
	}
	return &Triangle {
		SideA: uint(a),
		SideB: uint(b),
		SideC: uint(c),
		CosA:  cosZ(b, c, a),
		CosB:  cosZ(c, a, b),
		CosC:  cosZ(a, b, c),
	}
}

func (t *Triangle) Gcd() int {
	return meccano.Gcd(int(t.SideA), meccano.Gcd(int(t.SideB), int(t.SideC)))
}

func (t *Triangle) String() string {
	return fmt.Sprintf("a=%d,b=%d,c=%d,cosA=%s,cosB=%s,cosC=%s",
		t.SideA, t.SideB, t.SideC, t.CosA, t.CosB, t.CosC)
}
