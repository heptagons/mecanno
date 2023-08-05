package alg

import (
	"math"
)

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


const N32_MAX = N(0xffffffff)

type N32 uint32 // range 0 - 0xffffffff

// gcd returns the greatest common divisor of 
// this natural and the other given
func (a N32) gcd(b N32) N32 {
	if b == 0 {
		return a
	}
	return b.gcd(a % b)
}

// NatGCD returns the greatest common divisor of two naturals
func NatGCD(a, b N32) N32 {
	if b == 0 {
		return a
	}
	return NatGCD(b, a % b)
}

func (a *N32) Reduce2(b *N32) {
	if g := NatGCD(*a, *b); g > 1 {
		*a /= g
		*b /= g
	}
}

func (a *N32) Reduce3(b, c *N32) {
	if g := NatGCD(NatGCD(*a, *b), *c); g > 1 {
		*a /= g
		*b /= g
		*c /= g
	}
}


// N32s is factory
type N32s struct {
	primes []N32
}

func NewN32s() *N32s {
	value := 0xffff
    f := make([]bool, value)
    for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
        if f[i] == false {
            for j := i * i; j < value; j += i {
                f[j] = true
            }
        }
    }
    primes := make([]N32, 0)
    for i := N32(2); i < N32(value); i++ {
        if f[i] == false {
            primes = append(primes, i)
        }
    }
	return &N32s{
		primes: primes,
	}
}

// CosC returns the rational cosine of the angle C using the law of cosines:
//	       a² + b² - c²
//	cosC = ------------
//	           2ab
func (n *N32s) CosC(a, b, c N32) (num Z32, den N32, overflow bool) {
	num64 := Z(a)*Z(a) + Z(b)*Z(b) - Z(c)*Z(c)
	den64 := 2*N(a)*N(b)
	if den32, nums32, overflow := n.reduceQ(den64, num64); overflow {
		return 0, 0, true
	} else {
		return nums32[0], den32, false
	}
}

// SinC return the algebraic sine of the angle C using the law of sines:
//	       math.Sqrt(4a²b² - (a²+b²-c²)²)
//	sinC = ------------------------------
//	                  2ab 
func (n *N32s) SinC(a, b, c N32) (out, in Z32, den N32, overflow bool) {
	//	stable := N(a+(b+c)) * N(c-(a-b)) * N(c+(a-b)) * N(a+(b-c))
	//	if out, in, ok := t.algs.roiN(1, stable); ok {
	//		next.area = NewAlg(NewRat(int(out), 4), in)
	//	}
	ab := 2*N(a)*N(b)
	aa := Z(a)*Z(a)
	bb := Z(b)*Z(b)
	cc := Z(c)*Z(c)
	i2 := aa + bb - cc
	out, in, overflow = n.reduceRoot(1, Z(ab)*Z(ab) - i2*i2)
	if overflow {
		return
	}
	var nums32 []Z32
	den, nums32, overflow = n.reduceQ(ab, Z(out))
	if overflow {
		return
	}
	out = nums32[0]
	return
}

// reduceQ reduces the quotient (± num0 ± num1 ± num2 ± ... ± numN) / den
func (n *N32s) reduceQ(den N, nums ...Z) (den32 N32, n32s []Z32, overflow bool) {
	if den == 0 {
		return 0, nil, true // infinite
	}
	allNum0 := true
	for _, n := range nums {
		if n != 0 {
			allNum0 = false
		}
	}
	if allNum0 {
		return 1, make([]Z32, len(nums)), false
	}
	// minPos points to the smallest num to reduce primes use
	var minPos int
	var min N
	ns := make([]N, len(nums))
	for p, n := range nums {
		if n > 0 {
			ns[p] = N(+n)
		} else {
			ns[p] = N(-n) // correct sign
		}
		// update minimum always at start p==0
		// or for n not zero and being smaller than prev smallest
		if p == 0 || (n != 0 && ns[p] < min) {
			minPos, min = p, ns[p] // new smallest
		}
	}
	for _, prime := range n.primes {
		p := N(prime)
		if ns[minPos] < p {
			break // done: no more primes to check
		}
		for {
			all := true
			for _, n := range ns {
				if n % p != 0 {
					all = false
					break // prime not common to all
				}
			}
			if all && (den % p == 0) { // reduce
				den /= p
				for pos := range ns {
					ns[pos] /= p
				}
				continue // check same prime again
			}
			break // check next prime
		}
	}
	if den > AZ_MAX {
		return 1, nil, true // overflow
	}
	n32s = make([]Z32, len(ns))
	for p, num := range nums {
		if n := ns[p]; n > AZ_MAX {
			return 0, nil, true // overflow
		} else if num > 0 { // original sign
			n32s[p] = Z32(+n)
		} else {
			n32s[p] = Z32(-n)
		}
	}
	return N32(den), n32s, false
}

