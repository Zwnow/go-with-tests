package helloworld

import "testing"

func TestHello(t *testing.T) {
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Sven", "")
        want := "Hello, Sven!"
        assertCorrectMessage(t, got, want)
    })
    t.Run("say 'Hello, world!' when an empty string is supplied", func(t *testing.T) {
        got := Hello("", "")
        want := "Hello, world!"
        assertCorrectMessage(t, got, want)
    })
    t.Run("say 'Moin, Sven!' when language 'de' was provided", func(t *testing.T) {
        got := Hello("Sven","de")
        want := "Moin, Sven!"
        assertCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
