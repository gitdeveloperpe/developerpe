package main 
import "fmt"



func (this *Nodo)Imprimir_Real() {
	
		for i := len(this.pila_real)-1; i>=0; i-- {
			if this.pila_real[i]!=0 {
				fmt.Println(this.auxiliar_real[i])
				this.pila_real=append(this.pila_real,this.auxiliar_real[i])

			}
		}
	
}

func (this *Nodo2)Imprimir_Entero() {
	for i := len(this.pila_entero)-1; i>=0; i-- {
			if this.pila_entero[i]!=0 {
				fmt.Println(this.auxiliar_entero[i])
				this.pila_entero=append(this.pila_entero,this.auxiliar_entero[i])

			}
		}
}
func (this *Nodo)Desapilar_real() {
	for i := len(this.pila_real)-1; i>=0; i-- {
		if this.pila_real[i]!=0 {
			this.auxiliar_real=append(this.auxiliar_real,this.pila_real[i])
		}
	}
}


func (this *Nodo) Apilar_real(n float64) {
		this.pila_real=append(this.pila_real,n)
}



func (this *Nodo2)Desapilar_entero() {
	for i := len(this.pila_entero)-1; i>=0; i-- {
		if this.pila_entero[i]!=0 {
			this.auxiliar_entero=append(this.auxiliar_entero,this.pila_entero[i])
		}
	}

}

func (this *Nodo2)Apilar_entero(n int) {
	this.pila_entero=append(this.pila_entero,n)
}

type Nodo struct{	
	pila_real []float64
	auxiliar_real []float64
	
}
type Nodo2 struct{
	pila_entero [] int
	auxiliar_entero [] int
}

func Ayuda_Apilar_real() float64{
	var ayuda float64 
	fmt.Println("Ingrese el valor: ")
	fmt.Scanf("%d\n",&ayuda)
	return ayuda
}
func Ayuda_Apilar_enteros() int{
	var ayuda int 
	fmt.Println("Ingrese el valor: ")
	fmt.Scanf("%d\n",&ayuda)
	return ayuda
}

func Menu() {
	
	var aux int
	var aux1 int
	objeto:=new(Nodo)
	objeto2:=new(Nodo2)
	if aux!=3 {
		fmt.Println("-------Bienvenido-----------")
		fmt.Println("Elija una opcion---> ")
		fmt.Println("1.- Pila de real\n2.- Pila de enteros\n3.- Salir")
		fmt.Scanf("%d\n",&aux)
		if aux==1{
			for aux1!=3{
				fmt.Println("-------Bienvenido al menu de Pilas Reales-----------")				
				fmt.Println("Elija una opcion---> ")
				fmt.Println("1.-Apilar \n2.-Imprimir \n3.-Salir")
				fmt.Scanf("%d\n",&aux1)
				if aux1==1 {
					objeto.Apilar_real(Ayuda_Apilar_real())
				}else if aux1==2 {
					objeto.Desapilar_real()
					objeto.Imprimir_Real()
				}
			}
				
					
		}else{
			for aux1!=3 {
				fmt.Println("-------Bienvenido al menu de Pilas Enteras-----------")				
				fmt.Println("Elija una opcion---> ")
				fmt.Println("1.-Apilar \n2.-Imprimir \n3.-Salir")
				fmt.Scanf("%d\n",&aux1)
				if aux1==1 {
					objeto2.Apilar_entero(Ayuda_Apilar_enteros())
				}else if aux1==2 {
					objeto2.Desapilar_entero()
					objeto2.Imprimir_Entero()
				}
			}
			
		}			
	
	}
}

func main() {
	Menu()
		
	/*prueba:=new(Nodo2)

	prueba.Apilar_entero(Ayuda_Apilar_enteros())
	prueba.Apilar_entero(Ayuda_Apilar_enteros())

	prueba.Desapilar_entero()
	prueba.Imprimir_Entero()

	


	*/
}	