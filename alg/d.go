package alg

type D struct {
	ab *B
	cd *R32
}

func NewD(nats *N32s, a, b, c, d Z) *D {

	(&a).Reduce3(&b, &c)

	if ab := NewB(b, a); ab == nil {
		return nil // overflow
	} else if d < 0 {
		return nil // imaginary
	} else if cd := NewR32(nats, c, uint64(d)); cd != nil {
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

func (d *D) Str(s *Str) {
	if d == nil || d.ab == nil || d.cd == nil {
		s.Infinite()
	} else {
		a := d.ab.a // denominator
		if a > 1 {
			s.WriteString("(")
		}
		if b := d.ab.b; b != nil {
			b.Str(s)
		}
		d.cd.Str(s)
		if a > 1 {
			s.N32(a)
			s.WriteString(")")
		}
	}
}
