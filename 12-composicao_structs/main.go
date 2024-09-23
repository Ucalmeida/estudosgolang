package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero int
	Cidade string
	Estado string
}

type Cliente struct {
	Nome string
	Idade int
	Ativo bool
	// Compondo structs
	// Endereco

	// Criando uma propriedade do tipo Endereco
	Address Endereco
}

func main() {
	urian := Cliente {
		Nome: "Urian",
		Idade: 47,
		Ativo: true,
	}
	
	// Posso declarar Cidade dessa forma, asim como o Estado
	// urian.Cidade = "Aracaju"
	// urian.Estado = "SE"
	// fmt.Printf("Nome: %s | Idade: %d | Ativo: %t | Estado: %s | Cidade: %s", urian.Nome, urian.Idade, urian.Ativo, urian.Estado, urian.Cidade)

	// Mas também posso fazer da seguinte forma:
	// urian.Endereco.Cidade = "Aracaju"
	// urian.Endereco.Estado = "SE"
	// fmt.Printf("Nome: %s | Idade: %d | Ativo: %t | Estado: %s | Cidade: %s", urian.Nome, urian.Idade, urian.Ativo, urian.Endereco.Estado, urian.Endereco.Cidade)
	
	// Atribuindo Endereco como tipo a uma propriedade ba struct, declaro dessa forma:
	urian.Address.Cidade = "Aracaju"
	urian.Address.Estado = "SE"
	fmt.Printf("Nome: %s | Idade: %d | Ativo: %t | Estado: %s | Cidade: %s", urian.Nome, urian.Idade, urian.Ativo, urian.Address.Estado, urian.Address.Cidade)

}