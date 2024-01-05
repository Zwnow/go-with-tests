package iteration

import (
    "testing"
)

func TestRepeat(t *testing.T) {
    t.Run("repeat without n times parameter", func (t *testing.T) {
        repeated := Repeat("a")
        expected := "aaaaa"

        if repeated != expected {
            t.Errorf("expected %q but got %q", expected, repeated)
        }
    })
    t.Run("repeat with n times parameter.", func (t *testing.T) {
        repeated := Repeat("a", 12)
        expected := "aaaaaaaaaaaa"

        if repeated != expected {
            t.Errorf("expected %q but got %q", expected, repeated)
        }
    })
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 100)
    }
}
