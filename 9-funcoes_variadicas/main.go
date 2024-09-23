package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
}

// Usado quando não se sabe o número exato de parâmetros
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}