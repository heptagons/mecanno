package nest

import (
	"fmt"
)

// A32s is a factory to produce algebraic numbers A32
// uses Z32s factory which produce 32 bit integers
type A32s struct {
	*Z32s
}

// NewA32s creates A32s factory
func NewA32s() *A32s {
	return &A32s{
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
// aNew produces a A32 number trying to simplify and reduce.
func (qs *A32s) aNew(d N, n ...Z) (q *A32, err error) {
	switch len(n) {
	case 1:
		// b/a
		return qs.aNew1(d, n[0])
	case 3:
		// (b+c√d)/a
		return qs.aNew3(d, n[0], n[1], n[2])
	case 5:
		// (b+c√d+e√f)/a
		return qs.aNew5(d, n[0], n[1], n[2], n[3], n[4])
	case 7:
		// (b+c√d+e√f+g√h)/a
		return qs.aNew7(d, n[0], n[1], n[2], n[3], n[4], n[5], n[6])
	default:
		return nil, ErrInvalid
	}
}

// aNew1 produces a most a number of size 1, that is number of the form:
//	
//	b / a
//
func (qs *A32s) aNew1(a N, b Z) (*A32, error) {
	if A, B, err := qs.zFrac(a, b); err != nil {
		return nil, err
	} else {
		return newA32(A, B), nil
	}
}

// aNew3 produces at most a number of size 3, that is a number of the form
//	
//	(b + c√d) / a
//	
// Where c != 0 and d != 1.
// Otherwise returns aNew1 versions.
func (qs *A32s) aNew3(a N, b, c, d Z) (*A32, error) {
	if c == 0 || d == 0 {
		return qs.aNew1(a, b)                                   // b/a
	} else if d == 1 {
		return qs.aNew1(a, b + c)                               // (b+c)/a
	} else if C, D, err := qs.zSqrt(c, d); err != nil {         // (b + C√D)/a
		return nil, err
	} else if D == 1 {
		return qs.aNew1(a, b + Z(C))                            // (b+c)/a
	} else if A, BC, err := qs.zFracN(a, b, Z(C)); err != nil { // (B + C√D)/A
		return nil, err
	} else {
		B, C := BC[0], BC[1]
		return newA32(A, B, C, D), nil
	}
}

// aNew5 produces at most a number of size 5, that is a number of the form:
//
//	(b + c√d + e√f) / a
//
// Where c != 0, d != 1, e != 0, f != 1 and d > f.
// Otherwise returns aNew3 versions.
func (qs *A32s) aNew5(a N, b, c, d, e, f Z) (*A32, error) {
	if e == 0 || f == 0 {                              // (b + c√d)/a
		return qs.aNew3(a, b, c, d)
	} else if f == 1 {                                 // (b+e + c√d)/a
		return qs.aNew3(a, b+e, c, d)
	} else if d == f {                                 // (b + (c+e)√d)/a
		return qs.aNew3(a, b, c+e, d)
	} else if C, D, err := qs.zSqrt(c, d); err != nil { // (b + C√D + e√f)/a
		return nil, err
	} else if D == +1 {                                // (b+C + e√f)/a
		return qs.aNew3(a, b+Z(C), e, f) 
	} else if E, F, err := qs.zSqrt(e, f); err != nil { // (b + C√D + E√F)/a
		return nil, err
	} else if F == 1 {                                 // (b+E, C√D)/a
		return qs.aNew3(a, b+Z(E), Z(C), Z(D))
	} else if D == F {                                 // (b, (C+E)√D)/a)
		return qs.aNew3(a, b, Z(C) + Z(E), Z(D))
	} else if A, BCE, err := qs.zFracN(a, b, Z(C), Z(E)); err != nil { // (B + C√D + E√F)/A
		return nil, err
	} else {
		B, C, E := BCE[0], BCE[1], BCE[2]
		if D < F {
			return newA32(A, B, C, D, E, F), nil
		} else {
			return newA32(A, B, E, F, C, D), nil
		}
	}
}

// aNew7 produces at most a number of size 7, that is a number of the form:
//
//	(b + c√d + e√(f+g√h)) / a
//
// Where e != 0, f != 0, g != 0, h != 1.
// Otherwise returns aNew5 or aNew3 versions.
func (qs *A32s) aNew7(a N, b, c, d, e, f, g, h Z) (*A32, error) {
	if g == 0 {                                         // (b+c√d+e√f)/a
		return qs.aNew5(a, b, c, d, e, f)            
	} else if h == +1 {                                 // (b+c√d+e√f+g)/a
		return qs.aNew5(a, b, c, d, e, f+g)          
	} else if e == 0 {                                  // (b+c√d)/a
		return qs.aNew3(a, b, c, d)               
	} else if G, H, err := qs.zSqrt(g, h); err != nil { // (b+c√d+e√(f+G√H))/a
		return nil, err
	} else if H == + 1 {                                // (b+c√d+e√f+G)/a
		return qs.aNew5(a, b, c, d, e, f+Z(G))

	} else if E, FG, err := qs.zSqrtN(e, f, Z(G)); err != nil { // (b+c√d+E√(F+G√H))/a
		return nil, err
	} else if C, D, err := qs.zSqrt(c, d); err != nil {         // (b+C√D+E√(F+G√H))/a
		return nil, err
	} else if A, BCE, err := qs.zFracN(a, b, Z(C), Z(E)); err != nil { // (B + C√E + D√(F+G√H))/A
		return nil, err
	} else {
		F, G := FG[0], FG[1]
		// try to denest √(F+G√H)
		if den, num, err := qs.zSqrtDenest3(Z(F), Z(G), Z(H)); err != nil {
			return nil, err
		} else {
			a := N(A)*N(den)
			b, c, e := Z(BCE[0])*Z(den), Z(BCE[1])*Z(den), Z(BCE[2])*Z(den)
			d := Z(D)
			switch len(num) {
			default: // and case 0
				// cannot denest
				break
			case 1:
				// Denested √(F+G√H) = s
				s := Z(num[0])
				// b + c√d + e(s)
				// b + e + c√d
				return qs.aNew3(a, b + e*s, c, d)
			case 3:
				// Denested √(F+G√H) = s + t√u
				s    := Z(num[0])
				t, u := Z(num[1]), Z(num[2])
				if d == u {
					// b + c√d + e(s + t√d))
					// b + es + (c+et)√d
					return qs.aNew3(a, b + e*s, c + e*t, d)
				} else {
					// b = c√d + e(s + t√u)
					// b + es + c√d + et√u
					return qs.aNew5(a, b + e*s, c, d, e*t, u)
				}
			case 5:
				// Denested √(F+G√H) = s + t√u + v√w
				s    := Z(num[0])
				t, u := Z(num[1]), Z(num[2])
				v, w := Z(num[3]), Z(num[4])
				if d == u {
					// b + c√d + e(s + t√d + v√w)
					// b + es + (c + et)√d + ev√w
					return qs.aNew5(a, b + e*s, c + e*t, d, e*v, w)
				} else if d == w {
					// b + c√d + e(s + t√u + v√d)
					// b + es + (c + ev)√d + et√u
					return qs.aNew5(a, b + e*s, c + e*v, d, e*t, u)
				} else { // u != e && w != e
					// forget denest we dont want to have a sum
					// like B + x√e + y√u + z√w with three radicals √
					// which forces us to go with a more complicated aNew9
					break
				}
			}
		}
		B, C, E := BCE[0], BCE[1], BCE[2]
		// √(F+G√H) cannot be denested
		return newA32(A, B, C, D, E, F, G, H), nil
	}
}

// aAddN adds the given numbers and returns a new result number.
// Call multiple times this function aAdd and follow its limitations.
// Return an error if summand sizes cannot be added.
func (qs *A32s) aAddN(q ...*A32) (s *A32, err error) {
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
		if s, err = qs.aAdd(max, min); err != nil {
			return
		}
		max = s
	}
	return
}

// aMulN multiplies the given numbers and return a new result number.
// Call multiple times this function aMul and follows its limitations.
// Return an error if multiplands sizes cannot be multiplied.
func (qs *A32s) aMulN(q ...*A32) (s *A32, err error) {
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
		if s, err = qs.aMul(max, min); err != nil {
			return
		}
		max = s
	}
	return
}

