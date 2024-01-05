package helloworld

import "fmt"

func main() {
    fmt.Println(Hello("Sven"))
}

func Hello(name string) string {
    return "Hello, " + name + "!"
}
