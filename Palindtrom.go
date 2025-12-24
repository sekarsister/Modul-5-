package main

import "fmt"

func main() {
    var a, c int
    var cond bool
    
    fmt.Scan(&a)
    
    reversed := 0
    c = a
    
    for a != 0 {
        digit := a % 10
        reversed = reversed*10 + digit
        a = a / 10
    }
    
    cond = c == reversed
    fmt.Println(cond)
}
