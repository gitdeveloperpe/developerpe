package main

import "fmt"


func calculo(a,b,c int)(r int){
	a = b*c
	b = a-c
	c = a+b
	if a>b {
		panic("Detente!")
	}
	r = a+b+c
	return c
}


func main() {
	fmt.Println(calculo(5,6,7))
}