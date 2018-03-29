package main 

import "fmt"

func main() {
    primos := [6]int{2, 3, 5, 7, 11, 13}
    var s []int = primos[1:4]
    fmt.Println(s)

    const a string="mi constante"

    fmt.Println(a)
    fmt.Println(suma(3,4))

} 

func suma(a int,b int) int {
	return a+b
}