package main 

import "fmt"

func main() {
	var arreglo [5] int
	fmt.Println(arreglo)

	arreglo2 := [5]int{1,2,3,4,5}
	fmt.Println(arreglo2)
	for i := 0; i < len(arreglo2); i++ {
		fmt.Println(arreglo2[i])
	}

	var matriz [3][4]int
	fmt.Println(matriz)
	for i := 0; i < len(matriz); i++ {
		matriz[i][2]=i+1
	}
	fmt.Println(matriz)
}