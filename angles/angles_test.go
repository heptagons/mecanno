package angles

import (
	"fmt"
	"testing"
)

func TestQ(t *testing.T) {

	nan   := NewQ(  1,  0) //   ""
	zero  := NewQ(  0,  1) //   0
	one   := NewQ(  1,  1) //   1
	half  := NewQ( -1, -2) //   1/2
	_1_4  := NewQ( -1,  4) //  -1/4
	_four := NewQ( -4,  1) //  -4
	_ten  := NewQ(-10,  1) // -10
	one_17 := NewQ(2*3*5*7*11*13, 2*3*5*7*11*13*17)
	for _, r := range []struct { q *Q; exp string } {
		{ q:nan,         exp:   ""     },
		{ q:zero,        exp:   "0"    },
		{ q:NewQ(44,33), exp:   "4/3"  },
		{ q:NewQ( 5, 7), exp:   "5/7"  },
		{ q:NewQ(99, 6), exp:  "33/2"  },
		{ q:NewQ(66,33), exp:   "2"    },
		{ q:NewQ(14,21), exp:   "2/3"  },
		{ q:NewQ(20,23), exp:  "20/23" },

		{ q:NewQ(-1,10), exp:  "-1/10" },
		{ q:_ten,        exp: "-10"    },
		{ q:NewQ(-1,-1), exp:   "1"    },

		{ q:one_17,      exp:"   1/17" },
	} {
		if got := r.q.String(); got != r.exp {
			t.Fatalf("new got: %s exp: %s", got, r.exp)	
		}
	}
	// Test Add
	for _, r := range []struct { a *Q; b *Q; exp string } {
		{ a:nan,    b:nan,    exp:   ""     },
		{ a:nan,    b:zero,   exp:   ""     },
		{ a:zero,   b:zero,   exp:   "0"    },
		{ a:one,    b:zero,   exp:   "1"    },
		{ a:one,    b:one,    exp:   "2"    },
		{ a:half,   b:half,   exp:   "1"    },
		{ a:_1_4,   b:_1_4,   exp:  "-1/2"  },
		{ a:_1_4,   b:half,   exp:   "1/4"  },
		{ a:_1_4,   b:_four,  exp: "-17/4"  },
		{ a:_four,  b:_four,  exp: "-8"     },
		{ a:_ten,   b:_1_4,   exp: "-41/4"  },
		{ a:one_17, b:one_17, exp:   "2/17" },
	} {
		c := r.a.Add(r.b)
		if got := c.String(); got != r.exp {
			t.Fatalf("add got: %s exp: %s", got, r.exp)	
		}
	}

	// Test Times
	for _, r := range []struct { a *Q; b *Q; exp string } {
		{ a:nan,    b:nan,    exp:  ""      },
		{ a:nan,    b:zero,   exp:  ""      },
		{ a:zero,   b:zero,   exp:  "0"     },
		{ a:one,    b:zero,   exp:  "0"     },
		{ a:one,    b:one,    exp:  "1"     },
		{ a:half,   b:half,   exp:  "1/4"   },
		{ a:_1_4,   b:_1_4,   exp:  "1/16"  },
		{ a:_1_4,   b:half,   exp: "-1/8"   },
		{ a:_1_4,   b:_four,  exp:  "1"     },
		{ a:_four,  b:_four,  exp: "16"     },
		{ a:_ten,   b:_1_4,   exp:  "5/2"   },
		{ a:one_17, b:one_17, exp:  "1/289" },
	} {
		c := r.a.Times(r.b)
		if got := c.String(); got != r.exp {
			t.Fatalf("times got: %s exp: %s", got, r.exp)	
		}
	}
}

func TestTriangle(t *testing.T) {
	// test errors
	for _, triangle := range []*Triangle{
		NewTriangle(0,1,1), // a = 0
		NewTriangle(1,0,1), // b = 0
		NewTriangle(1,1,0), // c = 0
		NewTriangle(1,2,1), // b > a
		NewTriangle(1,1,2), // c > b
		NewTriangle(2,1,1), // a >= c+b (area zero)
	} {
		if triangle != nil {
			t.Fatalf("exp nil got %v", t)
		}
	}
}

