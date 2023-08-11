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
func (q *Q32) String() string {
	s := NewStr()
	a := q.den
	if a == 0 {
		return "NaN"
	}
	switch len(q.num) {
	default:
		return ErrInvalid.Error()	
	case 1:
		b := q.num[0]
		if b == 0 {
			return "0"
		}
		s.z(b)                             // b
	
	case 3: // b + c√d
		b, c, d := q.num[0], q.num[1], q.num[2]
		if b == 0 {
			q.bcd(s, 0, c, d)
		} else {
			s.par(a > 1, func(s *Str) { // (
				q.bcd(s, b, c, d)       // b+c√d
			})                          // )
		}
	
	case 5: // b + c√d + e√f
		b, c, d, e, f := q.num[0], q.num[1], q.num[2], q.num[3], q.num[4]
		x := b!=0
		y := c!=0 && d!=0
		z := e!=0 && f!=0
		var par bool
		if x && y || y && z || z && x { par = true }
		if !x && !y && !z {
			return "0"
		}
		s.par(a > 1 && par, func(s *Str) {     // (
			if x {
				s.z(b)                         // b
			}
			if y {
				if x && c > 0 { s.pos() }      // +
				q.cd(s, c, d)                  // c√d | -c√d
			}
			if z {
				if (x||y) && e > 0 { s.pos() } // +
				q.cd(s, e, f)                  // e√f | -e√f
			}
		})                                     // )
	
	case 7: // b + c√d + e√(f+g√h)
		b, c, d, e, f, g, h := q.num[0], q.num[1], q.num[2], q.num[3], q.num[4], q.num[5], q.num[6]
		x := b!=0
		y := c!=0 && d!=0
		z := e!=0 && (f!=0 || (g!=0 && h!=0))
		var par bool
		if x && y || y && z || z && x { par = true }
		if !x && !y && !z {
			return "0"
		}
		s.par(a > 1 && par, func(s *Str) {     // (
			if x {
				s.z(b)                         // b
			}
			if y {
				if x && c > 0 { s.pos() }      // +
				q.cd(s, c, d)                  // c√d | -c√d
			}
			if z {
				if (x||y) && e > 0 { s.pos() } // +
				if e != 0 {
					if e == -1 {
						s.neg()
					} else if -1 > e || e > 1 {
						s.z(e)                 // e
					}
					s.sqrtP(func(s *Str){      // √(
						q.bcd(s, f, g, h)      // MAX f+g√h
					});                        // )
				}
			}
		})                                     // )
	}
	if a > 1 {
		s.over(a) // /a
	}
	return s.String()
}

// bcd simplifies printing of x + y√z
// preventing printing unnecessary plus signs, zeros and ones.
func (q *Q32) bcd(s *Str, x, y, z Z32) {
	if x == 0 {
		if y == 0 || z == 0 {
			s.z(0) // return "0"
			return
		}
		if y > 0 {
			q.cd(s, y, z) // add "y√z
		} else {
			q.cd(s, y, z) // add "y√z" or "-y√z"
		}
	} else {
		s.z(x)           // add x or -x
		if y == 0 || z == 0 {
			return
		}
		if y > 0 {
			s.pos()      // add "+"
		}
		q.cd(s, y, z)    // add "y√z" or "-y√z"
	}
}

// cd simplifies printing y√z
func (q *Q32) cd(s *Str, y, z Z32) {
	if y == 0 || z == 0 {
		s.z(0)             // add 0 and done.
	} else if z == 1 {
		s.z(y)
	} else if z == -1 {
		if y == -1 {
			s.neg()
		} else {
			s.z(y)
		}
		s.WriteString("i")
	} else if z >= 1 {
		if -1 > y || y > 1 {
			s.z(y)         // add y | -y and done.
		}
		if z != 1 {
			s.sqrt(z)      // add √z
		}
	} else if z <= -1 {
		if -1 > y || y > 1 {
			s.z(y)         // add y | -y and done.
		}
		s.WriteString("i") // add yi | -yi and donde.
		if z != -1 {
			s.sqrt(-z)     // add √z
		}
	}
}



