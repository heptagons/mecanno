package alg

type H struct {
	gh *AI32
	ef *AI32
	cd *AI32
	ab *B
}

type Hs struct {
	*AI32s
}

func NewHs(rs *AI32s) *Hs {
	return &Hs{
		AI32s: rs,
	}
}

func (s *Hs) NewH(b, c, d, e, f, g, h Z, a N) *H {
	gh := s.AI(g, h, nil)
	if gh == nil {
		return nil // overflow
	}
	ef := s.AI(e, f, nil) // TODO
	if ef == nil {
		return nil
	}
	ab := NewBnotReduce(b, a)
	return &H {
		gh: gh,
		ab: ab,
	}


	/*
	(&a).Reduce4(&b, &c, &e)
	if ab == nil {
		return nil // infinite
	}
	if c == 0 || d == 0 { // degenerated D is B
		return &D{
			ab: ab,
			cd: NewR32zero(),
		}
	}
	if cd := s.NewR32(c, d); cd == nil {
		return nil // overflow
	} else {
		// after the d simplification, c was increased
		// specially when b is 0, we need to try reduce a and c
		ab.Reduce3(cd.out)
		return &D{
			ab: ab,
			cd: cd,
		}
	}*/
}

func (x *H) String() string {
	str := &Str{}
	x.gh.Str(str)
	return str.String()
}
