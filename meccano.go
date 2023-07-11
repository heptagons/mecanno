package meccano

import (
	"fmt"
)

type Sols struct {
	sols [][]int
}

func (s *Sols) Add(rods ...int) {
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
	fmt.Println()
}
