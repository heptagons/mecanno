package alg

type Z int64


func gcd64(a, b uint64) uint64 {
	if b == 0 {
		return a
	}
	return gcd64(b, a % b)
}

func (a *Z) Reduce2(b *Z) {
	var aa, bb uint64
	if 0 > *a {
		aa = uint64(-*a)
	} else {
		aa = uint64(*a)
	}
	if 0 > *b {
		bb = uint64(-*b)
	} else {
		bb = uint64(*b)
	}
	if g := Z(gcd64(aa, bb)); g > 1 {
		*a /= g
		*b /= g
	}
}

func (a *Z) Reduce3(b, c *Z) {
	var aa, bb, cc uint64
	if 0 > *a {
		aa = uint64(-*a)
	} else {
		aa = uint64(*a)
	}
	if 0 > *b {
		bb = uint64(-*b)
	} else {
		bb = uint64(*b)
	}
	if 0 > *c {
		cc = uint64(-*c)
	} else {
		cc = uint64(*c)
	}
	if g := Z(gcd64(gcd64(aa, bb), cc)); g > 1 {
		*a /= g
		*b /= g
		*c /= g
	}
}



