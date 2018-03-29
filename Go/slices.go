package main
import "fmt"

func main() {
	
	var slice [] int
 
 	if slice==nil{
 		fmt.Println("Esta vacio")
 	}else{
 		fmt.Println(slice)	
 	}


 	// Un slice tiene dentro
 	//-Un puntero a un arreglo
 	//-Capacidad
 	//-Longitud del arreglo al que apunta


 	//------------------------
 //Creaion de un Slice a partir de un arreglo
 	//Creamos un arreglo
 	arreglo := [4] int {1,2,3}
 	//Creamos un slice a partir de un arreglo
 	slices:= arreglo[0:2]
 	//El primer valor del corchete es a partir de donde se tomaran los datos del arreglo, es decir la posicion inicial a tomar
 	//El segundo valor es la posicion hasta la que tomara datos,sin tomar esa posicion, es decir tomara hasta la posicion n-1
 	fmt.Println(slices)
}