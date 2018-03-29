package main

import "fmt"

func main() {
	
	//Copy copia el minimo de las capacidades de ambos arreglos o slices


	//ice=append(slice,445)
	var n,aux int

	//entrada:=bufio.NewReader(os.Stdin)
	
	//n,err:=entrada.ReadString('\n')
	//aux,_:=strconv.Atoi(n)
	
	fmt.Println("Cuantos elementos desea ingresar")
	fmt.Scanf("%d\n",&aux)

	slice:= make([] int ,aux,aux)

	for i:=0;i<aux;i++{
		fmt.Printf("Ingrese el elemento %d del arreglo \n",i)
		//arreglo[i],err=entrada.ReadString('\n')
		fmt.Scanf("%d\n",&n)
		//slice=append(slice,n)
		slice[i]=n

	}
//for i :=0;i<lent(slice);i++{
	//fmt.Println("\n",slice[i])
//}


//Copiar cantidad y cuales son los datos elegidos por el usuario 
	var aux2,aux3 int
	fmt.Println("Cuantos elementos desea copiar?")
	fmt.Scanf("%d\n",&aux2)
	slice2:=make([] int,aux2)


	fmt.Println("Cuales desea copiar?")
	

	for i := 0; i < aux2; i++ {
		fmt.Scanf("%d\n",&aux3)	
		for j := 0; j < aux; j++ {
			
			if aux3==slice[j] {
				slice2[i]=aux3
			}

		}
	}
	

	//copy(slice2,slice)


	fmt.Println(slice)
	fmt.Println(slice2)

}