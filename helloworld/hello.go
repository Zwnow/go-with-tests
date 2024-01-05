package helloworld

import "fmt"

func main() {
    fmt.Println(Hello("Sven"))
}

func Hello(name string) string {
    if len(name) == 0 {
        name = "world"
    }
    return "Hello, " + name + "!"
}
