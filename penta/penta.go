package penta

import (
	"math"

	"github.com/heptagons/meccano"
	"github.com/heptagons/meccano/nest"
)

func Type_1(max int) *meccano.Sols {

	sols := &meccano.Sols{}

	check := func(a, b, c int) {
		f := float64(a*a + b*b + c*c - a*c)
		if f < 0 {
			return
		}
		if d := int(math.Sqrt(f)); math.Pow(float64(d), 2) == f {
			sols.Add(a, b, c, d)
		}
	}

	for a := 1; a < max; a++ {
		for b := 1; b <= a; b++ {
			for c := 0; c <= a; c++ {
				if a*c == (a + c)*b {
					check(a, b, c)
				}
			}
		}
	}
	return sols
}

func Type_2(max int) *meccano.Sols {

	sols := &meccano.Sols{}

	check := func(a, b, c, d int) {
		f := float64(a*a + b*b + c*c + d*d - a*c - b*d - a*b)
	    if f < 0 {
	    	return
	    }
		if e := int(math.Sqrt(f)); math.Pow(float64(e), 2) == f {
			sols.Add(a, b, c, d, e)
		}
	}

    for a := 1 ; a < max; a++ {
    	for b := 1; b < a; b++ {
        	for c := 0; c < a; c++ {
          		for d := 1; d < a; d++ {
            		if ((a - b)*(c - d) + a*b == c*d) {
              			check(a, b, c, d)
              		}
              	}
            }
        }
    }
    return sols
}

// this function fails to find what pentagons_type_2 find (roudoff errors)
func Type_2b(max int) {
	sols := &meccano.Sols{}
	cosA, sinA := math.Cos(2*math.Pi/5), math.Sin(2*math.Pi/5)
	cosB, sinB := math.Cos(  math.Pi/5), math.Sin(  math.Pi/5)
	for a := 1; a <= max; a++ {
		ax, ay := float64(a)*cosA, float64(a)*sinA
		for b := 1; b < a; b++ {
			bx, by := float64(b)*cosB, float64(b)*sinB
			for d := 1; d < (a-b); d++ {
				dx, dy := float64(d)*cosB, float64(d)*sinB
				for c := 1; c < a; c++ {
					cx := float64(c)
					ex := cx + ax - bx - dx
					ey := ay + by - dy
					f := ex*ex + ey*ey
					e := int(math.Sqrt(f))
					if (math.Pow(float64(e), 2) - f) == 0 {
						sols.Add(a, b, c, d, e)
					}
				}
			}
		}
	}
}

func Type_2_Half(max int) *meccano.Sols {
	
	sols := &meccano.Sols{}
	
	aa, a_b, ab, bb, dd, ad, bc, c_d, cd, cc := 0,0,0,0,0,0,0,0,0,0
	for a := 1; a <= max; a++ {
		aa = a*a
		for b := 1; b < a; b++ {
			a_b, ab, bb = a - b, a*b, b*b
			for d := 1; d < (a-b); d++ {
				dd, ad = d*d, a*d
				for c := 0; c < a; c++ {
					bc, c_d, cd, cc = b*c, c - d, c*d, c*c
					if a_b * c_d + ab == cd {
						if f := float64(aa + bb + cc + dd - ad - bc - cd); f > 0 {
							if e := int(math.Sqrt(f)); math.Pow(float64(e), 2) == f {
								sols.Add(a, b, c, d, e)
							}
						}
					}
				}
			}
		}
	}
	return sols
}

func Type_2_HalfWithConjecture(max int) *meccano.Sols {
	
	sols := &meccano.Sols{}
	
	aa, a_b, ab, bb, dd, ad, bc, c_d, cd, cc := 0,0,0,0,0,0,0,0,0,0
	for a := 1; a <= max; a++ {
		aa = a*a
		for b := 1; b < a; b++ {
			a_b, ab, bb = a - b, a*b, b*b
			for d := 1; d < (a-b); d++ {
				dd, ad = d*d, a*d
				for c := 0; c < a; c++ {
					bc, c_d, cd, cc = b*c, c - d, c*d, c*c
					if a_b * c_d + ab == cd {

						e2 := aa + bb + cc + dd - ad - bc - cd
						x := 1
						for {
							if e := 10*x + 1; e*e == e2 {
								sols.Add(a, b, c, d, e)
								break
							} else if e*e > e2 {
								break
							}
							x++
						}
					}
				}
			}
		}
	}
	return sols
}

type Diagonals struct {
	*nest.A32s
}

func NewDiagonals() *Diagonals {
	return &Diagonals{
		nest.NewA32s(),
	}
}

func (d *Diagonals) Get(min, max int, callback func(a, b, c int, surd *nest.A32)) {
	for a := min; a <= max; a++ {
		for b := 1; b <= a; b++ {
			for c := 0; c <= b; c++ {
				if surd, err := d.GetOne(a,b,c); err == nil {
					callback(a, b, c, surd)
				}
			}
		}
	}
}

func (d *Diagonals) GetOne(a, b, c int) (*nest.A32, error) {
	p := (a-b)*(a-b) + (a-c)*(a-c) + (b-c)*(b-c) + 2*a*a + 2*b*b + 2*c*c
	q := 2*(a*b + a*c - b*c)
	// (b + c√d + e√(f+g√h)) / a
	F := nest.Z(p)
	G := nest.Z(q)
	return d.ANew7(2, 0, 0, 1, 1, F, G, 5)
}

