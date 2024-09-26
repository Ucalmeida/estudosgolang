package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum4(50, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
	
	fmt.Println(sum(1, 2))
	fmt.Println(sum2(4, 2))
	fmt.Println(sum3(51, 2))
	fmt.Println(sum4(51, 2))
}

func sum(a int, b int) int {
	return a + b
}

// Se for do mesmo tipo, pode declarar assim
func sum2(a, b int) int {
	return a + b
}

// Se for retornar mais de um valor, pode declarar assim
func sum3(a, b int) (int, bool) {
	if a + b >= 50 {
		return a + b, true
	}
	return a + b, false
}

// Se for retornar com erro 
func sum4(a, b int) (int, error) {
	if a + b > 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}