func TestTriangles(t *testing.T) {

	triangles := NewTriangles()

	for a := 1; a <= 5; a++ {
		for b := 1; b <= a; b++ {
			for c := 1; c <= b; c++ {
				if next := triangles.Add(a, b, c); next != nil {
					fmt.Println(len(triangles), next)
				}
			}
		}
	}
	exp := []string {
		"a=1,b=1,c=1,cosA=1/2,cosB=1/2,cosC=1/2,sin2A=3/4,sin2B=3/4,sin2C=3/4",
		"a=2,b=2,c=1,cosA=1/4,cosB=1/4,cosC=7/8,sin2A=15/16,sin2B=15/16,sin2C=15/64",
		"a=3,b=2,c=2,cosA=-1/8,cosB=3/4,cosC=3/4,sin2A=63/64,sin2B=7/16,sin2C=7/16",
		"a=3,b=3,c=1,cosA=1/6,cosB=1/6,cosC=17/18,sin2A=35/36,sin2B=35/36,sin2C=35/324",
		"a=3,b=3,c=2,cosA=1/3,cosB=1/3,cosC=7/9,sin2A=8/9,sin2B=8/9,sin2C=32/81",
		"a=4,b=3,c=2,cosA=-1/4,cosB=11/16,cosC=7/8,sin2A=15/16,sin2B=135/256,sin2C=15/64",
		"a=4,b=3,c=3,cosA=1/9,cosB=2/3,cosC=2/3,sin2A=80/81,sin2B=5/9,sin2C=5/9",
		"a=4,b=4,c=1,cosA=1/8,cosB=1/8,cosC=31/32,sin2A=63/64,sin2B=63/64,sin2C=63/1024",
		"a=4,b=4,c=3,cosA=3/8,cosB=3/8,cosC=23/32,sin2A=55/64,sin2B=55/64,sin2C=495/1024",
		"a=5,b=3,c=3,cosA=-7/18,cosB=5/6,cosC=5/6,sin2A=275/324,sin2B=11/36,sin2C=11/36",
		"a=5,b=4,c=2,cosA=-5/16,cosB=13/20,cosC=37/40,sin2A=231/256,sin2B=231/400,sin2C=231/1600",
		"a=5,b=4,c=3,cosA=0,cosB=3/5,cosC=4/5,sin2A=1,sin2B=16/25,sin2C=9/25",
		"a=5,b=4,c=4,cosA=7/32,cosB=5/8,cosC=5/8,sin2A=975/1024,sin2B=39/64,sin2C=39/64",
		"a=5,b=5,c=1,cosA=1/10,cosB=1/10,cosC=49/50,sin2A=99/100,sin2B=99/100,sin2C=99/2500",
		"a=5,b=5,c=2,cosA=1/5,cosB=1/5,cosC=23/25,sin2A=24/25,sin2B=24/25,sin2C=96/625",
		"a=5,b=5,c=3,cosA=3/10,cosB=3/10,cosC=41/50,sin2A=91/100,sin2B=91/100,sin2C=819/2500",
		"a=5,b=5,c=4,cosA=2/5,cosB=2/5,cosC=17/25,sin2A=21/25,sin2B=21/25,sin2C=336/625",
	}
	if got := len(triangles); got != len(exp) {
		t.Fatalf("triangles got:%d exp:%d", got, len(exp))
	}
	for i, exp := range exp {
		if got := triangles[i].String(); got != exp {
			t.Fatalf("got %s exp:%s", got, exp)
		}
	}
}

func TestPrimes(t *testing.T) {
	primes := PrimesList(0xffff) // max squared uint 32 bits
	if got, exp := len(primes), 6542; got != exp {
		t.Fatalf("primes-squared got %d exp:%d", got, exp)
	}
	for _, s := range []struct{ pos int; prime uint } {
		{ pos:     0, prime:      2 },
		{ pos:     1, prime:      3 },
		{ pos:     2, prime:      5 },
		{ pos:     3, prime:      7 },
		{ pos:    10, prime:     31 },
		{ pos:   100, prime:    547 },
		{ pos: 1_000, prime:  7_927 },
		{ pos: 6_541, prime: 65_521 },
	} {
		if got, exp := primes[s.pos], s.prime; got != exp {
			t.Fatalf("primes-squared pos=%d got=%d exp=%d", s.pos, got, exp)
		}
	}
}

