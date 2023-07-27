package alg

import (
	"fmt"
	"strings"
)

type D struct {
	ab *B
	cd *R32
}

func NewD(nats *N32s, a, b, c, d Z) *D {

	if a == 0 {
		return nil // infinite
	} else if d < 0 {
		return nil // imaginary
	}

	(&a).Reduce3(&b, &c)

	if ab := NewB(b, a); ab == nil {
		return nil // overflow
	} else if cd := NewR32(nats, c, uint64(d)); cd != nil {
		return nil // overflow
	} else {
		// after the d simplification, c was increased
		// specially when b is 0, we need to try reduce a and c
		ab.Reduce3(cd.out)
		return &D{
			ab: ab,
			cd: cd,
		}
	}
}

func (d *D) WriteString(sb *strings.Builder) {
	a := d.ab.a // denominator
	if a > 1 {
		sb.WriteString("(")
	}
	if b := d.ab.b; b != nil {
		b.WriteString(sb)
	}
	d.cd.WriteString(sb)
	if a > 1 {
		sb.WriteString(fmt.Sprintf("%d)", a))
		sb.WriteString(")")
	}
}
