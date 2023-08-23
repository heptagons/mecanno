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
    for t := 0; t < 2; t++ { // two tables, 16 and 32 squares
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

// nSqrtFloorCeil returns for the given number the "floor" and "ceiling" square roots
// Example for given n=133 return floor=144 (12²) and ceil=169 (13²).
func (n *N32s) nSqrtFloorCeil(num N) (floor, ceil N, err error) {
	floor = N(0)
	for _, pow2 := range n.pow2s {
		ceil = pow2[len(pow2) - 1]
		if num == ceil {
			// num is a square, return ASAP floor=ceil=n
			floor = ceil
			return
		} else if num < ceil {
			floor, ceil = nSqrtFloorCeil(floor, num, pow2)
			return
		}
		// look in next table
		// next floor is this ceil
		floor = ceil // pass next table this as the min
	}
	err = ErrOverflow
	return
}

func nSqrtFloorCeil(floorPrev, num N, table []N) (floor, ceil N) {
	c := len(table) // 16
	d := c/2        // 8 -> 4 -> 2 -> 1
	curPos := c-1-d // start with 7
	floor = floorPrev // start with pos -1
	ceil = table[c-1] // start with pos 15
	for {
		if d == 0 {
			return
		}
		cur := table[curPos]
		if num == cur {
			// num is a square, return ASAP floor=ceil=n
			return num, num
		}
		d /= 2
		if num > cur {
			floor = cur
			curPos += d
		} else {
			ceil = cur
			curPos -= d
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

                                         (-8)|<--------------------------------------------------max
                                             ?                                                           step 2 (c=4)
                                            min ---------------------->|(+4)
                                                                       ?                                 step 3 (c=2)
                                                                      min ---------->|(+2)
                                                                                     ?                   step 4 (c=1)
                                                                           (-1)|<---max
                                                                               ?                         step 5 (c=0)
                                                                              min 


Example request sqrts_floor_ceil of 133:
First we test 133 < table[0].max = 225 on success we jump to table "0":

step 1) Is 133 > a[14]=196 no  c=8 test a[15-c] (a[7])
step 2) Is 133 >  a[7]= 49 yes c=4 test a[7+c] (a[11])
step 3) Is 133 < a[11]=121 yes c=2 goto a[11+c] (a[13])
step 4) Is 133 > a[13]=169 no  c=1 goto a[13-c] (a[12])
step 5) Is 133 < a[12]=144 no  c=0 return min=a[11]=121, max=a[12]=144




*/

