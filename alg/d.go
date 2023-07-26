package alg

import (
	"fmt"
)

type Root32 struct {
	out *I32 // external 
	in  N32  // internal
}

func NewRoot32(nats *N32s, out Z, in uint64) *Root32 {
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
		return &Root32{ // +out√in
			out: o,
			in:  in32,
		}
	}
}

func (r *Root32) String() string {
	if r == nil {
		return "" // overflow
	} else if r.out == nil || r.out.n == 0 || r.in == 0 {
		return "0"
	} else if r.in == 1 {
		return r.out.String(false) // Just integer
	} else {
		return fmt.Sprintf("%s√%d", r.out.String(true), r.in)
	}

}
