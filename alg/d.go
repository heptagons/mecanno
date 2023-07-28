package alg

type D struct {
	ab *B
	cd *R32
}

func NewD(nats *N32s, b, c, d, a Z) *D {

	// Reduce a, b and c
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
	if cd := NewR32(nats, c, d); cd == nil {
		return nil // overflow
	} else {
		// after the d simplification, c was increased
		// specially when b is 0, we need to try reduce a and c
		
		//ab.Reduce3(cd.out)
		
		return &D{
			ab: ab,
			cd: cd,
		}
	}
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
		s.Divisor()
		s.N32(a)
		if !abZero && !cdZero {
			s.WriteString(")")
		}
	}
}

func (d *D) String() string {
	s := NewStr()
	d.Str(s)
	return s.String()
}
