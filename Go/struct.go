package main 
import "fmt"
//import "strconv"



type Persona struct{
	nombre,apellido,sexo string
	edad int 
	Calcula  func()
}

func main() {
	oliver:=new(Persona)
	//fmt.Println(oliver.nombre)
	var aux string
	fmt.Println("Escriba su nombre: ")
	fmt.Scanf("%s\n",&aux)
	oliver.nombre=aux

	
	fmt.Println(oliver.nombre)

}