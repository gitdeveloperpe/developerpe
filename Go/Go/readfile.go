package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	contenido,err := ioutil.ReadFile("archivo.txt")
	if err != nil {
		panic("Noooooooooooooooooooooooooooooooooo :(")
	}
	fmt.Println(string(contenido))
}