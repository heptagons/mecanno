package alg

import (
//	"fmt"
)

type D struct {
	ab *B
	cd *R32
}

func NewD(rs *R32s, b, c, d Z, a N) *D {
	if a == 0 {
		return nil // infinite
	}
	if b == 0 {
		(&a).Reduce2(&c)	
	} else {
		(&a).Reduce3(&b, &c)
	}
	ab := NewB(b, a)
	if ab == nil {
		return nil // infinite
	}
	if c == 0 || d == 0 { // degenerated D is B
		return &D{
			ab: ab,
			cd: NewR32zero(),
		}
	}
	if cd := rs.NewR32(c, d); cd == nil {
		return nil // overflow
	} else {
		// after the d simplification, c was increased
		// specially when b is 0, we need to try reduce a and c
		
		ab.Reduce3(cd.out)
		
		return &D{
			ab: ab,
			cd: cd,
		}
	}
}

//	 _____    ____
//	√(b/a) = √(ab)/a
func NewDsqrtB(rs *R32s, b1 Z, a1 N) *D {
	b2 := Z(0)
	c2 := Z(1)
	d2 := Z(a1) * Z(b1)
	return NewD(rs, b2, c2, d2, a1)
}

func (d *D) Str(s *Str) {
	if d == nil || d.ab == nil || d.cd == nil {
		s.Infinite()
		return
	}
	abZero := d.ab.IsZero()
	cdZero := d.cd.IsZero()
	if abZero && cdZero {
		s.Zero()
		return
	}
	a := d.ab.a // denominator
	if a > 1 {
		if !abZero && !cdZero {
			s.WriteString("(")
		}
	}
	if !abZero {
		d.ab.b.Str(s)
	}
	if !cdZero {
		d.cd.Str(s)
	}
	if a > 1 {
		if !abZero && !cdZero {
			s.WriteString(")")
		}
		s.Divisor()
		s.N32(a)
	}
}

func (d *D) String() string {
	s := NewStr()
	d.Str(s)
	return s.String()
}
