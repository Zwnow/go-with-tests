package helloworld

import "testing"

func TestHello(t *testing.T) {
    got := Hello("Sven")
    want := "Hello, Sven!"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
