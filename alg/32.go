package alg

import (
	"fmt"
	//"strings"
)





type I32 struct {
	s bool // sign: true means negative
	n N32  // positive value
}

type AI32 struct {
	o *I32  // outside
	i *I32  // inside
	e *AI32 // inside extension
}

//type AQ32 struct {
//	nums []*AI32 // numerators
//	den  N32     // denominator
//}











func (r *AI32) outVal() Z {
	return r.o.val()
}

func (r *AI32) outSet(out *I32) {
	r.o = out.clone()
}

func (r *AI32) outValPow2() (*I32, bool) {
	return r.o.mul(r.o)
}

func (r *AI32) inVal() Z {
	return r.i.val()
}

func (a *AI32) isZero() bool {
	if a == nil || a.o == nil || a.o.n == 0 {
		return true
	} else if a.e == nil {
		if a.i == nil || a.i.n == 0 {
			return true
		}
	}
	return false
}

/*func (a *AI32) Str(s *Str) { 
	if a.isZero() {
		// For 0√x return +0
		s.Zero()
	} else if a.e.isZero() {
		if a.i == nil || a.i.n == 0 {
			// For ±x√0 return +0
			s.Zero()
		} else if a.i.n == 1 && a.i.s == false {
			// For ±x√+1 return ±x
			a.o.Str(s)
		} else if a.i.s {
			// For ±x√-y return ±xi√y
			a.o.Str(s)
			s.WriteString("i")
			if a.i.n > 1 {
				s.WriteString(fmt.Sprintf("√%d", a.i.n))
			}
		} else {
			// For ±x√+y return ±xi√y
			a.o.Str(s)
			s.WriteString(fmt.Sprintf("√%d", a.i.n))
		}
	} else {
		// For ±x√±y ± e return ±xi√(±y RECURSE )
		a.o.Str(s)
		s.WriteString("√(")
		// internal
		if a.i.isZero() {
			//s.WriteString("0")
		} else {
			if a.i.s {
				s.WriteString("-")	
			}
			s.WriteString(fmt.Sprintf("%d", a.i.n))
		}
		// extension
		a.e.Str(s) // recurse
		s.WriteString(")")
	}
}*/

// WriteString appends to given buffer very SIMPLE format:
// For nil, out or in zero appends "+0"
// For n > 0 always appends +n or -n including N=1
// For in > 1 appends √ and then in (always positive)
//func (a *AI32) Str(s *Str) {
//	s.AI32(a)
//}
/*
func (r *AI32) String() string {
	s := NewStr()
	r.Str(s)
	return s.String()
}*/











// newI32 returns a 32 bit integer
// for z = 0 returns nil, false
// for overflow return nil, true
// for 0 < z <= N32_MAX return positive i32, false
// for 0 < -z <= N32_MAX return negative i32, false
func newI32(z Z) (i *I32, overflow bool) {
	if z == 0 {
		return nil, false // zero
	}
	if z > 0 {
		if N(z) > N32_MAX {
			return nil, true // overflow
		}
		return &I32{ s:false, n:N32(z) }, false // positive
	}
	if N(-z) > N32_MAX {
		return nil, true // overflow
	}
	return &I32{ s:true, n:N32(-z) }, false // negative
}

func (i *I32) isZero() bool {
	if i == nil || i.n == 0 {
		return true
	}
	return false
}

func (i *I32) clone() *I32 {
	if i == nil {
		return nil
	}
	return &I32 {
		s: i.s,
		n: i.n,
	}
}

func (i *I32) add(j *I32) (*I32, bool) {
	return newI32(i.val() + j.val())
}

func (i *I32) addN(j N32) (*I32, bool) {
	return newI32(i.val() + Z(j))
}

func (i *I32) mul(j *I32) (*I32, bool) {
	return newI32(i.val() * j.val())
}

func (i *I32) mulN(j N32) (*I32, bool) {
	return newI32(i.val() * Z(j))
}


func (i *I32) val() Z {
	if i == nil || i.n == 0 {
		return 0
	}
	if i.s {
		return -Z(i.n)
	}
	return +Z(i.n)
}

func (i *I32) Str(s *Str) {
	if i == nil {
		s.WriteString("+0")
	} else if i.s {
		s.WriteString(fmt.Sprintf("-%d", i.n))
	} else {
		s.WriteString(fmt.Sprintf("+%d", i.n))
	}
}

func (i *I32) String() string {
	s := &Str{}
	i.Str(s)
	return s.String()
}


