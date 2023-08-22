package nest

import (
	"testing"
)

func TestS32s(t *testing.T) {
	factory := NewZ32s()
	s1 := NewS32s(factory)
	s1.addSurd(2)  // √2
	s1.addSurd(12) // √2 + √12 = √2 + 2√3
	s1.addSurd(8)  // √2 + 2√3 + √8 = √2 + 2√3 + 2√2 = 3√2 + 2√3
	if got, exp := s1.String(), "3√2+2√3"; got != exp {
		t.Fatalf("got %s exp %s", got, exp)
	}

	s2 := NewS32s(factory)
	s2.addSurd(2) // √2
	s2.addSurd(3) // √2+√3
	if s3, err := s2.newPow2(); err != nil {
		t.Fatalf("err %v", err)
	} else if got, exp := s3.String(), "5+2√6"; got != exp {
		t.Fatalf("got %s exp %s", got, exp)
	} else if s4, err := s3.newSqrtRoot(); err != nil {
		t.Fatalf("err %v", err)
	} else if got, exp := s4.String(), "√2+√3"; got != exp {
		t.Fatalf("got %s exp %s", got, exp)
	}

	s2.addSurd(5) // √2+√3+√5 
	if s3, err := s2.newPow2(); err != nil {
		t.Fatalf("err %v", err)
	} else if got, exp := s3.String(), "10+2√6+2√10+2√15"; got != exp {
		t.Fatalf("got %s exp %s", got, exp)
	}
	s2.addSurd(7) // √2+√3+√5+√7
	if s3, err := s2.newPow2(); err != nil {
		t.Fatalf("err %v", err)
	} else if got, exp := s3.String(), "17+2√6+2√10+2√14+2√15+2√21+2√35"; got != exp {
		t.Fatalf("got %s exp %s", got, exp)
	}
}