package alg

import (
	"fmt"
	"strings"
)

// I is an integer of 32 bits
type I32 struct {
	s bool
	n N32
}

func newI32plus(n N32) *I32 {
	return &I32{ s:false, n:n }
}

func newI32minus(n N32) *I32 {
	return &I32{ s:true, n:n }	
}

func (i *I32) mul(n N32) Z {
	if i.s {
		return -Z(n) * Z(i.n)
	} else {
		return +Z(n) * Z(i.n)
	}
}

/*func (i *I32) String(omitone bool) string {
	if i == nil {
		return ""
	} else if i.n == 0 {
		return "0"
	} else if i.s {
		if i.n == 1 && omitone {
			return "-"
		}
		return fmt.Sprintf("-%d", i.n)
	} else {
		if i.n == 1 && omitone {
			return ""
		}
		return fmt.Sprintf("%d", i.n)
	}
}*/

func (i *I32) WriteString(sb *strings.Builder) {
	if i == nil || i.n == 0 {
		sb.WriteString("+0")
	} else {
		if i.s {
			sb.WriteString("-")
		} else {
			sb.WriteString("+")
		}
		sb.WriteString(fmt.Sprintf("%d", i.n))
	}
}







