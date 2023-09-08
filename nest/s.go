package nest

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// S32 represents 32 bit surd
// from minimum 0√0 to maximum 0xffffffff√0xffffffff
type S32 struct {
	out Z32
	in  N32
}

// S32s is a factory to operates over sum of surds
// Uses factory Z32s
type S32s struct {
	*Z32s
	surds map[int]Z32 // map[in]out
	den   N32
}

// NewS32s creates a new S32s factory
// with an initial sum equals to zero
func NewS32s(factory *Z32s) *S32s {
	return &S32s{
		Z32s:  factory,
		surds: make(map[int]Z32, 0),
		den:   1,
	}
}

// sAdd add/subtract integers according surds signs
func (s *S32s) sAddZ(surds []Z) error {
	for _, surd := range surds {
		if surd > 0 {
			if err := s.sAdd(N(surd)); err != nil {
				return err
			}
		} else {
			if err := s.sSub(N(-surd)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *S32s) sAddSqrt(out, in Z) error {
	if o32, i32, err := s.zSqrt(out, Z(in)); err != nil {
		return err
	} else {
		k := int(i32)
		if out, ok := s.surds[k]; !ok {
			s.surds[k] = o32
		} else {
			s.surds[k] = out + o32
		}
	}
	return nil
}

func (s *S32s) sDivide(den N) error {
	ins  := make([]int, len(s.surds))
	nums := make([]Z, len(s.surds))
	i := 0
	for in, out := range s.surds {
		ins[i] = in
		nums[i] = Z(out)
		i++
	}
	if newDen, newNums , err := s.zFracN(den, nums...); err != nil {
		return err
	} else {
		for p, num := range newNums {
			in := ins[p]
			s.surds[in] = Z32(num)
		}
		s.den = newDen
		return nil
	}
}

// sAdd add the given surd √in to the current sum of surds
func (s *S32s) sAdd(in N) error {
	// reduce 1√in -> o32√in32
	if o32, i32, err := s.zSqrt(1, Z(in)); err != nil {
		return err
	} else {
		k := int(i32)
		if out, ok := s.surds[k]; !ok {
			s.surds[k] = o32
		} else {
			s.surds[k] = out + o32
		}
	}
	//fmt.Println("sAdd", in, s.surds)
	return nil
}

// sSub substract the given surd √in from the current sum of surds
func (s *S32s) sSub(in N) error {
	// reduce 1√in -> o32√in32
	if o32, i32, err := s.zSqrt(1, Z(in)); err != nil {
		return err
	} else {
		k := int(i32)
		if out, ok := s.surds[k]; !ok {
			s.surds[k] = -o32
		} else {
			s.surds[k] = out - o32
		}
	}
	return nil
}

// sNewPow2 returns a new S32s with its surd sum equals to 
// this surds sum elevated to the second power
func (s *S32s) sNewPow2() (*S32s, error) {
	surds := make(map[int]Z32, 0)
	keys := s.Keys()
	for _, k1 := range keys {
		for _, k2 := range keys {
			if k1 == k2 {
				// x√a * x√a = xxa√1 = o√i
				o, err := s.zMul(Z(s.surds[k1]), Z(s.surds[k2]), Z(k1))
				if err != nil {
					return nil, err
				}
				i := int(1)
				if out, ok := surds[i]; !ok {
					surds[i] = o
				} else {
					surds[i] = out + o
				}
			} else {
				// x√a * y√b = xy√(ab) = o√i
				x, y := Z(s.surds[k1]), Z(s.surds[k2])
				o32, i32, err := s.zSqrt(x*y, Z(k1)*Z(k2))
				if err != nil {
					return nil, err
				}
				i := int(i32)
				if out, ok := surds[i]; !ok {
					surds[i] = o32
				} else {
					surds[i] = out + o32
				}
			}
		}
	}
	return &S32s{
		Z32s:  s.Z32s,
		surds: surds,
	}, nil
}

// sNewSqrt tries to return a new S32s with value equals to the square root of
// this sum of surds. Returns nil or error when can't perform the operation.
func (s *S32s) sNewSqrt() (*S32s, error) {
	keys := s.Keys()
	if len(keys) != 2 || keys[0] != 1 {
		// suspend for sum other than format b√1 + c√d
		return nil, fmt.Errorf("Cant sqrt of keys %v", keys)
	}
	b := Z(s.surds[keys[0]])
	c := Z(s.surds[keys[1]])
	d := Z(keys[1])
	// For √(b + c√d) look if b² - c²d = x²
	// In other words, look a x such that 1√(b²-c²d) = x√1
	// Case example: √(6+2√5) = 1+√5
	if x, r, err := s.zSqrt(1, b*b - c*c*d); err != nil{
		return nil, err
	} else if r != 1 {
		// cannot denest
		return nil, fmt.Errorf("Can sqrt b*b - c*c*d x=%d r=%d", x, r)
	} else {
		// √(b + c√d) = (√(2b+2x) + √(2b-2x))/2
		// √(b - c√d) = (√(2b+2x) - √(2b-2x))/2
		if o1, i1, err := s.zSqrt(1, 2*(b + Z(x))); err != nil {
			return nil, err
		} else
		if o2, i2, err := s.zSqrt(1, 2*(b - Z(x))); err != nil {
			return nil, err
		} else {
			// o1√i1 = √(b+x)
			// o2√i2 = √(b-x)
			if o1 % 2 == 0 && o2 % 2 == 0 {
				// numerators are all even. divide them by 2, left denominator as 1.
				o1 /= 2
				o2 /= 2
			} else {
				//den = 2
				// numerators have some odd number, set denominator as 2.
				// Here we cannot return a denominator, we abort.
				return nil, nil
			}
			if c < 0 {
				// correct second numerator
				//o2 = -o2
				o1 = -o1
			}
			surds := make(map[int]Z32, 0)
			if i1 == +1 && i2 != +1 {
				// √(b+x) is integer, √(b-x) is not
				// return (o1 + o2√i2)
				surds[int(1)] = o1
				surds[int(i2)] = o2
			} else
			if i1 != +1 && i2 == +1 {
				// √(b+x) is not integer, √(b-x) is.
				// return ( o2 + o1√i1)
				surds[int(1)] = o2
				surds[int(i1)] = o1
			} else {
				// return (o1√i1 + o2√i2)
				surds[int(i1)] = o1
				surds[int(i2)] = o2
			}
			return &S32s{
				Z32s:  s.Z32s,
				surds: surds,
			}, nil
		}
	}
}

func (s *S32s) Keys() []int {
	keys := make([]int, len(s.surds))
	i := 0
	for k := range s.surds {
		keys[i] = int(k)
		i++
	}
	sort.Ints(keys[:])
	return keys
}

func (s *S32s) Tex() string {
	return s.string(true)
}

func (s *S32s) String() string {
	return s.string(false)
}

func (s *S32s) string(tex bool) string {
	var sb strings.Builder
	if s.den > 1 {
		if tex {
			sb.WriteString("\\frac{")
		} else {
			sb.WriteString("(")
		}
	}
	for pos, key := range s.Keys() {
		out := s.surds[key]
		if pos == 0 {
			if out == 1 {
				// do nothing
			} else if out == -1 {
				sb.WriteString("-")
			} else if out != 0 {
				sb.WriteString(fmt.Sprintf("%d", out))
			}
		} else {
			if out == 1 {
				sb.WriteString("+")
			} else if out == -1 {
				sb.WriteString("-")
			} else {
				sb.WriteString(fmt.Sprintf("%+d", out))
			}
		}
		if key != 1 {
			if tex {
				sb.WriteString(fmt.Sprintf("\\sqrt{%d}", key))
			} else {
				sb.WriteString(fmt.Sprintf("√%d", key))
			}
		} else if out == 1 || out == -1 {
			sb.WriteString("1")
		}
	}
	if s.den > 1 {
		if tex {
			sb.WriteString(fmt.Sprintf("}{%d}", s.den))
		} else {
			sb.WriteString(fmt.Sprintf(")/%d", s.den))
		}
	}

	if s := sb.String(); s != "" {
		return s
	}
	return "0"
}

func (s *S32s) sCmp(t *S32s) (greater, equal bool) {
	return false, false 
	// a√b + c√d > √N
}

func (s *S32s) sFloorCeil() (floor, ceil Z32, err error) {
	for in, out := range s.surds {
		if in == 1 {
			floor += out
			ceil += out
		} else if f, c, e := s.nSqrtFloorCeil(N(out)*N(out)*N(in)); err != nil {
			err = e
			return
		} else {
			floor += Z32(f)
			ceil += Z32(c)
		}
	}
	return
}

// sFloorCeil2 first raise this surbs to pow2, then calculate floor and ceil
// which then both are squared
func (s *S32s) sFloorCeil2() (floor, ceil Z32, err error) {

	if t, err := s.sNewPow2(); err != nil {
		return 0, 0, err
	} else if floor1, ceil1, err := t.sFloorCeil(); err != nil {
		return 0, 0, err
	} else if floor, _, err := s.nSqrtFloorCeil(N(floor1)); err != nil {
		return 0, 0, err
	} else if _, ceil, err := s.nSqrtFloorCeil(N(ceil1)); err != nil {
		return 0, 0, err
	} else {
		return Z32(floor), Z32(ceil), nil
	}
}


func (s *S32s) sFloat() float64 {
	f := float64(0)
	for in, out := range s.surds {
		if in == 1 {
			f += float64(out)
		} else {
			f += float64(out) * math.Sqrt(float64(in))
		}
	}
	return f
}


