package main

import (
	"fmt"
	"strconv"
)

func main() {
	numero := 32
	//numero_str := strconv.Itoa(numero)

	numero2 := "45"
	numero2_int,err := strconv.Atoi(numero2)
	fmt.Println(numero+numero2_int , err)
}