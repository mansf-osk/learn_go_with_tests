package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Name supplied", func(t *testing.T) {
		got := Hello("Oskar", "English")
		want := "Hello, Oskar"

		assertCorrectMessage(t, got, want)
	})

	t.Run("No name supplied -> say 'Hello World'.", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Antoine", "French")
		want := "Bonjour, Antoine"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T) {
		got := Hello("Elodie", "German")
		want := "Hallo, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
