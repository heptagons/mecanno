package nest

import (
	"fmt"
	"strings"
	"testing"
)

func TestTcos(t *testing.T) {

	factory := NewT32s()

	frac := func(num, den Z) string {
		a, _ := factory.aNew1(N(den), num)
		return a.String()
	}
	surdFrac := func(surd, den Z32) string {
		a, _ := factory.aNew3(N(den), 0, 1, Z(surd))
		return a.String()
	}

	for _, r := range []struct { a, b, c N32; cosines, sines string } {
		{ 1, 1, 1, "1/2 1/2 1/2",       "√3/2 √3/2 √3/2"          }, // √3
		{ 2, 2, 1, "1/4 1/4 7/8",       "√15/4 √15/4 √15/8"       }, // √15
		{ 3, 2, 2, "-1/8 3/4 3/4",      "3√7/8 √7/4 √7/4"         }, // √7
		{ 3, 3, 1, "1/6 1/6 17/18",     "√35/6 √35/6 √35/18"      }, // √35
		{ 3, 3, 2, "1/3 1/3 7/9",       "2√2/3 2√2/3 4√2/9"       }, // √2
		{ 4, 3, 2, "-1/4 11/16 7/8",    "√15/4 3√15/16 √15/8"     }, // √15
		{ 4, 3, 3, "1/9 2/3 2/3",       "4√5/9 √5/3 √5/3"         }, // √5
		{ 4, 4, 1, "1/8 1/8 31/32",     "3√7/8 3√7/8 3√7/32"      }, // √7
		{ 4, 4, 3, "3/8 3/8 23/32",     "√55/8 √55/8 3√55/32"     }, // √55
		{ 5, 3, 3, "-7/18 5/6 5/6",     "5√11/18 √11/6 √11/6"     }, // √11
		{ 5, 4, 2, "-5/16 13/20 37/40", "√231/16 √231/20 √231/40" }, // √231
		{ 5, 4, 3, "0 3/5 4/5",         "1 4/5 3/5"               }, // √1
		{ 5, 4, 4, "7/32 5/8 5/8",      "5√39/32 √39/8 √39/8"     }, // √39
		{ 5, 5, 1, "1/10 1/10 49/50",   "3√11/10 3√11/10 3√11/50" }, // √11
		{ 5, 5, 2, "1/5 1/5 23/25",     "2√6/5 2√6/5 4√6/25"      }, // √6
		{ 5, 5, 3, "3/10 3/10 41/50",   "√91/10 √91/10 3√91/50"   }, // √91
		{ 5, 5, 4, "2/5 2/5 17/25",     "√21/5 √21/5 4√21/25"     }, // √21

		{ 7, 6, 5, "1/5 19/35 5/7",     "2√6/5 12√6/35 2√6/7"     }, // √6

	} {
		tr := newT(r.a, r.b, r.c)
		if tr == nil {
			t.Fatalf("T: nil for %d %d %d", r.a, r.b, r.c)
		}
		if cosines := fmt.Sprintf("%s %s %s", 
			frac(factory.cos(tr, TangA)),
			frac(factory.cos(tr, TangB)),
			frac(factory.cos(tr, TangC))); cosines != r.cosines {
			t.Fatalf("T.cosines got %s exp %s", cosines, r.cosines)
		}
		if sines := fmt.Sprintf("%s %s %s",
			surdFrac(factory.sin(tr, TangA)), 
			surdFrac(factory.sin(tr, TangB)),
			surdFrac(factory.sin(tr, TangC))); sines != r.sines {
			t.Fatalf("T.sines got %s exp %s", sines, r.sines)
		}
	}
}

func TestTdiag1(t *testing.T) {

	factory := NewT32s()

	for _, tr := range []*T { 
		newT(3,3,3), // equilateral
		newT(4,3,3), // isoceles
		newT(7,6,5), // scalene
	} {
		fmt.Printf("[%s] diagonals:\n", tr.String())
		for _, ang := range []Tang { TangC, TangB, TangA } {
			diagsAng, den := factory.tDiagsAng(tr, ang)
			fmt.Printf("  ang=%c den=%d:\n", ang, den)
			for d, diags := range diagsAng {
				var surds strings.Builder
				for pos, diag := range diags {
					if pos > 0 {
						surds.WriteString(", ")
					}
					if a, err := factory.aNew3(den, 0, 1, Z(diag)); err != nil {
						continue
					} else {
						surds.WriteString(a.String())
					}
				}
				fmt.Printf("    %d: %v -> %s\n", d, diags, surds.String())
			}
		}
	}
}


