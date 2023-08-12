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

func (s *Str) par(par bool, f func(*Str)) {
	if par { s.WriteString("(") }
	f(s)
	if par { s.WriteString(")") }
	
}

func (s *Str) neg() {
	s.WriteString("-")
}

func (s *Str) pos() {
	s.WriteString("+")	
}

// zS prints a integer without forced positive sign
func (s *Str) z(z Z) {
	s.WriteString(fmt.Sprintf("%d", z))
}

// zS prints a integer with forced positive sign
func (s *Str) zS(z Z) {
	s.WriteString(fmt.Sprintf("%+d", z))             
}

func (s *Str) sqrt(z Z) {
	s.WriteString(fmt.Sprintf("√%d", z))
}

func (s *Str) sqrtP(f func(s *Str)) {
	s.WriteString("√(")     
	f(s)
	s.WriteString(")")
}

func (s *Str) over(n N) {
	s.WriteString(fmt.Sprintf("/%d", n))
}
