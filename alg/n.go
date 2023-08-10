package alg

import (
	"fmt"
	"math"
)

var ErrOverflow = fmt.Errorf("Overflow")
var ErrInfinite = fmt.Errorf("Infinite")
var ErrInvalid  = fmt.Errorf("Invalid")

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

// nFrac returns simplified denominator and numerator
func (n *N32s) nFrac(den N, num Z) (den32 N32, n32 Z32, err error) {
	if den == 0 {
		return 0, 0, ErrInfinite
	} else if num == 0 {
		return 1, 0, nil // zero
	}
	dn := den
	nn := N(num); if num < 0 { nn = N(-num) } // convert numerator to natural
	min := dn; if nn < min { min = nn }       // min(den, N(num))
	for _, prime := range n.primes {
		p := N(prime)
		if min < p {
			break // done: no more primes to check
		}
		for {
			if dn % p == 0 && nn % p == 0 { // reduce
				dn /= p
				nn /= p
				continue // check same prime again
			}
			break // check next prime
		}
	}
	if dn > AZ_MAX || nn > AZ_MAX {
		return 0, 0, ErrOverflow
	}
	if num > 0 { // original sign
		n32 = Z32(+nn)
	} else {
		n32 = Z32(-nn)
	}
	return N32(dn), n32, nil
}

// nFracN reduces the quotient (± num0 ± num1 ± num2 ± ... ± numN) / den
func (n *N32s) nFracN(den N, nums ...Z) (den32 N32, n32s []Z32, err error) {
	if den == 0 {
		return 0, nil, ErrInfinite
	}
	if len(nums) == 0 {
		return 0, nil, nil
	}
	allNum0 := true
	for _, n := range nums {
		if n != 0 {
			allNum0 = false
		}
	}
	if allNum0 {
		return 1, make([]Z32, len(nums)), nil
	}
	// minPos points to the smallest num to reduce primes use
	minPos := 0
	min := Z(N32_MAX)
	ns := make([]N, len(nums))
	for p, n := range nums {
		if n >= 0 {
			ns[p] = N(+n)
		} else {
			ns[p] = N(-n) // correct sign
		}
		// update minimum always at start p==0
		// or for n not zero and being smaller than prev smallest
		if n != 0 && n < min {
			minPos, min = p, n // new smallest
		}
	}
//fmt.Println("reduceQn1", den, ns, ns[minPos])
	for _, prime := range n.primes {
		p := N(prime)
		if ns[minPos] < p {
			break // done: no more primes to check
		}
		for {
			all := true
			for _, n := range ns {
				if (n != 0) && (n % p != 0) {
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
		return 1, nil, ErrOverflow
	}
	n32s = make([]Z32, len(ns))
	for p, num := range nums {
		if n := ns[p]; n > AZ_MAX {
			return 0, nil, ErrOverflow
		} else if num > 0 { // original sign
			n32s[p] = Z32(+n)
		} else {
			n32s[p] = Z32(-n)
		}
	}
//fmt.Println("reduceQn2", den, n32s)
	return N32(den), n32s, nil
}

// nSqrt reduce the number o√i 
func (n *N32s) nSqrt(o, i Z) (o32, i32 Z32, err error) {
	if o == 0 || i == 0 {
		return 0, 0, nil
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
		return 0, 0, ErrOverflow
	}
	o32 = Z32(on); if o < 0 { o32 = Z32(-o32) }
	i32 = Z32(in); if i < 0 { i32 = Z32(-i32) }
	return o32, i32, nil
}

func (a *N32s) nSqrtN(o Z, is ...Z) (o32 Z32, i32s []Z32, err error) {
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
		return 0, nil, ErrOverflow
	} else if o > 0 { // origin sign
		o32 = Z32(+on)
	} else {
		o32 = Z32(-on)
	}
	i32s = make([]Z32, len(ins))
	for p := range is {
		if i := ins[p]; i > AZ_MAX {
			return 0, nil, ErrOverflow
		} else if is[p] > 0 { // original sign
			i32s[p] = Z32(+i)
		} else {
			i32s[p] = Z32(-i)
		}
	}
	return o32, i32s, nil
}




/*
func (a *N32s) Reduce(p ...Z) (r []Z32, err error) {

	n := len(p)
	if n == 2 {
		// reduce c√d
		if c1, d1, err := a.reduceRoot(p[0], p[1]); err != nil {
			return nil, err
		} else if d1 == 1 { // convert x√+1 into x
			return []Z32{ c1 }, nil
		} else {
			return []Z32{ c1, d1 }, nil
		}
	} else if n == 4 {
		// first reduce g1√h1
		if g1, h1, err := a.reduceRoot(p[2], p[3]); err != nil {
			return nil, err // g√h overflow
		} else if g1 == 0 {
			// F2 degerates to F1(e, f)
			return a.Reduce(p[0], p[1]) // Go to reduce e√h
		} else if h1 == +1 {
			// F2 degerates into F1(e, f + g1)
			return a.Reduce(p[0], p[1] + Z(g1))
		} else if e1, fg, err := a.roie(p[0], p[1], Z(g1)); err != nil { // reduced 
			// reduction e1√(f1+g2) = e√(f+g1) overflow
			return nil, err
		} else {
			f1 := fg[0]
			g2 := fg[1]
			return []Z32 { e1, f1, g2, h1 }, nil
		}
	} else if n == 8 {
		// ijklmnop
		if mnop, err := a.Reduce(p[4], p[5], p[6], p[7]); err != nil {
			// m√(n+o√p) oveflow
			return nil, err
		} else if m1 := mnop[0]; m1 == 0 {
			// Degenerates to i√(j+k√l)
			return a.Reduce(p[0], p[1], p[2], p[3])
		} else if n1 := mnop[1]; n1 == 0 { // means n1 = +1
			// Degenerates to i√(j+m1+k√l)
			return a.Reduce(p[0], p[1] + Z(m1), p[2], p[3])

		} else if kl, err := a.Reduce(p[2], p[3]); err != nil {
			// k√l overflow
			return nil, err

		} else if i1, jkm, err := a.roie(p[0], p[1], Z(kl[0]), Z(m1)); err != nil { // reduced 
			// reduction i√(j+k+m1) overflow
			return nil, err
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
			}, nil
		}
	}
	return nil, nil
}
*/



func (a *N32s) roie(o Z, is ...Z) (o32 Z32, i32s []Z32, err error) {
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
		return 0, nil, ErrOverflow
	} else if o > 0 { // origin sign
		o32 = Z32(+on)
	} else {
		o32 = Z32(-on)
	}
	i32s = make([]Z32, len(ins))
	for p := range is {
		if i := ins[p]; i > AZ_MAX {
			return 0, nil, ErrOverflow
		} else if is[p] > 0 { // original sign
			i32s[p] = Z32(+i)
		} else {
			i32s[p] = Z32(-i)
		}
	}
	return o32, i32s, nil
}














