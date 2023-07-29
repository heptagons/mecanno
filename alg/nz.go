package alg

type N uint64

func Ngcd(a, b N) N {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a > b {
		return Ngcd(b, a % b)
	}
	return Ngcd(a, b % a)
}

func (a *N) Reduce2(b *Z) N {
	var bb N
	if *b < 0 {
		bb = N(-*b)
	} else {
		bb = N(*b)
	}
	g := Ngcd(*a, bb)
	if g > 1 {
		*a /= g
		*b /= Z(g)
	}
	return g
}

func (a *N) Reduce3(b, c *Z) N {
	var bb, cc N
	if *b < 0 {
		bb = N(-*b)
	} else {
		bb = N(*b)
	}
	if *c < 0 {
		cc = N(-*c)
	} else {
		cc = N(*c)
	}
	g := Ngcd(Ngcd(*a, bb), cc)
	if g > 1 {
		*a /= g
		*b /= Z(g)
		*c /= Z(g)
	}
	return g
}

func (a *N) Reduce4(b, c, e *Z) N {
	var bb, cc, ee N
	if *b < 0 {
		bb = N(-*b)
	} else {
		bb = N(*b)
	}
	if *c < 0 {
		cc = N(-*c)
	} else {
		cc = N(*c)
	}
	if *e < 0 {
		ee = N(-*e)
	} else {
		ee = N(*e)
	}
	g := Ngcd(Ngcd(Ngcd(*a, bb), cc), ee)
	if g > 1 {
		*a /= g
		*b /= Z(g)
		*c /= Z(g)
		*e /= Z(g)
	}
	return g
}

type Z int64




