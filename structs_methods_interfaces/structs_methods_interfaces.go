package structsmethodsinterfaces

import "math"


type Shape interface {
    Area() float64
}


type Rectangle struct {
    Height float64
    Width float64
}

func (r Rectangle) Area() float64 {
    return r.Height * r.Width
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return  math.Pow(c.Radius, 2) * math.Pi
}

func Perimeter(r Rectangle) float64 {
    return 2 * (r.Height + r.Width)
}
