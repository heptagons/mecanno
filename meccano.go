package meccano

import (
	"fmt"
)

type Sols struct {
	sols  [][]int
	chars string
}

func (s *Sols) Chars(chars string) {
	s.chars = chars
}

func (s *Sols) Add(vals ...int) {
	s.add("", vals...)
}

func (s *Sols) Add2(last string, vals ...int) {
	s.add(last, vals...)
}

func (s *Sols) Compare(expected [][]int ) error {
	if exp, got := len(expected), len(s.sols); exp != got {
		return fmt.Errorf("size expected: %d, got:%d", exp, got)
	}
	for i, exp := range expected {
		sol := s.sols[i]
		if e, g := len(exp), len(sol); e != g {
			return fmt.Errorf("Pos:%d size expected: %d, got:%d", i, e, g)		
		} else {
			for j, v := range exp {
				if v != sol[j] {
					return fmt.Errorf("Pos:%d size expected: %v, got:%v", i, exp, sol)
				}
			}
		}
	}
	return nil
}

func (s *Sols) add(last string, vals ...int) {
	if len(vals) < 0 {
		return
	}
	if s.chars == "" {
		s.chars = "abcdefhijkl"
	}
	for _, s := range s.sols {
		a := vals[0]
		if a % s[0] != 0 { 
			continue
		}
		// new a is a factor of previous a
		f := a / s[0]
		cont := false
		for r := 1; r < len(vals); r++ {
			if s[r] == 0 {
				continue
			}
			b := vals[r]
			if t := b % s[r] == 0 && b / s[r] == f; !t {
				cont = true
				break
			}
		}
		if cont {
			continue // scaled solution already found (reject)
		}
		return
	}
	// prime combination
	if s.sols == nil {
		s.sols = make([][]int, 0)
	}
	s.sols = append(s.sols, vals)
	fmt.Printf("%4d) ", len(s.sols))
	for i, r := range vals {
		fmt.Printf(" %c=%3d", s.chars[i], r)
	}
	if last != "" {
		fmt.Printf(" %s", last)
	}
	fmt.Println()
	return
}

// Gcd returns the greatest common divisor of two numbers
func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a % b)
}


