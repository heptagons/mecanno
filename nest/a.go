package nest

import (
	"fmt"
	"strings"
)


// 0   1   2   3   4   5   6   7   8   9  10  11  12  13  14 ...
//                                  ________________________
//                 __________      /               _________
//        ___     /       ___     /       ___     /       __
// b + c √ d + e √ f + g √ h + i √ j + k √ l + m √ n + o √ p ...
// =-----------=---------------=-----------------------------
//                           a
// A32 is a sum of nested algebraic numbers in the numerator
// of increasing complexity and a natural number in the denominator.
// There are A32 sizes according the number of integers in the numerator:
//	 1: b
//	 3: b + c √ d 
//	 5: b + c √ d + e √ f
//	 7: b + c √ d + e √(f + g √ h)
//	 9: b + c √ d + e √(f + g √ h)+ i √ j
//	11: b + c √ d + e √(f + g √ h)+ i √(j + k √ l)
//	13: b + c √ d + e √(f + g √ h)+ i √(j + k √ l + m √ n)
//	15: b + c √ d + e √ f + g √ h + i √(j + k √ l + m √(n + o √ p))
type A32 struct {
	den N32   // denominator a
	num []Z32 // numerator parts b,c,d,e,f,g,h
}

func (a *A32) Num(pos int) (Z32, bool) {
	if len(a.num) >= pos {
		return a.num[pos], true
	}
	return 0, false
}

func newA32(den N32, num ...Z32) *A32 {
	return &A32{
		den: den,
		num: num,
	}
}

func (a *A32) Equals(den N32, num ...Z32) bool {
	if n := len(num); n != len(a.num) {
		return false
	} else {
		if den != a.den {
			return false
		}
		for i := 0; i < n; i++ {
			if num[i] != a.num[i] {
				return false
			}
		}
		return true
	}
}

func newA32Int(z Z32) *A32 {
	return &A32{
		den: 1, // a
		num: []Z32{
			z, // b
		},
	}
}

func newA32Rat(num Z32, den N32) *A32 {
	return &A32{
		den: den, // a
		num: []Z32{
			num, // b
		},
	}
}

func newA32Surd(z Z32) *A32 {
	return &A32{
		den: 1, // a
		num: []Z32{
			0, // b
			1, // c
			z, // d
		},
	}
}

func (a *A32) addInt(z Z32) *A32 {
	if a == nil {
		return nil
	}
	if len(a.num) == 0 {
		return nil
	}
	a.num[0] += z*Z32(a.den)
	return a
}

func (q *A32) IsInteger() (Z32, bool) {
	if q.den == 1 {
		return q.num[0], true
	}
	return 0, false
}