// aAdd return the addition of the two given numbers.
func (qs *A32s) aAdd(q, r *A32) (s *A32, err error) {
	if q == nil || r == nil {
		return nil, nil
	}
	max, min := q, r
	if len(q.num) < len(r.num) {
		max, min = r, q
	}
	A, B := max.ab()
	a, b := min.ab()
	// Use lcm always to prevent overflows
	w := (A / Ngcd(A, a)) * a // lcm denominator
	U := Z(w / A) // upper factor for max numerator terms
	u := Z(w / a) // upper factor for min numerator terms
	switch len(max.num) {
	case 1:
		//  B     b     BU + bu    x
		// --- + --- = -------- = ---
		//  A     a        w       w
		return qs.aNew1(w,
			B*U + b*u) // x

	case 3:
		C, D := max.cd()
		switch len(min.num) {
		case 1:
			// B + C√D    b    BU + bu + CU√D    x + y√z
			// ------- + --- = -------------- = --------
			//    A       a          w             w
			return qs.aNew3(w,
				B*U + b*u, // x
				C*U, D)    // y√z

		case 3:
			c, d := Z(min.num[1]), Z(min.num[2])
			if D == d {
				// simpler case
				// B + C√D   b + c√D    BU + bu + (CU + cu)√D     x + y√z
				// ------- + --------- = ---------------------- = -------
				//    A         a               w                   w
				return qs.aNew3(w,
					B*U + b*u,    // x
					C*U + c*u, D) // y√z
			}
			//  B + C√D   b + c√d   BU + bu + CU√D + cu√d   x + y√z + o√i
			// -------- + ------- = --------------------- = -------------
			//      A        a                 w                  w
			return qs.aNew5(w, 
				B*U + b*u, // x
				C*U, D,    // y√z
				c*u, d)    // o√i
		}

	case 5:
		C, D, E, F := max.cdef()
		switch len(min.num) {
		case 1:
			// B + C√D + E√F    b    BU + bu + CU√D + EU√F    x + y√z + o√i
			// ------------- + --- = ---------------------- = -------------
			//        A         a               w                   w
			return qs.aNew5(w, 
				B*U + b*u, // x
				C*U, D,    // y√z
				E*U, F)    // o√i

		case 3:
			c, d := Z(min.num[1]), Z(min.num[2])
			if D == d {
				// B + C√D + E√F    b + c√D    BU + bu + (CU+cu)√D + EU√F    x + y√z + o√i
				// ------------- + -------- = --------------------------- = --------------
				//        A            a                   w                      w
				return qs.aNew5(w, 
					B*U + b*u,    // x
					C*U + c*u, D, // y√z
					E*U, F)       // o√i
			}
			if F == d {
				// B + C√D + E√F    b + c√F    BU + bu + CU√D + (EU+cu)√F    x + y√z + o√i
				// ------------- + -------- = --------------------------- = --------------
				//        A            a                   w                      w
				return qs.aNew5(w, 
					B*U + b*u,    // x
					C*U, D,       // y√z
					E*U + c*u, F) // o√i
			}
		}
	}
	return nil, fmt.Errorf("Can't add pair %s and %s", q, r)
}

