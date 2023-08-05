package alg

import (
	"fmt"
)

const AZ_MAX = 0x7fffffff

// AZ32 represent an algebraic integer number
type AZ32 struct {
	o int32
	i []*AZ32
}

// AQ32 represent an algebraic rational number
type AQ32 struct {
	den uint32
	num []*AZ32
}

func newAZ32(p ...int32) *AZ32 {
	if n := len(p); n == 0 {
		return nil
	} else {
		a := &AZ32{}
		if n >= 1 {
			a.o = p[0] // b, c, e, i, ...
		}
		if n >= 2 {
			a.i = make([]*AZ32, 0)
			a.i = append(a.i, newAZ32(p[1])) // d, f, j, ...
		}
		if n >= 4 {
			a.i = append(a.i, newAZ32(p[2:4]...)) // gh, kl, ...
		}
		if n >= 8 {
			a.i = append(a.i, newAZ32(p[4:8]...)) // mnop, ...
		}
		if n >= 16 {
			a.i = append(a.i, newAZ32(p[8:16]...)) // 
		}
		return a
	}
}









// CosC returns the rational cosine of the angle C using the law of cosines:
//	       a² + b² - c²
//	cosC = ------------
//	           2ab
/*
func (a *A32s) cosC(a, b, c N32) *AQ32 {
	num := Z(a)*Z(a) + Z(b)*Z(b) - Z(c)*Z(c)
	den := 2*Z(a)*Z(b)
	return NewRat(num, den)
}*/







func (a *AZ32) out() int32 {
	if a == nil {
		return 0 // was 1
	}
	return a.o
}

func (a *AZ32) in() int32 {
	if len(a.i) == 0 {
		return +1
	}
	return a.i[0].out()
}

func (r *AZ32) String() string {
	if r == nil {
		return ""
	}
	s := NewStr()
	r.Str(s,true)
	return s.String()
}

func (a *AZ32) Str(s *Str, sign bool) {
	if sign {
		s.WriteString(fmt.Sprintf("%+d", a.o))
	} else{
		s.WriteString(fmt.Sprintf("%d", a.o))
	}
	if a.o == 0 {
		return
	}
	if n := len(a.i); n == 0 {
		return
	} else if n == 1 {
		if o := a.i[0].o; o < 0 {
			s.WriteString("i")
			if o != -1 {
				s.WriteString(fmt.Sprintf("√%d", -o))
			}
		} else if o > 1 {
			s.WriteString(fmt.Sprintf("√%d", o))
		}
	} else {
		s.WriteString("√(")
		for p, i := range a.i {
			i.Str(s, p != 0)
		}
		s.WriteString(")")
	}
}

func (a *AQ32) String() string {
	s := NewStr()
	s.WriteString("(")
	for pos, num := range a.num {
		num.Str(s, pos != 0)
	}
	s.WriteString(fmt.Sprintf(")/%d", a.den))
	return s.String()
}



/*
func (a *A32s) rQ(num []*AZ32, den uint64) (q *AQ32, overflow bool) {
	if den32 == 0 {
		return nil, true // infinite
	}
	if n := len(num); n == 1 {
		num32 := num[0].o
		if num32 == 0 {
			return num0, 1, false
		n := N(num); if num < 0 { n := N(-num) }
		d := N(den); if den < 0 { d := N(-den) }
		// greatest common divisor (GCD) via Euclidean algorithm
		for d != 0 {
			t := d
			d = n % d
			n = t
		}
		num /= uint64(n)
		den /= uint64(n)


	}
	return &AQ32 {
		num: num,
		den: den,
	}
}*/



type A32Poly struct { // Polygon
	sides []N32
}

func (a *A32Poly) String() string {
	return fmt.Sprintf("sides:%v", a.sides)
}

type A32Tri struct { // Triangle
	*A32Poly
	cosines []*AQ32
	sines   []*AQ32
}

type A32Tris struct { // Triangles
	list []*A32Tri
}

func NewA32Tris(max int) *A32Tris {
	ts := &A32Tris {
		list: make([]*A32Tri, 0),
	}
	for a := N32(1); a <= N32(max); a++ {
		for b := N32(1); b <= a; b++ {
			for c := N32(1); c <= b; c++ {
				ts.add(a, b, c)
			}
		}
	}
	return ts
}

func (t *A32Tris) add(a, b, c N32) {
	gcd := NatGCD(a, NatGCD(b, c))
	ga, gb, gc := a / gcd, b / gcd, c / gcd
	for _, p := range t.list {
		if p.sides[0] == ga && p.sides[1] == gb && p.sides[2] == gc {
			// scaled version already stored don't append
			return
		}
	}
	t.list = append(t.list, &A32Tri{
		A32Poly: &A32Poly{
			sides: []N32{ a, b, c },
		},
	})
	/*next.cosA = t.algs.CosC(b, c, a)
	next.cosB = t.algs.CosC(c, a, b)
	next.cosC = t.algs.CosC(a, b, c)
	next.sinA = t.algs.SinC(b, c, a)
	next.sinB = t.algs.SinC(c, a, b)
	next.sinC = t.algs.SinC(a, b, c)
	stable := N(a+(b+c)) * N(c-(a-b)) * N(c+(a-b)) * N(a+(b-c))
	if out, in, ok := t.algs.roiN(1, stable); ok {
		next.area = NewAlg(NewRat(int(out), 4), in)
	}
	*/
}
/*
func (t *A32Tris) Cos(a32s *A32s) {
	for _, p := range t.list {
		a, b, c := p.sides[0], p.sides[1], p.sides[2]
		num := Z(a)*Z(a) + Z(b)*Z(b) - Z(c)*Z(c)
	}
}


// SinC return the algebraic sine of the angle C using the law of sines:
//	       math.Sqrt(4a²b² - (a²+b²-c²)²)
//	sinC = ------------------------------
//	                  2ab 
func (algs *Algs) SinC(a, b, c N32) *Alg {
	p := int(4*a*a*b*b)
	q := int((a*a + b*b - c*c))
	d := int(2*a*b)
	if rat := NewRat(p - q*q, d*d); rat == nil {
		return nil
	} else {
		return rat.Sqrt(algs.Red32)
	}
}*/

