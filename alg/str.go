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

func (s *Str) Sign(sign bool) {
	if sign {
		s.WriteString("-")
	} else {
		s.WriteString("+")
	}
}

func (s *Str) Root() {
	s.WriteString("√")	
}
