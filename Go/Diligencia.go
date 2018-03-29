package main 

import "fmt"



func Cantidad_Estados(n int) string {
	
	clase_estado:=""

	if n==1 {
		clase_estado="Etapa 1"
	}else if n==2 {
		clase_estado="Etapa 2"
	}else if n==3 {
		clase_estado="Etapa 3"
	}else{
		clase_estado="Etapa 4"
	}

	return clase_estado
}



type  Nodo_estado1 struct{

	anterior1 *Nodo_estado1
	distancia1 int
	distancia2 int
	distancia3 int
	
}

type Nodo_estado2 struct{
	anterior1 *Nodo_estado1
	
	distancia1 int
	distancia2 int
	distancia3 int
}


type Nodo_estado3 struct{

	anterior1 *Nodo_estado2
	anterior2 *Nodo_estado2
	anterior3 *Nodo_estado2
	distancia1 int
	distancia2 int
	distancia3 int	
}

type Nodo_estado4 struct{
	anterior1 *Nodo_estado3
	anterior2 *Nodo_estado3
	anterior3 *Nodo_estado3
	distancia1 int
	distancia2 int

}

type Nodo_estado5 struct{
	anterior1 *Nodo_estado4
	anterior2 *Nodo_estado4
	distancia1 int
	distancia2 int

}

