package main

import "fmt"

func main() {
	salarios := map[string]int{"Joao": 1000, "Maria": 2000, "Jose": 3000}
	fmt.Println(salarios["Joao"])
	fmt.Println(salarios)
	delete(salarios, "Joao")
	salarios["Urian"] = 18000
	fmt.Println(salarios["Urian"])
	fmt.Println(salarios)

	sal := make(map[string]int)
	sal["Jenser"] = 3500
	fmt.Println(sal)

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	// Blank identifier
	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}
}