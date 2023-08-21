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
	surds map[N32]Z32 // map[in]out
}

// NewS32s creates a new S32s factory
func NewS32s() *S32s {
	return &S32s{
		Z32s:  NewZ32s(),
		surds: make(map[N32]Z32, 0),
	}
}

func (s *S32s) addN(in N) error {
	// reduce 1√in -> o32√in32
	o32, i32, err := s.zSqrt(1, Z(in))
	if err != nil {
		return err
	}
	key := N32(i32)
	if out, ok := s.surds[key]; !ok {
		s.surds[key] = o32
	} else {
		s.surds[key] = out + o32
	}
	return nil
}

func (s *S32s) String() string {
	keys := make([]int, len(s.surds))
	i := 0
	for k := range s.surds {
		keys[i] = int(k)
		i++
	}
	sort.Ints(keys[:])
	var sb strings.Builder
	for pos, key := range keys {
		out := s.surds[N32(key)]
		if pos == 0 {
			if out == 1 {
				sb.WriteString(fmt.Sprintf("√%d", key))
			} else if out == -1 {
				sb.WriteString(fmt.Sprintf("-√%d", key))
			} else if out != 0 {
				sb.WriteString(fmt.Sprintf("%d√%d", out, key))
			}
		} else {
			if out == 1 {
				sb.WriteString(fmt.Sprintf("+√%d", key))
			} else if out == -1 {
				sb.WriteString(fmt.Sprintf("-√%d", key))
			} else {
				sb.WriteString(fmt.Sprintf("%+d√%d", out, key))
			}
		}
	}
	if s := sb.String(); s != "" {
		return s
	}
	return "0"
}
