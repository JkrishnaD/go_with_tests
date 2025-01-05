package itretion

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaa"

	if repeated != expected {
		t.Errorf("expected %q but repeated %q", expected, repeated)
	}
}
