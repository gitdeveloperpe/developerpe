package main

import "fmt"

func main() {
	mapa := make(map[string]string)
	fmt.Println(mapa)

	mapa2 := make(map[string]string,10)
	fmt.Println(mapa2)

	edades := map[string]int{
		"oliver":21,
		"anyi":20,
		"jesus":40,
		"luis":24,
	}

	fmt.Println(edades)

	mapa["piura"] = "Peru"

	fmt.Println(mapa)

	delete(mapa,"piura")
	fmt.Println(mapa)
}