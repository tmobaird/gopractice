package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()

		if got != want {
			t.Errorf("Got '%s', Wanted '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello, World when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

// You can have as many tests as you want in a go test file
// The function name must be TestXXX with the same signature
func TestHello2(t *testing.T) {
	got := Hello("Thomas")
	want := "Hello, Thomas"

	if got != want {
		t.Errorf("Got '%s', Wanted '%s'", got, want)
	}
}
