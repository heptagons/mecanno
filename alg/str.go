package alg

import (
	"fmt"
	"strings"
)

type Str struct {
	strings.Builder
}

func NewStr() *Str {
	return &Str{}
}

func (s *Str) Infinite() {
	s.WriteString("∞")
}

func (s *Str) Zero() {
	s.WriteString("+0")	
}

func (s *Str) Divisor() {
	s.WriteString("/")
}

func (s *Str) N32(n N32) {
	s.WriteString(fmt.Sprintf("%d", n))
}

func (s *Str) Integer32(i *I32) {
	if s == nil {
		s.WriteString("+0")
	} else if i.s {
		s.WriteString("-")
		s.N32(i.n)
	} else {
		s.WriteString("+")
		s.N32(i.n)
	}
}

func (s *Str) Radical32(i *I32) {
	if s == nil {
		return
	} else if i.n == 0 {
		s.Zero()
	} else {
		if i.s { // Imaginary
			s.WriteString("i")	
		}
		if i.n > 1 {
			s.WriteString("√")
			s.N32(i.n)
		}
	}
}
