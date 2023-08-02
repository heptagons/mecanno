package alg

import (
	"fmt"
)

// A32s is factory to create AZ32 and AQ32
type A32s struct {
	
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

func (a *A32s) F0(b int32) *AZ32 {
	return &AZ32{
		o: b,
	}
}

func (a *A32s) F1(c, d int32) *AZ32 {
	return &AZ32 {
		o: c,
		i: []*AZ32{
			a.F0(d),
		},
	}
}

func (a *A32s) F2(e, f, g, h int32) *AZ32 {
	return &AZ32 {
		o: e,
		i: []*AZ32{
			a.F0(f),
			a.F1(g, h),
		},
	}
}

func (a *A32s) F3(i, j, k, l, m, n, o, p int32) *AZ32 {
	return &AZ32 {
		o: i,
		i: []*AZ32{
			a.F0(j),
			a.F1(k, l),
			a.F2(m, n, o, p),
		},
	}
}

func (a *A32s) Q(num []*AZ32, den uint32) *AQ32 {
	return &AQ32 {
		num: num,
		den: den,
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
	s.WriteString("√")
	if n == 1 {
		a.i[0].Str(s, false)
	} else {
		s.WriteString("(")
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



