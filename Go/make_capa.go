package main 

import "fmt"


//Len o longitud es la cantidad de elementos que hay en ese momento en el arreglo al que apunta
// Cap o capacidad es la cantidad total de elementos que entran en el arreglo al que apunta

func main() {

	
	//ice=append(slice,445)
	var n,aux int

	//entrada:=bufio.NewReader(os.Stdin)
	
	//n,err:=entrada.ReadString('\n')
	//aux,_:=strconv.Atoi(n)
	
	fmt.Println("Cuantos elementos desea ingresar")
	fmt.Scanf("%d\n",&aux)

	slice:= make([] int ,0)

	for i:=0;i<aux;i++{
		fmt.Printf("Ingrese el elemento %d del arreglo \n",i)
		//arreglo[i],err=entrada.ReadString('\n')
		fmt.Scanf("%d\n",&n)
		slice=append(slice,n)

	}
//for i :=0;i<lent(slice);i++{
	//fmt.Println("\n",slice[i])
//}


	fmt.Println(slice)

	
}