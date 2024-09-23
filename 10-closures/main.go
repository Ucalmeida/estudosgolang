package main

import "fmt"

func main() {
	total := func() int {
		return soma(1, 3, 45, 55, 654, 7645, 534, 543) * 2
	}()
	fmt.Println(total)

	func() {
		fmt.Println("Rodar Web Service")
	}()
}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}