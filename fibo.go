package main

import "fmt"

func main() {
    var a int
    fmt.Scan(&a)
    
    x, y := 0, 1
    
    for i := 0; i < a; i++ {
        fmt.Print(x, " ")
        x, y = y, x+y
    }
    
    fmt.Println()
}
