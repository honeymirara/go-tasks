// 3. Выведите день недели по номеру (1-7) с использованием switch
package main

import "fmt"

var weekDays string = "Понедельник"

func main() {
	switch weekDays {
	case "Понек":
		fmt.Println("1")
	case "Вторник":
		fmt.Println("2")
	case "Cреда":
		fmt.Println("3")
	case "Четверг":
		fmt.Println("4")
	case "Пятница":
		fmt.Println("5")
	case "Суббота":
		fmt.Println("6")
	case "Воскресенье":
		fmt.Println("7")
	default:
		fmt.Println("не день неделеи")
	}
}
