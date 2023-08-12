package alg

import (
	"fmt"
)

type Q32s struct {
	*Z32s
}

func NewQ32s() *Q32s {
	return &Q32s{
		Z32s: NewZ32s(),
	}
}

// 1   2   3   4   5   6   7   8   9  10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  26  27  28  29  30  31  32
//                                                                                                                                    ,-
//                                                                   ,-------------------------------------------------------        /
//                                  ,------------------------       /                               ,------------------------       /
//                 ,---------      /               ,---------      /               ,---------      /               ,---------      /
//        ,--     /       ,--     /       ,--     /       ,--     /       ,--     /       ,--     /       ,--     /       ,--     /
// b + c √ d + e √ f + g √ h + i √ j + k √ l + m √ n + o √ p + q √ r + s √ t + u √ v + w √ x + y √ z + A √ B + C √ D + E √ F + G √ ...
// =-----------=---------------=-------------------------------=---------------------------------------------------------------=------
//                                                             a
// qNew returns reduced rational number
func (qs *Q32s) qNew(d N, n ...Z) (q *Q32, err error) {
	switch len(n) {
	case 1:
		// b/a
		return qs.qNew1(d, n[0])
	case 3:
		// (b+c√d)/a
		return qs.qNew3(d, n[0], n[1], n[2])
	case 5:
		// (b+c√d+e√f)/a
		return qs.qNew5(d, n[0], n[1], n[2], n[3], n[4])
	case 7:
		// (b+c√d+e√f+g√h)/a
		return qs.qNew7(d, n[0], n[1], n[2], n[3], n[4], n[5], n[6])
	default:
		return nil, ErrInvalid
	}
}

// qNew1 simplifies given number b/a
func (qs *Q32s) qNew1(a N, b Z) (*Q32, error) {
	if A, B, err := qs.zFrac(a, b); err != nil {
		return nil, err
	} else {
		return newQ32(A, B), nil
	}
}

// qNew3 simplifies given number (b + c√d)/a
func (qs *Q32s) qNew3(a N, b, c, d Z) (*Q32, error) {
	if c == 0 || d == 0 {
		return qs.qNew1(a, b)                                   // b/a
	} else if d == 1 {
		return qs.qNew1(a, b + c)                               // (b+c)/a
	} else if C, D, err := qs.zSqrt(c, d); err != nil {         // (b + C√D)/a
		return nil, err
	} else if D == 1 {
		return qs.qNew1(a, b + Z(C))                            // (b+c)/a
	} else if A, BC, err := qs.zFracN(a, b, Z(C)); err != nil { // (B + C√D)/A
		return nil, err
	} else {
		B, C := BC[0], BC[1]
		return newQ32(A, B, C, D), nil
	}
}

// qNew5 simplifies given number (b + c√d + e√f)/a
func (qs *Q32s) qNew5(a N, b, c, d, e, f Z) (*Q32, error) {
	if e == 0 || f == 0 {                              // (b + c√d)/a
		return qs.qNew3(a, b, c, d)
	} else if f == 1 {                                 // (b+e + c√d)/a
		return qs.qNew3(a, b+e, c, d)
	} else if d == f {                                 // (b + (c+e)√d)/a
		return qs.qNew3(a, b, c+e, d)
	} else if C, D, err := qs.zSqrt(c, d); err != nil { // (b + C√D + e√f)/a
		return nil, err
	} else if D == +1 {                                // (b+C + e√f)/a
		return qs.qNew3(a, b+Z(C), e, f) 
	} else if E, F, err := qs.zSqrt(e, f); err != nil { // (b + C√D + E√F)/a
		return nil, err
	} else if F == 1 {                                 // (b+E, C√D)/a
		return qs.qNew3(a, b+Z(E), Z(C), Z(D))
	} else if D == F {                                 // (b, (C+E)√D)/a)
		return qs.qNew3(a, b, Z(C) + Z(E), Z(D))
	} else if A, BCE, err := qs.zFracN(a, b, Z(C), Z(E)); err != nil { // (B + C√D + E√F)/A
		return nil, err
	} else {
		B, C, E := BCE[0], BCE[1], BCE[2]
		if D < F {
			return newQ32(A, B, C, D, E, F), nil
		} else {
			return newQ32(A, B, E, F, C, D), nil
		}
	}
}

