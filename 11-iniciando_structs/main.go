package main

import "fmt"

type Cliente struct {
	Nome string
	Idade int
	Ativo bool
}

func main() {
	urian := Cliente {
		Nome: "Urian",
		Idade: 47,
		Ativo: true,
	}
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", urian.Nome, urian.Idade, urian.Ativo)
	
	// Modificando valores da struct
	urian.Idade = 100
	urian.Ativo = false
	fmt.Printf("\nNome: %s, Idade: %d, Ativo: %t", urian.Nome, urian.Idade, urian.Ativo)
}