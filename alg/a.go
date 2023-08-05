package alg

import (
	"fmt"
)

const AZ_MAX = 0x7fffffff

// AZ32 represent an algebraic integer number
type AZ32 struct {
	o Z32
	i []*AZ32
}


// AQ32 represent an algebraic rational number
type AQ32 struct {
	num []*AZ32
	den N32
}

func newAQ32(num Z32, den N32) *AQ32 {
	return &AQ32{
		num: []*AZ32 {
			newAZ32(num),
		},
		den: den,
	}
}

func newAQ32Root(out, in Z32, den N32) *AQ32 {
	return &AQ32{
		num: []*AZ32 {
			newAZ32(out, in),
		},
		den: den,
	}
}


func newAZ32(p ...Z32) *AZ32 {
	if n := len(p); n == 0 {
		return nil
	} else {
		a := &AZ32{}
		if n >= 1 {
			a.o = p[0] // b, c, e, i, ...
		}
		if n >= 2 {
			a.i = make([]*AZ32, 0)
			a.i = append(a.i, newAZ32(p[1])) // d, f, j, ...
		}
		if n >= 4 {
			a.i = append(a.i, newAZ32(p[2:4]...)) // gh, kl, ...
		}
		if n >= 8 {
			a.i = append(a.i, newAZ32(p[4:8]...)) // mnop, ...
		}
		if n >= 16 {
			a.i = append(a.i, newAZ32(p[8:16]...)) // 
		}
		return a
	}
}

func (a *AZ32) out() Z32 {
	if a == nil {
		return 0 // was 1
	}
	return a.o
}

func (a *AZ32) in() Z32 {
	if len(a.i) == 0 {
		return +1
	}
	return a.i[0].out()
}

func (r *AZ32) String() string {
	if r == nil {
		return ""
	}
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
	for len(a.num) > 1 {
		s.WriteString("(")
	}
	for pos, num := range a.num {
		num.Str(s, pos != 0)
	}
	for len(a.num) > 1 {
		s.WriteString(")")
	}
	if a.den > 1 {
		s.WriteString(fmt.Sprintf("/%d", a.den))
	}
	return s.String()
}



