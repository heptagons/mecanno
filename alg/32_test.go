package alg

import (
	"fmt"
	"testing"
)

func Test32(t *testing.T) {

	i0, _ := newI32(0)
	i1, _ := newI32(1)
	i2, _ := newI32(2)
	i4, _ := newI32(4)

	// AI32 (not irreducible)
	for _, s := range []struct{ exp string; ai *AI32 } {
		{ "+0", &AI32{} },
		{ "+0", &AI32{ o:i0 }},
		{ "+0", &AI32{ o:i1 }},
		{ "+1", &AI32{ o:i1, i:i1 }}, // 1√1
		{ "+2", &AI32{ o:i2, i:i1 }}, // 2√1
		
		{ "+1√2", &AI32{ o:i1, i:i2 }}, // 1√2
		{ "+1√4", &AI32{ o:i1, i:i4 }}, // 1√4

		{ "+1√4", &AI32{o:i1, i:i4, e:&AI32{} }},
		{ "+1√4", &AI32{o:i1, i:i4, e:&AI32{ o:i0 } }}, // 1√(4+0)
		{ "+1√4", &AI32{o:i1, i:i4, e:&AI32{ o:i1 } }}, // 1√(4+1√0)
		
		{ "+1√(4+1)",   &AI32{o:i1, i:i4, e:&AI32{ o:i1, i:i1 } }}, // 1√(4+1√1)
		{ "+1√(4+1√2)", &AI32{o:i1, i:i4, e:&AI32{ o:i1, i:i2 } }}, // 1√(4+1√2)
		{ "+1√(4+1√4)", &AI32{o:i1, i:i4, e:&AI32{ o:i1, i:i4 } }}, // 1√(4+1√4)
	} {
		if got := s.ai.String(); got != s.exp {
			t.Fatalf("AI32 got=%s exp=%s", got, s.exp)
		}
	}

	//t.Skip()


	r := NewAI32s()

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
			t.Fatalf("primes pos=%d got=%d exp=%d", s.pos, got, exp)
		}
	}


	// roi
	for _, s := range []struct{ o, i Z; exp string; overflow bool } {
		{ o: 0, i: 0, exp:"+0" },
		{ o: 0, i: 1, exp:"+0" },
		{ o: 0, i:-1, exp:"+0" },
		{ o:-1, i: 0, exp:"+0" },
		{ o: 1, i: 0, exp:"+0" },
		{ o: 1, i: 1, exp:"+1" },
		{ o: 1, i: 2, exp:"+1√2" },
		{ o: 1, i: 4, exp:"+2" },
		{ o: 2, i: 2, exp:"+2√2" },
		{ o: 3, i: 3, exp:"+3√3" },
		{ o: 4, i: 4, exp:"+8" },
		{ o:-4, i:-4, exp:"-8i" },
		{ o:-3, i:-3, exp:"-3i√3" },
		{ o:-2, i:-2, exp:"-2i√2" },
		{ o:-1, i:-1, exp:"-1i" },

		{ o:1, i:3*5*7*11*13*17*19*23*29, exp:"+1√3234846615" },
		{ o:1, i:25*9*4*4*3*25*9*4*7,     exp:"+1800√21" },
		{ o:9, i:3*3*3*3*5*3*3*3*3*3,     exp:"+729√15"  },

		{ o:1,            i:3*5*7*11*13*17*19*23*29*31, overflow:true },
		{ o:Z(N32_MAX+1), i:1,                          overflow:true },
		{ o:1,            i:Z(N32_MAX+2),               overflow:true },/**/
	} {
		if ai, overflow := r.roi(s.o , s.i); overflow {
			if !s.overflow {
				t.Fatalf("roi overflow expected")
			}
		} else if got := ai.String(); got != s.exp {
			t.Fatalf("roi got=%s exp=%s", got, s.exp)
		}
	}

	// roie with irreducible extra in e.i = +17
	for _, s := range []struct{ o, inA, inB Z; exp string } {

		{ o:0, inA:0,  inB: 0, exp:"+0" },
		{ o:1, inA:0,  inB: 0, exp:"+0" },
		{ o:1, inA:1,  inB: 0, exp:"+1" },
		{ o:1, inA:2,  inB: 0, exp:"+1√2" },
		{ o:1, inA:3,  inB: 0, exp:"+1√3" },
		{ o:1, inA:4,  inB: 0, exp:"+2" },
		{ o:1, inA:1,  inB: 1, exp:"+1√(1+1√17)" },

		{ o:4, inA: 4, inB: 4, exp:"+8√(1+1√17)" },
		{ o:8, inA: 8, inB: 8, exp:"+16√(2+2√17)" },
		{ o:5, inA:12, inB:56, exp:"+10√(3+14√17)" },

		{ o:1, inA:25*9*4*4*3, inB:25*9*4*7,  exp:"+30√(12+7√17)" },
		{ o:9, inA:3*3*3*3*5,  inB:3*3*3*3*3, exp:"+81√(5+3√17)"  },

		{ o:-4, inA:-4,  inB: -4, exp:"-8√(-1-1√17)" },
	} {
		if o, i, eo, ok := r.roie(s.o, s.inA, s.inB); !ok {
			t.Fatalf("unexpected infinite")
		} else {
			ext, _ := r.AI(eo, +17, nil)
			alg, _ := r.AI(o.val(), i.val(), ext)
			if got := alg.String(); got != s.exp {
				t.Fatalf("reduceExtra got=%s exp=%s", got, s.exp)
			}
		}
	}

	// A1 ±c√±d
	for pos, s := range []struct { o, a, b Z; e string; overflow bool } {
		{ o: 1, a: 0,  b: 0,  e: "+0" },
		{ o: 1, a: 0,  b: 1,  e: "+0" },
		{ o: 1, a: 1,  b: 0,  e: "+0" },
		{ o: 1, a: 1,  b: 1,  e: "+1" },
		{ o: 1, a: 2,  b: 1,  e: "+1√2" },
		{ o: 1, a: 3,  b: 1,  e: "+1√3" },
		{ o:-1, a: 3,  b: 1,  e: "-1√3" },
		{ o: 1, a: 4,  b: 1,  e: "+2" },

		{ o: 1, a:11*11, b: 10*11, e: "+11√110" },
		{ o: 1, a:1024,  b:   512, e: "+512√2" },
		{ o: 1, a:3*3*3, b: 7*7*7, e: "+21√21" },
		{ o: 1, a:3*3*5, b: 5*7*7, e: "+105" },
		{ o:-1, a:12345, b: 12345, e: "-12345" },

		{ o:1, a:  0xfffff, b:  0xfffff, e:fmt.Sprintf("+%d", 0xfffff)   },
		{ o:1, a: 0xffffff, b: 0xffffff, e:fmt.Sprintf("+%d", 0xffffff)  },
		{ o:1, a:0x1000000, b:0x1000000, e:fmt.Sprintf("+%d", 0x1000000) },
		{ o:1, a:0x1ffffff, b:0x1ffffff, e:fmt.Sprintf("+%d", 0x1ffffff) },

		{ o:1, a:0xffffffff, b:         1, e:"+1√4294967295" }, // max uint32 is prime
		{ o:1, a:0xffffffff, b:       0xf, e:"+15√286331153" }, // product ok
		{ o:1, a:0xffffffff, b:      0xff, e:"+255√16843009" }, // product ok
		{ o:1, a:0xffffffff, b:     0xfff, overflow: true    }, // prime overflow
		{ o:1, a:0xffffffff, b:    0xffff, e:"+65535√65537"  }, // product ok
		{ o:1, a:0xffffffff, b:   0xfffff, overflow: true    }, // overflow
		{ o:1, a:0xffffffff, b:  0xffffff, overflow: true    }, // overflow
		{ o:1, a:0xffffffff, b: 0xfffffff, overflow: true    }, // overflow
		{ o:1, a:0xffffffff, b:0xffffffff, overflow: true    }, // overflow

		{ o:+1, a:1, b:-1, e: "+1i"   }, // imaginary
		{ o:+2, a:1, b:-1, e: "+2i"   },
		{ o:+2, a:1, b:-2, e: "+2i√2" },
		{ o:+2, a:1, b:-4, e: "+4i"   },
		{ o:-1, a:1, b:-1, e: "-1i"   },
		{ o:-2, a:1, b:-1, e: "-2i"   },
		{ o:-2, a:1, b:-2, e: "-2i√2" },
		{ o:-2, a:1, b:-4, e: "-4i"   },
	} {
		in := s.a * s.b
		if ai, overflow := r.AI(s.o, in, nil); overflow  {
			if !s.overflow {
				t.Fatalf("Reduce1 pos=%d a=%d, b=%d overflow exp=true", pos, s.a, s.b)
			}
		} else if got := ai.String(); got != s.e {
			t.Fatalf("Reduce1 pos=%d a=%d, b=%d got:%s exp=%s", pos, s.a, s.b, got, s.e)
		}
	}

	// A2 ±e√(±f±g√±h)
	for _, s := range[]struct { o, i, eo, ei Z; exp string; overflow bool } {

		{ o: Z(N32_MAX+1), i:1,              overflow:true },
		{ o: 1,            i:Z(N32_MAX+2),   overflow:true },

		{ o: 1, i:1, eo:Z(N32_MAX+1), ei:1,            overflow:true },
		{ o: 1, i:1, eo:1,            ei:Z(N32_MAX+2), overflow:true },

		{                       exp:"+0" }, // +0√(0+0√0) = +0
		{ o:1,                  exp:"+0" }, // +1√(0+0√0) = +1√(0+0) = +1√0 = +0
		{ o:1, i:1,             exp:"+1" }, // +1√(1+0√0) = +1√(1+0) = +1√1 = +1
		{ o:1, i:1, eo:1,       exp:"+1" }, // +1√(1+1√0) = +1√(1+0) = +1√1 = +1
		{ o:1, i:1, eo:1, ei:9, exp:"+2" }, // +1√(1+1√9) = +1√(1+3) = +1√4 = +2

		{ o: 1, i: 1, eo: 1, ei:-2,  exp:"+1√(1+1i√2)" }, // +1√(1+1√-2) = +1√(1+2i)
		{ o: 1, i: 1, eo: 1, ei:-1,  exp:"+1√(1+1i)"   }, // +1√(1+1√-1) = +1√(1+1i)
		{ o: 1, i: 1, eo: 1, ei: 1,  exp:"+1√2"        }, // +1√(1+1√1) = +1√(1+1) = +1√2
		{ o: 1, i: 1, eo: 1, ei: 2,  exp:"+1√(1+1√2)"  }, 
		{ o: 1, i: 1, eo: 1, ei: 3,  exp:"+1√(1+1√3)"  }, 
		{ o: 1, i: 1, eo: 1, ei: 4,  exp:"+1√3"        }, // +1√(1+1√4) = +1√(1+2) = +1√3
		{ o: 1, i: 1, eo: 1, ei: 5,  exp:"+1√(1+1√5)"  }, 

		{ o: 1, i: 1, eo: 2, ei:-1,  exp:"+1√(1+2i)"   }, // +1√(1+1√-1) = +1√(1+1i)
		{ o: 1, i: 1, eo: 2, ei: 1,  exp:"+1√3"        }, // +1√(1+2√1) = +1√(1+2) = +1√3
		{ o: 1, i: 1, eo: 2, ei: 2,  exp:"+1√(1+2√2)"  }, // +1√(1+2√2)
		{ o: 1, i: 1, eo: 2, ei: 3,  exp:"+1√(1+2√3)"  }, 
		{ o: 1, i: 1, eo: 2, ei: 4,  exp:"+1√5"        }, // +1√(1+2√4) = +1√(1+4) = +1√5
		{ o: 1, i: 1, eo: 2, ei: 5,  exp:"+1√(1+2√5)"  }, 

		{ o: 1, i: 1, eo: 3, ei:-1,  exp:"+1√(1+3i)"   }, // +1√(1+1√-1) = +1√(1+1i)
		{ o: 1, i: 1, eo: 3, ei: 1,  exp:"+2"          }, // +1√(1+3√1) = +1√(1+3) = +1√4 = +2
		{ o: 1, i: 1, eo: 3, ei: 2,  exp:"+1√(1+3√2)"  }, // +1√(1+3√2)
		{ o: 1, i: 1, eo: 3, ei: 3,  exp:"+1√(1+3√3)"  }, 
		{ o: 1, i: 1, eo: 3, ei: 4,  exp:"+1√7"        }, // +1√(1+3√4) = +1√(1+6) = +1√7
		{ o: 1, i: 1, eo: 3, ei: 5,  exp:"+1√(1+3√5)"  }, 

		{ o: 2, i: 3, eo: 4, ei:5,  exp:"+2√(3+4√5)" },
		{ o:10, i:20, eo:30, ei:40, exp:"+20√(5+15√10)" },

		{ o: 1, i: 3, eo: 1, ei:1,  exp:"+2"   }, // +1√(3+1√1) = +1√(3+1) = +1√4 = +2
		{ o: 1, i: 4, eo: 1, ei:1,  exp:"+1√5" }, // +1√(4+1√1) = +1√(4+1) = +1√5
		
		{ o:-3, i:-5, eo:-7, ei:-11,  exp:"-3√(-5-7i√11)" },

		{ o: 1, i: 1, eo:-2, ei: 1,  exp:"+1i" }, // +1√(1-2√1) = +1√(1-2) = +1i
		{ o: 1, i: 1, eo:-1, ei: 1,  exp:"+0" }, // +1√(1-1√1) = +1√(1-1) = 0
		{ o: 1, i:-1, eo: 1, ei: 1,  exp:"+0" }, // +1√(-1+1√1) = +1√(-1+1) = 0


		//{ o:1, i:0, eo:1, ei:1, exp:"+1"  }, // +1√(0+1√1) = +1√(0+1) = +1√1 = +1
		//{ o:2, i:0, eo:2, ei:1, exp:"..."  }, // +1√(0+1√2) = +1√(0+1√2) 2√√2
		//{ o:1, i:0, eo:1, ei:2, exp:"..."  }, // +1√(0+1√2) = +1√(0+1√2) = TODO!!!




	} {
		ext, o1 := r.AI(s.eo, s.ei, nil)
		ai, o2 := r.AI(s.o, s.i, ext)
		if o1 || o2 {
			if !s.overflow {
				t.Fatalf("A2 overflow expected for %s", ai)
			}
		} else if got := ai.String(); got != s.exp {
			t.Fatalf("A2 got=%s exp=%s for %d√(%d %d√%d)", got, s.exp, s.o, s.i, s.eo, s.ei)
		}
	}
}
