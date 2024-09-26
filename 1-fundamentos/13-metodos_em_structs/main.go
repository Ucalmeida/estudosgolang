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
	Endereco
}

// O * em Cliente indica que estamos passando um ponteiro para o Cliente
// Sem esse *, Cliente estava recebendo uma cópia do struts original, o que não altera o valor de urian.Desativar() para false
// Somente era alterado o valor para false dentro da função. Quando chamava o valor dentro da func main, o valor 
// permanecia true
func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func main() {
	urian := Cliente {
		Nome: "Urian",
		Idade: 47,
		Ativo: true,
	}
	urian.Desativar()
	fmt.Printf("Nome do Cliente: %s | Idade do Cliente: %d | Cliente está ativo? %t", urian.Nome, urian.Idade, urian.Ativo)
}
