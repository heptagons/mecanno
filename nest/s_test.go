package nest

import (
	"testing"
)

func TestS32s(t *testing.T) {
	s := NewS32s()
	s.addN(2)  // √2
	s.addN(12) // √2 + √12 = √2 + 2√3
	s.addN(8)  // √2 + 2√3 + √8 = √2 + 2√3 + 2√2 = 3√2 + 2√3
	if got, exp := s.String(), "3√2+2√3"; got != exp {
		t.Fatalf("got %s exp %s", got, exp)
	}
}