package alg

import (
	"fmt"
)

type Q32 struct {
	num []Z32
	den N32
}

func (q *Q32) oid() (out, in Z32, den N32) {
	if n := len(q.num); n == 0 {
		return 0, 1, q.den
	} else if n == 1 {
		return q.num[0], 1, q.den
	} else {
		return q.num[0], q.num[1], q.den
	}
}

// newQ32 creates num/den not reduced
func newQ32(num Z32, den N32) *Q32 {
	return &Q32 {
		num: []Z32 {
			num,
		},
		den: den,
	}
}

// newQ32 creates out√in/den not reduced
func newQ32Root(out, in Z32, den N32) *Q32 {
	return &Q32 {
		num: []Z32 {
			out,
			in,
		},
		den: den,
	}
}

func (q *Q32) String() string {
	n := len(q.num)
	if n == 0 {
		return ""
	} else if q.den == 0 {
		return "∞"
	}
	s := NewStr()
	a := q.num[0]
	if n == 1 {
		s.WriteString(fmt.Sprintf("%d", a))
	} else if n == 2 {
		if a == 1 {
			// dont put 1
		} else if a == -1 {
			s.WriteString("-") // don't put -1
		} else {
			s.WriteString(fmt.Sprintf("%d", a))
		}
		if in := q.num[1]; in != +1 {
			s.WriteString(fmt.Sprintf("√%d", in))
		} else if a == 1 || a == -1 {
			s.WriteString("1") // yes, put 1
		}
	}
	if q.den > 1 {
		s.WriteString(fmt.Sprintf("/%d", q.den))
	}
	return s.String()
}



type Q32s struct {
	*N32s
}

func (qs *Q32s) reduceQMul(p ...*Q32) (q *Q32, overflow bool) {
	if n := len(p); n == 0 {
		return nil, false
	} else if n == 1 {
		return p[0], false
	}
	// o1√i1    o2√i2    o1o2√i1i2
	// ----- x ------ = ----------
	//  d1       d2       d1d2
	o1, i1, d1 := p[0].oid()
	o2, i2, d2 := p[1].oid()
	den := N(d1) * N(d2)
	if o32, i32, overflow := qs.reduceRoot(Z(o1)*Z(o2), Z(i1)*Z(i2)); overflow {
		return nil, true
	} else if den32, n32s, overflow := qs.reduceQ(den, Z(o32)); overflow {
		return nil, true
	} else {
		return newQ32Root(n32s[0], i32, den32), false
	}
}


