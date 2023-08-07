package alg

import (
	"fmt"
)

type Q32 struct {
	den N32   // a
	num []Z32 // b, c, d, e...
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
	} else if n == 3 {
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
	}
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
func (qs *Q32s) newQ32(den N, num ...Z) (q *Q32, err error) {
	if den == 0 {
		return nil, ErrInfinite
	} else if den > N32_MAX {
		return nil, ErrOverflow
	}
	a := den
	switch len(num) {
	case 1:
		return qs.ab(a, num[0])
	case 3:
		return qs.abcd(a, num[0], num[1], num[2]) // (b+c√d)/a
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
	aa := qa * ra
	qbra := qb*Z(ra)
	qarb := Z(qa)*rb
	switch len(q.num) {
	case 1:
		// qb/aq + rb/ra =  (qb*ra + qa*rb)/(qa*ra) = (qb*ra + qa*rb)/aa
		return qs.newQ32(aa, qbra + qarb)

	case 3:
		qc, qd := Z(q.num[1]), Z(q.num[2])
		qcra := qc*Z(ra)
		switch len(r.num) {
		case 1:
			// qb + qc√qd   rb   qb*ra + qa*rb + (qc*ra)√qd   b + c√qd
			// ---------- + -- = -------------------------- = --------
			//    qa        ra           aa                     aa
			return qs.newQ32(aa, qbra + qarb, qcra, qd)
		case 3:
			rc, rd := Z(r.num[1]), Z(r.num[2])
			qarc := Z(qa)*rc
			if qd == rd { // simpler case
				// qb + qc√qd   rb + rc√qd  qb*ra + qa*rb + (qc*ra + qa*rc)√qd
				// ---------- + --------- = ----------------------------------
				//     qa          ra                     aa
				return qs.newQ32(aa, qbra + qarb, qcra + qarc, qd)
			}
			if qb == rb && qb == 0 {
				// qc√qd   rc√rd   qc*ra√qd/a + rc*qa√rd/a   x√qd + y√rd   √(x*x*qd + y*y*rd +2*x*y√qd*rd)   e√f + g√h)
				// ----- + ----- = ----------------------- = ----------- = ------------------------------- = ----------
				//   qa     ra                a                   a                      a                       a
				// GCD =(qa,ra)
				a := NatGCD(qa, ra)
				x := qcra / Z(a)
				y := qarc / Z(a)
				f := x*x*qd + y*y*rd
				g := 2*x*y
				h := qd*rd
				if g32, h32, err := qs.reduceRoot(g, h); err != nil { // g√h
					return nil, err
				}


			}


		}
	}
	return nil, ErrInvalid
}












func (qs *Q32s) reduceQAdd2(p ...*Q32) (q *Q32, overflow bool) {
	if n := len(p); n == 0 {
		return nil, false
	} else if n == 1 {
		return p[0], false
	}
	return nil, false
	/*
	o1, i1, d1 := p[0].oid()
	o2, i2, d2 := p[1].oid()
	//       __      __
	//  / o1√i1   o2√i2  \2
	// (  ----- + ------  )
	//  \   d1      d2   /
	//                  ____                     _                 _         _
	//   o1o1i1   2oio2√i1i2    o2o2i2    a    c√i    f     m    c√i   x + y√i
	// = ------ + ----------- + ------ = --- + --- + --- = --- + --- = -------
	//    d1d1      d1d2         d2d2     b     d     g     n     d      z

	if b, a, overflow := qs.reduceQ1(N(d1)*N(d2), Z(o1)*Z(o1)*Z(i1)); overflow {
		return nil, true
	} else if c, i, overflow := qs.reduceRoot(2*Z(o1)*Z(o2), Z(i1)*Z(i2)); overflow {
		return nil, true
	} else if d, c, overflow := qs.reduceQ1(N(d1)*N(d2), Z(c)); overflow {
		return nil, true
	} else if g, f, overflow := qs.reduceQ1(N(d2)*N(d2), Z(o2)*Z(o2)*Z(i2)); overflow {
		return nil, true
	
	} else if m, n, overflow := qs.reduceQ1(N(b)*N(g), Z(a)*Z(g)+Z(f)*Z(b)); overflow {
		return nil, true
	
	} else if x, xy := qs.reduceQn(N(n)*N(d), Z(m)*Z(d), Z(c)*Z(n)); overflow {

		return nil
		//return newQ32Root(out, in Z32, den N32) 
	}*/
}