func TestTdiag2(t *testing.T) {
	factory := NewT32s()
	/*   *
		/|\
     5 / | \ 6
      /  |  \
     /   |   \
    x1   |8   x2  x1x2 = 3 + 2√5
     \   |   /
      \  |  /
     5 \ | / 6
        \|/
         *
	*/
    for _, pair := range []struct { t1, t2 *T; max, min N32 } {
    	{ newT(8,5,5), newT(8,6,6), 6, 5 },
    	{ newT(7,6,5), newT(6,5,4), 6, 5 },
    	{ newT(5,5,5), newT(5,5,5), 5, 3 },
    	{ newT(5,5,5), newT(5,5,4), 5, 4 },
    	{ newT(5,5,5), newT(5,4,4), 5, 4 },
    } {
    	fmt.Printf("Pair [%s] [%s]\n", pair.t1.String(), pair.t2.String())
		cosX, _ := factory.tCosAplusB(pair.t1, TangB, pair.t2, TangB)
		fmt.Printf("  cosBB = %s\n", cosX.String())
		for p1 := N32(1) ; p1 <= pair.max; p1++ {
			for p2 := N32(1) ; p2 <= pair.min; p2++ {
				if p1 >= p2 {
					x, _ := factory.tLawOfCos(p1, p2, cosX)
					fmt.Printf("    [%d,%d,%s]\n", p1, p2, x)
				}
			}
		}
    }
}

func TestT765diags(t *testing.T) {
	factory := NewT32s()
	a, b, c := N(7), N(6), N(5)
	aa, bb, cc := Z(a)*Z(a), Z(b)*Z(b), Z(c)*Z(c)
	
	m := make(map[string]*A32)

	bc := N(b)*N(c)
	bc2 := Z(bc)*Z(bc)
	b2_c2_a2 := bb + cc - aa
	for _, g := range []struct { y, z Z } {
		{ 1,1 },{ 2,2 },{ 3,3 },{ 4,4 },{ 5,5 },
		{ 2,1 },{ 3,2 },{ 4,3 },{ 5,4 },
		{ 3,1 },{ 4,2 },{ 5,3 },{ 6,4 },
		{ 4,1 },{ 5,2 },{ 6,3 },
		{ 5,1 },{ 6,2 },
		{ 6,1 },
	} {
		in := bc2*(g.y*g.y + g.z*g.z) - Z(bc)*g.y*g.z*b2_c2_a2
		if diag, err := factory.aNew3(bc, 0, 1, in); err == nil {
			m[fmt.Sprintf("b_%d,c_%d", g.y, g.z)] = diag
		}
	}

	ac := N(a)*N(c)
	ac2 := Z(ac)*Z(ac)
	a2_c2_b2 := aa + cc - bb
	for _, g := range []struct { x, z Z } {
		{ 1,1 },{ 2,2 },{ 3,3 },{ 4,4 },{ 5,5 },
		{ 2,1 },{ 3,2 },{ 4,3 },{ 5,4 },{ 6,5 },
		{ 3,1 },{ 4,2 },{ 5,3 },{ 6,4 },
		{ 4,1 },{ 5,2 },{ 6,3 },{ 7,4 },
		{ 5,1 },{ 6,2 },{ 7,3 },
		{ 6,1 },{ 7,2 },
		{ 7,1 },
	} {
		in := ac2*(g.x*g.x + g.z*g.z) - Z(ac)*g.x*g.z*a2_c2_b2
		if diag, err := factory.aNew3(ac, 0, 1, in); err == nil {
			m[fmt.Sprintf("a_%d,c_%d", g.x, g.z)] = diag
		}
	}

	ab := N(a)*N(b)
	ab2 := Z(ab)*Z(ab)
	a2_b2_c2 := aa + bb - cc
	for _, g := range []struct { x, y Z } {
		{ 1,1 },{ 2,2 },{ 3,3 },{ 4,4 },{ 5,5 },{ 6,6 },
		{ 2,1 },{ 3,2 },{ 4,3 },{ 5,4 },{ 6,5 },
		{ 3,1 },{ 4,2 },{ 5,3 },{ 6,4 },{ 7,5 },
		{ 4,1 },{ 5,2 },{ 6,3 },{ 7,4 },
		{ 5,1 },{ 6,2 },{ 7,3 },
		{ 6,1 },{ 7,2 },
		{ 7,1 },
	} {
		in := ab2*(g.x*g.x + g.y*g.y) - Z(ab)*g.x*g.y*a2_b2_c2
		if diag, err := factory.aNew3(ab, 0, 1, in); err == nil {
			m[fmt.Sprintf("a_%d,b_%d", g.x, g.y)] = diag	
		}
	}

	fmt.Println("A[b,c]")
	for x:=N(1); x <= c; x++ {
		for y:=N(1); y <= b; y++ {
			if diag, ok := m[fmt.Sprintf("b_%d,c_%d", y, x)]; ok {
				fmt.Printf("%12s ", diag)
			} else {
				fmt.Printf("%12s ", "x")
			}
		}
		fmt.Println()
	}
	fmt.Println("B[a,c]")
	for x:=N(1); x <= c; x++ {
		for y:=N(1); y <= a; y++ {
			if diag, ok := m[fmt.Sprintf("a_%d,c_%d", y, x)]; ok {
				fmt.Printf("%12s ", diag)
			} else {
				fmt.Printf("%12s ", "x")
			}
		}
		fmt.Println()
	}
	fmt.Println("C[a,b]")
	for x:=N(1); x <= b; x++ {
		for y:=N(1); y <= a; y++ {
			if diag, ok := m[fmt.Sprintf("a_%d,b_%d", y, x)]; ok {
				fmt.Printf("%12s ", diag)
			} else {
				fmt.Printf("%12s ", "x")
			}
		}
		fmt.Println()
	}
}

