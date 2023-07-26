package alg

import (
	"testing"
)

func TestNat(t *testing.T) {

	nats := NewN32s()

	// primes
	if got, exp := len(nats.primes), 6542; got != exp {
		t.Fatalf("nat-primes got %d exp:%d", got, exp)
	}
	for _, s := range []struct { pos int; prime N32 } {
		{ pos:     0, prime:      2 },
		{ pos:     1, prime:      3 },
		{ pos:     2, prime:      5 },
		{ pos:     3, prime:      7 },
		{ pos:    10, prime:     31 },
		{ pos:   100, prime:    547 },
		{ pos: 1_000, prime:  7_927 },
		{ pos: 6_541, prime: 65_521 },
	} {
		if got, exp := nats.primes[s.pos], s.prime; got != exp {
			t.Fatalf("nats.primes pos=%d got=%d exp=%d", s.pos, got, exp)
		}
	}

	// sqrt32
	for pos, s := range []struct { a, b, in, out N32; ok bool }	{
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
		m := uint64(s.a) * uint64(s.b)
		if out, in, ok := nats.Sqrt32(1, m); in != s.in || out != s.out || ok != s.ok {
			t.Fatalf("nats.sqrt32 pos=%d a=%d, b=%d got in:%d,out=%d,ok=%t exp in:%d,out=%d,ok=%t",
				pos,
				s.a, s.b,
				in, out, ok,
				s.in, s.out, s.ok)
		}
	}

	// Sqrt
	for _, s := range []struct { num, den int; exp string } {
		{ num: 0, den:1, exp: "0" },
		{ num: 1, den:1, exp: "1" },
		{ num:-1, den:1, exp:  "" }, // Imaginary
		{ num: 1, den:0, exp:  "" }, // NaN
		{ num: 2, den:2, exp: "1" },
		{ num: 4, den:1, exp: "2" },
		{ num: 1, den:4, exp: "1/2" },

		{ num: 2, den: 1, exp: "(1)√(2)"    }, // sqrt(2)
		{ num: 1, den: 2, exp: "(1/2)√(2)"  }, // 1/sqrt(2)
		{ num: 3, den: 1, exp: "(1)√(3)"    },
		{ num: 5, den: 7, exp: "(1/7)√(35)" },
		{ num: 7, den: 5, exp: "(1/5)√(35)" },
		{ num: 1, den:18, exp: "(1/6)√(2)"  },
		{ num:18, den: 1, exp: "(3)√(2)"    },
	} {
		if got := NewRat(s.num, s.den).Sqrt(nats).String(); got != s.exp {
			t.Fatalf("got %s exp:%s", got, s.exp)
		}
	}

	// Mult
	for _, s := range []struct { a, b *Alg; exp string } {
		{ a:NewAlg(NewRat(1,1), 1), b:NewAlg(NewRat(-1,1), 1), exp: "-1"    },
		{ a:NewAlg(NewRat(1,1), 1), b:NewAlg(NewRat(-1,1), 0), exp:  "0"    },
		{ a:NewAlg(NewRat(1,1), 0), b:NewAlg(NewRat(-1,0), 0), exp:  ""     },
		{ a:NewAlg(NewRat(1,2), 2), b:NewAlg(NewRat( 1,2), 2), exp:  "1/2"  },
		{ a:NewAlg(NewRat(1,3), 3), b:NewAlg(NewRat( 1,3), 3), exp:  "1/3"  },
		{ a:NewAlg(NewRat(7,1), 7), b:NewAlg(NewRat( 1,1), 7), exp: "49"    },
		{ a:NewAlg(NewRat(5,3), 7), b:NewAlg(NewRat( 1,3), 7), exp: "35/9"  },

		{ a:NewAlg(NewRat(1,2), 2), b:NewAlg(NewRat( 1,3), 3), exp:  "(1/6)√(6)" },
		{ a:NewAlg(NewRat(5,3), 8), b:NewAlg(NewRat( 3,5),10), exp:  "(4)√(5)"   },
 	} {
		if got := s.a.Multiply(s.b, nats).String(); got != s.exp {
			t.Fatalf("got %s exp %s", got, s.exp)
		}
	}
}


