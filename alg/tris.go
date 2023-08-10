package alg

type Tris struct {
	*Q32s

	list []*Tri
}

func NewTris(max int, factory *N32s) *Tris {
	ts := &Tris {
		Q32s: &Q32s{
			N32s: factory,
		},
		list: make([]*Tri, 0),
	}
	for a := N32(1); a <= N32(max); a++ {
		for b := N32(1); b <= a; b++ {
			for c := N32(1); c <= b; c++ {
				if b+c > a {
					ts.add(a, b, c)
				}
			}
		}
	}
	return ts
}

func (ts *Tris) add(a, b, c N32) {
	gcd := NatGCD(a, NatGCD(b, c))
	ga, gb, gc := a / gcd, b / gcd, c / gcd
	for _, t := range ts.list {
		if t.abc[0] == ga && t.abc[1] == gb && t.abc[2] == gc {
			// scaled version already stored don't append
			return
		}
	}
	ts.list = append(ts.list, &Tri{
		abc: []N32{ a, b, c },
	})
}

func (ts *Tris) setSinCos() error {
	for _, t := range ts.list {
		a, b, c := t.abc[0], t.abc[1], t.abc[2]

		if cosA, err := ts.cosC(b, c, a); err != nil {
			return err
		} else
		if cosB, err := ts.cosC(c, a, b); err != nil {
			return err
		} else
		if cosC, err := ts.cosC(a, b, c); err != nil {
			return err
		} else {
			t.cos = []*Q32 { cosA, cosB, cosC }
		}
		// https://en.wikipedia.org/wiki/Heron%27s_formula Numerical stability
		area2_4 := Z(a+(b+c)) * Z(c-(a-b)) * Z(c+(a-b)) * Z(a+(b-c))
		// area = √(area2_4)/4
		// https://en.wikipedia.org/wiki/Law_of_sines
		// Area = (absinC)/2 -> sinC = 2A/(a*b)
		if sinA, err := ts.newQ32(2*N(b*c), 0, 1, area2_4); err != nil {
			return err
		} else 
		if sinB, err := ts.newQ32(2*N(c*a), 0, 1, area2_4); err != nil {
			return err
		} else
		if sinC, err := ts.newQ32(2*N(a*b), 0, 1, area2_4); err != nil {
			return err
		} else {
			t.sin = []*Q32 { sinA, sinB, sinC }
		}
	}
	return nil
}

// cosC returns the rational cosine of the angle C using the law of cosines:
//	       a² + b² - c²
//	cosC = ------------
//	           2ab
func (ts *Tris) cosC(a, b, c N32) (*Q32, error) {
	den64 := 2*N(a)*N(b)
	num64 := Z(a)*Z(a) + Z(b)*Z(b) - Z(c)*Z(c)
	return ts.newQ32(den64, num64)
}

//func (ts *Tris) cosCQ(a *Q32, b, c N32) (*Q32, error) {
//	den64 :=
//}




// cosLaw return the third side (squared) cc. Squared to keep simple the Q32 returned.
// Uses the law of cosines to determine the rational algebraic side cc = aa + bb - 2abcosC
func (ts *Tris) cosLaw2(a, b N32, cosC *Q32) (*Q32, error) {
	if aa_bb, err := ts.newQ32(1, Z(a)*Z(a) + Z(b)*Z(b)); err != nil { // a*a + b*b
		return nil, err
	} else if ab, err := ts.newQ32(1, -2*Z(a)*Z(b)); err != nil { // -2a*b
		return nil, err
	} else if abCosC, err := ts.MulQ(ab, cosC); err != nil { // -2a*b*cosC
		return nil, err
	} else {
		return ts.AddQ(aa_bb, abCosC) // a*a + b*b - 2*a*b*cosC
	}
}

func (ts *Tris) addPair(tA, tB *Tri, pA, pB int) (*TriPair, error) {
	pair, err := newTriPair(tA, tB, pA, pB)
	if err != nil {
		return nil, err
	}
	sinA, cosA := tA.sin[pA], tA.cos[pA]
	sinB, cosB := tB.sin[pB], tB.cos[pB]
	// sin(A+B) = sinAcosB + cosAsinB
	if sinAcosB, err := ts.MulQ(sinA, cosB); err != nil {
		return nil, err 
	} else if sinBcosA, err := ts.MulQ(sinB, cosA); err != nil {
		return nil, err
	} else if sinAB, err := ts.AddQ(sinAcosB, sinBcosA); err != nil {
		return nil, err
	} else {
		pair.sin = sinAB
	}
	// cos(A+B) = cosAcosB - sinAsinB
	if cosAcosB, err := ts.MulQ(cosA, cosB); err != nil {
		return nil, err 
	} else if sinAsinB, err := ts.MulQ(sinA, sinB); err != nil {
		return nil, err
	} else if cosAB, err := ts.AddQ(cosAcosB, sinAsinB.Neg()); err != nil {
		return nil, err
	} else {
		pair.cos = cosAB
	}
	return pair, nil
}

func (ts *Tris) AddPairs(results func(pair *TriPair, err error)) {
	n := len(ts.list)
	for p1 := 0; p1 < n; p1++ {
		t1 := ts.list[p1]
		a1s := make(map[N32]bool, 0)
		for a1, s1 := range t1.abc {
			if _, repeated := a1s[s1]; repeated {
				continue
			}
			a1s[s1] = true
			for p2 := p1; p2 < n; p2++ {
				t2 := ts.list[p2]
				a2s := make(map[N32]bool, 0)
				for a2, s2 := range t2.abc {
					if _, repeated := a2s[s2]; repeated {
						continue
					}
					a2s[s2] = true
					if p1 == p2 && a1 < a2 {
						continue
					}
					results(ts.addPair(t1, t2, a1, a2))
				}
			}
		}
	}
}

func (ts *Tris) setTriqs(t *TriPair) error {
	// build tris, triangles with two sides natural and one side Q
	triqs := make([]*TriQ, 0)
	for _, a := range t.tA.otherSides(t.pA) {
		for _, b := range t.tB.otherSides(t.pB) {
			max, min := a, b
			if max < min {
				max, min = b, a
			}
			repeated := false
			for _, triq := range triqs {
				if max == triq.max && min == triq.min {
					repeated = true
				}
			}
			if !repeated {
				if cc, err := ts.cosLaw2(max, min, t.cos); err != nil {
					return err
				} else if c, err := ts.sqrtQ(cc); err != nil {
					return err
				} else if len(c.num) <= 1 && c.den == 1 {
					// prevent a triq with all naturals like below [4 3 2]
						// 4 [2 2 1]'0 [4 3 2]'2 sin=√15/4 cos=-1/4
					// tris=[[2√6 4 2] [4 3 2] [√19 4 1] [√46/2 3 1]]
					continue
				} else if triq, err := newTri32Q(max, min, cc, c); err != nil {
					return err
				} else {
					triqs = append(triqs, triq)
				}
			}
		}
	}
	t.triqs = triqs
	return nil
}

