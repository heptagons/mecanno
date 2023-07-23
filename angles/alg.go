package angles

import (
	"fmt"
)

type Alg struct {
	*Rat
	In Nat
}

func NewAlg(rat *Rat, in Nat) *Alg {
	if rat == nil {
		return nil
	}
	return &Alg{
		Rat: rat,
		In: in,
	}
}

func (s *Alg) String() string {
	if s == nil || s.Rat == nil {
		return ""
	} else if s.In == 0 {
		return "0"
	} else if s.In == 1 {
		return s.Rat.String()
	} else {
		return fmt.Sprintf("(%s)√(%d)", s.Rat, s.In)
	}
}




