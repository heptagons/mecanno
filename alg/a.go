package alg

import (
	"fmt"
	"math"
)

const AZ_MAX = 0x7fffffff

// A32s is factory to create AZ32 and AQ32
type A32s struct {
	primes []N32
}

// AZ32 represent an algebraic integer number
type AZ32 struct {
	o int32
	i []*AZ32
}

// AQ32 represent an algebraic rational number
type AQ32 struct {
	num []*AZ32
	den uint32
}


func NewA32s() *A32s {
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
	return &A32s{
		primes: primes,
	}
}


func (a *A32s) f(b int32, cd, efgh, ijklmnop []int32) *AZ32 {
	o := b
	i := make([]*AZ32, 0)
	if n := len(cd); n != 0 {
		if n != 2 {
			return nil
		}
		i = append(i, a.f1(cd[0],cd[1]))
	}
	return &AZ32{ o:o, i:i }
}

func (a *A32s) f0(b int32) *AZ32 {
	return &AZ32{
		o: b,
	}
}

func (a *A32s) f1(p ...int32) *AZ32 {
	if len(p) != 2 {
		return nil
	}
	return &AZ32 {
		o: p[0], // c
		i: []*AZ32{
			a.f0(p[1]), // d
		},
	}
}

func (a *A32s) f2(p ...int32) *AZ32 {
	if len(p) != 4 {
		return nil
	}
	return &AZ32 {
		o: p[0], // e
		i: []*AZ32{
			a.f0(p[1]), // f
			a.f1(p[2], p[3]), // g,h
		},
	}
}

func (a *A32s) f3(p ...int32) *AZ32 {
	if len(p) != 8 {
		return nil
	}
	return &AZ32 {
		o: p[0], // i
		i: []*AZ32{
			a.f0(p[1]), // j
			a.f1(p[2], p[3]), // k,l
			a.f2(p[4], p[5], p[6], p[7]), // m,n,o,p
		},
	}
}

func (a *A32s) f4(p ...int32) *AZ32 {
	if len(p) != 16 {
		return nil
	}
	return &AZ32 {
		o: p[0],
		i: []*AZ32{
			a.f0(p[1]),
			a.f1(p[2], p[3]),
			a.f2(p[4], p[5], p[6], p[7]),
			a.f3(p[9], p[9], p[10], p[11], p[12], p[13], p[14], p[15]),
		},
	}
}

func (a *A32s) Q(num []*AZ32, den uint32) *AQ32 {
	return &AQ32 {
		num: num,
		den: den,
	}
}

func (a *A32s) F1(c, d int64) (*AZ32, bool) {
	if c, d, overflow := a.roi(c, d); overflow {
		return nil, true
	} else if d == 1 { // convert x√+1 into x
		return a.f0(c), false
	} else {
		return a.f1(c, d), false
	}
}

func (a *A32s) Reduce(p ...int64) (r ...int32, overflow bool) {

	if n := len(p); n == 2 {
		// reduce c√d
		if c1, d1, overflow := a.roi(p[0], p[1]); overflow {
			return nil, true
		} else if d1 == 1 { // convert x√+1 into x
			return []int32{ c1 }, false
		} else {
			return []int32{ c1, d1 }, false
		}
	} else if n == 4 {

		// first reduce g1√h1
		if g1, h1, overflow := a.roi(p[2], p[3]); overflow {
			return nil, true // g√h overflow

		} else if g1 == 0 {
			// F2 degerates to F1(e, f)
			return a.Reduce(p[0], p[1]) // Go to reduce e√h
		
		} else if h1 == +1 {
			// F2 degerates into F1(e, f + g1)
			return a.Reduce(e, f + int64(g1))

		} else if e1, fg, overflow := a.roie(e, f, int64(g1)); overflow { // reduced 
			// reduction e1√(f1+g2) = e√(f+g1) overflow
			return nil, true
		} else {
			f1 := fg[0]
			g2 := fg[1]
			return []int32 { e1, f1, g2, h1 }, false
		}
	} else if n == 8 {
		// ijklmnop
		if mnop, overflow := a.Reduce(p[4], p[5], p[6], p[7]); overflow {
			// m√(n+o√p) oveflow
			return nil, true
		} else if m1 := mnop[0]; m1 == 0 {
			// Degenerates to i√(j+k√l)
			return a.Reduce(p[0], p[1], p[2], p[3])
		} else if n1 := mnop[1] == 0 { // means n1 = +1
			// Degenerates to i√(j+m1+k√l)
			return a.Reduce(p[0], p[1] + m1, p[2], p[3])

		} if kl, overflow := a.Reduce(p[2], p[3]); overflow {
			// k√l overflow
			return nil, true

		} else if i1, jkm, overflow := a.roie(p[0], p[1], kl[0], m1); overflow { // reduced 
			// reduction i√(j+k+m1) overflow
			return nil, true
		} else {
			return []int32 {
				i1,
				jkm[0],  // j1
				jkm[1],  // k1
				kl[1],   // l1
				jkm[2],  // m2
				mnop[1], // n1
				mnop[2], // o1
				mnop[3], // p1
			}, false
		}
	}
}


