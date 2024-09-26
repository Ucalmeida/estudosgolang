package main

import "fmt"

func main() {
	// Memória -> Endereço -> Valor
	// variavel -> ponteiro que tem um endereco na memoria -> valormain()
	a := 10
	fmt.Printf("Ponteiros: %d\n", a)
	fmt.Println("Endereço de Memória:", &a)
	var ponteiro *int = &a
	fmt.Println("Ponteiro:", ponteiro)
	*ponteiro = 20
	fmt.Printf("Valor: %d\n", a)
	b := &a
	fmt.Println("Endereço de Memória de B:", b)
	fmt.Println("Valor na Memória de B:", *b)

	*b = 30

	fmt.Printf("Valor de A: %d\n", a)
	fmt.Printf("Valor de B: %d\n", *b)
}
