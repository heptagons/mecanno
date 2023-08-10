package alg

import (
	"fmt"
)

type Q32 struct {
	den N32   // a
	num []Z32 // b, c, d, e...
}

func newQ32(den N32, num ...Z32) *Q32 {
	return &Q32{
		den: den,
		num: num,
	}
}

// Equal returns true it the given r is identical to this one.
func (q *Q32) Equal(r *Q32) bool {
	if q == nil || r == nil {
		return false
	}
	if q.den != r.den {
		return false
	}
	if len(q.num) != len(r.num) {
		return false
	}
	for p, qn := range q.num {
		if qn != r.num[p] {
			return false
		}
	}
	return true
}

// Neg changes the signs of b, c and e.
func (q *Q32) Neg() *Q32 {
	switch len(q.num) {
	case 1:
		q.num[0] = -q.num[0] // b = -b
	case 3:
		q.num[0] = -q.num[0] // b = -b
		q.num[1] = -q.num[1] // c = -c
	case 5:
		q.num[0] = -q.num[0] // b = -b
		q.num[1] = -q.num[1] // c = -c
		q.num[3] = -q.num[3] // e = -e
	}
	return q
}

// GreatherThanN returns true iff this q is type 1 and greater than given n
func (q *Q32) GreaterThanZ(num Z) (bool, error) {
	if q == nil {
		return false, nil
	}
	switch len(q.num) {
	case 1:
		// q.num[0]    n
		// -------- > ---- ; q.num[0] > n * q.den
		//   q.den     1
		return Z(q.num[0]) > num*Z(q.den), nil
	}
	return false, fmt.Errorf("Can't compare %s and %d", q, num)

}

//     a    b     c     d     e    f    g    h    i
//    ---  ===   ---   ---   ===  ---  ---  ---  ===
// 1A  >0  !=0 |___________                                            b/a
// 3A  >0   0    !=0   !=1 |                                         c√d/a
// 3B  >0  !=0   !=0   !=1 |__________                           (b+c√d)/a
// 5A  >0   0    !=0   !=1   !=0  !=1 |                        (c√d+e√f)/a 
// 5B  >0  !=0   !=0   !=1   !=0  !=1 |_________             (b+c√d+e√f)/a
// 7A  >0   0     0     x    !=0   0   !=1  !=0 |              (e√(g√h))/a
// 7B  >0   0     0     x    !=0  !=0  !=1  !=0 |            (e√(f+g√h))/a
// 7C  >0   0    !=0   !=1   !=0   0   !=1  !=0 |          (c√d+e√(g√h))/a
// 7C  >0   0    !=0   !=1   !=0  !=0  !=1  !=0 |        (c√d+e√(f+g√h))/a
// 7A  >0  !=0    0     x    !=0   0   !=1  !=0 |            (b+e√(g√h))/a
// 7B  >0  !=0    0     x    !=0  !=0  !=1  !=0 |          (b+e√(f+g√h))/a
// 7C  >0  !=0   !=0   !=1   !=0   0   !=1  !=0 |        (b+c√d+e√(g√h))/a
// 7C  >0  !=0   !=0   !=1   !=0  !=0  !=1  !=0 |      (b+c√d+e√(f+g√h))/a
//
func (q *Q32) class() string {
	switch len(q.num) {
	case 1:
		return "b/a"
	case 3: // a>0, c!=0, d!=+1
		if b := q.num[0]; b == 0 {
			return "c√d/a"
		} else {
			return "(b+c√d)/a"
		}
	case 5: // a>0, c!=0, e!=0, d!=+1 d!=f
		if b := q.num[0]; b == 0 {
			return "(c√d+e√f)/a"
		} else {
			return "(b+c√d+e√f)/a"
		}
	case 7:
		if b := q.num[0]; b == 0 {
			if c := q.num[1]; c == 0 {
				if f := q.num[4]; f == 0 {
					return "(e√(g√h))/a"
				} else {
					return "(e√(f+g√h))/a"
				}
			} else {
				if f := q.num[4]; f == 0 {
					return "(c√d+e√(g√h))/a"
				} else {
					return "(c√d+e√(f+g√h))/a"
				}
			}
		} else {
			if c := q.num[1]; c == 0 {
				if f := q.num[4]; f == 0 {
					return "(b+e√(g√h))/a"
				} else {
					return "(b+e√(f+g√h))/a"
				}
			} else {
				if f := q.num[4]; f == 0 {
					return "(b+c√d+e√(g√h))/a"
				} else {
					return "(b+c√d+e√(f+g√h))/a"
				}
			}
		}
	}
}

func (q *Q32) String() string {
	if q == nil {
		return "+0"
	}
	n := len(q.num)
	if n == 0 {
		return ""
	} else if q.den == 0 {
		return "∞"
	}
	s := NewStr()
	a, b := q.den, q.num[0]
	if n == 1 {
		// b/a
		s.WriteString(fmt.Sprintf("%d", b))
	}

	if n == 3 {
		// (b + c√d)/a
		c, d := q.num[1], q.num[2]
		par := a > 1 && b*c != 0
		if par {
			s.WriteString("(")
		}
		q.bcdStr(s, b, c, d)
		if par {
			s.WriteString(")")
		}
	} else if n == 5 {
		// (b + c√d + e√f)/a
		c, d, e, f := q.num[1], q.num[2], q.num[3], q.num[4]
		if a > 1 {
			s.WriteString("(")
		}
		q.bcdefStr(s, b, c, d, e, f)
		if a > 1 {
			s.WriteString(")")
		}
	}
	// denominator
	if a > 1 {
		s.WriteString(fmt.Sprintf("/%d", a))
	}
	return s.String()
}

func (q *Q32) bcdStr(s *Str, b, c, d Z32) { // b + c√d
	if b != 0 {
		s.WriteString(fmt.Sprintf("%d", b))
	}
	if c == 1 {
		if b != 0 {
			s.WriteString("+")
		}
	} else if c == -1 {
		s.WriteString("-") // don't put -1
	} else {
		if b == 0 {
			s.WriteString(fmt.Sprintf("%d", c))
		} else {
			s.WriteString(fmt.Sprintf("%+d", c))
		}
	}
	s.WriteString(fmt.Sprintf("√%d", d))
}

func (q *Q32) bcdefStr(s *Str, b, c, d, e, f Z32) { // b + c√d + e√f 
	if b != 0 {
		s.WriteString(fmt.Sprintf("%d", b))
	}
	if c == 1 {
		if b != 0 {
			s.WriteString("+")
		}
	} else if c == -1 {
		s.WriteString("-") // don't put -1
	} else {
		if b == 0 {
			s.WriteString(fmt.Sprintf("%d", c))
		} else {
			s.WriteString(fmt.Sprintf("%+d", c))
		}
	}
	s.WriteString(fmt.Sprintf("√%d", d))
	if e == 1 {
		s.WriteString("+")
	} else if e == -1 {
		s.WriteString("-")
	} else {
		s.WriteString(fmt.Sprintf("%+d", e))
	}
	s.WriteString(fmt.Sprintf("√%d", f))
}