// aMul return the multiplication of the two given numbers.
// Limitations:
//	- Both numbers sizes=1 are multiplied ok.
//	- One number size=3 and the other size=1 are multiplied ok.
//	- Both numbers sizes=3 are multiplied ok.
//	- For other sizes combination "Can't mul pair" error is returned.
func (qs *A32s) aMul(q, r *A32) (s *A32, err error) {
	if q == nil || r == nil {
		return nil, nil
	}
	max, min := q, r
	if len(q.num) < len(r.num) {
		max, min = r, q
	}
	switch len(max.num) {
	case 1:
		A, B := max.ab()
		a, b := min.ab()
		//  B * b = Bb 
		return qs.aNew(A*a,
			B*b) // x

	case 3:
		A, B := max.ab()
		C, D := max.cd()
		switch len(min.num) {
		case 1:
			a, b := min.ab()
			// = (B + C√D) * b
			// = Bb + Cb√D
			// =  x +  y√z
			return qs.aNew3(A*a,
				B*b,    // x
				C*b, D) // y√z

		case 3:
			a, b := min.ab()
			c, d := min.cd()
			if D < d {
				A, B = min.ab(); C, D = min.cd()
				a, b = max.ab(); c, d = max.cd()
			}
			if B == 0 {
				if b == 0 {
					// C√D * c√d = Cc√Dd = y√z
					return qs.aNew3(A*a,
						0,        // x
						C*c, D*d) // y/z
				}
				// = C√D * (b + c√d)
				// = Cc√Dd + bC√D
				// =  y√z  +  e√f
				return qs.aNew5(A*a,
					0,        // x
					C*c, D*d, // y√z
					b*C, D)   // e√f
			}
			if b == 0 {
				// = (B + C√D) * c√d
				// = Bc√d + Cc√Dd
				// =  y√z +  e√f
				return qs.aNew5(A*a,
					0,          // x
					B*c, d,     // y√z
					C*c,   D*d) // e√f
			}
			if D == d {
				// = (B + C√D) * (b * c√D)
				// = Bb + CcD + Bc√D + bC√D
				// = Bb + CcD + (Bc+bC)√D
				// =    x     +    y   √z
				return qs.aNew3(A*a,
					B*b + C*c*D,  // x
					B*c + b*C, D) // y√z
			}//  ;fmt.Println("5")
			// = (B + C√D) * (b * c√d) 
			// = Bb + Bc√d + bC√D + Cc√Dd
			// = Bb + Bc√d + (C√D)*(b+c√d)
			// = Bb + Bc√d + (C√D)*√((b+c√d)(b+c√d))
			// = Bb + Bc√d + (C√D)*√(b²+2bc√d+c²d)
			// = Bb + Bc√d + (C√D)*√(b²+c²d + 2bc√d)
			// = Bb + Bc√d + C√(bD²+c²Dd + 2bcD√d)
			// =  x +  y√z + e√(    f    +   g √h)
			return qs.aNew7(A*a,
				B*b,             // x
				B*c, d,          // y√z
				C,               // e
				b*D*D + c*c*D*d, // f
				2*b*c*D, d)      // g√h
		}
	}
	return nil, fmt.Errorf("Can't mul pair %s and %s", q, r)
}

