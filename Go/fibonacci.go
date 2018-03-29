package main 

import "fmt"


func main() {
	
	var n int
	a:=0
	b:=1

	fmt.Println("Cuantos valores desea ingresar?")
	fmt.Scanf("%d",n)
	fmt.Println(a,b)

	suma:=a+b

	for suma<n {
		fmt.Println(suma)
		a=b
		b=suma
		suma=a+b
	}
	


	
}