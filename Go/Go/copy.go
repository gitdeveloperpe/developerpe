package main

import "fmt"

func main() {
	origen := []int{1,2,3,5,6,7,8}
	destino := make([]int,len(origen))
	fmt.Println(destino)
	copy(destino,origen)

	fmt.Println(destino)
}