func (a *A32s) F2(e, f, g, h int64) (*AZ32, bool) {
	if g1, h1, overflow := a.roi(g, h); overflow {
		// reduction g1√h1 = g√h overflow
		return nil, true
	} else if g1 == 0 {
		// F2 degerates into F1(e, f)
		return a.F1(e, f)
	} else if h1 == +1 {
		// F2 degerates into F1(e, f + g1)
		return a.F1(e, f + int64(g1))

	} else if e1, fg, overflow := a.roie(e, f, int64(g1)); overflow { // reduced 
		// reduction e1√(f1+g2) = e√(f+g1) overflow
		return nil, true
	} else {
		f1 := fg[0]
		g2 := fg[1]
		return a.f2(e1, f1, g2, h1), false
	}
}



/*
func (a *A32s) F3(i, j, k, l, m, n, o, p int64) (*AZ32, bool) {
	// get mnop
	if mnop, overflow := a.F2(m, n, o, p); overflow {
		// reduction m√(n+o√p) overflow
		return nil, true
	} else if m1 := int64(mnop.out()); m1 == 0 { // reduced m
		// Degenerates to i√(j+k√l)
		return a.F2(i, j, k, l)
	} else if len(mnop.i) == 0 { // means n1 = +1
		// Degenerates j2=j+m1 to i√(j2+k√l)
		return a.F2(i, j + m1, k, l)

	// get kl
	} if kl, overflow := a.F1(k, l); overflow {
		// reduction g1√h1 = g√h overflow
		return nil, true
	}

	} else if i1, jkm, overflow := a.roie(i, j, k, m1); overflow { // reduced 
		// reduction i√(j+k+m1) overflow
		return nil, true
	} else {
		j1 := jkm[0]
		k1 := jkm[1]
		m2 := jkm[2]
		n2 := mnop.in()
		o2 := int32(0)
		p2 := int32(0)
		return a.f3(i1, j1, k1, m2, n2, o2, p2), false
	}
}*/


func (a *AZ32) out() int32 {
	if a == nil {
		return 0 // was 1
	}
	return a.o
}

func (a *AZ32) in() int32 {
	if len(a.i) == 0 {
		return +1
	}
	return a.i[0].out()
}

func (r *AZ32) String() string {
	s := NewStr()
	r.Str(s,true)
	return s.String()
}

func (a *AZ32) Str(s *Str, sign bool) {
	if sign {
		s.WriteString(fmt.Sprintf("%+d", a.o))
	} else{
		s.WriteString(fmt.Sprintf("%d", a.o))
	}
	if a.o == 0 {
		return
	}
	if n := len(a.i); n == 0 {
		return
	} else if n == 1 {
		if o := a.i[0].o; o < 0 {
			s.WriteString("i")
			if o != -1 {
				s.WriteString(fmt.Sprintf("√%d", -o))
			}
		} else if o > 1 {
			s.WriteString(fmt.Sprintf("√%d", o))
		}
	} else {
		s.WriteString("√(")
		for p, i := range a.i {
			i.Str(s, p != 0)
		}
		s.WriteString(")")
	}
}

func (a *AQ32) String() string {
	s := NewStr()
	s.WriteString("(")
	for pos, num := range a.num {
		num.Str(s, pos != 0)
	}
	s.WriteString(fmt.Sprintf(")/%d", a.den))
	return s.String()
}

func (a *A32s) roi(o, i int64) (o32, i32 int32, overflow bool) {
	if o == 0 || i == 0 {
		return 0, 0, false
	}
	on := N(o); if o < 0 { on = N(-on) }
	in := N(i); if i < 0 { in = N(-in) }
	for _, prime := range a.primes {
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
	o32 = int32(on); if o < 0 { o32 = int32(-o32) }
	i32 = int32(in); if i < 0 { i32 = int32(-i32) }
	return o32, i32, false
}

func (a *A32s) roie(o int64, is ...int64) (o32 int32, i32s []int32, overflow bool) {
	allIsZero := true
	for _, i := range is {
		if i != 0 {
			allIsZero = false
		}
	}
	if o == 0 || allIsZero {
		return // zero
	}
	on := N(+o)
	if o < 0 {
		on = N(-on)
	}
	ins := make([]N, len(is))
	// insMaxPos points to the greatest in to reduce primes use
	insMaxPos, insMax := 0, N(0)
	for p, i := range is {
		if i > 0 {
			ins[p] = N(+i)
		} else {
			ins[p] = N(-i)
		}
		if ins[p] > insMax { 
			insMaxPos, insMax = p, ins[p]
		}
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
				for p := range ins { ins[p] /= pp }
				continue // check same prime again
			}
			break // check next prime
		}
	}
	if on > AZ_MAX {
		return 0, nil, true // overflow
	} else if o > 0 { // origin sign
		o32 = int32(+on)
	} else {
		o32 = int32(-on)
	}
	i32s = make([]int32, len(ins))
	for p := range is {
		if i := ins[p]; i > AZ_MAX {
			return 0, nil, true // overflow
		} else if is[p] > 0 { // original sign
			i32s[p] = int32(+i)
		} else {
			i32s[p] = int32(-i)
		}
	}
	return o32, i32s, false
}
