package main

// Posso criar meus tipo de variáveis personalizado
// Tornando meu código mais expressivo
type ID int
type NOME string

var (
	i ID = 1
	meuNome NOME = "Urian"
)

func main() {
	println(i)
	println("Meu nome é", meuNome)
}
