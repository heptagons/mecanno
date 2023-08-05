package alg

import (
	"fmt"
)

type Q32 struct {
	num []Z32
	den N32
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


