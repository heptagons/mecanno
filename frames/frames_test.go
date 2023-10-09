package frames

import (
	"fmt"
	"os"
	"testing"

	. "github.com/heptagons/meccano/nest"
)

func TestTrianglesSurdExt(t *testing.T) {
	tri := NewTriangles()
	surd := Z(7)
	max := N32(10)
	frames := tri.SurdExt(surd, max)
	fmt.Printf("√%d max=%d qty=%d:\n", surd, max, len(frames))
	for f, frame := range frames {
		fmt.Fprintf(os.Stdout, "% 3d) ", f+1)
		frame.WriteString(os.Stdout)
		fmt.Println()
	}
}

//       A
//      /|
//     / |
//    B  |
//   / \ |
//  /   \| 
// C     D------E
// 
// Let strips AB = BC = DB = a
// Let bar AD = b
// Let bar DE = c
// Fix angle DAE to be right
// Then 
// CD = sqrt((2a)^2 - b^2) = sqrt(d) when 2a > b
// and CE = c + sqrt(d)
func TestFrameY(t *testing.T) {
	factory := NewA32s()
	min := N32(1)
	max := N32(10)
	for a := min; a <= max; a++ {
		for b := min; b < 2*a; b++ {
			ab := 4*Z(a)*Z(a) - Z(b)*Z(b)
			if o, i, err := factory.ZSqrt(Z(1), ab); err != nil {

			} else if i != 1 {
				fmt.Printf("a=% 3d b=% 3d c=%d√%d\n", a, b, o, i)
			}
		}
	}
}

//    C-_                    a^2 + b^2 - c^2
//    |  -_           cosC = ----------------
//  a |    -_ b                   2*a*b
//    |      -_                     _
//    B---___  -_      x^2 = ( a + √n )^2 + b^2 - 2*(a + √n)*b*cosC
//    |    c ---_A                      /      a^2 + b^2 - c^2 \  _
//    |       _/           = n + c^2 + ( 2*a - ---------------  )√n
// √n |     _/                          \            a         /
//    |   _/  x                         a^2 - b^2 + c^2 _
//    | _/                 = n + c^2 + ----------------√n
//    |/                                       a
//    N
//
//
func TestFrameX(t *testing.T) {
	min := 1
	max := 500
	n := 100*5 // 10√5
	for a := min; a <= max; a++ {
		for b := min; b <= max; b++ {
			for c := min; c <= max; c++ {
				if a + b <= c || b + c <= a || c + a <= b {
					continue
				}
				if d := a*a - b*b + c*c; d > 0 && d % a == 0 {
					x1 := n + c*c
					x2 := d/a
					if x1 % x2 == 0 && x1 / x2 == 25 {
					fmt.Printf("a=% 3d b=% 3d c=% 3d x^2= %d%+d√%d\n", a, b, c, x1, x2, n)
						fmt.Println("XXX")
					}
				}
			}
		}
	}	
}