// Equal returns true it the given number is identical to this one.
func (q *A32) Equal(r *A32) bool {
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

// Neg changes the signs of numerator parts b, c, e and i.
func (q *A32) Neg() *A32 {
	if len(q.num) > 7 { q.num[7] = -q.num[7] } // i = -i
	if len(q.num) > 3 { q.num[3] = -q.num[3] } // e = -e
	if len(q.num) > 1 { q.num[1] = -q.num[1] } // c = -c
	if len(q.num) > 0 { q.num[0] = -q.num[0] } // b = -b
	return q
}

// Deeper return the numbers for deeper order. q and r should be valid.
// Deep depends in which number has a more complicated numerator, more nested terms.
//	- maxThis true if q is deeper than r.
//	- max is q if q is deeper than r.
//	- min is r if q is deeper than r.
func (q *A32) Deeper(r *A32) (maxThis bool, max, min *A32) {
	maxThis = true
	max, min = q, r
	if len(q.num) < len(r.num) {
		maxThis = false
		max, min = r, q
	}
	return
}

//	- w is the LCD denominator
//	- U is the numerator of deeper number.
//	- u is the numerator of not deeper number.
func (max *A32) LCM(min *A32) (w N, U, u, B, b Z) {
	A, a := N(max.den), N(min.den)
	// Use lcm always to prevent overflows
	w = (A / Ngcd(A, a)) * a // lcm denominator
	U = Z(w / A) // upper factor for max numerator terms
	u = Z(w / a) // upper factor for min numerator terms
	B = Z(max.num[0])
	b = Z(min.num[0])
	return
}

// ab returns the natural denominator and the numerator part b. Panic for smaller sizes.
func (q *A32) ab() (N, Z) {
	return N(q.den), Z(q.num[0])
}

// cd returns the numerator part c√d of this number. Panic for smaller sizes.
func (q *A32) cd() (Z, Z) {
	return Z(q.num[1]), Z(q.num[2])
}

// cdef returns the numerator part c√d + e√f. Panic for smaller sizes.
func (q *A32) cdef() (Z, Z, Z, Z) {
	return Z(q.num[1]), Z(q.num[2]), 
		Z(q.num[3]), Z(q.num[4])
}

// cdefgh returns the numerator part c√d + e√(f+g√h). Panic for smaller sizes.
func (q *A32) cdefgh() (Z, Z, Z, Z, Z, Z) {
	return Z(q.num[1]), Z(q.num[2]),
		Z(q.num[3]), Z(q.num[4]),
		Z(q.num[5]), Z(q.num[6])
}

// GreatherThanN returns true iff this number is size 1 and greater than given n
func (q *A32) GreaterThanZ(num Z) (bool, error) {
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

// Key is a simpler identifier than String
func (a *A32) Key() string {
	var sb strings.Builder 
	sb.WriteString(fmt.Sprintf("%d", a.den))
	for _, n := range a.num {
		sb.WriteString(fmt.Sprintf(",%d", n))
	}
	return sb.String()
}

// size   a    b     c     d     e    f    g    h    i
// ----  ---  ===   ---   ---   ===  ---  ---  ---  ===
//   1A   >0  !=0 |___________                                            b/a
//   3A   >0   0    !=0   !=1 |                                         c√d/a
//   3B   >0  !=0   !=0   !=1 |__________                           (b+c√d)/a
//   5A   >0   0    !=0   !=1   !=0  !=1 |                        (c√d+e√f)/a 
//   5B   >0  !=0   !=0   !=1   !=0  !=1 |_________             (b+c√d+e√f)/a
//   7A   >0   0     0     x    !=0   0   !=1  !=0 |              (e√(g√h))/a
//   7B   >0   0     0     x    !=0  !=0  !=1  !=0 |            (e√(f+g√h))/a
//   7C   >0   0    !=0   !=1   !=0   0   !=1  !=0 |          (c√d+e√(g√h))/a
//   7C   >0   0    !=0   !=1   !=0  !=0  !=1  !=0 |        (c√d+e√(f+g√h))/a
//   7A   >0  !=0    0     x    !=0   0   !=1  !=0 |            (b+e√(g√h))/a
//   7B   >0  !=0    0     x    !=0  !=0  !=1  !=0 |          (b+e√(f+g√h))/a
//   7C   >0  !=0   !=0   !=1   !=0   0   !=1  !=0 |        (b+c√d+e√(g√h))/a
//   7C   >0  !=0   !=0   !=1   !=0  !=0  !=1  !=0 |      (b+c√d+e√(f+g√h))/a
//
func (q *A32) String() string {
	if q == nil {
		return ""
	}
	s := NewStr()
	if q.den == 0 {
		return "NaN"
	}
	switch len(q.num) {
	default:
		return ErrInvalid.Error()	
	case 1:
		b := Z(q.num[0])
		if b == 0 {
			return "0"
		}
		s.z(b) // b
	
	case 3: // b + c√d
		a, b := q.ab()
		c, d := q.cd()
		if b == 0 {
			q.sbcd(s, 0, c, d)
		} else {
			s.par(a > 1, func(s *Str) { // (
				q.sbcd(s, b, c, d)      // b+c√d
			})                          // )
		}
	
	case 5: // b + c√d + e√f
		a, b := q.ab()
		c, d, e, f := q.cdef()
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
				q.scd(s, c, d)                 // c√d | -c√d
			}
			if z {
				if (x||y) && e > 0 { s.pos() } // +
				q.scd(s, e, f)                 // e√f | -e√f
			}
		})                                     // )
	
	case 7: // b + c√d + e√(f+g√h)
		a, b := q.ab()
		c, d, e, f, g, h := q.cdefgh()
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
				q.scd(s, c, d)                 // c√d | -c√d
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
						q.sbcd(s, f, g, h)     // MAX f+g√h
					});                        // )
				}
			}
		})                                     // )
	}
	if a := N(q.den); a > 1 {
		s.over(a) // /a
	}
	return s.String()
}