func TestRat(t *testing.T) {

	nan   := NewRat(  1,  0) //   ""
	zero  := NewRat(  0,  1) //   0
	one   := NewRat(  1,  1) //   1
	half  := NewRat( -1, -2) //   1/2
	_1_4  := NewRat( -1,  4) //  -1/4
	_four := NewRat( -4,  1) //  -4
	_ten  := NewRat(-10,  1) // -10
	one_17 := NewRat(2*3*5*7*11*13, 2*3*5*7*11*13*17)
	for _, r := range []struct { q *Rat; exp string } {
		{ q:nan,           exp:   ""     },
		{ q:zero,          exp:   "0"    },
		{ q:NewRat(44,33), exp:   "4/3"  },
		{ q:NewRat( 5, 7), exp:   "5/7"  },
		{ q:NewRat(99, 6), exp:  "33/2"  },
		{ q:NewRat(66,33), exp:   "2"    },
		{ q:NewRat(14,21), exp:   "2/3"  },
		{ q:NewRat(20,23), exp:  "20/23" },

		{ q:NewRat(-1,10), exp:  "-1/10" },
		{ q:_ten,          exp: "-10"    },
		{ q:NewRat(-1,-1), exp:   "1"    },

		{ q:one_17,        exp:   "1/17" },
	} {
		if got := r.q.String(); got != r.exp {
			t.Fatalf("new got: %s exp: %s", got, r.exp)	
		}
	}
	// additions
	for _, r := range []struct { a, b *Rat; exp string } {
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
	// multiplications
	for _, r := range []struct { a, b *Rat; exp string } {
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
		c := r.a.Mul(r.b)
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

	nats := NewN32s()
	algs := NewAlgs(nats)
	triangles := NewTriangles(algs)
	triangles.Find(5)
	exp := []string {
		"a=1,b=1,c=1,cosA=1/2,cosB=1/2,cosC=1/2,sinA=(1/2)√(3),sinB=(1/2)√(3),sinC=(1/2)√(3)",
		"a=2,b=2,c=1,cosA=1/4,cosB=1/4,cosC=7/8,sinA=(1/4)√(15),sinB=(1/4)√(15),sinC=(1/8)√(15)",
		"a=3,b=2,c=2,cosA=-1/8,cosB=3/4,cosC=3/4,sinA=(3/8)√(7),sinB=(1/4)√(7),sinC=(1/4)√(7)",
		"a=3,b=3,c=1,cosA=1/6,cosB=1/6,cosC=17/18,sinA=(1/6)√(35),sinB=(1/6)√(35),sinC=(1/18)√(35)",
		"a=3,b=3,c=2,cosA=1/3,cosB=1/3,cosC=7/9,sinA=(2/3)√(2),sinB=(2/3)√(2),sinC=(4/9)√(2)",
		"a=4,b=3,c=2,cosA=-1/4,cosB=11/16,cosC=7/8,sinA=(1/4)√(15),sinB=(3/16)√(15),sinC=(1/8)√(15)",
		"a=4,b=3,c=3,cosA=1/9,cosB=2/3,cosC=2/3,sinA=(4/9)√(5),sinB=(1/3)√(5),sinC=(1/3)√(5)",
		"a=4,b=4,c=1,cosA=1/8,cosB=1/8,cosC=31/32,sinA=(3/8)√(7),sinB=(3/8)√(7),sinC=(3/32)√(7)",
		"a=4,b=4,c=3,cosA=3/8,cosB=3/8,cosC=23/32,sinA=(1/8)√(55),sinB=(1/8)√(55),sinC=(3/32)√(55)",
		"a=5,b=3,c=3,cosA=-7/18,cosB=5/6,cosC=5/6,sinA=(5/18)√(11),sinB=(1/6)√(11),sinC=(1/6)√(11)",
		"a=5,b=4,c=2,cosA=-5/16,cosB=13/20,cosC=37/40,sinA=(1/16)√(231),sinB=(1/20)√(231),sinC=(1/40)√(231)",
		"a=5,b=4,c=3,cosA=0,cosB=3/5,cosC=4/5,sinA=1,sinB=4/5,sinC=3/5",
		"a=5,b=4,c=4,cosA=7/32,cosB=5/8,cosC=5/8,sinA=(5/32)√(39),sinB=(1/8)√(39),sinC=(1/8)√(39)",
		"a=5,b=5,c=1,cosA=1/10,cosB=1/10,cosC=49/50,sinA=(3/10)√(11),sinB=(3/10)√(11),sinC=(3/50)√(11)",
		"a=5,b=5,c=2,cosA=1/5,cosB=1/5,cosC=23/25,sinA=(2/5)√(6),sinB=(2/5)√(6),sinC=(4/25)√(6)",
		"a=5,b=5,c=3,cosA=3/10,cosB=3/10,cosC=41/50,sinA=(1/10)√(91),sinB=(1/10)√(91),sinC=(3/50)√(91)",
		"a=5,b=5,c=4,cosA=2/5,cosB=2/5,cosC=17/25,sinA=(1/5)√(21),sinB=(1/5)√(21),sinC=(4/25)√(21)",
	}
	if got := len(triangles.list); got != len(exp) {
		t.Fatalf("triangles got:%d exp:%d", got, len(exp))
	}
	for i, exp := range exp {
		triangle := triangles.list[i]
		if got := triangle.String(); got != exp {
			t.Fatalf("got %s exp:%s", got, exp)
		}
	}

	triangles = NewTriangles(algs)
	triangles.Find(50)
	for pos, tr := range triangles.list {
		t.Logf("% 3d area %d,%d,%d: %s", pos+1, tr.a, tr.b, tr.c, tr.area)
	}
}