// aSqrt returns the square root of the given number.
// Limitations:
//	Square root of numbers of size 1 are returned ok.
//	Square root of numbers of size 3 are returned ok.
//	For other sizes a "Can't square root" error is returned
func (qs *A32s) aSqrt(q *A32) (s *A32, err error) {
	if q == nil {
		return nil, nil
	}
	switch len(q.num) {
	case 1:
		//   / b     √b     √ab    √ab    B + C√D
		//  / --- = ----- = ---- = ---- = -------
		// √   a     √a     √aa     a       A
		a, b := q.ab()
		return qs.aNew3(a,
			0,         // B
			1, Z(a)*b) // C√D

	case 3:
		a, b := q.ab()
		c, d := q.cd()
		if b == 0 {
			//   / c√d    √(c√d)   √a√(c√d)   √(ac√d)  B + C√D + E√(F+G√H)
			//  / ----- = ------ = -------- = ------ = ------------------
			// √    a       √a       √a√a        a             A
			return qs.aNew7(a, // A
				0,         // B = 0
				0, 1,      // C√D = 0
				1,         // E
				0,         // F
				Z(a)*c, d) // G√H
		}
		
		den, num, err := qs.zSqrtDenest3(b, c, d)
		if err != nil {
			return nil, err
		}
		a *= N(den)
		switch len(num) {
		case 1:
			// Denested √(b+c√d) = s
			s := Z(num[0])
			return qs.aNew1(a,
				s) // b
		case 3:
			// Denested √(b+c√d) = s + t√u
			s, t, u := Z(num[0]), Z(num[1]), Z(num[2])
			return qs.aNew3(a,
				s,    // s
				t, u) // t√u
		case 5:
			// Denested √(b+c√d) = s + t√u + v√w
			s, t, u, v, w := Z(num[0]), Z(num[1]), Z(num[2]), Z(num[3]), Z(num[4])
			return qs.aNew5(a,
				s,    // s
				t, u, // t√u
				v, w) // v√w
		default:
			break
		}
		

		//   / b + c√d    √(b + c√d)   √a√(b + c√d)   √(ab + ac√d)   √(F + G√H))
		//  / --------- = ---------- = ------------ = ------------ = -----------
		// √      a           √a           √a√a             a             A
		return qs.aNew7(a, // A
			0,      // B = 0
			0, 1,   // C√D = 0
			1,      // E
			Z(a)*b, // F
			Z(a)*c, // G
			d,      // H
		)
	}
	return nil, fmt.Errorf("Can't square root of %s", q.String())
}

