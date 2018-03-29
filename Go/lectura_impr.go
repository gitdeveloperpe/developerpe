package main

import 	"fmt"
import	"bufio"
import	"os"
//import  "strconv"



func main() {
	numero:=22
	apellido:="sandoval"
	fmt.Printf("El numero es %d \n",numero)
	fmt.Println("El nombre es "+apellido)

	//Impresion de booleanos

	bander:=true
	fmt.Printf("%t\n",bander)

	//%f cantidad de decimales abreviada
	//%e y %v nos daran decimales mas grandes, usado mas para datos cientificos o estadisticos
	decimal:=24.567
	fmt.Printf("%f\n",decimal)

	//Imprimira el numero dependiendo donde se coloque el %d,e,v,s , si lo colocamos en medio, la primera variable la imprimita en el medio

//---------------------------------------------------------------------------------------------------------
	//Primera forma de lectura de datos

	//Para enteros
	//var num int
	//fmt.Print("Ingresa un numero: ")
	//fmt.Scanf("%d\n",&num)
	//fmt.Println("El numero ingresado es: "+strconv.Itoa(num))


	//Para cadenas
	//var nombre string
	//fmt.Print("Escriba su nombre: ")
	//fmt.Scanf("%v\n",nombre)
	//fmt.Println("el nombre ingresado es :"+nombre)


//----------------------------------------------------------------------------------------------------------

	//Segunda forma de lectura de datos

	//declaramos un Reader 
	//STDIN es el lector que utiliza y es de la libreria OS
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese su nombre: ")
	//Reader retorna 2 valores, la lectura de datos y un error
	nombre,err:= reader.ReadString('\n')
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("Holaa "+nombre)
	}


}