package prime

import (
	"fmt"
	"testing"
)

func TestPrimesInRange(t *testing.T) {
	p := primesInRange(0, 1000000)
	if len(p) != 78498 {
		t.Error("Wrong number of prime numbers lower than 1M,", len(p))
	}
	if p[0] != 2 {
		t.Error("1st prime number is not 2,", p[0])
	}
	if p[1] != 3 {
		t.Error("2nd prime number is not 3,", p[1])
	}
	if p[23423] != 267391 {
		t.Error("23424th prime number is not 267391,", p[23423])
	}
	p = primesInRange(0, 100)
	if len(p) != 25 {
		t.Error("Wrong number of prime numbers lower than 100,", len(p))
	}
	p = primesInRange(1, 7)
	if p[3] != 7 {
		t.Error("If max is 7 last prime must be 7, but it is", p[3])
	}
	p = primesInRange(1, 1)
	if len(p) != 0 {
		t.Error("Edge case of 1 not correct,", len(p))
	}
	p = primesInRange(1, 2)
	if len(p) != 1 {
		t.Error("Edge case of 2 not correct,", len(p))
	}
}

func TestNumberOfDivisors(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{60, 12},
		{5, 2},
		{1, 1},
		{20, 6},
		{21213, 6},
		{95254, 8},
		{44100, 81},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := numberOfDivisors(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkPrimesInRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesInRange(0, 1000000)
	}
}