// aPow2 returns the given number raised to the power 2.
// Limitations:
//	Numbers of sizes 1, 3, 5 and 7 are raised ok.
//	For the rest of sizes a "Can't pow2" error is returned.
func (qs *A32s) aPow2(q *A32) (s *A32, err error) {
	if q == nil {
		return nil, nil
	}
	switch len(q.num) {
	case 1:
		// = b*b = b²
		// = B
		a, b := q.ab()
		return qs.aNew(a*a, // A
			b*b)            // B
	
	case 3:
		a, b := q.ab()
		c, d := q.cd()
		if b == 0 {
			// = (c√d)(c√d) = c²d
			// = B
			return qs.aNew(a*a, // A
				c*c*d)          // B
		}
		// x² = (b + c√d)(b + c√d) 
		//    = b² + 2bc√d + c²d
		//    = b² + c²d + 2bc√d
		//    = B + C√D
		return qs.aNew(a*a, // A
			b*b + c*c*d,    // B 
			2*b*c, d)       // C√D

	case 5: 
		a, b := q.ab()
		c, d, e, f := q.cdef()
		if b == 0 {
			// = (c√d + e√f)(c√d + e√f)
			// = c²d + e²f + 2ce√(df)
			// = B + C√D
			return qs.aNew(a*a, // A
				c*c*d + e*e*f,  // B
				2*c*e, d*f)     // C√D
		}
		// = (b + c√d + e√f)(b + c√d + e√f)
		// = (x + e√f)(x + e√f)
		// = x² + 2xe√f + e²f
		// = x² + e²f + 2√(b + c√d)*e√f
		// = x² + e²f + 2e√(bf + cf√d)
		// = b² + c²d + e²f + 2bc√d + 2e√(bf + cf√d)
		return qs.aNew(a*a,      // A
			b*b + c*c*d + e*e*f, // B 
			2*b*c, d,            // C√D
			2*e,                 // E
			b*f,                 // F
			c*f, d)              // G√H

	case 7:
		a, b := q.ab()
		c, d, e, f, g, h := q.cdefgh()
		if b == 0 {
			if c == 0 {
				// = e√(f+g√h) * e√(f+g√h)
				// = e²(f + g√h)
				// = e²f + e²g√h
				// = B + C√D
				return qs.aNew(a*a, // A
					e*e*f,          // B
					e*e*g, h)       // C√D
			} else {
				// = (c√d + e√(f+g√h)) * (c√d + e√(f+g√h))
				// = c²d + 2ce√(df + dg√h) + e²(f + g√h)
				// = c²d + e²f + e²g√h + 2ce√(df + dg√h)
				// = B + C√D + E√(F + G√H)
				return qs.aNew(a*a, // A
					c*c*d + e*e*f,  // B
					e*e*g, h,       // C√D
					2*c*e,          // E
					d*f,            // F
					d*g, h)         // G√H
			}
		} else {
			if c == 0 { // b != 0 and c == 0 is like b = 0 and c=b d=1
				// = (b + e√(f+g√h)) * (b + e√(f+g√h))
				// = b² + 2be√(f+g√h) + e²(f+g√h)
				// = b² + e²f + e²g√h + 2be√(f+g√h)
				// = B + C√D + E√(F + G√H)
				return qs.aNew(a*a, // A
					b*b + e*e*f,    // B
					e*e*g, h,       // C√D
					2*b*e,          // E
					f,              // F
					g, h)           // G√H
			}
		}
	}
	return nil, fmt.Errorf("Can't pow2 of %s", q.String())
}










