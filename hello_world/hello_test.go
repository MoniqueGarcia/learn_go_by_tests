package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
	
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("say 'Hello, world' when an empty string is supplied", func(t* testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("spanish as language", func(t *testing.T) {
		got := Hello("Monique", "Spanish")
		want := "Hola, Monique"

		assertCorrectMessage(t, got, want)
	})

	t.Run("french as language", func(t *testing.T) {
		got := Hello("Monique", "French")
		want := "Bonjour, Monique"

		assertCorrectMessage(t, got, want)
	})

	t.Run("portuguese as language", func(t *testing.T) {
		got := Hello("Monique", "Portuguese")
		want := "Olá, Monique"

		assertCorrectMessage(t, got, want)
	})

}