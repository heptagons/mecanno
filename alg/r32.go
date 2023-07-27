package alg

import (
	"fmt"
	"strings"
)

type R32 struct {
	out *I32 // external 
	in  N32  // internal
}

func NewR32(nats *N32s, out Z, in uint64) *R32 {
	outPos := out
	if out < 0 {
		outPos = -out
	}
	if out32, in32, ok := nats.Sqrt(uint64(outPos), in); !ok {
		return nil // reject overflows
	} else {
		var o *I32
		if out < 0 { // fast negative -out√in
			o = newI32minus(out32)
		} else { // fast positive +out√in
			o = newI32plus(out32)
		}
		return &R32{ // +out√in
			out: o,
			in:  in32,
		}
	}
}

// WriteString appends to given buffer very SIMPLE format:
// For nil, out or in zero appends "+0"
// For n > 0 always appends +n or -n including N=1
// For in > 1 appends √ and then in (always positive)
func (r *R32) WriteString(sb *strings.Builder) {
	if r == nil || r.out == nil || r.out.n == 0 || r.in == 0 {
		sb.WriteString("+0")
	} else {
		if r.out.s  {
			sb.WriteString("-")
		} else {
			sb.WriteString("+")
		}
		sb.WriteString(fmt.Sprintf("%d", r.out.n))
		if r.in > 1 {
			sb.WriteString("√")	
			sb.WriteString(fmt.Sprintf(r.in))
		}
	}
}
