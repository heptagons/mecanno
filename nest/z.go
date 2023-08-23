package nest

type Z int64

type Z32 int32

const Z32_MAX = 0x7fffffff

// Z32s is a factory to operates integers
// Uses factory N32s
type Z32s struct {
	*N32s
}

// NewZ32s creates a new Z32s factory
func NewZ32s() *Z32s {
	return &Z32s{
		N32s:  NewN32s(),
	}
}

// zFrac returns simplified denominator and numerator for rationals
// Example: zFrac(10, 5) returns 2,1,nil:
//    5     1
//	---- = ---
//   10     2
func (z *Z32s) zFrac(den N, num Z) (den32 N32, n32 Z32, err error) {
	if den == 0 {
		return 0, 0, ErrInfinite
	} else if num == 0 {
		return 1, 0, nil // zero
	}
	dn := den
	nn := N(num); if num < 0 { nn = N(-num) } // convert numerator to natural
	min := dn; if nn < min { min = nn }       // min(den, N(num))
	for _, prime := range z.primes {
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
	if dn > Z32_MAX || nn > Z32_MAX {
		return 0, 0, ErrOverflow
	}
	if num > 0 { // original sign
		n32 = Z32(+nn)
	} else {
		n32 = Z32(-nn)
	}
	return N32(dn), n32, nil
}

// zMul returns the result of multiplying of the given arguments.
func (z *N32s) zMul(nums ...Z) (Z32, error) {
	mul := Z(0)
	if len(nums) > 0 {
		mul = 1
	}
	for _, n := range nums {
		mul *= n
		if mul > Z32_MAX || mul < -Z32_MAX {
			return 0, ErrOverflow
		}
	}
	return Z32(mul), nil
}

// zFracN reduces the rationals (± num0 ± num1 ± num2 ± ... ± numN) / den
// Example: zFracN(8, 4, 2) return 4,[2,1],nil:
//   4x + 2y   2x + y
//	-------- = ------
//      8         4
func (z *N32s) zFracN(den N, nums ...Z) (den32 N32, n32s []Z32, err error) {
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
	for _, prime := range z.primes {
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
	if den > Z32_MAX {
		return 1, nil, ErrOverflow
	}
	n32s = make([]Z32, len(ns))
	for p, num := range nums {
		if n := ns[p]; n > Z32_MAX {
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

// zSqrt reduce the number o√i 
// Example zSqrt(3,8) return 6,2,nil:
//    _     _
//	3√8 = 6√2
func (z *Z32s) zSqrt(o, i Z) (o32, i32 Z32, err error) {
	if o == 0 || i == 0 {
		return 0, 0, nil
	}
	on := N(o); if o < 0 { on = N(-o) }
	in := N(i); if i < 0 { in = N(-i) }
	for _, prime := range z.primes {
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
			break // check next bigger prime
		}
	}
	if on > Z32_MAX || in > Z32_MAX {
		return 0, 0, ErrOverflow
	}
	o32 = Z32(on); if o < 0 { o32 = Z32(-o32) }
	i32 = Z32(in); if i < 0 { i32 = Z32(-i32) }
	return o32, i32, nil
}

// zSqrtN reduces the external o and several internal i
// Example zSqrt(3,8,16) return 6,[2,4],nil:
//    __________     _________
//	3√(8x + 16y) = 6√(2x + 4y)
func (z *Z32s) zSqrtN(o Z, is ...Z) (o32 Z32, i32s []Z32, err error) {
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
	for _, prime := range z.primes {
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
	if on > Z32_MAX {
		return 0, nil, ErrOverflow
	} else if o > 0 { // origin sign
		o32 = Z32(+on)
	} else {
		o32 = Z32(-on)
	}
	i32s = make([]Z32, len(ins))
	for p := range is {
		if i := ins[p]; i > Z32_MAX {
			return 0, nil, ErrOverflow
		} else if is[p] > 0 { // original sign
			i32s[p] = Z32(+i)
		} else {
			i32s[p] = Z32(-i)
		}
	}
	return o32, i32s, nil
}

// zSqrtDenest3 returns a denominator and array of denested numerators
// if √(b + c√d) can be denested. There are four results types:
//	- len(n) == 0:
//		arguments cannot be denest.
//
//	- len(n) == 1:
//		  ,-------    n[0]
//		 / b + c√d = ------
//		√             den
//	                where den > 0.
//
//	- len(n) == 3:                _____
//		  ,-------    n[0] + n[1]√ n[2]
//		 / b + c√d = -------------------
//		√                    den
//	                 where den > 0, n[1] != 0, n[2] != 1.
//
//	- len(n) == 5:                _____        _____
//		  ,-------    n[0] + n[1]√ n[2] + n[3]√ n[4]
//		 / b + c√d = --------------------------------
//		√                         den
//	                 where den > 0, n[0] = 0, n[1] != 0, n[3] != 0, n[2] != n[4] != 1.
//
func (z *Z32s) zSqrtDenest3(b, c, d Z) (den N32, n [] Z32, e error) {
	den = 1
	if b == 0 {
		// simpler case √(c√d)
		if c == 0 || d == 0 {
			// √(0) = 0
			n = []Z32 { 0 }
		} else if o, i, err := z.zSqrt(c, d); err != nil {
			e = err
		} else if i == 1 {
			// √(c√d) = √o
			if o2, i2, err := z.zSqrt(1, Z(o)); err != nil {
				e = err
			} else if i2 == 1 {
				// √(c√d) = √o = o2
				n = []Z32 { o2 }
			} else {
				// √(c√d) = 0 + √o
				n = []Z32 { 0, 1, o }
			}
		} else {
			// cannot denest
		}
		return
	}
	if c == 0 || d == 0 {
		// simpler case √b
		if o, i, err := z.zSqrt(1, b); err != nil {
			e = err
		} else if i == 1 {
			// √b = o√1
			n = []Z32 { o }
		} else {
			// √b = 0 + o√i
			n = []Z32 { 0, o, i }
		}
		return
	}
	if d == 1 {
		// simpler case √(b+c)
		if o, i, err := z.zSqrt(1, b+c); err != nil {
			e = err
		} else if i == 1 {
			// √(b+c) = o√1 = o
			n = []Z32 { o }
		} else {
			// √(b+c) = 0 + o√i
			n = []Z32 { 0, o, i }
		}
		return
	}
	// For √(b + c√d) look if b² - c²d = x²
	// In other words, look a x such that 1√(b²-c²d) = x√1
	// Case example: √(6+2√5) = 1+√5
	if x, r, err := z.zSqrt(1, b*b - c*c*d); err != nil{
		e = err
	} else if r != 1 {
		// cannot denest
	} else {
		// √(b + c√d) = (√(2b+2x) + √(2b-2x))/2
		// √(b - c√d) = (√(2b+2x) - √(2b-2x))/2
		if o1, i1, err := z.zSqrt(1, 2*(b + Z(x))); err != nil {
			e = err
		} else
		if o2, i2, err := z.zSqrt(1, 2*(b - Z(x))); err != nil {
			e = err
		} else {
			// o1√i1 = √(b+x)
			// o2√i2 = √(b-x)
			if o1 % 2 == 0 && o2 % 2 == 0 {
				// numerators are all even. divide them by 2, left denominator as 1.
				o1 /= 2
				o2 /= 2
			} else {
				// numerators have some odd number, set denominator as 2.
				den = 2
			}
			if c < 0 {
				// correct second numerator
				o2 = -o2 // ??? or o1?
			}
			if i1 == +1 && i2 != +1 {
				// √(b+x) is integer, √(b-x) is not
				// return (o1 + o2√i2) / den
				n = []Z32{ o1, o2, i2 }
			} else
			if i1 != +1 && i2 == +1 {
				// √(b+x) is not integer, √(b-x) is.
				// return ( o2 + o1√i1) / den
				n = []Z32 { o2, o1, i1 }
			} else
			if i1 < i2 {
				// return ( 0 + o1√i1 + o2√i2 ) / den
				n = []Z32 {
					0, o1, i1, o2, i2,
				}
			} else {
				// return ( 0 + o2√i2 + o1√i1 ) / den
				n = []Z32 {
					0, o2, i2, o1, i1,
				}
			}
		}
	}
	return
}



