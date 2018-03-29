// Range

package main 

import "fmt"

func main() {
	 numeros := []int{2,4,6}
	 suma := 0
	 for _, numero := range numeros {
	 	suma = suma+ numero
	 }

	 fmt.Println("suma:", suma)

	 for i, numero := range numeros {
	     if numero == 3{
	     	fmt.Println("index:", i)
	     }
	     	
	 } 
}

/*
Para un array, que  es una estructura tipo indice-valor es decir para 
 numeros := []int{2,4,6}  tiene
      posicion 0 = 2
      posición 1 = 4
      posición 2 = 6
La función range devuelve 0 y 2 luego 1 y 4, y luego 2 y 6.

Para un mapa devuelve la clave valor, por ejemplo para
  algo := map[string]string{"a": "auto", "b": "bebé"} tiene
       a = auto
       b = bebé
La función range devuelve "a" y "auto" , luego "b" y "bebé"

Entonces, como siempre devuelve 2 valores, sí en algún momento no vas a usar el primero de ellos, 
puedes usar ese "comodin", "_". Sí no lo usases, 
o bien sólo obtendrías el primer valor o bien tendrías el valor en una variable que no usas y el programa no funciona.
*/