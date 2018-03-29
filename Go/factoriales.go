package main 
import "fmt"



func factorial(aux int) {
	arreglo:=make([] int ,aux)
	fact:=1
	i:=1
	for  i < aux {
		fact=fact*i
		arreglo[i]=fact
		i++
	}
	
	fmt.Println(arreglo)
	
}




func main() {

	var aux int
	fmt.Println("cuantos elementos desea? ")
	fmt.Scanf("%d\n",&aux)	

	factorial(aux)

}