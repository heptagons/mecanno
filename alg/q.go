package alg

import (
	"fmt"
)

// 1   2   3   4   5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  26  27  28  29  30  31  32
//                                                                                                                                    ,-
//                                                                   ,-------------------------------------------------------        /
//                                  ,------------------------       /                               ,------------------------       /
//                 ,---------      /               ,---------      /               ,---------      /               ,---------      /
//        ,--     /       ,--     /       ,--     /       ,--     /       ,--     /       ,--     /       ,--     /       ,--     /
// b + c √ d + e √ f + g √ h + i √ j + k √ l + m √ n + o √ p + q √ r + s √ t + u √ v + w √ x + y √ z + A √ B + C √ D + E √ F + G √ ...
// ------------=---------------=-------------------------------=---------------------------------------------------------------=------
//                                                             a
type Q32 struct {
	den N32   // a
	num []Z32 // b, c, d, e...
}

// newQ32 creates num/den not reduced
func newQ32(den N32, num ...Z32) *Q32 {
	return &Q32 {
		num: num,
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
	b := q.num[0]
	if n == 1 {
		s.WriteString(fmt.Sprintf("%d", b))
	} else if n == 3 {
		if q.den > 1 {
			s.WriteString("(")
		}
		if b != 0 {
			s.WriteString(fmt.Sprintf("%d", b))
		}
		c, d := q.num[1], q.num[2]
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
		if q.den > 1 {
			s.WriteString(")")
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
	}*/
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





