// 5. Проверьте, имеет ли человек право на льготу (меньше 18 или старше 65 лет).
// Вход: 15 → "Льгота есть"
// Вход: 40 → "Льгота нет"
// Вход: 70 → "Льгота есть
package main

import "fmt"

func main() {
	var value int = 45
	switch {
	case value < 18:
		fmt.Println("Льгота есть")
	case value > 65:
		fmt.Println("Льгота есть")
	default:
		fmt.Println("Льготы нет")
	}
}