func TestProductRoots(t *testing.T) {
	squares := NewSquares32() // max squared uint 32 bits
	for pos, s := range []struct{ a, b, in, out uint; ok bool }	{
		{ a: 0,  b: 0,  in: 0, out: 0, ok: true },
		{ a: 0,  b: 1,  in: 0, out: 0, ok: true },
		{ a: 1,  b: 0,  in: 0, out: 0, ok: true },
		{ a: 1,  b: 1,  in: 1, out: 1, ok: true },
		{ a: 2,  b: 1,  in: 2, out: 1, ok: true },
		{ a: 3,  b: 1,  in: 3, out: 1, ok: true },
		{ a: 4,  b: 1,  in: 1, out: 2, ok: true },

		{ a:    11*11, b:    10*11, in:110, out:    11, ok:true },
		{ a:     1024, b:      512, in:  2, out:   512, ok:true },
		{ a:    3*3*3, b:    7*7*7, in: 21, out:    21, ok:true },
		{ a:    3*3*5, b:    5*7*7, in:  1, out: 3*5*7, ok:true },
		{ a:    12345, b:    12345, in:  1, out: 12345, ok:true },


		{ a:  0xfffff, b:  0xfffff, in: 1, out:  0xfffff, ok:true },
		{ a: 0xffffff, b: 0xffffff, in: 1, out: 0xffffff, ok:true },
		{ a:0x1000000, b:0x1000000, in: 1, out:0x1000000, ok:true },
		{ a:0x1ffffff, b:0x1ffffff, in: 1, out:0x1ffffff, ok:true },
		
		{ a:0xffffffff, b:         1, in:0xffffffff, out:      1, ok:true  }, // max uint32 is prime
		{ a:0xffffffff, b:       0xf, in: 286331153, out:     15, ok:true  }, // product ok
		{ a:0xffffffff, b:      0xff, in:  16843009, out:    255, ok:true  }, // product ok
		{ a:0xffffffff, b:     0xfff, in:         0, out:      0, ok:false }, // prime overflow
		{ a:0xffffffff, b:    0xffff, in:     65537, out: 0xffff, ok:true  }, // product ok
		{ a:0xffffffff, b:   0xfffff, in:         0, out:      0, ok:false }, // overflow
		{ a:0xffffffff, b:  0xffffff, in:         0, out:      0, ok:false }, // overflow
		{ a:0xffffffff, b: 0xfffffff, in:         0, out:      0, ok:false }, // overflow
		{ a:0xffffffff, b:0xffffffff, in:         0, out:      0, ok:false }, // overflow
	} {
		if in, out, ok := squares.productRoot(s.a, s.b); in != s.in || out != s.out || ok != s.ok {
			t.Fatalf("productRoot pos=%d a=%d, b=%d got in:%d,out=%d,ok=%t exp in:%d,out=%d,ok=%t",
				pos,
				s.a, s.b,
				in, out, ok,
				s.in, s.out, s.ok)
		}
	}
}

func TestQRoots(t *testing.T) {
	squares := NewSquares32() // max squared uint 32 bits
	for _, s := range []struct{ num, den int; exp string } {
		{ num: 0, den:1, exp:         "0" },
		{ num: 1, den:1, exp:         "1" },
		{ num:-1, den:1, exp:          "" }, // Imaginary
		{ num: 1, den:0, exp:          "" }, // NaN
		{ num: 2, den:1, exp:"sqrt(2)(1)" }, // sqrt(2)
		{ num: 2, den:2, exp:         "1" },
		{ num: 3, den:1, exp:"sqrt(3)(1)" },
		{ num: 4, den:1, exp:         "2" },

		{ num: 5, den: 7, exp: "sqrt(35)(1/7)" },
		{ num: 7, den: 5, exp: "sqrt(35)(1/5)" },
		{ num: 1, den:18, exp: "sqrt(2)(1/6)"  },
		{ num:18, den: 1, exp: "sqrt(2)(3)" },
	} {
		if got := squares.Root(NewQ(s.num, s.den)).String(); got != s.exp {
			t.Fatalf("got %s exp:%s", got, s.exp)
		}
	}
}