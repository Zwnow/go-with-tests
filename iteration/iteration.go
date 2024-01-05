package iteration

func Repeat(in string, iterations ...int) (result string) {
    switch len(iterations) {
        case 0: {
            for i := 0; i < 5; i++ {
                result += in
            }
            break
        }
        default: {
            for i := 0; i < iterations[0]; i++ {
                result += in
            }
            break
        }
    }
    return
}
