package main

import "fmt"

func main() {
	slice := make([] int,3,4)
	slice = append(slice,2)
	slice = append(slice,3)
	slice = append(slice,4)
	fmt.Println(slice)
	fmt.Println(cap(slice))
}