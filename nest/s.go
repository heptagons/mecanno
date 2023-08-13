package nest

import (
	"fmt"
	"strings"
)

// Str extends strings.Builder with method useful
// for nested algebraic numbers
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

// neg append negative sign "-"
func (s *Str) neg() {
	s.WriteString("-")
}

// pos append positive sign "+"
func (s *Str) pos() {
	s.WriteString("+")	
}

// zS append an integer without forced positive sign
func (s *Str) z(z Z) {
	s.WriteString(fmt.Sprintf("%d", z))
}

// zS append an integer with forced positive sign
func (s *Str) zS(z Z) {
	s.WriteString(fmt.Sprintf("%+d", z))             
}

// sqrt append a integer preceded by root symbol "√".
func (s *Str) sqrt(z Z) {
	s.WriteString(fmt.Sprintf("√%d", z))
}

// sqrtP append a root symbol "√" and then append
// parentesis "(" the calls given function f to
// end appending parentesis ")".
func (s *Str) sqrtP(f func(s *Str)) {
	s.WriteString("√(")     
	f(s)
	s.WriteString(")")
}

// over append a number preceded by bar symbol "/".
func (s *Str) over(n N) {
	s.WriteString(fmt.Sprintf("/%d", n))
}
