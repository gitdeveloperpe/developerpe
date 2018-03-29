package main

import "fmt"
import "strconv"
//import{
//	"fmt"
//	"strconv"
//}

func main() {
	
	//Conversion de enteros a cadena
	edad:=30
	edad_str:=strconv.Itoa(edad)

	fmt.Println("La edad es : "+edad_str)

	//conversion de cadena a enteros
	//Primer tipo de declaracion de vaariables
	var cadena string ="25"
	numero,_:=strconv.Atoi(cadena)
	//Usamos el Subguion para recibir el segundo valor que arroja el metodo ATOI y desecharlo porque no nos interesa
	fmt.Println(numero+200)
}