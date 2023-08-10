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

// root prints y√z -y√z or √z or -√z
func (s *Str) root(y, z Z32) {
	if y == -1 {
		s.neg()
	} else if y != +1 {
		s.z(y)
	}
	s.sqrt(z)
}

// plus_root prints x+y√z or x-y√z or x+√z or x-√z
func (s *Str) plus_root(x, y, z Z32) {
	s.z(x)
	if y == -1 {
		s.neg()
	} else if y == +1 {
		s.pos()
	} else {
		s.zS(y)
	}
	s.sqrt(z)
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
func (s *Str) z(z Z32) {
	s.WriteString(fmt.Sprintf("%d", z))
}

// zS prints a integer with forced positive sign
func (s *Str) zS(z Z32) {
	s.WriteString(fmt.Sprintf("%+d", z))             
}

func (s *Str) sqrt(z Z32) {
	s.WriteString(fmt.Sprintf("√%d", z))
}

func (s *Str) sqrtP(f func(s *Str)) {
	s.WriteString("√(")     
	f(s)
	s.WriteString(")")
}

func (s *Str) over(n N32) {
	s.WriteString(fmt.Sprintf("/%d", n))
}
