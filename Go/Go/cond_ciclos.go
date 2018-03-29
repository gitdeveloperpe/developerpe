package main

import "fmt"

func main() {
	
//condicionales
	edad:=20
	if edad <10{
		fmt.Println("La edad es menor a 10")
	}else{
		fmt.Println("La edad es mayor a 10")
	}

//ciclos
	i:=0
	for {
		fmt.Println("1")
		i++
		if (i>=5){
			break
		}else{			
			continue
		}
	}

}