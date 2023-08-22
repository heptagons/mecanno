package nest

import (
	"fmt"
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
}



// NewS32s creates a new S32s factory
// with an initial sum equals to zero
func NewS32s(factory *Z32s) *S32s {
	return &S32s{
		Z32s:  factory,
		surds: make(map[int]Z32, 0),
	}
}

// sAdd add the given value √in to the current sum
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
	fmt.Println("sAdd", in, s.surds)
	return nil
}

// sSub substract the given value √in to the current sum
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


// sNewPow2 returns a new S32s with the sum equals to this sum elevated to 2nd power
func (s *S32s) sNewPow2() (*S32s, error) {
	surds := make(map[int]Z32, 0)
	keys := s.Keys()
	for _, k1 := range keys {
		for _, k2 := range keys {
			if k1 == k2 {
				k := int(1)
				// √a * √a = a√1 
				if out, ok := surds[k]; !ok {
					surds[k] = Z32(k1)
				} else {
					surds[k] = out + Z32(k1)
				}
			} else if o32, i32, err := s.zSqrt(1, Z(k1)*Z(k2)); err != nil {
				return nil, nil
			} else {
				// x√a * y√b = xy√(a*b) = o32√i32
				o32 *= s.surds[k1] 
				o32 *= s.surds[k2]
				k := int(i32)
				if out, ok := surds[k]; !ok {
					surds[k] = o32
				} else {
					surds[k] = out + o32
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
		return nil, nil
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
		return nil, err
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

func (s *S32s) String() string {
	keys := s.Keys()
	var sb strings.Builder
	for pos, key := range keys {
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
			sb.WriteString(fmt.Sprintf("√%d", key))
		} else if out == 1 || out == -1 {
			sb.WriteString("1")
		}
	}
	if s := sb.String(); s != "" {
		return s
	}
	return "0"
}

/*func (s *S32s) GreaterOrEqualZ(z Z32) (ok bool, err error) {
	// first sort these surds and compare with integer argument
	// x√a + y√b + z√c >= z*z + 0 + 0
	if 
}*/
