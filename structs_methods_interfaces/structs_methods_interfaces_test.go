package structsmethodsinterfaces

import "testing"
func TestPerimeter(t *testing.T) {
    got := Perimeter(Rectangle{10.0, 10.0})
    want := 40.0
    
    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}

func TestArea(t *testing.T) {
    t.Run("Test area for rectangle", func(t *testing.T) {

        rectangle := Rectangle{10.0, 10.0}
        want := 100.0
        checkShape(t, rectangle, want)
    })

    t.Run("Test area for rectangle", func(t *testing.T) {
        circle := Circle{10.0}
        want := 314.1592653589793
        checkShape(t, circle, want)
    })
}

func checkShape(t testing.TB, shape Shape, want float64) {
    t.Helper()
    got := shape.Area()
    if got != want {
        t.Errorf("got %g want %g", got, want)
    }
}
