package alg

import (
	"fmt"
)

type Q32 struct {
	den N32   // a
	num []Z32 // b, c, d, e...
}

func newQ32(den N32, num Z32) *Q32 {
	return &Q32{
		den: den,
		num: []Z32{ num },
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


type Q32s struct {
	*N32s
}

func NewQ32s(factory *N32s) *Q32s {
	return &Q32s{
		N32s: factory,
	}
}

// 1   2   3   4   5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  26  27  28  29  30  31  32
//                                                                                                                                    ,-
//                                                                   ,-------------------------------------------------------        /
//                                  ,------------------------       /                               ,------------------------       /
//                 ,---------      /               ,---------      /               ,---------      /               ,---------      /
//        ,--     /       ,--     /       ,--     /       ,--     /       ,--     /       ,--     /       ,--     /       ,--     /
// b + c √ d + e √ f + g √ h + i √ j + k √ l + m √ n + o √ p + q √ r + s √ t + u √ v + w √ x + y √ z + A √ B + C √ D + E √ F + G √ ...
// ------------=---------------=-------------------------------=---------------------------------------------------------------=------
//                                                             a
// newQ32 returns reduced rational number
func (qs *Q32s) newQ32(den N, n ...Z) (q *Q32, err error) {
	if den == 0 {
		return nil, ErrInfinite
	} else if den > N32_MAX {
		return nil, ErrOverflow
	}
	a := den
	switch len(n) {
	case 1:
		return qs.ab(a, n[0])
	case 3:
		return qs.abcd(a, n[0], n[1], n[2]) // (b+c√d)/a
	case 5:
		return qs.abcdef(a, n[0], n[1], n[2], n[3], n[4]) // (b+c√d+e√f)/a


	default:
		return nil, ErrInvalid
	}
}

func (qs *Q32s) ab(a N, b Z) (*Q32, error) {
	if a32, n32, err := qs.reduceQ1(a, b); err != nil {
		return nil, err
	} else {
		return &Q32{
			den: a32, // a
			num: []Z32{ n32 }, // b
		}, nil
	}
}

func (qs *Q32s) abcd(a N, b, c, d Z) (*Q32, error) {
	if c == 0 || d == 0 {
		return qs.ab(a, b) // b/a
	} else if d == 1 {
		return qs.ab(a, b + c) // (b+c)/a
	} else if c32, d32, err := qs.reduceRoot(c, d); err != nil { // c√d
		return nil, err
	} else if d32 == 1 {
		return qs.ab(a, b + Z(c32)) // (b+c)/a
	} else if a32, bc32, err := qs.reduceQn(a, b, Z(c32)); err != nil { // (b,c)/a
		return nil, err
	} else {
		return &Q32{
			den: a32, // a
			num: []Z32{ bc32[0], bc32[1], d32 }, // b,c,d
		}, nil
	}
}

// abcdef simplifies and return (b + c√d + e√f)/a
func (qs *Q32s) abcdef(a N, b, c, d, e, f Z) (*Q32, error) {
	if e == 0 || f == 0 {
		return qs.abcd(a, b, c, d) // (b + c√d)/a
	} else if f == 1 {
		return qs.abcd(a, b+e, c, d) // (b+e + c√d)/a
	} else if d == f {
		//
		return qs.abcd(a, b, c+e, d) // (b + (c+e)√d)/a

	} else if c32, d32, err := qs.reduceRoot(c, d); err != nil { // c√d
		return nil, err
	} else if d32 == 1 {
		return qs.abcd(a, b+Z(c32), e, f) // (b+c + e√f)/a

	} else if e32, f32, err := qs.reduceRoot(e, f); err != nil { // e√f
		return nil, err
	} else if f32 == 1 {
		return qs.abcd(a, b+Z(e32), Z(c32), Z(d32)) // (b+e, c√d)/a
	
	} else if d32 == f32 {
		return qs.abcd(a, b, Z(c32) + Z(e32), Z(d32)) // (b, (c+e)√d)/a)

	} else if a32, bce32, err := qs.reduceQn(a, b, Z(c32), Z(e32)); err != nil { // (b,c,e)/a
		return nil, err
	} else {
		var num []Z32
		c, d, e, f := bce32[1], d32, bce32[2], f32
		if d < f {
			num = []Z32{ bce32[0],c,d,e,f } 
		} else {
			num = []Z32{ bce32[0],e,f,c,d }
		}
		return &Q32{ den:a32, num:num }, nil
	}
}

func (qs *Q32s) AddQ(q ...*Q32) (s *Q32, err error) {
	n := len(q)
	if n == 1 {
		return q[0], nil
	}
	max := q[0]
	for i := 1; i < n; i++ {
		min := q[i]
		if max == nil || min == nil {
			return nil, nil
		}
		if len(max.num) < len(min.num) {
			s = max; max = min; min = s
		}
		if s, err = qs.addQ2(max, min); err != nil {
			return
		}
		max = s
	}
	return
}

func (qs *Q32s) MulQ(q ...*Q32) (s *Q32, err error) {
	n := len(q)
	if n == 1 {
		return q[0], nil
	}
	max := q[0]
	for i := 1; i < n; i++ {
		min := q[i]
		if max == nil || min == nil {
			return nil, nil
		}
		if len(max.num) < len(min.num) {
			s = max; max = min; min = s
		}
		if s, err = qs.mulQ2(max, min); err != nil {
			return
		}
		max = s
	}
	return
}

func (qs *Q32s) mulQ2(q, r *Q32) (s *Q32, err error) {
	qa, qb := N(q.den), Z(q.num[0])
	ra, rb := N(r.den), Z(r.num[0])
	aa := qa * ra
	bb := qb * rb
	switch len(q.num) {
	case 1:
		// qb/aq * rb/ra =  (qb*rb)/(qa*ra) = aa/bb
		return qs.newQ32(aa, bb)

	case 3:
		qc, qd := Z(q.num[1]), Z(q.num[2])
		switch len(r.num) {
		case 1:
			// qb + qc√qd   rb   qb*rb + qc*rb√qd    bb + qcrb√qd   b32 + c32√d32
			// ---------- * -- = ----------------- = ------------ = -------------
			//    qa        ra       qa * ra              aa             a32
			return qs.newQ32(aa, bb, qc*rb, qd) // a, b, c, d
		case 3:
			rc, rd := Z(r.num[1]), Z(r.num[2])
			if qd == rd { // simpler case
				// qb + qc√qd   rb + rc√qd   (bb + qc*rc*qd) + (qb*rc + rb*qc)√qd
				// ---------- * --------- = -----------------------------------
				//     qa          ra                     aa
				return qs.newQ32(aa, bb + qc*rc*qd, qb*rc + rb*qc, qd)
			}
		}
	}
	return nil, ErrInvalid
}

func (qs *Q32s) addQ2(q, r *Q32) (s *Q32, err error) {
	qa, qb := N(q.den), Z(q.num[0])
	ra, rb := N(r.den), Z(r.num[0])

	// Use lcm always to prevent overflows
	a := (qa / Ngcd(qa, ra)) * ra // lcm
	aq := Z(a / qa)
	ar := Z(a / ra) 

	switch len(q.num) {
	case 1:
		//  qb     rb     qb*aq + rb*ar    b
		// ---- + ---- = -------------- = ---
		//  qa     qb           a          a
		return qs.newQ32(a, qb*aq + rb*ar)

	case 3:
		qc, qd := Z(q.num[1]), Z(q.num[2])
		switch len(r.num) {
		case 1:
			// qb + qc√qd   rb   qb*aq + rb*ar + (qc*aq)√qd    b + c√qd
			// ---------- + -- = -------------------------- = --------
			//    qa        ra              a                     aa
			return qs.newQ32(a, qb*aq + rb*ar, qc*aq, qd)			

		case 3:
			rc, rd := Z(r.num[1]), Z(r.num[2])
			if qd == rd { // simpler case
				// qb + qc√qd   rb + rc√qd  qb*aq + rb*ar + (qc*aq + rc*ar)√qd
				// ---------- + --------- = ----------------------------------
				//     qa          ra                     a 
				return qs.newQ32(a, qb*aq + rb*ar, qc*aq + rc*ar, qd)
			}
			if qb == rb && qb == 0 { // simpler case both b's=0
				// qc√qd   rc√rd   qc*aq√qd + rc*ar√rd
				// ----- + ----- = -------------------
				//   qa     ra              a
				return qs.newQ32(a, 0, qc*aq, qd, rc*ar, rd)
			}
		}
	}
	return nil, fmt.Errorf("Can't add %s and %s", q.String(), r.String())
}

func (qs *Q32s) sqrtQ(q *Q32) (s *Q32, err error) {
	if q == nil {
		return nil, nil
	}
	switch len(q.num) {
	case 1:
		a, b := q.den, q.num[0]
		//  b      1√ab   0 + C√D
		// --- -> ----- = -------
		//  a       a       A 
		return qs.newQ32(N(a), 0, 1, Z(a)*Z(b))

	case 3:
		// b + c√d     √(ab + ac√d)    0 + C√(D + E√F)
		// ------- -> ------------- = ----------------
		//    a             a               A
		// TODO finish and return a Q of size 5 a,b,c,d,e

	}
	return nil, fmt.Errorf("Can't square root of %s", q.String())
}
















