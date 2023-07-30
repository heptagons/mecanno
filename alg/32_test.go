package alg

import (
	"fmt"
	"testing"
)

func Test32R(t *testing.T) {

	r := NewR32s()

	// primes
	if got, exp := len(r.primes), 6542; got != exp {
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
		if got, exp := r.primes[s.pos], s.prime; got != exp {
			t.Fatalf("n32s.primes pos=%d got=%d exp=%d", s.pos, got, exp)
		}
	}

	// reduce2(out, inA, inB N) (N32, N32, N32)
	for _, s := range []struct{ out, inA, inB Z; exp string } {

		{ out:0, inA:0,  inB: 0, exp:"+0√(+0+0√x)" },

		{ out:1, inA: 1, inB: 1, exp:"+1√(+1+1√x)" },
		{ out:4, inA: 4, inB: 4, exp:"+8√(+1+1√x)" },
		{ out:8, inA: 8, inB: 8, exp:"+16√(+2+2√x)" },
		{ out:5, inA:12, inB:56, exp:"+10√(+3+14√x)" },

		{ out:1, inA:25*9*4*4*3, inB:25*9*4*7,  exp:"+30√(+12+7√x)" },
		{ out:9, inA:3*3*3*3*5,  inB:3*3*3*3*3, exp:"+81√(+5+3√x)"  },

		{ out:-4, inA:-4,  inB: -4, exp:"-8√(-1-1√x)" },
	} {
		o, a, b, _ := r.reduce2Z(s.out, s.inA, s.inB)
		if got := fmt.Sprintf("%s√(%s%s√x)", o, a, b); got != s.exp {
			t.Fatalf("n32s.reduceIns got=%s exp=%s", got, s.exp)
		}
	}

	// ±c√±d
	for pos, s := range []struct { out, a, b Z; e string }	{
		{ out: 1, a: 0,  b: 0,  e: "+0" },
		{ out: 1, a: 0,  b: 1,  e: "+0" },
		{ out: 1, a: 1,  b: 0,  e: "+0" },
		{ out: 1, a: 1,  b: 1,  e: "+1" },
		{ out: 1, a: 2,  b: 1,  e: "+1√2" },
		{ out: 1, a: 3,  b: 1,  e: "+1√3" },
		{ out:-1, a: 3,  b: 1,  e: "-1√3" },
		{ out: 1, a: 4,  b: 1,  e: "+2" },

		{ out: 1, a:11*11, b: 10*11, e: "+11√110" },
		{ out: 1, a:1024,  b:   512, e: "+512√2" },
		{ out: 1, a:3*3*3, b: 7*7*7, e: "+21√21" },
		{ out: 1, a:3*3*5, b: 5*7*7, e: "+105" },
		{ out:-1, a:12345, b: 12345, e: "-12345" },

		{ out:1, a:  0xfffff, b:  0xfffff, e:fmt.Sprintf("+%d", 0xfffff)   },
		{ out:1, a: 0xffffff, b: 0xffffff, e:fmt.Sprintf("+%d", 0xffffff)  },
		{ out:1, a:0x1000000, b:0x1000000, e:fmt.Sprintf("+%d", 0x1000000) },
		{ out:1, a:0x1ffffff, b:0x1ffffff, e:fmt.Sprintf("+%d", 0x1ffffff) },

		{ out:1, a:0xffffffff, b:         1, e:"+1√4294967295" }, // max uint32 is prime
		{ out:1, a:0xffffffff, b:       0xf, e:"+15√286331153" }, // product ok
		{ out:1, a:0xffffffff, b:      0xff, e:"+255√16843009" }, // product ok
		{ out:1, a:0xffffffff, b:     0xfff, e:"∞"             }, // prime overflow
		{ out:1, a:0xffffffff, b:    0xffff, e:"+65535√65537"  }, // product ok
		{ out:1, a:0xffffffff, b:   0xfffff, e:"∞"             }, // overflow
		{ out:1, a:0xffffffff, b:  0xffffff, e:"∞"             }, // overflow
		{ out:1, a:0xffffffff, b: 0xfffffff, e:"∞"             }, // overflow
		{ out:1, a:0xffffffff, b:0xffffffff, e:"∞"             }, // overflow

		{ out:+1, a:1, b:-1, e: "+1i" }, // imaginaries
		{ out:+2, a:1, b:-1, e: "+2i" },
		{ out:+2, a:1, b:-2, e: "+2i√2" },
		{ out:+2, a:1, b:-4, e: "+4i"   },
		{ out:-1, a:1, b:-1, e: "-1i"   },
		{ out:-2, a:1, b:-1, e: "-2i"   },
		{ out:-2, a:1, b:-2, e: "-2i√2" },
		{ out:-2, a:1, b:-4, e: "-4i"   },
	} {
		in := s.a * s.b
		if got := r.Radical(s.out, in, nil).String(); got != s.e {
			t.Fatalf("Reduce1 pos=%d a=%d, b=%d got:%s exp=%s", pos, s.a, s.b, got, s.e)
		}
	}

	// ±e√(±f±g√±h)
	for _, s := range[]struct { out, in Z; extra *R32; exp string } {

		{ out: 2, in: 3, extra:r.Radical( 4, 5,nil), exp:"+2√(3+4√5)" },
		{ out:10, in:20, extra:r.Radical(30,40,nil), exp:"+20√(5+15√10)" },

		{ out: 1, in: 1, extra:r.Radical( 1, 1,nil), exp:"+1√(1+1)" },
		{ out: 1, in: 3, extra:r.Radical( 1, 1,nil), exp:"+1√(3+1)" }, // TODO is +2
	} {
		if got := r.Radical(s.out, s.in, s.extra).String(); got != s.exp {
			t.Fatalf("Reduce2 got=%s exp=%s", got, s.exp)
		}
	}
}
