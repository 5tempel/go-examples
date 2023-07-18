package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func TestHelloEmptyString(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloEmptyString("Chris") //calling function for testing
		want := "Hello, Chris"           //expected value

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := HelloEmptyString("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func TestHelloLanguage(t *testing.T) {
	t.Run("in Spanish", func(t *testing.T) {
		got := HelloLanguage("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Polish", func(t *testing.T) {
		got := HelloLanguage("Zosia", "Polish")
		want := "Czesc, Zosia"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in English if empty string", func(t *testing.T) {
		got := HelloLanguage("Zosia", "")
		want := "Hello, Zosia"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in English if unknow language", func(t *testing.T) {
		got := HelloLanguage("Zosia", "w")
		want := "Hello, Zosia"
		assertCorrectMessage(t, got, want)
	})
}

func TestHelloLSwitch(t *testing.T) {
	t.Run("in Spanish", func(t *testing.T) {
		got := HelloLSwitch("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Polish", func(t *testing.T) {
		got := HelloLSwitch("Zosia", "Polish")
		want := "Czesc, Zosia"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in English if empty string", func(t *testing.T) {
		got := HelloLSwitch("Zosia", "")
		want := "Hello, Zosia"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in English if unknow language", func(t *testing.T) {
		got := HelloLSwitch("Zosia", "w")
		want := "Hello, Zosia"
		assertCorrectMessage(t, got, want)
	})
}