func main() {
	
	fmt.Println("\n\n------------Bienvenidos a  la solucion del problema de la Diligencia----------")
	fmt.Println("# La cantidad de estados que tomaremos para este ejercicio sera de ",4)
	//variable n que es la cantidad de estados que tiene el problema
	//variable x que es el estado inmediato en la etapa n
	//variable c que es el costo
	//variable suma que  es el costo total de la ruta
	ayudante:=Cantidad_Estados(4)	
	fmt.Println("#Por lo tanto empezaremos a solucionar el problema por la ",ayudante)



//------------------------------------------
//CREACION DE ESTADOS 
	//ESTADO INICIAL TENEMOS SOLO 1
	A:=new(Nodo_estado1)

	//ETAPA 2 TENEMOS 3 ESTADOS
	B:=new(Nodo_estado2)
	C:=new(Nodo_estado2)
	D:=new(Nodo_estado2)

	//ETAPA 3 TENEMOS 3 ESTADOS
	E:=new(Nodo_estado3)
	F:=new(Nodo_estado3)
	G:=new(Nodo_estado3)

	//ETAPA 4 TENEMOS 2 ESTADOS
	H:=new(Nodo_estado4)
	I:=new(Nodo_estado4)


	//ESTADO FINAL ES SOLO 1
	J:=new(Nodo_estado5)

//-------------------------------------
//ASIGNACION DE COSTOS A LOS CAMINOS EN CADA ETAPA

	
  //ENLAZAMOS ESTADOS POR ETAPA
	//ETAPA 4
	J.anterior1=H
	J.anterior2=I

	//ETAPA 3
	H.anterior1=E
	H.anterior2=F
	H.anterior3=G

	I.anterior1=E
	I.anterior2=F
	I.anterior3=G

	//ETAPA 2
	E.anterior1=B
	E.anterior2=C
	E.anterior3=D

	F.anterior1=B
	F.anterior2=C

	F.anterior3=D

	G.anterior1=B
	G.anterior2=C
	G.anterior3=D

	//ETAPA 1
	B.anterior1=A
	C.anterior1=A
	D.anterior1=A
	A.anterior1=nil




  //ASIGNAMOS COSTOS A DISTANCIAS ENTRE ETAPAS
	//Etapa 1 valores:2		3	4
	//Etapa 2 valores: 7,4,6	3,2,4	4,1,5
	//Etapa 3 valores: 1,4	6,3		3,3
	//Etapa 4 valores: 3	4

	A.distancia1=2
	A.distancia2=4
	A.distancia3=3

	B.distancia1=7
	B.distancia2=4
	B.distancia3=6

	C.distancia1=3
	C.distancia2=2
	C.distancia3=4

	D.distancia1=4
	D.distancia2=1
	D.distancia3=5

	E.distancia1=1
	E.distancia2=4

	F.distancia1=6
	F.distancia2=3

	G.distancia1=3
	G.distancia2=3

	H.distancia1=3
	I.distancia1=4


//EMPEZAMOS A BUSCAR RECORRIDO MINIMO
	
	suma:=0
	//ETAPA 4
	fmt.Println("\n1.- Al estar en el estado J elegimos el camino que menos nos cuesta , H o I \n")
	if H.distancia1< I.distancia1 {
		
		suma=suma+H.distancia1
		fmt.Printf("\n2.- Elegimos el camino de H a J por ser el menos costoso, por lo tanto hasta \nahora el coste es %d ,ahora elegimos el camino menos costoso de E,F,G hacia H \n",suma)
		//ETAPA 3

		if  E.distancia1<F.distancia1 && E.distancia1<G.distancia1{
			suma=suma+E.distancia1
			fmt.Printf("\n3.- Elegimos el camino de E a H por ser el menos costoso, por lo tanto hasta \nahora el coste es %d ,ahora elegimos el camino menos costoso de B,C,D hacia E \n",suma)
			//Etapa 2
			
			if B.distancia1<C.distancia1 && B.distancia1<D.distancia1 {
				suma=suma+B.distancia1
				fmt.Printf("\n4.- Elegimos el camino de E a B por ser el menos costoso, por lo tanto hasta \nahora el coste es %d ,ahora elegimos el camino menos costoso de B hacia A \n",suma)
			//ETAPA 1
				suma=suma+A.distancia1
				fmt.Printf("\n5.- Elegimos el camino de B a A por ser el menos costoso, por lo tanto\n el coste total es %d ,y asi  ya hemos llegado a nuestro destino \n",suma)
			}else{
				if C.distancia1<B.distancia1 && C.distancia1<D.distancia1 {
					suma=suma+C.distancia1
					fmt.Printf("\n4.- Elegimos el camino de E a C por ser el menos costoso, por lo tanto hasta \nahora el coste es %d ,ahora elegimos el camino menos costoso de C hacia A \n",suma)
				//ETAPA 1
					suma=suma+A.distancia2
					fmt.Printf("\n5.- Elegimos el camino de C a A por ser el menos costoso, por lo tanto\n el coste total es %d ,y asi  ya hemos llegado a nuestro destino \n",suma)
				}else{
					suma=suma+D.distancia1
					fmt.Printf("\n4.- Elegimos el camino de E a D por ser el menos costoso, por lo tanto hasta \nahora el coste es %d ,ahora elegimos el camino menos costoso de B hacia A \n",suma)
				//ETAPA 1
					suma=suma+A.distancia3
					fmt.Printf("\n5.- Elegimos el camino de D a A por ser el menos costoso, por lo tanto\n el coste total es %d ,y asi  ya hemos llegado a nuestro destino \n",suma)
				}
			}
		}else{
			if F.distancia1<E.distancia1 && F.distancia1<G.distancia1{
				suma=suma+F.distancia1
				fmt.Printf("\n2.- Elegimos el camino de F a J por ser el menos costoso, por lo tanto hasta \nahora el coste es %d ,ahora elegimos el camino menos costoso de E,F,G hacia H \n",suma)
				//ETAPA 2
				if B.distancia2<C.distancia2 && B.distancia2<D.distancia2 {
					suma=suma+B.distancia2
				//ETAPA 1
					suma=suma+A.distancia1
				}else{
					if C.distancia2<B.distancia2 && C.distancia2<D.distancia2 {
						suma=suma+C.distancia2
					//ETAPA 1
						suma=suma+A.distancia2
					}else{
						suma=suma+D.distancia2
					//ETAPA 1
						suma=suma+A.distancia3
					}
				}
			}else{
				suma=suma+G.distancia1
				//ETAPA 2
				if B.distancia3<C.distancia3 && B.distancia3<D.distancia3 {
					suma=suma+B.distancia3
				//ETAPA 1
					suma=suma+A.distancia1
				}else{
					if C.distancia3<B.distancia3 && C.distancia3<D.distancia3 {
						suma=suma+C.distancia3
					//ETAPA 1
						suma=suma+A.distancia2
					}else{
						suma=suma+D.distancia3
					//ETAPA 1
						suma=suma+A.distancia3
					}
				}
			}
		}
	}else{
		suma=suma+I.distancia1
		//ETAPA 3
		if E.distancia2<F.distancia2 && E.distancia2<G.distancia2 {
			suma=suma+E.distancia2
		//ETAPA 2
			if B.distancia1<C.distancia1 && B.distancia1<D.distancia1 {
				suma=suma+B.distancia1
		//ETAPA 1
				suma=suma+A.distancia1
			}else{
				if C.distancia1<B.distancia1 && C.distancia1<D.distancia1 {
					suma=suma+C.distancia1
				//ETAPA 1
					suma=suma+A.distancia2
				}else{
					suma=suma+D.distancia1
				//ETAPA 1
					suma=suma+A.distancia3
				}
			}
		}else{
			//ETAPA 3
			if F.distancia2<E.distancia2 && F.distancia2<G.distancia2{
				suma=suma+F.distancia2
			//ETAPA 2
				if B.distancia2<C.distancia2 && B.distancia2<D.distancia2 {
					suma=suma+B.distancia2
			//ETAPA 1
					suma=suma+A.distancia1
				}else{
					if C.distancia2<B.distancia2 && C.distancia2<D.distancia2 {
						suma=suma+C.distancia2
					//ETAPA 1
						suma=suma+A.distancia2
					}else{
						suma=suma+D.distancia2
					//ETAPA 1
						suma=suma+A.distancia3
					}
				}
			}else{
				suma=suma+G.distancia2
			//ETAPA 2
				if B.distancia3<C.distancia3 && B.distancia3<D.distancia3 {
					suma=suma+B.distancia3
			//ETAPA 1
					suma=suma+A.distancia1
				}else{
					if C.distancia3<B.distancia3 && C.distancia3<D.distancia3 {
						suma=suma+C.distancia3
					//ETAPA 1
						suma=suma+A.distancia2
					}else{
						suma=suma+D.distancia3
					//ETAPA 1
						suma=suma+A.distancia3
					}
				}
			}
		}
	}


}