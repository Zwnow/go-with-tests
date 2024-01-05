package helloworld

import "fmt"

func main() {
    fmt.Println(Hello("Sven",""))
}

func Hello(name string, language string) string {
    if len(name) == 0 {
        name = "world"
    }
    switch language {
    case "de":
        return "Moin, " + name + "!"
    default:
        return "Hello, " + name + "!"
    }
}
