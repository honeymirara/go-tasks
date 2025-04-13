// 7. Преобразуйте целое число в строку и строку в число.
// Вход: 123 → "123" (преобразовать в строку)
// Вход: "456" → 456 (преобразовать в число)

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input interface{}

	input = 123
	handleInput(input)

	input = "578"
	handleInput(input)
}

func handleInput(value interface{}) {
	switch v := value.(type) {
	case int:
		str := strconv.Itoa(v)
		fmt.Println("Число:", v, "→ строка:", str)
	case string:
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("String isn't a number")
		} else {
			fmt.Println("Строка:", v, "→ число:", num)
		}
	default:
		fmt.Println("Тип не поддерживается:", v)
	}
}
