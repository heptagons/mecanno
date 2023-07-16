package hepta

import (
	"math"

	"github.com/heptagons/meccano"
)

func Diagonals(max int) *meccano.Sols {

	sols := &meccano.Sols{}

	bMax, aa := 0, 0
	for a := 2; a <= max; a++ {
		bMax = int(math.Ceil(float64(a) / 2))
		aa = 2*a*a
		for b := 1; b < bMax; b++ {
			f := float64(aa - a*b + b*b)
			if c := int(math.Sqrt(f)); math.Pow(float64(c), 2) == f {
				sols.Add(a, b, c)
			}
		}
	}
	return sols
}