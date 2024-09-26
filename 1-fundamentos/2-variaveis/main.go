package main

const a string = "Hello World!"

var (
	b bool // Aqui foi inferido o valor false, mesmo que 
		   // o valor não tenha sido atribuído de forma intencional por nós
	c int // Aqui foi inferido o valor 0, mesmo que
		  // o valor não tenha sido atribuído de forma intencional por nós
	d float64 // Aqui foi inferido o valor 0.0, mesmo que
              // o valor não tenha sido atribuído de forma intencional por nós
	s string // Aqui foi inferido o valor "", mesmo que 
			  // o valor não tenha sido atribuído de forma intencional por nós
) // Declaração de escopo global. Todas as funções que eu possuir no programa tem acesso a elas
// Posso declarar e atribuir em seguida
// Lembrar também que var é mutável e const não é
// Ao atribuir um valor a uma variável, não pode mudar o 
// tipo de valor atribuído. Se uma var nasce bool, vai morrer bool

func main() {
	println("Variável b antes da atribuição:", b)
	b = true
	println("Variável b após da atribuição:", b)
	println(a)
	println("Valor inteiro antes da atribuição:", c)
	println("Valor Float antes da atribuição:", d)
	println("Valor String antes da atribuição:", s)
	c = 4
	d = 5.3
	s = "Urian"
	println("Valor inteiro após atribuição:", c)
	println("Valor Float após atribuição:", d)
	println("Valor String após atribuição:", s)
	// Se eu tentar acessar a variável e que foi declarada dentro de uma função, 
	// ela não será acessada. Terei como valor undefined
	x()
	xman()
}

func x() {
	println("Função x")
	var e string = "Olá" // Ao declarar uma var local, é obrigatório usar ela, senão o compilador não irá permitir rodar o programa
	var a string = "Hello, again"
	// usar um short hand
	x := "X" // Uso do := somente no momento da declaração da variável
	x = "XY" // Para mudar o valor da string somente usa o =
	// É como qualquer atribuição de variável. var x = "X" -> Depois x = "XY"
	println("Tenho acesso a variável local", e)
	println("A variável a local não é a constante a global. São diferentes:", a)
	println("Tenho acesso a variável global", s)
	println("Uso da SHort Hand:", x)
}

func xman() {
	// Declarando e atribuindo variáveis locais
	vampira := "Vampira"
	wolverine := "Logan"
	println(vampira)
	println("Nome do Wolverine:", wolverine)
}