package main

// Variáveis globais
var (
	s string = "Hello, World!"
	i int = 42
	f float64 = 3.14
	b bool
)

func main() {
	// Imprimindo as variáveis globais
	println(s)
	println(i)
	println(f)
	xman()
}

func xman() {
	// Declarando e atribuindo variáveis locais
	vampira := "Vampira"
	wolverine := "Logan"
	println(vampira)
	println("Nome do Wolverine:", wolverine)
	println("Global Booleana sem atribuição:", b)
	b = true
	println("Global Booleana com atribuição:", b)
}