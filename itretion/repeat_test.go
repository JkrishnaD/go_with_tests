package itretion

import "testing"

// checks is our code is working accordingly
func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaa"

	if repeated != expected {
		t.Errorf("expected %q but repeated %q", expected, repeated)
	}
}

// checks the performance of the code
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
