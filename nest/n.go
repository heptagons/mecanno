package nest

import (
	"fmt"
	"math"
)

var ErrOverflow = fmt.Errorf("Overflow")
var ErrInfinite = fmt.Errorf("Infinite")
var ErrInvalid  = fmt.Errorf("Invalid")

type N uint64

func Ngcd(a, b N) N {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a > b {
		return Ngcd(b, a % b)
	}
	return Ngcd(a, b % a)
}

func (a *N) Reduce2(b *Z) N {
	var bb N
	if *b < 0 {
		bb = N(-*b)
	} else {
		bb = N(*b)
	}
	g := Ngcd(*a, bb)
	if g > 1 {
		*a /= g
		*b /= Z(g)
	}
	return g
}

func (a *N) Reduce3(b, c *Z) N {
	var bb, cc N
	if *b < 0 {
		bb = N(-*b)
	} else {
		bb = N(*b)
	}
	if *c < 0 {
		cc = N(-*c)
	} else {
		cc = N(*c)
	}
	g := Ngcd(Ngcd(*a, bb), cc)
	if g > 1 {
		*a /= g
		*b /= Z(g)
		*c /= Z(g)
	}
	return g
}

func (a *N) Reduce4(b, c, e *Z) N {
	var bb, cc, ee N
	if *b < 0 {
		bb = N(-*b)
	} else {
		bb = N(*b)
	}
	if *c < 0 {
		cc = N(-*c)
	} else {
		cc = N(*c)
	}
	if *e < 0 {
		ee = N(-*e)
	} else {
		ee = N(*e)
	}
	g := Ngcd(Ngcd(Ngcd(*a, bb), cc), ee)
	if g > 1 {
		*a /= g
		*b /= Z(g)
		*c /= Z(g)
		*e /= Z(g)
	}
	return g
}


const N32_MAX = N(0xffffffff)

type N32 uint32 // range 0 - 0xffffffff

// gcd returns the greatest common divisor of 
// this natural and the other given
func (a N32) gcd(b N32) N32 {
	if b == 0 {
		return a
	}
	return b.gcd(a % b)
}

// NatGCD returns the greatest common divisor of two naturals
func NatGCD(a, b N32) N32 {
	if b == 0 {
		return a
	}
	return NatGCD(b, a % b)
}

func (a *N32) Reduce2(b *N32) {
	if g := NatGCD(*a, *b); g > 1 {
		*a /= g
		*b /= g
	}
}

func (a *N32) Reduce3(b, c *N32) {
	if g := NatGCD(NatGCD(*a, *b), *c); g > 1 {
		*a /= g
		*b /= g
		*c /= g
	}
}


// N32s is factory with a primes list to speed up
// some 32-bit nested algebraic rational numbers
type N32s struct {
	primes []N32
	pow2s  [][]N
}

func NewN32s() *N32s {
	value := 0xffff
    f := make([]bool, value)
    for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
        if f[i] == false {
            for j := i * i; j < value; j += i {
                f[j] = true
            }
        }
    }
    primes := make([]N32, 0)
    for i := N32(2); i < N32(value); i++ {
        if f[i] == false {
            primes = append(primes, i)
        }
    }

    pow2s := make([][]N, 0)
    last := 0   // first pow2
    count := 16 // first table size
    for t := 0; t < 8; t++ { // 16+32+64+128+256+512+1024+2048 squares
		table := make([]N, count)
		for i := 0; i < count; i++ {
			table[i] = N(last+i) * N(last+i)
		}
		pow2s = append(pow2s, table)
		last += count
		count *= 2
	}
	return &N32s{
		primes: primes,
		pow2s:  pow2s,
	}
}

func (n *N32s) nSqrtFloorCeil(num N) (floor, ceil N, err error) {
	if floor, floor2, ceil2, err := n.nSqrtFloor(num); err != nil {
		return 0, 0, err
	} else if floor2 == ceil2 {
		return floor, floor, nil
	} else {
		return floor, floor+1, nil
	}
}

// nSqrtFloor returns for the given number the squred "floor" and "ceiling" of num*num
// Example for given n=133 return floor=144 (12²) and ceil=169 (13²).
func (n *N32s) nSqrtFloor(num N) (sqrt, floor, ceil N, err error) {
	if num == 0 {
		return
	}
	pos := 0
	sqrt = N(0)
	floor = N(0)
	for _, pow2 := range n.pow2s {
		size := len(pow2)
		ceil = pow2[size-1]
		if num <= ceil {
			pos, floor, ceil = nSqrtFloor(floor, num, pow2)
			sqrt += N(pos)
			return
		}
		// look in next table
		// next floor is this ceil
		sqrt += N(size) // accumulate indices which are sqrt-floors
		floor = ceil // pass next table this as the min
	}
	err = ErrOverflow
	return
}

func nSqrtFloor(floorPrev, num N, table []N) (pos int, floor, ceil N) {
	c := len(table) // 16
	d := c / 2      // 8 -> 4 -> 2 -> 1
	pos = d-1       // start with 7
	if num == table[c-1] {
		return c-1, table[c-1], table[c-1]
	}
	for {
		d >>= 1 // d /= 2
		if cur := table[pos]; num == cur {
			// num is a square already in cells
			return pos, cur, cur
		} else if num > cur {
			if d == 0 {
				return pos, cur, table[pos+1]
			}
			pos += d // next look above
		} else {
			if d == 0 {
				if pos == 0 {
					return pos-1, floorPrev, cur
				}
				return pos-1, table[pos-1], cur
			}
			pos -= d // next look below
		}
	}
	return
}

/*
Sqrts_floor_ceil
Table[0] first 16 squares:

0          1/8         2/4         3/8         1/2         5/8           3/4           7/8            1
|           |           |           |           |           |             |             |             |
| a[0] a[1] | a[2] a[3] | a[4] a[5] | a[6] a[7] | a[8] a[9] | a[10] a[11] | a[12] a[13] | a[14] a[15] |
+-----------+-----------+-----------+-----------+-----------+-------------+-------------+-------------+
|   0    1      4    9     16   25     36   49     64   81     100   121     144   169     196   225 
                                             *
                                             |------------------------>|
                                                                       |------------>|

example: 133
 d=8
 pos=7

 d=4
 if 133 == table[7]=49 no
 if 133 > 49 yes:
   pos += 4 => 11
 
 d=2
 if 133 = table[11]=121 no
 if 133 > 121 yes:
   pos += 2 => 13

 d=1
 if 133 == table[13]=169 no
 if 133 > 169 no
 if 133 < 169 yes
   pos -= 1 => 12

 d=0
 if 133 == table[12]=144 no
 if 133 > 144 no
 if 133 < 144 yes
   if d==0 {
	 return 11, table[11], table[12]
   }

*/