func TestTalphas(t *testing.T) {
	factory := NewT32s()
	sqrt := N32(24)
	slur, _ := factory.aNew3(1, 0, 1, Z(sqrt))
	for _, tri := range newTalphas(sqrt) {
		if cosA, cosB, cosC, err := factory.tAlphaCosines(tri); err != nil {
			t.Fatalf("error %v", err)
		} else {
			t.Logf("(%s,%d,%d) %s %s %s",
				slur.String(), tri.b, tri.c, cosA, cosB, cosC)
		}
	}
}

func TestTbetas(t *testing.T) {
	factory := NewT32s()
	sqrt := N32(24)
	slur, _ := factory.aNew3(1, 0, 1, Z(sqrt))
	for _, tri := range newTbetas(sqrt,9) {
		if cosA, cosB, cosC, err := factory.tBetaCosines(tri); err != nil {
			t.Fatalf("error %v", err)
		} else {
			t.Logf("(%d,%s,%d) %s %s %s",
				tri.a, slur.String(), tri.c, cosA, cosB, cosC)
		}
	}
}

func TestTgammas(t *testing.T) {
	factory := NewT32s()
	sqrt := N32(24)
	slur, _ := factory.aNew3(1, 0, 1, Z(sqrt))
	for _, tri := range newTgammas(sqrt,9) {
		if cosA, cosB, cosC, err := factory.tGammaCosines(tri); err != nil {
			t.Fatalf("error %v", err)
		} else {
			t.Logf("(%d,%d,%s) %s %s %s",
				tri.a, tri.b, slur.String(), cosA, cosB, cosC)
		}
	}
}

func TestTCosAplusB(t *testing.T) {

	factory := NewT32s()

	ts := &Ts{}
	ts.AddTris(3)

	type Sum struct {
		t string
		a *A32
	}

	sums := make(map[string]*Sum, 0)
	list := make([]string, 0)
	
	for i:=0; i < len(ts.tris); i++ {
		m := ts.tris[i]
		mCos := &tRats{}
		for _, ang := range []Tang{ TangA, TangB, TangC } {
			num, den := factory.cos(m, ang)
			mCos.addRat(ang, num, den)
		}
		for j:=0; j <= i; j++ {
			n := ts.tris[j]
			nCos := &tRats{}
			for _, ang := range []Tang{ TangA, TangB, TangC } {
				num, den := factory.cos(n, ang)
				nCos.addRat(ang, num, den)
			}
			for _, mRat := range mCos.rats {
				dA, nA := mRat.den, mRat.num
				for _, nRat := range nCos.rats {
					dB, nB := nRat.den, nRat.num
					a := N(dA)*N(dB)
					b := nA*nB
					c := Z(-1)
					d := (dA + nA)*(dA - nA) * (dB + nB)*(dB - nB)
					if cosAB, err := factory.aNew3(a,b,c,d); err == nil {
						//if len(cosAB.num) > 2 && cosAB.num[2] == 5 {
							key := cosAB.Key()
							if _, ok := sums[key]; !ok {
								ts := fmt.Sprintf("(%s)[%c] & (%s)[%c]", m, mRat.angle, n, nRat.angle)
								sums[key] = &Sum{
									t: ts,
									a: cosAB,
								}
								list = append(list, fmt.Sprintf("%s & $%s$\\\\ %%%s", ts, cosAB.Tex(), cosAB.String()))
							}
						//}
					}
				}
			}
		}
	}
	for _, l := range list {
		fmt.Println(l)
	}
}