// reduceRoot reduce the number o√i 
func (n *N32s) reduceRoot(o, i Z) (o32, i32 Z32, overflow bool) {
	if o == 0 || i == 0 {
		return 0, 0, false
	}
	on := N(o); if o < 0 { on = N(-o) }
	in := N(i); if i < 0 { in = N(-i) }
	for _, prime := range n.primes {
		p := N(prime)
		pp := p*p
		if in < pp {
			break // done: no more primes to check
		}
		for {
			if in % pp == 0 { // reduce
				on *= p
				in /= pp
				continue // check same prime again
			}
			break // check next prime
		}
	}
	if on > AZ_MAX || in > AZ_MAX {
		return 0, 0, true // overflow
	}
	o32 = Z32(on); if o < 0 { o32 = Z32(-o32) }
	i32 = Z32(in); if i < 0 { i32 = Z32(-i32) }
	return o32, i32, false
}











func (a *N32s) Reduce(p ...Z) (r []Z32, overflow bool) {

	n := len(p)
	if n == 2 {
		// reduce c√d
		if c1, d1, overflow := a.reduceRoot(p[0], p[1]); overflow {
			return nil, true
		} else if d1 == 1 { // convert x√+1 into x
			return []Z32{ c1 }, false
		} else {
			return []Z32{ c1, d1 }, false
		}
	} else if n == 4 {
		// first reduce g1√h1
		if g1, h1, overflow := a.reduceRoot(p[2], p[3]); overflow {
			return nil, true // g√h overflow
		} else if g1 == 0 {
			// F2 degerates to F1(e, f)
			return a.Reduce(p[0], p[1]) // Go to reduce e√h
		} else if h1 == +1 {
			// F2 degerates into F1(e, f + g1)
			return a.Reduce(p[0], p[1] + Z(g1))
		} else if e1, fg, overflow := a.roie(p[0], p[1], Z(g1)); overflow { // reduced 
			// reduction e1√(f1+g2) = e√(f+g1) overflow
			return nil, true
		} else {
			f1 := fg[0]
			g2 := fg[1]
			return []Z32 { e1, f1, g2, h1 }, false
		}
	} else if n == 8 {
		// ijklmnop
		if mnop, overflow := a.Reduce(p[4], p[5], p[6], p[7]); overflow {
			// m√(n+o√p) oveflow
			return nil, true
		} else if m1 := mnop[0]; m1 == 0 {
			// Degenerates to i√(j+k√l)
			return a.Reduce(p[0], p[1], p[2], p[3])
		} else if n1 := mnop[1]; n1 == 0 { // means n1 = +1
			// Degenerates to i√(j+m1+k√l)
			return a.Reduce(p[0], p[1] + Z(m1), p[2], p[3])

		} else if kl, overflow := a.Reduce(p[2], p[3]); overflow {
			// k√l overflow
			return nil, true

		} else if i1, jkm, overflow := a.roie(p[0], p[1], Z(kl[0]), Z(m1)); overflow { // reduced 
			// reduction i√(j+k+m1) overflow
			return nil, true
		} else {
			var l1, n1, o1, p1 Z32
			if len(kl)   > 1 { l1 = kl[1]   }
			if len(mnop) > 1 { n1 = mnop[1] }
			if len(mnop) > 2 { o1 = mnop[2] }
			if len(mnop) > 3 { p1 = mnop[3] }
			return []Z32 {
				i1,
				jkm[0],  // j1
				jkm[1],  // k1
				l1,
				jkm[2],  // m2
				n1,
				o1,
				p1,
			}, false
		}
	}
	return nil, false
}




func (a *N32s) roie(o Z, is ...Z) (o32 Z32, i32s []Z32, overflow bool) {
	ins0 := true
	for _, i := range is {
		if i != 0 {
			ins0 = false
		}
	}
	if o == 0 || ins0 {
		return // zero
	}
	on := N(+o); if o < 0 { on = N(-on) }
	ins := make([]N, len(is))
	// insMaxPos points to the greatest in to reduce primes use
	insMaxPos, insMax := 0, N(0)
	for p, i := range is {
		if i > 0 { ins[p] = N(+i) } else { ins[p] = N(-i) } // correct sign
		if ins[p] > insMax { insMaxPos, insMax = p, ins[p] } // new greatest
	}
	for _, prime := range a.primes {
		p := N(prime)
		pp := p*p
		if ins[insMaxPos] < pp {
			break // done: no more primes to check
		}
		for {
			all := true
			for _, i := range ins {
				if i % pp != 0 {
					all = false
					break // at least one has no this pp factor
				}
			}
			if all { // reduce
				on *= p
				for x := range ins { 
					ins[x] /= pp
				}
				continue // check same prime again
			}
			break // check next prime
		}
	}
	if on > AZ_MAX {
		return 0, nil, true // overflow
	} else if o > 0 { // origin sign
		o32 = Z32(+on)
	} else {
		o32 = Z32(-on)
	}
	i32s = make([]Z32, len(ins))
	for p := range is {
		if i := ins[p]; i > AZ_MAX {
			return 0, nil, true // overflow
		} else if is[p] > 0 { // original sign
			i32s[p] = Z32(+i)
		} else {
			i32s[p] = Z32(-i)
		}
	}
	return o32, i32s, false
}














