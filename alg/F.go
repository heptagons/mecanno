package alg

import (
	"fmt"
)

type Z int64

const MaxN = Z(0xffffffff)

func gcd(a, b Z) Z {
	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}

// I is an integer of 32 bits
type I struct {
	s bool
	n N32
}

func newIplus(n N32) *I {
	return &I{ s:false, n:n }
}

func newIminus(n N32) *I {
	return &I{ s:true, n:n }	
}

func (x *I) mul(n N32) Z {
	if x.s {
		return -Z(n) * Z(x.n)
	} else {
		return +Z(n) * Z(x.n)
	}
}

func (x *I) string() string {
	if x == nil {
		return ""
	} else if x.n == 0 {
		return "0"
	} else if x.s {
		return fmt.Sprintf("-%d", x.n)
	} else {
		return fmt.Sprintf("%d", x.n)
	}
}






