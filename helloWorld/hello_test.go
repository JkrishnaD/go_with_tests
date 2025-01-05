package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jaya", " english")
		want := "Hello, Jaya english"
		errorMessage(t, got, want)

	})
	t.Run("saying 'Hello, world' when a empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		errorMessage(t, got, want)
	})
}

func errorMessage(t testing.TB, got, want string) {
	// t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