// sbcd simplifies printing of x + y√z
// preventing printing unnecessary plus signs, zeros and ones.
func (q *A32) sbcd(s *Str, x, y, z Z) {
	if x == 0 {
		if y == 0 || z == 0 {
			s.z(0) // return "0"
			return
		}
		if y > 0 {
			q.scd(s, y, z) // add "y√z
		} else {
			q.scd(s, y, z) // add "y√z" or "-y√z"
		}
	} else {
		s.z(x)             // add x or -x
		if y == 0 || z == 0 {
			return
		}
		if y > 0 {
			s.pos()        // add "+"
		}
		q.scd(s, y, z)     // add "y√z" or "-y√z"
	}
}

// scd simplifies printing y√z
func (q *A32) scd(s *Str, y, z Z) {
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

func (a *A32) IsZero() bool {
	if a == nil {
		return true
	}
	if len := len(a.num); len == 0 {
		return true
	} else if len == 2 {
		if a.num[0] == 0 {
			return true
		}
	}
	return false // TODO
}

func (a *A32) Tex() string {
	var sb strings.Builder
	if len(a.num) == 0 {
		return "0"
	}
	b := a.num[0]
	if den := a.den; den == 1 { // b
		if len(a.num) == 1 {
			return fmt.Sprintf("%d", b)
		} else if len(a.num) == 3 {
			if b != 0 {
				sb.WriteString(fmt.Sprintf("%d", b))
			}
			c, d := a.num[1], a.num[2]
			if c == 1 {
				// nothing
			} else if c == -1 {
				sb.WriteString("-") // -
			} else {
				if b != 0 {
					sb.WriteString(fmt.Sprintf("%+d", c)) // c
				} else {
					sb.WriteString(fmt.Sprintf("%d", c)) // c
				}
			}
			if c != 0 && d != 1 {
				sb.WriteString(fmt.Sprintf("\\sqrt{%d}", d))
			}
		} else {
			return "pending..."
		}
	} else {
		if len(a.num) == 1 {  // b/a
			if b == 0 {
				return "0"
			} else if b > 0 {
				sb.WriteString(fmt.Sprintf("\\frac{%d", b))
			} else {
				sb.WriteString(fmt.Sprintf("-\\frac{%d", -b))
			}
		} else if len(a.num) == 3 {
			c, d := a.num[1], a.num[2]
			if b < 0 { // -(b+c√d)/a
				sb.WriteString(fmt.Sprintf("-\\frac{%d", -b)) // b
				c = -c
				if c == 1 {
					// nothing
				} else if c == -1 {
					sb.WriteString("-") // -
				} else {
					sb.WriteString(fmt.Sprintf("%+d", c)) // c
				}
			} else if b == 0 { // c√d/a
				sb.WriteString("\\frac{")
				if c == 1 {
					// nothing
				} else if c == -1 {
					sb.WriteString("-") // -
				} else {
					sb.WriteString(fmt.Sprintf("%d", c)) // c
				}
			} else { // (b+c√d)/a
				sb.WriteString(fmt.Sprintf("\\frac{%d", b)) // b
				if c == 1 {
					// nothing
				} else if c == -1 {
					sb.WriteString("-") // -
				} else {
					sb.WriteString(fmt.Sprintf("%+d", c)) // +c
				}
			}
			if c != 0 && d != 1 {
				sb.WriteString(fmt.Sprintf("\\sqrt{%d}", d))
			}
		} else {
			return "pending..."
		}
		sb.WriteString(fmt.Sprintf("}{%d}", den)) // a
	}
	return sb.String()
}


