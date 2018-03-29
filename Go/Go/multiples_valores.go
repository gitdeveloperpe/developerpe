package main

import "fmt"

func calculo(n int)(c,d,e,f int){
	c = n*10
	d = n*20
	e = n*30
	f = n*50
	return
}



func main() {
	fmt.Println(calculo(20))
	c1,c2,c3,c4 := calculo(10)
	fmt.Println(c1,c2,c3,c4)
	_,d2,_,_ := calculo(5)
	fmt.Println(d2)
}