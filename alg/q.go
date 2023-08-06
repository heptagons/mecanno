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

		if b != 0 {
			if a > 1 {
				s.WriteString("(")
			}
			s.WriteString(fmt.Sprintf("%d", b))
		}
		if c == 1 {
			// dont put 1
		} else if c == -1 {
			s.WriteString("-") // don't put -1
		} else {
			s.WriteString(fmt.Sprintf("%d", c))
		}
		if d != +1 {
			s.WriteString(fmt.Sprintf("√%d", d))
		} else if c == 1 || c == -1 {
			s.WriteString("1") // yes, put 1
		}
		if b != 0 && a > 1 {
			s.WriteString(")")
		}
	}
	if a > 1 {
		s.WriteString(fmt.Sprintf("/%d", a))
	}
	return s.String()
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
	switch len(num) {
	case 1:
		return qs.ab(den, num[0])
	case 3:
		if num[1] == 0 || num[2] == 0 { // c == 0 || d == 0
			return qs.ab(den, num[0]) // return b/a
		} else if num[2] == 1 { // d == 1
			return qs.ab(den, num[0] + num[1]) // return (b+c)/a
		} else {
			return qs.abcd(den, num[0], num[1], num[2]) // return (b+c√d)/a
		}
	default:
		return nil, ErrInvalid
	}
}

func (qs *Q32s) ab(a N, b Z) (*Q32, error) {
	if a32, n32, err := qs.reduceQ1(a, b); err != nil {
		return nil, err
	} else {
		return &Q32{
			den: a32,
			num: []Z32{ n32 },
		}, nil
	}
}

func (qs *Q32s) abcd(a N, b, c, d Z) (*Q32, error) {
	if c32, d32, err := qs.reduceRoot(c, d); err != nil { // c√d
		return nil, err
	} else if d32 == 1 {
		return qs.ab(a, b * Z(c32)) // d wasn't square free
	} else if a32, bc32, err := qs.reduceQn(a, b, Z(c32)); err != nil { // (b,c)/a
		return nil, err
	} else {
		return &Q32{
			den: a32,
			num: []Z32{ bc32[0], bc32[1], d32 }, // b,c,d
		}, nil
	}
}





/*
func (qs *Q32s) reduceQMulN(p ...*Q32) (q *Q32, err error) {
	if n := len(p); n >= 0 {
		q = p[0]
		for i=1; i < n; i++ {
			q, err = qs.reduceQMul2(q, p[i])
			if err != nil {
				return
			}
		}
	}
	return
}*/

/*
func (qs *Q32s) reduceQMul2(q, r *Q32) (s *Q32, err error) {
	if q != nil || r != nil {
		n1, n2 := len(q.num), len(q.num)
		if n2 > n1 {
			q = s; q = r; r = s
		}
		if n1 == 0 {
			return
		}
		qa, qb := N(q.den), Z(q.num[0])
		ra, rb := N(r.den), Z(r.num[0])
		aa := qa * ra
		bb := qb * rb
		switch len(q) {
		case 1: // len(r) should be 1
			// qb   rb   qb * rb    bb    b32
			// -- * -- = ------- = ---- = ---
			// qa   ra   qa * qb    aa    a32
			return qs.newQ32(aa, bb)
		case 2:
			qc, qd := Z(q.num[1]), Z(q.num[2])
			switch len(r) {
			case 1:
				// qb + qc√qd   rb   qb*rb + qc*rb√qd    bb + qcrb√qd   b32 + c32√d32
				// ---------- * -- = ----------------- = ------------ = -------------
				//    qa        ra       qa * ra              aa             a32
				qcrb := qc * rb
				return qs.newQ32(aa, bb, qc*rb, qd) // a, b, c, d
			case 2:
				// qb + qc√qd   rb + rc√rd
				// ---------- * ----------
				//     qa           ra
				// q.b*r.b + r.c*q.b√r.d + q.c*r.b√q.d + q.c*r.c√(q.d*r.d) / (q.a*r.a)
				// (bb + o1i1 + o2i2 + o3i3) / aa
				if o1, i1, err := qs.ReduceRoot(Z(r.num[1])*Z(q.num[1]))


			}
		}
	}
	return return
}

	} else if n == 2 {
		return p[0].mul(p[1])
	}
	return nil, false
	/*
	o1, i1, d1 := p[0].oid()
	o2, i2, d2 := p[1].oid()
	//    __      __         ____
	// o1√i1   o2√i2    o1o2√i1i2
	// ----- x ------ = ---------
	//  d1      d2       d1d2
	den := N(d1) * N(d2)
	if o32, i32, overflow := qs.reduceRoot(Z(o1)*Z(o2), Z(i1)*Z(i2)); overflow {
		return nil, true
	} else
	if den32, n32s, overflow := qs.reduceQ(den, Z(o32)); overflow {
		return nil, true
	} else {
		return newQ32Root(n32s[0], i32, den32), false
	} *  /
*/




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