// qNew7 returns a form like (b+c√d+e√(f+g√h))/a
func (qs *Q32s) qNew7(a N, b, c, d, e, f, g, h Z) (*Q32, error) {
	if g == 0 {                                         // (b+c√d+e√f)/a
		return qs.qNew5(a, b, c, d, e, f)            
	} else if h == +1 {                                 // (b+c√d+e√f+g)/a
		return qs.qNew5(a, b, c, d, e, f+g)          
	} else if e == 0 {                                  // (b+c√d)/a
		return qs.qNew3(a, b, c, d)               
	} else if G, H, err := qs.zSqrt(g, h); err != nil { // (b+c√d+e√(f+G√H))/a
		return nil, err
	} else if H == + 1 {                                // (b+c√d+e√f+G)/a
		return qs.qNew5(a, b, c, d, e, f+Z(G))

	} else if E, FG, err := qs.zSqrtN(e, f, Z(G)); err != nil { // (b+c√d+E√(F+G√H))/a
		return nil, err
	} else if C, D, err := qs.zSqrt(c, d); err != nil {         // (b+C√D+E√(F+G√H))/a
		return nil, err
	} else if A, BCE, err := qs.zFracN(a, b, Z(C), Z(E)); err != nil { // (B + C√E + D√(F+G√H))/A
		return nil, err
	} else {
		B, C, E := BCE[0], BCE[1], BCE[2]
		F, G := FG[0], FG[1]
		return newQ32(A, B, C, D, E, F, H, G), nil
	}
}

// qAdd adds the given numbers and returns the result
func (qs *Q32s) qAdd(q ...*Q32) (s *Q32, err error) {
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
		if s, err = qs.qAddPair(max, min); err != nil {
			return
		}
		max = s
	}
	return
}

func (qs *Q32s) qMul(q ...*Q32) (s *Q32, err error) {
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
		if s, err = qs.qMulPair(max, min); err != nil {
			return
		}
		max = s
	}
	return
}

func (qs *Q32s) qSqrt(q *Q32) (s *Q32, err error) {
	if q == nil {
		return nil, nil
	}
	switch len(q.num) {
	case 1:
		a, b := q.den, q.num[0]
		//  b      1√ab   0 + C√D
		// --- -> ----- = -------
		//  a       a       A 
		return qs.qNew(N(a), 0, 1, Z(a)*Z(b))

	case 3:
		// b + c√d     √(ab + ac√d)    0 + C√(D + E√F)
		// ------- -> ------------- = ----------------
		//    a             a               A
		// TODO finish and return a Q of size 5 a,b,c,d,e

	}
	return nil, fmt.Errorf("Can't square root of %s", q.String())
}

func (qs *Q32s) qAddPair(q, r *Q32) (s *Q32, err error) {
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
		return qs.qNew(a, qb*aq + rb*ar)

	case 3:
		qc, qd := Z(q.num[1]), Z(q.num[2])
		switch len(r.num) {
		case 1:
			// qb + qc√qd   rb   qb*aq + rb*ar + (qc*aq)√qd    b + c√qd
			// ---------- + -- = -------------------------- = --------
			//    qa        ra              a                     aa
			return qs.qNew(a, qb*aq + rb*ar, qc*aq, qd)			

		case 3:
			rc, rd := Z(r.num[1]), Z(r.num[2])
			if qd == rd { // simpler case
				// qb + qc√qd   rb + rc√qd  qb*aq + rb*ar + (qc*aq + rc*ar)√qd
				// ---------- + --------- = ----------------------------------
				//     qa          ra                     a 
				return qs.qNew(a, qb*aq + rb*ar, qc*aq + rc*ar, qd)
			}
			if qb == rb && qb == 0 { // simpler case both b's=0
				// qc√qd   rc√rd   qc*aq√qd + rc*ar√rd
				// ----- + ----- = -------------------
				//   qa     ra              a
				return qs.qNew(a, 0, qc*aq, qd, rc*ar, rd)
			}
		}
	}
	return nil, fmt.Errorf("Can't add pair %s and %s", q, r)
}

func (qs *Q32s) qMulPair(q, r *Q32) (s *Q32, err error) {
	qa, qb := N(q.den), Z(q.num[0])
	ra, rb := N(r.den), Z(r.num[0])
	aa := qa * ra
	bb := qb * rb
	switch len(q.num) {
	case 1:
		// qb/aq * rb/ra =  (qb*rb)/(qa*ra) = aa/bb
		return qs.qNew(aa, bb)

	case 3:
		qc, qd := Z(q.num[1]), Z(q.num[2])
		switch len(r.num) {
		case 1:
			// qb + qc√qd   rb   qb*rb + qc*rb√qd    bb + qcrb√qd   b32 + c32√d32
			// ---------- * -- = ----------------- = ------------ = -------------
			//    qa        ra       qa * ra              aa             a32
			return qs.qNew(aa, bb, qc*rb, qd) // a, b, c, d
		case 3:
			rc, rd := Z(r.num[1]), Z(r.num[2])
			if qd == rd { // simpler case
				// qb + qc√qd   rb + rc√qd   (bb + qc*rc*qd) + (qb*rc + rb*qc)√qd
				// ---------- * --------- = -----------------------------------
				//     qa          ra                     aa
				return qs.qNew(aa, bb + qc*rc*qd, qb*rc + rb*qc, qd)
			}
		}
	}
	return nil, fmt.Errorf("Can't mul pair %s and %s", q, r)
}






