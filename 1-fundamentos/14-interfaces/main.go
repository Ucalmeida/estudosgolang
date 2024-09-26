package main

import "fmt"

// Pesquisar com IA para aprender mais sobre isso - Interface
type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

type Empresa struct {
	Nome string
}

func (e *Empresa) Desativar() {
}

func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func Desativacao( pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	urian := Cliente{
		Nome:  "Urian",
		Idade: 47,
		Ativo: true,
	}
	Desativacao(&urian)
	fmt.Println(urian)
	minhaEmpresa := Empresa{
		Nome: "Empresa",
	}
	fmt.Println(minhaEmpresa)
	Desativacao(&minhaEmpresa)
	fmt.Println(minhaEmpresa)
}
