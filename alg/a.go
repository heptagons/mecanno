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

func (a *A32s) f0(b int32) *AZ32 {
	return &AZ32{
		o: b,
	}
}

func (a *A32s) f1(c, d int32) *AZ32 {
	return &AZ32 {
		o: c,
		i: []*AZ32{
			a.f0(d),
		},
	}
}

func (a *A32s) f2(e, f, g, h int32) *AZ32 {
	return &AZ32 {
		o: e,
		i: []*AZ32{
			a.f0(f),
			a.f1(g, h),
		},
	}
}

func (a *A32s) f3(i, j, k, l, m, n, o, p int32) *AZ32 {
	return &AZ32 {
		o: i,
		i: []*AZ32{
			a.f0(j),
			a.f1(k, l),
			a.f2(m, n, o, p),
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

func (r *AZ32) String() string {
	s := NewStr()
	r.Str(s,true)
	return s.String()
}

func (a *AZ32) Str(s *Str, positiveSign bool) {
	if positiveSign && a.o >= 0 {
		s.WriteString("+")
	}
	s.WriteString(fmt.Sprintf("%d", a.o))
	if a.o == 0 {
		return
	}
	n := len(a.i)
	if n == 0 {
		return
	}
	if n == 1 {
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

func (a *A32s) roie(o int64, is []int64) (o32 int32, i32s []int32, overflow bool) {
	if o == 0 || len(is) == 0 {
		return // nothing to change
	}
	for _, i := range is {
		if i == 0 {
			return // nothing change
		}
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
