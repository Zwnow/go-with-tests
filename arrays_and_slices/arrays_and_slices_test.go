package arraysandslices

import (
    "reflect"
    "testing"
)

func TestSum(t *testing.T) {
    numbers := []int{1, 2, 3, 4, 5}

    got := Sum(numbers)
    want := 15

    if got != want {
        t.Errorf("got %d want %d given, %v", got, want, numbers)
    }
}

func TestSumAllTails(t *testing.T) {
    got := SumAllTails([]int{},[]int{1,2,3},[]int{0,9},[]int{3})
    want := []int{0,5,9,0}

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}
