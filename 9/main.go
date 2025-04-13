package main

import "fmt"

// 9. Выполните операции с входными числами: Сложение Вычитание Умножение Деление
// Остаток от делени

func main() {
	var input1 int64 = 567
	var input2 int64 = 342
	var operation string = "+"

	switch operation {
	case "+":
		fmt.Println("Сложение:", input1+input2)
	case "-":
		fmt.Println("Вычитание:", input1-input2)
	case "*":
		fmt.Println("Умножение:", input1*input2)
	case "/":
		fmt.Println("Деление:", input1/input2)
	case "%":
		fmt.Println("Остаток от деления:", input1%input2)
	default:
		fmt.Println("Неизвестная операция")
	}
}
