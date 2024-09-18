package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30
	// meuArray[3] = 40 // invalid argument: index 3 out of bounds [0:3]
	
	fmt.Println("Hello", meuArray)
	fmt.Println("Tamanho:", len(meuArray))
	fmt.Println("Último Elemento:", meuArray[len(meuArray)-1])

	for i, v := range meuArray {
		fmt.Printf("Índice: %d, Valor: %d\n", i, v) // %d é um placeholder para inteiros e %v é um placeholder para todos os tipos
		// %s é um placeholder para strings
	}
}