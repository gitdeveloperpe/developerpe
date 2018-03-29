package main

import (
	"fmt"
	"unsafe"
)

func main() {


	//var entero_8 int8   //Todos los enteros de 8-bits con signo (-128 a 127)
	//var entero_16 int16 //Todos los enteros de 16-bits con signo (-32768 a 32767)
	var entero_32 int32 //Todos los enteros de 32-bits con signo (-2147483648 a 2147483647)
	var entero_64 int64 //Todos los enteros de 64-bits con signo (-9223372036854775808 a 9223372036854775807)
	


	//var uentero_8 uint8   //Todos los enteros de 8-bits sin signo (0 a 255)
	//var uentero_16 uint16 //Todos los enteros de 16-bits sin signo (0 a 65535)
	//var uentero_32 uint32 //Todos los enteros de 32-bits sin signo (0 a 4294967295)
	//var uentero_64 uint64 //Todos los enteros de 64-bits sin signo (0 a 18446744073709551615)
	


	//var entero_Byte byte //para uint8
	var entero_Rune rune //para int32
		


	entero_32 = 100 
	entero_64 = 32

	//fmt.Println(entero_32 + entero_64)
	fmt.Println(entero_32 + int32(entero_64))

	//****************************************
	//Suma int32 y rune

	entero_Rune = 67 //rune
	fmt.Println(entero_32 + entero_Rune)

	//****************************************
	//Suma int32 y int

	var entero_Int int 
	entero_Int = 78
	//fmt.Println(entero_32 + entero_Int)
	fmt.Println(entero_32 + int32(entero_Int))

	fmt.Println(unsafe.Sizeof(entero_Int), unsafe.Sizeof(entero_32))

}