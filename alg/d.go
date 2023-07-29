package alg

type D struct {
	ab *B
	cd *R32
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


type Ds struct {
	*R32s
}

func NewDs(rs *R32s) *Ds {
	return &Ds{
		R32s: rs,
	}
}

func (ds *Ds) NewD(b, c, d Z, a N) *D {
	if a == 0 {
		return nil // infinite
	}
	(&a).Reduce3(&b, &c)
	ab := NewBnotReduce(b, a)
	if ab == nil {
		return nil // infinite
	}
	if cd := ds.NewR32(c, d); cd == nil {
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
func (ds *Ds) NewDsqrtB(b Z, a N) *D {
	b2 := Z(0)
	c2 := Z(1)
	d2 := Z(a) * Z(b)
	return ds.NewD(b2, c2, d2, a)
}

// inversion is
//            _
// +- ab -+ac√d
// ------------
//   bb - ccd
func (ds *Ds) NewInvD(dd *D) *D {
	ab     := dd.ab.aVal() * dd.ab.bVal()             // will be the new b
	ac     := dd.ab.aVal() * dd.cd.outVal()           // will be the new c
	d      := dd.cd.inVal()                           // the same d pass as d
	bb_ccd := dd.ab.bValPow2() - dd.cd.outValPow2()*d // will be the new a
	if bb_ccd == 0 {
		return nil
	} else if bb_ccd > 0 {
		return ds.NewD(+ab, -ac, d, N(+bb_ccd))
	} else {
		return ds.NewD(-ab, +ac, d, N(-bb_ccd))
	}
}

