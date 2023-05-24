package main

import (
	"fmt"
	"os"
)

func suma(num1, num2 int) int {
	return num1 + num2
}

func resta(num1, num2 int) int {
	return num1 - num2
}

func mul(num1, num2 int) int {
	return num1 * num2
}

func div(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("no se puede dividir entre cero")
	}
	return num1 / num2, nil
}

func main() {
	var num1, num2 int
	var oper string

	fmt.Print("Qué operación quiere hacer +, -, *, /: ")
	fmt.Scan(&oper)

	fmt.Print("Ingrese el número 1: ")
	fmt.Scan(&num1)

	fmt.Print("Ingrese el número 2: ")
	fmt.Scan(&num2)

	var resultado interface{}
	var err error

	switch oper {
	case "+":
		resultado = suma(num1, num2)
	case "-":
		resultado = resta(num1, num2)
	case "*":
		resultado = mul(num1, num2)
	case "/":
		resultado, err = div(float64(num1), float64(num2))
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Operación no válida")
		os.Exit(1)
	}

	fmt.Println("Resultado:", resultado)
}
