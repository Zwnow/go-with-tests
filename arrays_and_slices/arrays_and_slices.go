package arraysandslices

func Sum(numbers []int) (sum int) {
    for _, n := range numbers {
        sum += n
    }
    return
}

func SumAllTails(numbers ...[]int) []int {
    var result []int
    for _, slice := range numbers {
        if len(slice) > 0 {
            result = append(result, Sum(slice[1:]))
        } else {
            result = append(result, 0)
        }
    }
    return result
}
