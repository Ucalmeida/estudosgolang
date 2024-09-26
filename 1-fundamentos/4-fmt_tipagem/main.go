package main

// Utilizamos o () para diversas importações
import (
	"fmt"
)

type ID int

var (
	e float64 = 3.14
	i ID = 1
)

func main() {
	fmt.Printf("O tipo de E é %T", e) // %T mostra o tipo da variável
	fmt.Printf("\nO valor de E é %v", e) // %v mostra o valor da variável(Usar %v para todos os tipos)
	fmt.Printf("\nO tipo de I é %T", i)
	fmt.Printf("\nO valor de I é %v", i)
	// %s é um placeholder para strings
}