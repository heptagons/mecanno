package alg

type N uint64

func Ngcd(a, b N) N {
	if b == 0 {
		return a
	}
	return Ngcd(b, a % b)
}

func (a *N) Reduce2(b *Z) {
	var bb N
	if 0 > *b {
		bb = N(-*b)
	} else {
		bb = N(*b)
	}
	if g := Ngcd(*a, bb); g > 1 {
		*a /= g
		*b /= Z(g)
	}
}

func (a *N) Reduce3(b, c *Z) {
	var bb, cc N
	if 0 > *b {
		bb = N(-*b)
	} else {
		bb = N(*b)
	}
	if 0 > *c {
		cc = N(-*c)
	} else {
		cc = N(*c)
	}
	if g := Ngcd(Ngcd(*a, bb), cc); g > 1 {
		*a /= g
		*b /= Z(g)
		*c /= Z(g)
	}
}

type Z int64




