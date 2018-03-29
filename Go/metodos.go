package main 

import "fmt"



type Persona struct {
	nombre,apellido string
	edad int

}


//un metodo clona un objeto, por lo tanto al editar, editas no el original sino la copia
func (this Persona) Imprimir_Nombre()string {
	return this.nombre+" "+this.apellido
}

//por eso crearemos un puntero para solo apuntar al objeto creado y asi editar el original y no la copia
func (this *Persona)set_Nombre(n string) {
	this.nombre=n
}	

func main() {
	
	var aux string
	
	objeto:=new(Persona) 
	fmt.Println("Ingrese su nombre : ")
	fmt.Scanf("%s\n",&aux)
	objeto.nombre=aux
	fmt.Println("Ingrese su apellido: ")
	fmt.Scanf("%s\n",&aux)
	objeto.apellido=aux
	fmt.Println(objeto.Imprimir_Nombre())

	var bander bool 
	fmt.Println("Desea editar su nombre? ")
	fmt.Scanf("%s\n",&aux)
	if aux=="si" {
		bander=true
	}
		
	if bander==true {
		fmt.Println("Escriba su nombre :")
		fmt.Scanf("%s\n",&aux)
		objeto.set_Nombre(aux)
		fmt.Println(objeto.Imprimir_Nombre())
	}

}