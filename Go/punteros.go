package main

import "fmt"


func main() {

/*
	Es una direccion de memoria
	en lugar del valor,tenemos la direccion en la que esta el valor
	X,Y =>as13221d =>5
	X=>as13221d=>6
	Y¿?=>6


*/

	var x,y *int
	entero:=5
	//asignamos al puntero x la direccion en la que se encuentra nuestro valor 5
	x=&entero
	y=&entero
	//imprimimos las direcciones de memoria de x,y
	fmt.Println(*x,*y)

	//cambiamos valor al que apunta y sin usar y , ademas que cambiamos el valor al que apunta entero sin usarlo
	*x=10

	fmt.Println(*x,*y,entero)


}