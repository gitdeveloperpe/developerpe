package main

import (
	"fmt"
	"os"
	"bufio"
)


func main(){

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre : ")
	text,err := reader.ReadString('\n')
	if(err != nil){
		fmt.Println(err)
	}else{
		fmt.Println("Hola "+text)
	}
}