package nest

import (
	"testing"
)

func TestS32s(t *testing.T) {
	factory := NewZ32s()
	s1 := NewS32s(factory)
	s1.sAdd(2)  // √2
	s1.sAdd(12) // √2 + √12 = √2 + 2√3
	s1.sAdd(8)  // √2 + 2√3 + √8 = √2 + 2√3 + 2√2 = 3√2 + 2√3
	if got, exp := s1.String(), "3√2+2√3"; got != exp {
		t.Fatalf("got %s exp %s", got, exp)
	}

	s2 := NewS32s(factory)
	s2.sAdd(2) // √2
	s2.sAdd(3) // √2+√3
	if s3, err := s2.sNewPow2(); err != nil {
		t.Fatalf("err %v", err)
	} else if got, exp := s3.String(), "5+2√6"; got != exp {
		t.Fatalf("got %s exp %s", got, exp)
	} else if s4, err := s3.sNewSqrt(); err != nil {
		t.Fatalf("err %v", err)
	} else if got, exp := s4.String(), "√2+√3"; got != exp {
		t.Fatalf("got %s exp %s", got, exp)
	}

	s2.sAdd(5) // √2+√3+√5 
	if s3, err := s2.sNewPow2(); err != nil {
		t.Fatalf("err %v", err)
	} else if got, exp := s3.String(), "10+2√6+2√10+2√15"; got != exp {
		t.Fatalf("got %s exp %s", got, exp)
	}
	s2.sAdd(7) // √2+√3+√5+√7
	if s3, err := s2.sNewPow2(); err != nil {
		t.Fatalf("err %v", err)
	} else if got, exp := s3.String(), "17+2√6+2√10+2√14+2√15+2√21+2√35"; got != exp {
		t.Fatalf("got %s exp %s", got, exp)
	}

	s3 := NewS32s(factory)
	s3.sAdd(1)
	s3.sSub(5)
	if got, exp := s3.String(), "1-√5"; got != exp {
		t.Fatalf("got %s exp %s", got, exp)
	} else if s, err := s3.sNewPow2(); err != nil {
		t.Fatalf("err %v", err)		
	} else if got, exp := s.String(), "6-2√5"; got != exp {
		t.Fatalf("got %s exp %s", got, exp)
	} else if s, err := s.sNewSqrt(); err != nil {
		t.Fatalf("err %v", err)
	} else if got, exp := s.String(), "1-√5"; got != exp {
		t.Fatalf("got %s exp %s", got, exp)
	}
}