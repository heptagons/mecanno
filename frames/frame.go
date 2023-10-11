package frames

import (
	"fmt"
	"io"
	. "github.com/heptagons/meccano/nest"
)

// FrameSurd is a three strips meccano frame resembling a uppercase letter A:
//
//          C-_  b
//      a  /   -_
//        /   __ A_
//       B__--     -_  e
//   d  /     c      -_
//     /               -D
//    E
// 
// The frame has exactly three strips:
// * a+d = distance node C to node E when d > 0 or distance node C to node B when d = 0.
// * b+e = distance node C to node D when e > 0 or distance node C to node A when e = 0.
// * c = distance node A to node B.
// The frame has at most five nodes: A,B,C,D,E or four: A,B,C,E or A,B,C,D.
// We store five positive integer distances:
// a = distance nodes C and B >= 1
// b = distance nodes A and C >= 1
// c = distance nodes A and B >= 1
// d = distance nodes B and E >= 0
// e = distance nodes A and D >= 0
type FrameSurd struct {
	a,b,c,d,e N32
 	cos       *A32
}

func (f *FrameSurd) WriteString(w io.Writer) {
	if f.d == 0 {
		if f.e == 0 {
			fmt.Fprintf(w, "a=%d b=%d", f.a, f.b)
		} else {
			// triangle extension with only 4 bolts not 5
			fmt.Fprintf(w, "a=%d b=%d+%d", f.a, f.b, f.e)
		}
	} else if f.e == 0 {
		fmt.Fprintf(w, "a+d=%d+%d b=%d", f.a, f.d, f.b)
	} else {
		fmt.Fprintf(w, "a+d=%d+%d b=%d+%d", f.a, f.d, f.b, f.e)
	}
	fmt.Fprintf(w, " c=%d cos=%v", f.c, f.cos.String())
}


//    C-_                           C-_ 
//    |  -_                         *  -_
//    |a   -_ b                    /|a   -_
//    |      -_                   / |      -_
//    B---___  -_             *--*--B---___  -_
//    .    c ----A             \   /       ----A
//    .       _.                \ /         _.   
// √s .     _.                   x  √s    _.    ____
//    .   _.  nest                \     _.     /   _
//    . _.                         \  _.      √x+y√z
//    N                             N         ------
//                                               w
type FrameNest struct {      
	a,b,c N32
	surd  N32
	nest  *A32
}

