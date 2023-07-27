package alg

// I is an integer of 32 bits
type I32 struct {
	s bool
	n N32
}

func newI32plus(n N32) *I32 {
	return &I32{ s:false, n:n }
}

func newI32minus(n N32) *I32 {
	return &I32{ s:true, n:n }	
}

func (i *I32) mul(n N32) Z {
	if i.s {
		return -Z(n) * Z(i.n)
	} else {
		return +Z(n) * Z(i.n)
	}
}

func (i *I32) Str(s *Str) {
	if i == nil || i.n == 0 {
		s.Zero()
	} else {
		s.Sign(i.s)
		s.N32(i.n)
	}
}







