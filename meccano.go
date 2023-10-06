package meccano

import (
	"fmt"
)

type Sols struct {
	sols [][]int
}

func (s *Sols) add(last string, rods ...int) {
	if len(rods) < 0 {
		return
	}
	const RODS = "abcdefhijkl"
	for _, s := range s.sols {
		a := rods[0]
		if a % s[0] != 0 { 
			continue
		}
		// new a is a factor of previous a
		f := a / s[0]
		cont := false
		for r := 1; r < len(rods); r++ {
			if s[r] == 0 {
				continue
			}
			b := rods[r]
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
	// solution!
	if s.sols == nil {
		s.sols = make([][]int, 0)
	}
	s.sols = append(s.sols, rods)
	fmt.Printf("%3d ", len(s.sols))
	for i, r := range rods {
		fmt.Printf(" %c=%3d", RODS[i], r)
	}
	if last != "" {
		fmt.Printf(" %s", last)
	}
	fmt.Println()
	return
}

func (s *Sols) Add(rods ...int) {
	s.add("", rods...)
}

func (s *Sols) Add2(last string, rods ...int) {
	s.add(last, rods...)
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

// Gcd returns the greatest common divisor of two numbers
func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a % b)
}


