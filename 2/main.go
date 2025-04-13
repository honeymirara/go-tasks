package main

// 2. Оцените успеваемость входных данных:
// 90-100 — "Отлично".
// 75-89 — "Хорошо".
// 50-74 — "Удовлетворительно".
// < 50 — "Неудовлетворительно".
import "fmt"

func main() {
	var value int = 89
	if value >= 90 && value <= 100 {
		fmt.Println("Отлично")
	} else if value >= 75 && value <= 89 {
		fmt.Println("Удовлетворительно")
	} else if value <= 50 {
		fmt.Println("Неудовлетворительно")
	}

}
