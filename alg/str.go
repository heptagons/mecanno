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

func (s *Str) Over(n N32) {
	s.WriteString(fmt.Sprintf("/%d", n))
}

func (s *Str) I32(i *I32) {
	if s == nil {
		s.WriteString("+0")
	} else if i.s {
		s.WriteString(fmt.Sprintf("-%d", i.n))
	} else {
		s.WriteString(fmt.Sprintf("+%d", i.n))
	}
}

func (s *Str) Radical32(i *I32, ext func(s *Str)) {
	if ext == nil {
		if s == nil {
			return // do nothing
		} else if i.n == 0 {
			s.Zero()
			return
		}
		if i.s {
			s.WriteString("i")
		}
		if i.n > 1 {
			s.WriteString(fmt.Sprintf("√%d",i.n))
		}
	} else {
		s.WriteString("√(")
		if i.s {
			s.WriteString(fmt.Sprintf("-%d", i.n))
		} else {
			s.WriteString(fmt.Sprintf("%d", i.n))
		}
		ext(s)
		s.WriteString(")")
	}
}

