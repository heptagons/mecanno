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
	switch len(q.num) {
	default:
		return "?"
	
	case 1:
		b := q.num[0]
		s.z(b)                             // b
	
	case 3: // a>0, c!=0, d!=+1
		b, c, d := q.num[0], q.num[1], q.num[2]
		if b == 0 {
			s.root(c, d)
		} else {
			s.par(a > 1, func(s *Str) {    // (
				s.plus_root(b, c, d)       // b+c√d
			})                             // )
		}
	
	case 5: // a>0, c!=0, e!=0, d!=+1 d!=f
		s.par(a > 1, func(s *Str) {        // (
			c := q.num[1]
			if b := q.num[0]; b == 0 {
				s.z(c)                     // c
			} else {
				s.z(b); s.zS(c)            // b+c
			}
			d, e, f := q.num[2], q.num[3], q.num[4]
			s.sqrt(d); s.zS(e); s.sqrt(f)  // √d+e√f
		})                                 // )
	
	case 7:
		fgh := func(s *Str) {
			f, g, h := q.num[4], q.num[5], q.num[6]
			if f == 0 {
				s.root(g, h)         // g√h
			} else {
				s.plus_root(f, g, h) // f+g√h
			}
		}
		c, d, e := q.num[1], q.num[2], q.num[3]
		if b := q.num[0]; b == 0 {
			if c == 0 { // MAX e√(f+g√h)
				s.z(e);                     // e
				s.sqrtP(func(s *Str){       // √(
					fgh(s) // MAX f+g√h
				});                         // )
			} else { // MAX (c√d+e√(f+g√h))
				s.par(a > 1, func(s *Str) { // (
					s.z(c); s.sqrt(d)       // c√d
					s.zS(e);                // +e
					s.sqrtP(func(s *Str){   // √(
						fgh(s) // MAX f+g√h
					});                     // )
				})                          // )
			}
		} else { // MAX (b+e√(f+g√h))
			s.par(a > 1, func(s *Str) {     // (
				s.z(b)                      // b
				if c == 0 {
					s.zS(e);                // +e
				} else { // MAX (b+c√d+e√(f+g√h))
					s.zS(c); s.sqrt(d)      // +c√d
					s.zS(e);                // +e
				}
				s.sqrtP(func(s *Str){       // √(
					fgh(s) // MAX f+g√h   
				});                         // )
			})                              // )
		}
	}
	if a > 1 {
		s.over(a) // /a
	}
	return s.String()
}


