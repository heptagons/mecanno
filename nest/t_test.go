package nest

import (
	"fmt"
	"strings"
	"testing"
)

func TestTT(t *testing.T) {

	factory := NewT32s()

	frac := func(num, den Z) string {
		if den32, num32, err := factory.zFrac(N(den), num); err != nil {
			return err.Error()
		} else if den32 == 1 {
			return fmt.Sprintf("%d", num32)
		} else {
			return fmt.Sprintf("%d/%d", num32, den32)
		}
	}

	surdFrac := func(surd, den Z32) string {
		if out, in, err := factory.zSqrt(1, Z(surd)); err != nil {
			return err.Error()
		} else if den32, num32, err := factory.zFrac(N(den), Z(out)); err != nil {
			return err.Error()
		} else {
			var sb strings.Builder
			if in == 1 {
				sb.WriteString(fmt.Sprintf("%d", num32))
			} else if num32 == -1 {
				sb.WriteString("-")
			} else if num32 != +1 {
				sb.WriteString(fmt.Sprintf("%d", num32))
			}
			if in > 1 {
				sb.WriteString(fmt.Sprintf("√%d", in))
			}
			if den32 > 1 {
				sb.WriteString(fmt.Sprintf("/%d", den32))	
			}
			return sb.String()
		}
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
			frac(tr.cosA()),
			frac(tr.cosB()),
			frac(tr.cosC())); cosines != r.cosines {
			t.Fatalf("T.cosines got %s exp %s", cosines, r.cosines)
		}
		if sines := fmt.Sprintf("%s %s %s",
			surdFrac(tr.sinA()), 
			surdFrac(tr.sinB()),
			surdFrac(tr.sinC())); sines != r.sines {
			t.Fatalf("T.sines got %s exp %s", sines, r.sines)
		}
	}

	diagsF := func(diagsXY [][]N, den N) {
		for d, diags := range diagsXY {
			var surds strings.Builder
			for pos, diag := range diags {
				if pos > 0 {
					surds.WriteString(" ")
				}
				if num, den, err := factory.zFrac(diag, Z(den)); err != nil {
					t.Fatalf("err %v", err)
				} else if out, in, err := factory.zSqrt(1, Z(num)); err != nil {
					t.Fatalf("err %v", err)
				} else {
					if out == 1 {
						if in > 1 {
							surds.WriteString(fmt.Sprintf("√%d", in))
						} else {
							surds.WriteString("1")
						}
					} else {
						surds.WriteString(fmt.Sprintf("%d", out))
						if in > 1 {
							surds.WriteString(fmt.Sprintf("√%d", in))
						}
					}
					if den > 1 {
						surds.WriteString(fmt.Sprintf("/%d", den))
					}
				}
			}
			fmt.Printf("    diags %d %v -> %s\n", d, diags, surds.String())
		}
	}

	t765 := newT(7,6,5)
	fmt.Println("7,6,5")
	for _, r := range []struct { sides string; f func(t *T) ([][] N, N) } {
		{ "a=7", factory.aDiags },
		{ "b=6", factory.bDiags },
		{ "c=5", factory.cDiags },
	} {
		diags, den := r.f(t765)
		fmt.Printf("  %s den=%d\n", r.sides, den)
		diagsF(diags, den)
	}
}