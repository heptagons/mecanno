package alg

type D struct {
	ab *B
	cd *AI32
}

/*func (d *D) Str(s *Str) {
	if d == nil || d.ab == nil || d.cd == nil {
		s.Zero()
		return
	}
	abZero := d.ab.isZero()
	cdZero := d.cd.isZero()
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
		s.Over(a)
	}
}

func (d *D) String() string {
	s := NewStr()
	d.Str(s)
	return s.String()
}*/


type Ds struct {
	*Red32
}

func NewDs(rs *Red32) *Ds {
	return &Ds{
		Red32: rs,
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
	if cd, ok := ds.AI(c, d, nil); !ok {
		return nil // overflow
	} else {
		// after the d simplification, c was increased
		// specially when b is 0, we need to try reduce a and c
		ab.Reduce3(cd.o)
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
/*
func (ds *Ds) NewInvD(dd *D) *D {

	if ab, overflow := dd.ab.a.mul(dd.ab.b); overflow {
		return nil
	}



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
}*/

